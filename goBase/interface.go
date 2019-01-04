// Go 接口（Interface）。Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。
// interface是一组method的组合，我们通过interface来定义对象的一组行为。
// 总结：结构体类似php中的类，可以给类挂载一个方法，如果方法和之前声明的接口的方法同名，说明这个结构体实现了该接口
// 使用结构体和接口就实现了类似php中的类和对象的关系。如果用结构体来调用方法就类似php的静态方法，如果用接口来调用方法的话就类似php的动态方法（需要new才能调用）
// 接口对象new出来以后，貌似只能调用方法，不能做其他更多的事情，这也是很蛋疼的。
// 其实简单理解下，和php的是一样的，接口只是用来规范代码的，告诉使用者你要么不要实现这个接口，如果你实现了这个接口就一定要实现接口中的所有方法
package main

import ("fmt")
import "strconv"

type Phone interface {  // 定义接口
    call(); // 这点和PHP类似，接口中的方法
    say(); // 其实这点和php是一样的，实现了该接口的结构体，就需要全部实现接口中的方法，如果不全部实现就会报错
}

type NokiaPhone struct { // 定义一个诺基亚的结构体（结构体可以有属性，也可以没有，其实可以简单的理解为php中的类）
}

type IphonePhone struct { // 定义一个苹果手机结构体
    // var message string = "我是世界第一好用的手机iPhone"; // 这样声明结构体的属性是错误的，不能使用var
    // message string = "我是世界第一好用的手机iPhone"; // 这样也是错误的，结构体中不能赋初始值；报错：syntax error: unexpected =, expecting semicolon or newline or }
    message string;
}

func (nokia NokiaPhone) call() { // 定义一个方法，该方法是结构体NokiaPhone的方法，且方法名为call，方法名和接口中的方法名一样，意味着这个结构体实现了phone的方法
    fmt.Println("I am Nokia, I can call you!");
}

func (nokia NokiaPhone) say() {

}

func (iphone IphonePhone) call() { // 给结构体定义方法，可以简单的理解为，把一个方法挂载到一个结构体中
    fmt.Println("I am iPhone, I can call you!!");
}

func (iphone IphonePhone) say() {
    iphone.message = "你皮都不是"; // 这里的参数是按非引用传递的，所以这里的赋值并不会影响外部结构体的属性
    fmt.Println(iphone.message);
}

func main() {
    var phone Phone; // 定义一个变量为接口类型
    phone = new(NokiaPhone); // new一个结构体对象出来，赋给一个接口类型（和php new一个对象差不多）
    phone.call(); // 结构体（接口）【对象】调用方法
    
    //phone = new(IphonePhone);
    //phone.call();
    
    var iPhone Phone; // 其实声明不同的变量也是可以的
    iPhone = new(IphonePhone);
    iPhone.call();
    // iPhone.message = "XXX"; // 这样是不行的，报错：iPhone.message undefined (type Phone has no field or method message)
    // IphonePhone.message = "XXX"; // 这样也是不行的，报错：IphonePhone.message undefined (type IphonePhone has no method message)
    var iphoneStruct IphonePhone; // 声明一个结构体变量，相当于new了一个对象
    iphoneStruct.message = "我是世界第一好用的手机iPhone";
    // phone = new(iphoneStruct); 这样写也是错误的，报错：iphoneStruct is not a type。new方法括号中一定是定义结构体的名称
    //phone = new(IphonePhone);
    //phone.say();
    fmt.Println(iphoneStruct.message)
    iphoneStruct.say(); 
    fmt.Println(iphoneStruct.message)

    print_r()
}

// 下面演示空接口（interface{}）的使用
func print_r() {
    ouhaixiong := Human{"欧阳海雄", 35, "150-1926-1350"}
    fmt.Println("This Human is : ", ouhaixiong)
}

type Human struct {
    name string
    age int
    phone string
}

// 通过这个方法 Human 实现了 fmt.Stringer (任何实现了String方法的类型都能作为参数被fmt.Println调用)
func (h Human) String() string { // String() 是 fmt包里面的一个接口： type Stringer interface { String() string }
    return "< " + h.name + " - " + strconv.Itoa(h.age) + " years - tel:" + h.phone + " >"
}
