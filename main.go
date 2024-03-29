package main

import (
    "os"
    "service"
    flag "github.com/spf13/pflag"
)

const (
    PORT string = "8080" //默认的端口为8080
)

func main() {

    port := os.Getenv("PORT") //如果没有监听到端口，那么端口为8080
    if len(port) == 0 {
        port = PORT
    }
    //设置端口
    pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
    flag.Parse()
    if len(*pPort) != 0 {
        port = *pPort
    }

    service.NewServer(port)//启动服务器
}
