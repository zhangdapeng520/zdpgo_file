# zdpgo_file
Golang操作文件或者文件夹的快捷组件

项目地址：https://github.com/zhangdapeng520/zdpgo_file

## 版本历史
- 2022年2月7日：版本0.1.0
- 2022年4月3日：版本0.1.1 项目结构优化

## 使用示例
### 读写csv
```go
package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_file"
)

func main() {
	f := zdpgo_file.New()
	data := [][]string{
		{"a", "b", "c"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
		{"111", "222", "333"},
	}

	// 写入
	f.Csv.Write("test.csv", data)

	// 读取
	dataNew, err := f.Csv.Read("test.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dataNew)
}
```

### 文件夹操作
```go
package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_file"
)

func main() {
	f := zdpgo_file.New()
	fmt.Println(f.Directory.GetDirectorySize("../download"))
	fmt.Println(f.Directory.IsExist("D:\\BaiduNetdiskWorkspace\\文档"))
	fmt.Println(f.Directory.CreateMultiDir("logs/zdpgo"))
}
```

### 下载网络文件
```go
package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_file"
)

func main() {
	f := zdpgo_file.New()
	savePath := "./"

	// 单个下载
	url := "https://avatar.csdnimg.cn/2/9/0/1_togolife.jpg"
	err := f.Download.Download(savePath, url)
	if err != nil {
		fmt.Println("下载失败：", err)
	} else {
		fmt.Println("下载成功")
	}

	// 批量下载
	urls := []string{
		"https://alifei04.cfp.cn/creative/vcg/nowarter800/new/VCG41N695593548.jpg",
		"https://tenfei02.cfp.cn/creative/vcg/nowarter800/new/VCG41N1014325904.jpg",
		"https://tenfei05.cfp.cn/creative/vcg/nowater800/new/VCG41545444880.jpg",
	}
	f.Download.Downloads(savePath, urls)
}
```

## 文件操作
```go
package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_file"
)

func main() {
	f := zdpgo_file.New()
	fmt.Println(f.File.Exists("./file.go"))
	fmt.Println(f.File.Exists("./file111.go"))

	fmt.Println(f.File.Size("./file.go"))
	fmt.Println(f.File.Size("./file1.go"))
}
```