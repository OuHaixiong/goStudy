package main

import (
    abcP "./abc" // 两种写法都可以。 把包赋值给abcP，如果省略掉abcP，默认就是返回的包名【一般来讲我们都是写成和文件夹同名】（提倡省略的写法，除非本页中包存在同名的情况下）
	//"beegoStudy/abc"
	_ "beegoStudy/routers" // _ ：代表这个包只执行里面的常量、变量和init函数，在本页面中并不使用（舍弃掉）
	"github.com/astaxie/beego"
	"beegoStudy/abc/hehe" // 特别注意了，这里返回的包名并不是hehe，而是hehehe。
	// 特别注意了：如果一个包被引入了多次，那么它就会初始化多少次（初始化包括：执行里面的常量、变量和init函数）
)

const MM string = "main的常量";

func main() {
	var str = "欧欧欧";
	abcP.Wl(str);
	// abc.wor(); // 这样调用是会出错的：cannot refer to unexported name abc.wor
	abcP.EchoMeile();
	abcP.EchoNiyehao();
	println(abcP.MM);
	println(MM);
	println(hehehe.MM);
	
	beego.Run()
}

