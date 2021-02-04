## 内置单元测试框架
* Fail,Error：该测试失败，该测试继续，其它测试继续执行
* FailNow,Fatal：该测试失败，该测试终止，其它测试继续执行

## Benchmark
```$go
func BenchmarkConcatStringByAdd(b *testing.B){
    //与性能测试无关的代码
    b.ResetTimer()
    for i:=0; i< b.N; i++{
        //测试代码
    }
    b.StopTimer()
    //与性能测试无关的代码
}
```

## BDD in Go
1. [项目网站](https://github.com/smartystreets/goconvey)  
` https://github.com/smartystreets/goconvey `
2. 安装  
` go get -u github.com/smartystreets/goconvey/convey `
3. 启动WEB UI  
` $GOPATH/bin/goconvey `
