// 下面演示接口的一个应用：通过接口和结构体，实现继承、重载
// go中的接口（interface）：一个接口可以被多个结构体实现，一个结构体也可以实现多个接口。这点和php的接口是一样的
// 任意的类型都实现了空interface(我们这样定义：interface{})，也就是包含0个method的interface
package main

import (
	"fmt"
)

type Human struct {
    name string
    age int
    phone string
}
type Student struct {
    Human //匿名字段Human
    school string
    loan float32
}

type Employee struct {
    Human //匿名字段Human  ,  这样写相当于：Human Human。 同时这样写的一个好处就是，相当于Employee继承了Human的所有属性，在进行访问时可以直接用启属性也可以间接使用，如：
    // .phone 或 .Human.phone ； 但是赋值时却不可以这样，只能通过结构体赋值，如：var e2 = Employee{company:"深圳烁科科技有限公司", Human:h1} ， h1为Human结构体对象
    // 最重要的是，这样写的一个好处就是：连实现的方法也继承过来了。和PHP的继承是一样的
    company string
    money float32
}

//Human对象实现Sayhi方法
func (h *Human) SayHi() {
    fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Human对象实现Sing方法
func (h *Human) Sing(lyrics string) {
    fmt.Println("La la, la la la, la la la la la...", lyrics)
}

//Human对象实现Guzzle方法
func (h *Human) Guzzle(beerStein string) {
    fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

// Employee重载Human的Sayhi方法
func (e *Employee) SayHi() {
    fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,  // Employee结构体的name属性是通过继承Human过来的
        e.company, e.Human.phone) //此句可以分成多行
}

//Student实现BorrowMoney方法
func (s *Student) BorrowMoney(amount float32) {
    s.loan += amount // (again and again and...)
}

//Employee实现SpendSalary方法
func (e *Employee) SpendSalary(amount float32) {
    e.money -= amount // More vodka please!!! Get me through the day!
}

// 定义interface
type Men interface {
    SayHi()
    Sing(lyrics string)
    Guzzle(beerStein string)
}

type YoungChap interface {
    SayHi() // 定义了抽象方法，和php是一样的，没有实现
    Sing(song string)
    BorrowMoney(amount float32)
}

type ElderlyGent interface {
    SayHi()
    Sing(song string)
    SpendSalary(amount float32)
}

func main() {
	var i1 Men
	var h1 = Human{name:"何春华", phone:"13433543184"}
    // i1 = h1  // 不能这样写，会报错：cannot use h1 (type Human) as type Men in assignment: Human does not implement Men (Guzzle method has pointer receiver)
    i1 = &h1
	i1.SayHi()
    var i2 Men
    // var e2 = Employee{name:"欧阳海雄", company:"深圳烁科科技有限公司", phone:"15019261350"} // 不能这样写:unknown field 'name' in struct literal of type Employee
    var e2 = Employee{company:"深圳烁科科技有限公司", Human:h1}
    // i2 = new(e2) // 也不能这样写，报错：e2 is not a type，只有声明一个空的结构体时才可以这样写：i2 = new(Employee)
    i2 = &e2
	i2.SayHi()
    i2.Sing("SSSS")
    i2.Guzzle("GGGG")

    var y3 YoungChap
    var s3 Student = Student{loan:1.11}
    y3 = &s3
    fmt.Println("call before BorrowMoney Function, loan is : ", s3.loan)
    y3.BorrowMoney(8.88)
    fmt.Println("call after BorrowMoney Function, loan is : ", s3.loan)
    var e4 ElderlyGent
    var e5 = Employee{money:9.99}
    e4 = &e5
    fmt.Println("Before call the SpendSalary method, mone is : ", e5.money)
    e4.SpendSalary(7.7)
    fmt.Println("After call the SpendSalary method, mone is : ", e5.money)
    e5.Guzzle("e4 e4 e4") // 如果写成: e4.Guzzle 会报错
    e4.Sing("e4 e4 e4")
    e6 := &e5;
    e6.Guzzle("e6 e6 e6 e6") // 对结构体已实现的方法，不管是通过结构体对象还是通过结构体对象的引用（地址、指针）都可以调用此方法（因为go自动进行转换了）

}