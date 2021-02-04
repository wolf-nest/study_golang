# Go 的错误机制
*与其它主要编程语言的差异*
1. 没有异常机制
2. error类型实现了 error 接口
```$go 
type error interface{
    Error() string
}
```
3. 可以通过errors.New 来快速创建错误实例
` errors.New(n must be in the range [0,10]) `

## 最佳实践

1. 定义不同的错误变量，以便于判定错误类型
```$go
var LessThanTwoError = errors.New("n must be greater than 2")
var GreaterThanHundredError = errors.New("n must be  less than 100")
....
func TestGetFibonacci(t *testing.T) {
	if list, err := GetFibonacci(-5); err != nil {
		if err == LessThanTwoError {
			t.Error("Need a larger number.")
		}

		if err == GreaterThanHundredError {
			t.Error("Need a less number.")
		}

	} else {
		t.Log(list)
	}
}

```
2. 及早失败 避免嵌套

## Panic
* panic 用于不可以恢复的错误
* panic 退出前会执行 defer 指定的内容

## Panic vs os.Exit
* os.Exit 退出时不会调用 defer 指定的函数
* os.Exit 退出时不输出当前调用栈信息

## 最常见的“错误恢复”
```
defer func(){
    if err := recover(); err !=nil {
        log.Error("recovered panic", err)
    }
}()
```

## 当心！ recover 成为恶魔😈
* 形成僵尸服务进程，导致health check 失效。
* "Let it Ceash!" 往往是我们恢复不正确性错误的最好方法
