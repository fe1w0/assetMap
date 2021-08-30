package utils

import (
	"assetMap/scanner"
	"assetMap/vars"
	"fmt"
	"github.com/malfunkt/iprange"
	"github.com/urfave/cli/v2"
	"net"
	"os"
	"strconv"
	"strings"
)

func GetPorts(selection string) ([]int, error) {
	ports := []int{}
	if selection == "" {
		return ports, nil
	}

	ranges := strings.Split(selection, ",")
	for _, r := range ranges {
		r = strings.TrimSpace(r)
		if strings.Contains(r, "-") {
			parts := strings.Split(r, "-")
			if len(parts) != 2 {
				return nil, fmt.Errorf("Invalid port selection segment: '%s'", r)
			}

			p1, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, fmt.Errorf("Invalid port number: '%s'", parts[0])
			}

			p2, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, fmt.Errorf("Invalid port number: '%s'", parts[1])
			}

			if p1 > p2 {
				return nil, fmt.Errorf("Invalid port range: %d-%d", p1, p2)
			}

			for i := p1; i <= p2; i++ {
				ports = append(ports, i)
			}

		} else {
			if port, err := strconv.Atoi(r); err != nil {
				return nil, fmt.Errorf("Invalid port number: '%s'", r)
			} else {
				ports = append(ports, port)
			}
		}
	}
	return ports, nil
}


func GetIpList(ips string) ([]net.IP, error) {
	addressList, err := iprange.ParseList(ips)
	if err != nil {
		fmt.Println(err)
		fmt.Println(err)
		return nil, err
	}

	list := addressList.Expand()
	return list, err
}

func IsRoot() bool {
	return os.Geteuid() == 0
}

func CheckRoot() {
	if !IsRoot() {
		fmt.Println("must run with root")
		os.Exit(0)
	}
}

func Scan(ctx *cli.Context) error {
	if ctx.IsSet("ip") {
		vars.Host = ctx.String("ip")
	}

	if ctx.IsSet("port") {
		vars.Port = ctx.String("port")
	}

	if ctx.IsSet("mode") {
		vars.Mode = ctx.String("mode")
	}

	if ctx.IsSet("timeout") {
		vars.Timeout = ctx.Int("timeout")
	}

	if ctx.IsSet("concurrency") {
		vars.ThreadNum = ctx.Int("concurrency")
	}

	if strings.ToLower(vars.Mode) == "syn" {
		CheckRoot()
	}

	// 初始化 指纹库
	if err := vars.InitNamp(); err != nil {
		return err
	}

	ips, err := GetIpList(vars.Host)
	ports, err := GetPorts(vars.Port)
	tasks, n := scanner.GenerateTask(ips, ports)
	_ = n
	scanner.RunTask(tasks)
	scanner.GetInformation()
	scanner.PrintResult()

	return err
}
