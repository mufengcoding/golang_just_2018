package main

import (
	"fmt"
	"net"
	"net/http"
	"runtime"
	"strconv"
)

func Hello(response http.ResponseWriter, request *http.Request) {
	//操作系统
	fmt.Println("GOOS:", runtime.GOOS)
	//架构
	fmt.Println("GOARCH:", runtime.GOARCH)
	//GOROOT
	fmt.Println("GOROOT:", runtime.GOROOT())
	//go版本
	fmt.Println("Version:", runtime.Version())
	//cpu数
	fmt.Println("NumCPU:", runtime.NumCPU())
	fmt.Println()

	//MAC和IP地址
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Poor soul, here is what you got: " + err.Error())
	}
	var numbers string
	for _, inter := range interfaces {
		fmt.Println(inter.Name, inter.HardwareAddr)

		addrs, _ := inter.Addrs()
		for _, addr := range addrs {
			fmt.Println("  ", addr.String())
			numbers += addr.String() + "\n"
		}

	}
	var s string
	s += "操作系统:" + runtime.GOOS + "\n"
	s += "架构:" + runtime.GOARCH + "\n"
	s += "GOROOT:" + runtime.GOROOT() + "\n"
	s += "go版本:" + runtime.Version() + "\n"
	s += "cpu:" + strconv.Itoa(runtime.NumCPU()) + "\n"

	s += numbers
	fmt.Fprintf(response, s)
}

func main() {
	http.HandleFunc("/", Hello)

	http.ListenAndServe(":8087", nil)
}
