// 与字符处理相关的
package main
import (
	"fmt"
	"strconv"
	"crypto/md5"
	"crypto/sha1" // Go 在多个 crypto/* 包中实现了一系列散列函数。
	"encoding/hex"
    b64 "encoding/base64" // 这个语法引入了 encoding/base64 包, 并使用名称 b64 代替默认的 base64。
)

func main() {
    fmt.Println(reverseString("!selpmaxe oG , olleH"));
    str := "ab测试中文字符反(翻)转！嗯en8"
	fmt.Println(reverseString(str));
	
	s := "321"
	d, err := strconv.Atoi(s) // 字符串转整型
    if (err != nil) {
        fmt.Println(err.Error());
	}
	fmt.Println(d); // 转整成功，返回 321
	s = "3.12"
	d, err = strconv.Atoi(s) // 字符串转整型
    if (err != nil) {
		fmt.Println(err.Error()); // 这里转整失败，打印：strconv.Atoi: parsing "3.12": invalid syntax
	}
	fmt.Println(d); // 失败后，这里返回 0

	// 下面演示数字转字符串
	var i int = 123
	ss := strconv.Itoa(i) // 数字转字符串（int转string）
	ss = ss + " is string" // 连接两个字符串
	fmt.Println(ss); // 返回：123 is string

	var i64 int64 = 9887
    strings := strconv.FormatInt(i64, 10) // int64 转 string
	fmt.Printf("type: %T, value: %s\n", strings, strings) // type: string, value: 9887
	
	// 下面演示任意类型的数据，需要判断处理
	var e interface{}
	e = 101
	e = "wo s 欧海雄"
	switch v := e.(type) {
	case int : fmt.Println("整形：", v); break;
	case string : fmt.Println("字符串：", v); break;
	}

	var str1 = "hello"
    str2 := "world"
    str3 := fmt.Sprintf("%s %s", str1, str2) // Sprintf 相当于连接两个字符串，和用“+”是一样的
    n := len(str3)
    fmt.Println(str3) // hello world
	fmt.Printf("len(str3)=%d\n", n) // len(str3)=11
	for index, value := range str3 { // 字符串的下标是从0开始的
        fmt.Printf("str[%d]=%c\n", index, value) // %c:为单个字符；%s:为字符串；%d:为整形数字
    }
    substr := str3[0:5] // hello str[X:X] 相当于截取字符串
    fmt.Println(substr)
    substr = str3[5:] // ( world)
	fmt.Println(substr)

	str3 = "我是欧海雄"

	// str3[0] = "0" // 不能像php一样直接修改字符，这里会报错：./string.go:54:10: cannot assign to str3[0]
	var byteSlice []byte // 声明一个二进制的切片
	byteSlice = []byte(str3) // 将字符串str3转为二进制切片
    byteSlice[0] = 'a' // 转为二进制切片后就可以修改其中的值了 （我靠什么鬼呀，这里只能用单引号，而且值只能用数字或字母，不能为中文）
    str = string(byteSlice) // string() 二进制切片转为字符串
	fmt.Println("after modify:", str) // after modify: a是欧海雄
		
	n = len(str3) // len 计算字符串的长度和php是一样的，在utf-8中英文字母一个算一个长度，中文字一个算三个长度
	fmt.Printf("字符串‘%s’的长度为%d\n", str3, n) // 字符串‘我是欧海雄’的长度为15
    // result := reverse(str3)
    // fmt.Println(result)
    // result = reverse1(result)
    // fmt.Println(result)

	var b rune = '中' // 这里不能用双引号，不然报错：./string.go:71:15: cannot use "中" (type string) as type rune in assignment
    fmt.Printf("b=%c\n", b) // b=中

	var runeSlice []rune // 声明一个rune类型的切片（rune类似int32）
	runeSlice = []rune(str) // 将字符串转为rune的切片
	fmt.Printf("str为：%s， rune长度为：%d， 普通长度leng(X)=%d \n", str, len(runeSlice), len(str))
	// str为：a是欧海雄， rune长度为：7， 普通长度leng(X)=15           (这里的字符是经过了替换的，所以是这样)
	var strstr = "w是你AA"
	var runeSlice2 []rune
	runeSlice2 = []rune(strstr) // rune切片有点类似php中的mb_strlen；一个中文和一个英文都算一个字
    fmt.Printf("strstr为：%s， rune长度为：%d， 普通长度leng(X)=%d \n", strstr, len(runeSlice2), len(strstr))
	// strstr为：w是你AA， rune长度为：5， 普通长度leng(X)=9            （注意和上面的区别）

	fmt.Printf("strstr原：%s， 反转后为：%s \n", strstr, reverseStringV2(strstr));
	var str11 = "上海自来水来自海上";
	var str12 = reverseStringV2(str11)
	if str12 == str11 { // 判断连个字符串是否相等
		fmt.Println("是回文");
	} else {
        fmt.Println("不是回文");
	}

	var sstr string = "我是欧海雄！Ou year!"
	fmt.Println(GetMd5Hash(sstr));
	fmt.Println(GetSha1Hash(sstr));

	var text string = "abc123!?$*&()'-=@~"
	var standardEncode string = EncodeBase64(text, true)
	var urlEncode string = EncodeBase64(text, false)
	fmt.Println("standard base64 is:", standardEncode) // standard base64 is: YWJjMTIzIT8kKiYoKSctPUB+
	fmt.Println("URL compatible base64 is:", urlEncode) // URL compatible base64 is: YWJjMTIzIT8kKiYoKSctPUB-
	decodeString, err := DecodeBase64(standardEncode, true)
	if (err != nil) {
		fmt.Println("Base64 Decode fail :", err.Error())
	} else {
		fmt.Println("standard base64 decode is:", decodeString)
	}
	decodeString, err = DecodeBase64(urlEncode, false)
    if err != nil {
		fmt.Println("Base64 Decode fail :", err.Error())
	} else {
        fmt.Println("URL Compatible Base64 decode is:", decodeString)
	}
}

