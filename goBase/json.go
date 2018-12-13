// 练习json和go结构体的互转
package main

import (
	"fmt"
	"encoding/json"
)

type Info struct {
	Name string `json:"name"`  // 如果没有后面的tag，那么转为json字符串时，字段名为大写
	Age int `json:"age,string"`  // 如果类型也需要转换的话，就需要在tag中进行声明。当然如果json字符串中的类型和tag中的不一样，转结构体时会报错
	Sex string  // 如果字段名对不上的话，也需要声明tag；如： OrderId int `json:"order_id"`
	test string  // 此处的字段名为小写，这个字段在json中无法被获取，因为未导出
}

type InfoList struct {
	Infos []Info
}

type User struct {
	Name string `json:"-"` // 字段的tag是“-”（英文中划线），那么这个字段不会输出到json
	Age int `json:"age,string"` // 类型转换
	Sex string `json:"sex,omitempty"` // tag中如果带有”omitempty”，那么如果该字段值为空，就不会输出到json串中。
}

type UserList struct {
	Infos []User
}

func main() {
	// 多行字符串可以用``括起来
	var binary = []byte(`
	    [
			{"name":"欧阳海雄", "age":"35", "sex":"male"},
			{"name":"yan妍妍", "age":"23", "sex":"female"}
		]
	`) // []byte:字符串转二进制
	var infos []Info;
	err := json.Unmarshal(binary, &infos); // Unmarshal(data []byte, v interface{}) error ： 将json字符串流（二进制）解析为go语言中对应的类型
	// json解析的时候只会解析能找得到的字段，找不到的字段会被忽略
	// json字符串转struct时，字段的查找顺序是：
	// 1, 首先查找tag含有name的可导出的struct字段
	// 2, 其次查找字段名是Name的导出字段
	// 3, 最后查找类似NaMe或NAME这样的，除了首字母之外其他大小写不敏感的导出字段
	if err != nil {
        fmt.Println(err.Error())
	} else {
		fmt.Printf("%v \n", infos); // [{欧阳海雄 35 male } {yan妍妍 23 female }]
		println(infos[1].Name)
	}

	var infoList InfoList
	infoList.Infos = append(infoList.Infos, Info{Name:"萌萌", Age:4, Sex:"female"});
	infoList.Infos = append(infoList.Infos, Info{Name:"翔翔", Age:0, Sex:"male"});
	b, err := json.Marshal(infoList); // 结构体转json字符串
	if (err != nil) {
        fmt.Println("json encoded error : ", err)
	} else {
		fmt.Println(string(b)) // string ：二进制转字符串
		// {"Infos":[{"Name":"萌萌","Age":4,"Sex":"female"},{"Name":"翔翔","Age":0,"Sex":"male"}]}
		// 修改tag后：{"Infos":[{"name":"萌萌","age":"4","Sex":"female"},{"name":"翔翔","age":"0","Sex":"male"}]}
	}


	var userList UserList
	// userList.Infos[0] = User{Name:"huahua", Age:35, Sex:"female"} // 没有声明数组长度就不能这样写，会报错的。
	userList.Infos = append(userList.Infos, User{Name:"huahua", Age:35, Sex:"female"})
	userList.Infos = append(userList.Infos, User{Name:"铭铭", Age:1, Sex:""})
	buff, err := json.Marshal(userList)
	if err != nil {
        fmt.Println("json err : ", err)
	} else {
        fmt.Println(string(buff)) // {"Infos":[{"age":"35","sex":"female"},{"age":"1"}]}
	}
    // Marshal函数只有在转化成功的时候才会返回数据，在转化的过程中需要注意：
    // 嵌套的数据是不能编码的，不然会让json编码进入死循环；
    // bool类型，转化为json的boolean；
    // 整数类型，浮点数等数值类型，转化为json的number；
    // string类型，转化为json的字符串（带”“引号）；
    // struct类型，转化为json的object，再根据各个成员的类型递归；
    // 数组或切片，转化为json的array；
    // []byte类型，会先进行base64编码然后转化为json字符串；
    // map类型，转化为json的object，必须是map[string]T这种类型（T是go语言中任意的类型）；
    // interface{}按照内部的实际类型进行转化；
    // nil转化为json的null；
    // channel，complex，func等类型是不能被编码成json的，会返回UnsupportedTypeError

    test()
}


// 测试json转go的结构体
func test() {
	var binary = []byte(`{"section": [{"end": "21.129.111", "start": "0.00.1"}], "specify": ["78.23.1", "21.129.111", "81.2659.50", "789.1.0"]}`) // []byte:字符串转二进制
	var m Mandatory;
	err := json.Unmarshal(binary, &m); // Unmarshal(data []byte, v interface{}) error ： 将json字符串流（二进制）解析为go语言中对应的类型
	if err != nil {
        fmt.Println(err.Error())
	} else {
		println(m.Section[0].Start)
	}
}

type Mandatory struct {
	Section []SectionStruct
	Specify []string
}

type SectionStruct struct {
	Start string
	End string
}