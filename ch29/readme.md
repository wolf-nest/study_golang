# http

## Default Router
```$go
func (handler serverHandler) ServeHTTP(rw http.)
```

## 路由规则
* URL 分为两种，末尾是/：表示一个子树，后面可以跟其他子路径；末尾不是/，表示一个叶子，固定路径  
    ##### *以/结尾的URL 可以匹配它的任何子路径 如：/images 会匹配/images/cute-cat.jpg*  
* 它采用最长匹配原则，如果有多个匹配，一定采用匹配路径最长的那个进行处理  
* 如果没有找到任何匹配项，会返回404错误  