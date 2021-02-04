# Package
1. 基本复用模块单元  
    以首字母大写来表明可被包外代码访问
2. 代码的 package 可以和所在的目录不一致
3. 同一目录里的 Go 代码的 package 要保持一致

## init 方法
* 在 main 被执行前，所有依赖的 package 的init方法都会被执行
* 不同的包 的init函数按照包导入的依赖关系决定执行顺序
* 每个包可以有多个init函数
* 包的 每个源文件也可以有多个init函数 *这点比较特殊*

## vendor 路径
### 查找依赖包路径的解决方案如下
1. 当前包下的vendor目录
2. 向上级目录查找，直到找到src下的vendor目录
3. 在GOPATH下面查找依赖包
4. 在GOPATH目录下查找

## 常用的依赖管理工具

godep https://github.com/tools/godep  
glide https://github.com/masterminds/glide  
dep https://github.com/golang/dep