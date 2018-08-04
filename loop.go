// Go 循环语句
package main;

import "fmt";

func main() {
    var a int;
    var b int = 15;
    
    numbers := [6]int{1, 2, 3, 5}; // 声明数组
    
    for a := 0; a < 10; a++ { // for 循环 。 我靠在循环里面a := 0不会报错，且一定要这样写，且一定要先声明：var a int;
        fmt.Printf("a 的值为：%d\n", a);
    }
    
    for a < b { // for循环的另一种形式。 for true  {} 可以无限循环
        a++;
        fmt.Printf("a 的值为：%d\n", a);
    }
    
    for k, v := range numbers {
        fmt.Printf("数组numbers第 %d 位的值= %d \n", k, v);
        // 想不到吧，还会打印如下：    (如果声明了数组的长度，没有赋值的元素会给默认值)
        // 数组numbers第 4 位的值= 0 
        // 数组numbers第 5 位的值= 0 
    }

    primeNumber();
    
    gotoLoop();
}

func primeNumber() { // 质数（prime number）又称素数，有无限个。
    /* 定义局部变量 */
    var i, j int

    // for (i=2; i < 100; i++) { // 这种格式是错误的：syntax error: unexpected =, expecting )
    for i=2; i < 100; i++ {
        for j=2; j <= (i/j); j++ { // 嵌套循环
            if(i%j==0) {
                break; // 如果发现因子，则不是素数
            }
        }
        if(j > (i/j)) {
            fmt.Printf("%d  是素数\n", i);
        }
    } 
}

func gotoLoop() {
    var a int = 10;
    LOOP: for a < 20 { // lable标志，
        if (15 == a) { // if条件可以没有括号：if 15 == a {}
            a++;
            goto LOOP; // 使用goto进行跳转。但是，在结构化程序设计中一般不主张使用goto语句， 以免造成程序流程的混乱，使理解和调试程序都产生困难
        }
        fmt.Printf("a 的值为：%d\n", a);
        a++;
    }
}
