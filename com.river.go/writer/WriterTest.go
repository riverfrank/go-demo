package main

import (
	"bytes"
	"fmt"
)

func main() {
	//定义零值Buffer类型变量b
	var b bytes.Buffer
	//使用Write方法为写入字符串
	b.Write([]byte("你好"))

	//这个是把一个字符串拼接到Buffer里
	fmt.Fprint(&b, ",", "http://www.flysnow.org")
	//把Buffer里的内容打印到终端控制台
	//	b.WriteTo(os.Stdout)

	p := [100]byte{}
	fmt.Println(p[0])

	read, err := b.Read(p[:])
	fmt.Println(string(p[:]), read, err)

	//var p = []byte{'a', 'b', 'c'}
	//var testb = bytes.NewBuffer([]byte{'a', 'b', 'c'})
	//testb.ReadFrom(&b)
	//n, err := testb.Read(p[:])
	//fmt.Println("--------->")
	//fmt.Println(n, err, string(p[:n]))

}
