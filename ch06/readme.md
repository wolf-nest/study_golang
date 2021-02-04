# 数组 array
* 数组是固定长度的数组，使用前必须确定数组长度
* 数组是值类型,也就是说，如果你将一个数组赋值给另外一个数组，那么，实际上就是整个数组拷贝了一份
* 如果数组作为函数的参数，那么实际传递的参数是一份数组的拷贝，而不是数组的指针
* 数组的长度也是Type的一部分，这样就说明[10]int和[20]int是不一样的。

# 切片 slice
* 切片是一个引用类型，是一个动态的指向数组切片的指针。
* 切片是一个不定长的，总是指向底层的数组array的数据结构。
* 切片是共享存储结构
* 切片长度可变 cap

# 数组 vs 切片 不同之处
1. 容量是否可伸缩
    * 数组容量不可伸缩
    * 切片容量可以伸缩
        1. 如果 增加的 len < 切片的cap 则 新切片的cap 为当前cap*2
        2. 如果 增加的 len > 切片的cap 则 新切片的cap = 当前cap + 增加数据的 len 
2. 是否可以进行比较
    * 数组相同位数与相同长度可比较，只要数组中的每个元素都相同就会被认为数组相同
    * 切片只能与空进行比较（slice can only be compared to nil）
3. 类型声明
    * 声明数组时，方括号内写明了数组的长度或者...
    * 声明slice时候，方括号内为空
4. 作为函数参数时
    * array传递的是数组的副本
    * slice传递的是指针