package scanner

import (
	"assetMap/vars"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"time"
)

func Connect(ip string, port int) (string, int, error) {
	// 	普通 tcp 链接
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", ip, port), time.Duration(vars.Timeout)*time.Second)

	defer func() {
		if conn != nil {
			_ = conn.Close()
		}
	}()

	return ip, port, err
}

func GetTcpBanner(ip string, port int)(string, int, string, error){
	// 默认使用 tcp,返回ip,port,banner,error
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", ip, port), time.Duration(vars.Timeout)*time.Second)

	if err != nil {
		_ = conn
		return ip, port, "", err
	}

	tcpConn := conn.(*net.TCPConn)
	//
	// 设置读取的超时时间
	tcpConn.SetReadDeadline(time.Now().Add( time.Duration(vars.Timeout) * time.Second))
	reader := bufio.NewReader(conn)
	tcpBanner, err := reader.ReadString('\n')

	defer func() {
		if conn != nil {
			_ = conn.Close()
		}
	}()
	return ip, port, tcpBanner, err
}


func splitHttpHead(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// 用于分割出http头
	headEnd := bytes.Index(data, []byte("\r\n\r\n"))
	if headEnd == -1 {
		return 0, nil, nil
	}
	return headEnd + 4, data[:headEnd+4], nil
}

func GetHttpBanner(ip string, port int)(string, int, string, error){
	// 利用 GetHttpBanner 获取 Http 头信息
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", ip, port), time.Duration(vars.Timeout)*time.Second)
	if err != nil {
		_ = conn
		return ip, port, "", err
	}

	defer func() {
		if conn != nil {
			_ = conn.Close()
		}
	}()

	tcpConn := conn.(*net.TCPConn)
	if _, err := conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n")); err != nil {
		return ip, port, "", err
	}
	tcpConn.SetReadDeadline(time.Now().Add( time.Duration(vars.Timeout) * time.Second))
	httpBannerScanner := bufio.NewScanner(conn)
	httpBannerScanner.Split(splitHttpHead)
	if httpBannerScanner.Scan() {
		return ip, port, httpBannerScanner.Text(), nil
	}
	err = httpBannerScanner.Err()
	if err == nil {
		err = io.EOF
	}
	return ip, port, "", err
}