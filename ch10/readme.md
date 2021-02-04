# 函数：一等公民

### 与其它主要编程语言的差异
1. 函数可以有多个返回值
2. 所有参数都是值传递：slice，map，channel 会有传引用的错觉
3. 函数可以作为变量的值
4. 函数可以作为参数和返回值

### 可变长参数

### defer 函数（延迟执行函数）

```$go
func TestDefer(t *testing.t){
    defer func(){
        t.Log("Clear resource")
    }()
    t.Log("Started")
    panic("Fatal error") //defer 仍会执行
}
```