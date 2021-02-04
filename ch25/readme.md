# 反射编程

## reflect.TypeOf vs reflect.ValueOf
* reflect.TypeOf 返回类型
* reflect.ValueOf 返回值
* 可以从reflect.Value 获得类型
* 通过kind的来判断类型

## 利用反射编写灵活的代码
1. 按名字访问结构的成员  
`reflect.ValueOf(*e).FiledByName("name")`  
2. 按名字访问结构的方法  
`reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})`  

## 访问 StructTag
```$go
……
if nameField,ok := reflect.TypeOf(*e).FieldByName("Name");!ok {
    t.Error("Failed to get 'Name' field.")
}else{
    t.Log("Tag:format",nameField.Tag.Get("format"))
}
```
*reflect.Type 和 reflect.Value 都有 FieldByName 方法 使用时注意他们之间的区别*

## DeepEqual （比较切片和map）
```$go
func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 4: "three"}

	t.Log(reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}

	t.Log("s1 == s2?", reflect.DeepEqual(s1, s2))
	t.Log("s1 == s3?", reflect.DeepEqual(s1, s3))

	c1 := Customer{"1", "Mike", 40}
	c2 := Customer{"1", "Mike", 40}
	t.Log(c1 == c2)
	t.Log(reflect.DeepEqual(c1, c2))
}
```

## 关于 “反射” 你应该知道的
* 提高了程序的灵活性
* 降低了程序的可读性
* 降低了程序的性能



