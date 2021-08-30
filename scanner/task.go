package scanner

import (
	"assetMap/vars"
	"fmt"
	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
	"net"
	"strings"
	"sync"
)

// Discern 此包用于倒入 nmap-db.json,以用于 指纹识别
// 当前版本的 指纹识别，仅用于原型设计，验证相关可能性
func Discern(banner string) (string, error) {
patternMatch:
	for _, NmapStructValue := range vars.NmapProbes{
		for _, NmapStructMatch := range NmapStructValue.Matches {
			tmpPattern := ""
			if (NmapStructMatch.PatternFlag != ""){
				tmpPattern = NmapStructMatch.Pattern + "\\" + NmapStructMatch.PatternFlag
			}else {
				tmpPattern = NmapStructMatch.Pattern + NmapStructMatch.PatternFlag
			}
			regex := pcre.MustCompile(tmpPattern, 0)
			if regex.MatcherString(banner, 0).Matches() {
				// 匹配正常
				return NmapStructMatch.Versioninfo.Vendorproductname, nil
				break patternMatch
			}
		}
	}
	return "", nil
}

func GenerateTask(ipList []net.IP, ports []int) ([]map[string]int, int) {
	tasks := make([]map[string]int, 0)

	for _, ip := range ipList {
		for _, port := range ports {
			ipPort := map[string]int{ip.String(): port}
			tasks = append(tasks, ipPort)
		}
	}

	return tasks, len(tasks)
}

func RunTask(tasks []map[string]int) {
	wg := &sync.WaitGroup{}

	// 创建一个buffer为vars.threadNum 的channel
	taskChan := make(chan map[string]int, vars.ThreadNum * 2)

	// 创建vars.ThreadNum个协程
	for i := 0; i < vars.ThreadNum; i++ {
		go Scan(taskChan, wg)
	}

	// 生产者，不断地往taskChan channel发送数据，直接channel阻塞
	for _, task := range tasks {
		wg.Add(1)
		taskChan <- task
	}

	close(taskChan)
	wg.Wait()
}

func Scan(taskChan chan map[string]int, wg *sync.WaitGroup) {
	// 每个协程都从channel中读取数据后开始扫描并入库
	for task := range taskChan {
		for ip, port := range task {
			if strings.ToLower(vars.Mode) == "syn" {
				err := SaveResult(SynScan(ip, port))
				_ = err
			} else {
				err := SaveResult(Connect(ip, port))
				_ = err
			}
			wg.Done()
		}
	}
}

func SaveResult(ip string, port int, err error) error {
	// fmt.Printf("ip:%v, port: %v, goruntineNum: %v\n", ip, port, runtime.NumGoroutine())
	if err != nil {
		return err
	}
	if port > 0 {
		v, ok := vars.ConnResult.Load(ip)
		if ok {
			ports, ok1 := v.([]int)
			if ok1 {
				ports = append(ports, port)
				vars.ConnResult.Store(ip, ports)
			}
		} else {
			ports := make([]int, 0)
			ports = append(ports, port)
			vars.ConnResult.Store(ip, ports)
		}
	}

	return err
}

func GetInformation() {
	vars.ConnResult.Range(func(key, value interface{}) bool {
		ports := value.([]int)
		ip := key.(string)
		for _, port := range ports{
			_, _, tcpBanner, err := GetTcpBanner(ip, port)
			keyInform := fmt.Sprintf("%v:%v", ip, port)
			if tcpBanner != "" && err == nil {
				tmpInformation := vars.NodeInformation{}
				tmpInformation.Ip = ip
				tmpInformation.Port = port
				tmpInformation.Information, _ = Discern(tcpBanner)
				if tmpInformation.Information == "" {
					tmpInformation.Information = tcpBanner
				}
				vars.InformationResult[keyInform] = tmpInformation
			}else {
				_, _, httpBanner, err := GetHttpBanner(ip, port)
				if  err == nil{
					tmpInformation := vars.NodeInformation{}
					tmpInformation.Ip = ip
					tmpInformation.Port = port
					tmpInformation.Information, _ = Discern(httpBanner)
					if tmpInformation.Information == "" {
						tmpInformation.Information = httpBanner
					}
					vars.InformationResult[keyInform] = tmpInformation
				}else {
					tmpInformation := vars.NodeInformation{}
					tmpInformation.Ip = ip
					tmpInformation.Port = port
					tmpInformation.Information = "unknown"
					vars.InformationResult[keyInform] = tmpInformation
				}
			}
		}
		return true
	})
}

func PrintResult() {
	for _, value := range vars.InformationResult{
		fmt.Printf("ip:%v, port:%v, information:\r\n%v\n", value.Ip, value.Port, value.Information)
		fmt.Println(strings.Repeat("-", 20))
	}
}
