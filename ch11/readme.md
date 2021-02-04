
# 封装数据和行为
## 结构体定义
```$go
type Employee struct {
    Id      string
    Name    string
    Age     int
}
```

## 实例创建及初始化
```$go
e := Employee{"0","Bob",20}
e1 := Employee{Name:"Mike",Age:30}
e2 := new(Employee) // 注意这里返回的引用/指针，相当于 e:= &Employee{}
e2.Id = "2" // 与其它主要编程语言的差异
e2.Age = 22
e2.Name = "Rose"
```

## 行为（方法）定义
```$go
type Employee struct{
    Id      string
    Name    string
    Age     int
}
//第一种定义方式在实例对应方法被调用时，实例成员会进行值复制
func (e Employee) String() string{
    return fmt.Sprintf("ID:%s-Name:%s-Age:%d",e.Id,e.Name,e.Age)
}

//通常情况下为了避免内存拷贝我们使用第二种方式定义
func (e *Employee) String() string{
    return fmt.Sprintf("ID:%s/Name:%s/Age:%d",e.Id,e.Name,e.Age)
}
```

## 接口与依赖
```$go
//TODO
```

## Duck Type 式 接口实现
#### 接口定义
```$go
type Programmer interface{
    WriteHelloWorld() Code
}
```
#### 接口实现
```$go
type GoProgrammer struct{}

func (p *GoProgrammer) WriteHelloWorld() Code{
    return "fmt.Println(\"Hello World!\")"
}
```

## Go 接口
*与其它主要编程语言的差异*
1. 接口为非入侵性，实现不依赖于接口定义
2. 所有接口的定义可以包含在接口使用者包内

## 接口变量

## 自定义类型

## 空接口与duanyan
1. 空接口可以表示任何类型
2. 通过断言来将空接口转换为定制类型  
` v,ok := p.(int) // ok =true 时转换成功`

## Go 接口最佳实践

1. 倾向于使用最小的接口定义，很多接口只包含一个方法
```$go
type Reader interface{
    Read(p []byte)(n int,err error)
}

type Writer interface{
    Write(p []byte)(n int,err error)
}
```
2. 较大的接口定义，可以由多个小接口定义组合而成
```$go
type ReadWriter interface{
    Reader
    Writer
}
```
3. 只依赖于必要功能的最小接口
```$go
func StoreData(reader Reader) error{
    ...
}
```
