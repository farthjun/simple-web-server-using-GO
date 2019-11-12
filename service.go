package service

import (
    "github.com/martini-master"  //使用martini
)

//新建服务器
func NewServer(port string) {   
    m := martini.Classic()
    //添加参数[name]martini的参数中
    m.Get("/hello/:name", func(params martini.Params) string {
        return "Hello " + params["name"] + "!"
    })
    //对应mian函数中的端口
    m.RunOnAddr(":"+port)   
}
