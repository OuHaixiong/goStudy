// Go 结构体。Go 语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型。结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。
package main

import "fmt"

type Books struct { // 声明一个结构体
    title string
    author string // 结构体的属性（成员）
    subject string
    book_id int
}

func main() {
    var book1 Books; // 声明一个结构体变量
    book1.title = "PHP进阶"; // 给结构体的属性赋值
    book1.author = "欧阳海雄";
    book1.subject = "PHP语言教程";
    book1.book_id = 62018;
    
    book2 := Books{"Go入门", "maimengmei.com", "go语言教程", 60189}; // 声明一个结构体变量并初始化成员属性
    
    fmt.Printf("book1 title=>%s\n", book1.title); // 读取结构体的属性
    fmt.Printf("book1 author=>%s\n", book1.author);
    fmt.Printf("book1 subject=>%s\n", book1.subject);
    fmt.Printf("book1 book_id=>%d\n", book1.book_id);
    
    fmt.Printf("book2 title : %s\n", book2.title);
    fmt.Printf("book2 author : %s\n", book2.author);
    fmt.Printf("book2 subject : %s\n", book2.subject);
    fmt.Printf("book2 book_id : %d\n", book2.book_id);
    
    printBook(book2);
    
    printBook2(&book1); // 把结构体变量的地址传过去即可
}

func printBook(book Books) { // 结构体作为参数
    fmt.Printf("book title : %s\n", book.title);
    fmt.Printf("book author : %s\n", book.author);
    fmt.Printf("book subject : %s\n", book.subject);
    fmt.Printf("book book_id : %d\n", book.book_id);
}

func printBook2(book *Books) { // 结构体指针
    fmt.Printf("book title : %s\n", book.title);
    fmt.Printf("book author : %s\n", book.author);
    fmt.Printf("book subject : %s\n", book.subject);
    fmt.Printf("book book_id : %d\n", book.book_id);
}