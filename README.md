# zdpgo_file
Golang操作文件或者文件夹的快捷组件


## 一、概述

### 1.1 项目地址
- Github：https://github.com/zhangdapeng520/zdpgo_file

### 1.1 功能
- 判断文件是否存在
- 获取文件大小
- 获取文件夹大小

## 二、快速入门

### 2.1 判断文件是否存在
```go
func TestFileExists(t *testing.T) {
	fmt.Println(FileExists("./file.go"))
	fmt.Println(FileExists("./file111.go"))
}
```

### 2.2 获取文件大小
```go
func TestFileSize(t *testing.T){
	fmt.Println(FileSize("./file.go"))
	fmt.Println(FileSize("./file1.go"))
}
```

### 2.3 获取文件夹大小
```go
func TestDirectorySize(t *testing.T){
	fmt.Println(DirectorySize("D:\\BaiduNetdiskWorkspace\\文档"))
}
```
