# Map 与工厂模式

* Map 的 value 可以是一个方法
* 与 Go 的 Dock type 接口方式一起，可以方便的实现单一方法对象的工厂模式

# 实现 Set
### Go 的内置集合中没有 Set 实现，可以使用 map[type]bool

1. 元素的唯一性
2. 基本操作
    1) 添加元素
    2) 判断元素是否存在
    3) 删除元素
    4) 元素个数
