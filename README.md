# Go-CodeGen —— Go语言实现的代码生成器
作者：千石（掘金社区主页：[https://juejin.cn/user/3843548383816862](https://juejin.cn/user/3843548383816862)）
## 引言
本项目参加掘金社区举办的[码上掘金迎新年后端赛](https://juejin.cn/challenge/3)
## 特性
此程序可以帮助你生成代码，提高你的Coding效率，具体有如下特性：
- 不涉及第三方库，仅使用原生库
- 生成的二进制文件仅3MB大小
- 支持根据模板文件生成代码
- 支持json的解析
## 支持
程序目前支持以下Go代码的生成
- 结构体(struct)
- 函数(func)
- 方法
## 使用
### 打包
```go
go build -o GoCodeGen
```
### 帮助
```
GoCodeGen -h
```
### 参数
```
-dir string
    代码输出目录 (默认 "output")
-json string
    Json文件 (默认 "./json/default.json")
-name string
    实例名称
-pkg string
    包名
```
