package routers

import (
	"github.com/xxjwxc/ginrpc"
)

func init() {
	ginrpc.SetVersion(1575798798)
	ginrpc.AddGenOne("Hello.Hello", "/block", []string{"post", "get"})
	ginrpc.AddGenOne("Hello.Hello2", "hello.hello2", []string{"post"})
}