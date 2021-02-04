## 准备工作
* 安装graphviz
    * brew install graphviz
* 将$GOPATH/bin加入$PATH
    * Mac OS: 在.zshrc中修改或添加路径
* 安装 go-torch
    * go get github.com/uber-archive/go-torch
    * 下载并复制 flamegraph.pl(https://github.com/brendangregg/FlameGraph/blob/master/flamegraph.pl) 至 $GOPATH/bin 路径下  

## 通过文件方式输出Profile
* 灵活性高，适用于特定代码段的分析
* 通过手动调用runtime/pprof的API
* API相关文档
* go tool pprof [binary][binary.prof]

## 常见分析指标
* Wall Time
* CPU Time
* Block Time
* Memory allocation
* GC times/time spent