/**
 * 反转字符串
 */
func reverseString(s string) string {
	r := []rune(s) // 和C++语言不同，golang这边需要先将字符串string转换成rune类型，而后才能进行对调操作。
	for from, to := 0, len(r)-1; from < to; from, to = from+1, to-1 {
		r[from], r[to] = r[to], r[from]
	}
	return string(r)
}

/**
 * 反转字符串，这里和第一个的思想是一样的，写法不一样而已
 */
func reverseStringV2(s string) string {
	var r []rune = []rune(s)
	length := len(r)
	half := length/2;
	for i:=0; i < half; i++ {
		tmp := r[length-i-1] // 最后一个赋值给临时变量
		r[length-i-1] = r[i] // 第一个赋值给最后一个
		r[i] = tmp // 临时变量赋值给第一个
	}
	return string(r)
}

/**
 * 获取字符串的md5值
 * @param string text 输入字符串
 * @return string 32位的字符串
 */
func GetMd5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil)) 
}

/**
 * 获取字符串的sha1值
 * @param string text 输入的字符串
 * @return string 40位的字符串
 */
func GetSha1Hash(text string) string {
	h := sha1.New() // 产生一个散列值的方式是 sha1.New()，sha1.Write(bytes)，然后 sha1.Sum([]byte{})。这里我们从一个新的散列开始。
	h.Write([]byte(text)) // 写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。
	bs := h.Sum(nil) // 这个用来得到最终的散列值的字符切片。Sum 的参数可以用来都现有的字符切片追加额外的字节切片：一般不需要要。
    return fmt.Sprintf("%x", bs) // SHA1 值经常以 16 进制输出，例如在 git commit 中。使用%x 来将散列结果格式化为 16 进制字符串。
}

/**
 * 对字符串进行base64编码
 * @param string text 输入的字符串
 * @return string 24位的字符串
 * 标准 base64 编码和 URL 兼容 base64 编码的编码字符串存在稍许不同（后缀为 + 和 -; 其实是字符“~”编码后就成“+”），但是两者都可以正确解码为原始字符串
 */
func EncodeBase64(text string, isStandard bool) (string) { // 特别注意了：go语言是不支持默认参数和继承（重载）的
	data := []byte(text) // 编码需要使用 []byte 类型的参数，所以要将字符串转成字节类型。
	var stringEncode string
    if isStandard { // 使用标准的base64格式。（Go 同时支持标准的和 URL 兼容的 base64 格式。）
        stringEncode = b64.StdEncoding.EncodeToString(data) // 标准的和php的base64_encode是一样的，返回类似：YWJjMTIzIT8kKiYoKSctPUB+
	} else { // 使用 URL 兼容的 base64 格式进行编码。
        stringEncode = b64.URLEncoding.EncodeToString(data)
	}
	return stringEncode
}

/**
 * 对字符串进行base64解码
 * @param sting text 输入的字符串
 * @param bool isStandard 是否为标准解码
 * @return string 返回解码后的字符串
 */
func DecodeBase64(text string, isStandard bool) (stringDecode string, err error) {
	var b []byte
    if (isStandard) { // 标准的解码和php的base64_decode一样
		b, err = b64.StdEncoding.DecodeString(text) // 解码可能会返回错误，如果不确定输入信息格式是否正确，那么，你就需要进行错误检查了。
		if err != nil {
			return
		}
	} else {
		b, err = b64.URLEncoding.DecodeString(text)
		if err != nil {
			return
		}
	}
	stringDecode = string(b)
    return // 函数结束了一定要写return
}

