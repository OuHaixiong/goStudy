// Go 递归（recursion），就是在运行的过程中调用自己。
package main

import "fmt"

/**
 * 阶乘（factorial）
 *
 */
/*func Factorial(x int) (result int) {
    if (0 == x) {
        result = 1;
    } else {
        result = x * Factorial(x - 1);
    }
    return;
}*/
func Factorial(x int) int { // 注意：两种写法都是正确的
    if 1 == x { // if条件判断可以加括号，也可以不加
        return 1; 
    } else {
        return x * Factorial(x - 1);
    }
}

func main() {
    var i int = 8;
    fmt.Printf("%d 的阶乘是 %d \n", i, Factorial(i));
    
    // var i int; // 变量不允许重复声明（declaration）
    var k int;
    for k = 0; k < 10; k++ {
        fmt.Printf("%d \t", fibonacci(k));
    }

}

// 裴波纳契数
func fibonacci(n int) int { // fibonacci:[计] 斐波纳契   :  0   1   1   2   3   5   8   13  21  34 
    if 2 > n {
        return n;
    }
    return fibonacci(n-2) + fibonacci(n-1);
}
