// package mmm 从约束上来讲，这样是不可以的，一个文件夹下面只能有一个包名
package abc

import (
	"fmt"
	//  _ "./hehe" // 这里不用包名，舍弃掉
	 _ "beegoStudy/abc/hehe" // 使用路径时，最好用这种路径，不要使用上面那种路径
)

// const BB string = "bb" // 常量声明是全局的，也就是说整个包中的名字是需要唯一的。如果有重名是会报错的：BB redeclared in this block， previous declaration at xxx/xx.go:4:19
const MM int = 66;
var meilea = "没了啊";

func EchoMeile() {
	println("meilei.go function EchoMeile打印一个变量：", meilea);
}

func init() {
	fmt.Println("这里是meilei的init");
	fmt.Println("这里是meilei的init，打印外部的全局变量：", fullName);
}