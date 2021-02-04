# 字符串

1. string 是数据类型，不是引用或指针类型
2. string 是只读的 byte slice，len函数可以统计它所包含的byte数
3. string 的 byte 数组可以存放任何数据

### 编码与存储
    字符                “中”
    Unicode             0x4E2D
    UTF-8               0xE4B8AD
    string/[]byte       [0xE4,0xB8,0xAD]
### 常用字符串函数
1. string 包(https://golang.org/pkg/strings)
2. strconv 包(https://golang.org/pkg/strconv)

# Unicode UTF8
1. Unicode 是一种字符集（code point）
2. UTF8 是 unicode 的存储实现（转换为字节序列的规则）


