package main;

import "fmt"
import "strings"
import "strconv"

func Version4ToInt(version string) int {
  arr := strings.Split(version, ".")
  fix_data := 100000000
  ver := 0
  for _, num := range arr {
    //d, error := strconv.Atoi(num) // error declared and not used（error不能用来做标识符）
	d, _error := strconv.Atoi(num) // 字符串转整型
    if _error != nil { // 单个_表示不接受参数，且_不能用来作为变量（cannot use _ as value）
	    fmt.Println("字符串转换成整数失败")
	}
	ver += d*fix_data
    fix_data = fix_data / 100
  }
  return ver
}

func main() { // 如果是需要go来运行的一定不能少了main函数，除非你的文件是被包含进来的
    var version string = "1.1.0";
    var ver int = Version4ToInt(version);
    //fmt.Println(version, ver); // non-declaration statement outside function body
	println(version, ver); // 两点需要注意的：
	// 1，打印一定需要在函数里面，不能在外面
	// 2，打印是需要引入fmt包的，如果你不用fmt.Println而直接用println，其实是go帮你自动引入了包fmt
}
