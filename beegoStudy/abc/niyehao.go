package abc


var meileaN = "没了啊,这里是你也好"; // 在同一个包里的全局变量也是不能重复的

func EchoNiyehao() {
	println("niyehao.go function EchoNiyehao 打印外部的全局变量：", meilea);
	println("niyehao.go function EchoNiyehao 打印内部的全局变量：", meileaN);
	println("niyehao.go function EchoNiyehao 打印常量：", MM);
}

func kao() { // 在同一个包里的函数是不允许重复的，不管是私有还是公有的

}

func init() {
	println("这里是niyehao的init");
}