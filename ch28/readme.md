# Json解析

## 内置的JSON 解析
1. 利用反射实现，通过FeildTag 来标识对应的 json 值
```$go
type BasicInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type JobInfo struct {
	Skills []string `json:"skills"`
}

type Employee struct {
	BasicInfo BasicInfo `json:"basic_info"`
	JobInfo   JobInfo   `json:"job_info"`
}
```
## 更快的 JSON 解析

### EasyJSON 采用代码生成而非反射

1. 安装  
`go get -u github.com/mailru/easyjson/`  
2. 使用  
`easyjson -all <结构定义>.go`