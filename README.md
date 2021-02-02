# TigerScan
## 描述



	████████╗██╗ ██████╗ ███████╗██████╗ 
	╚══██╔══╝██║██╔════╝ ██╔════╝██╔══██╗
	   ██║   ██║██║  ███╗█████╗  ██████╔╝
	   ██║   ██║██║   ██║██╔══╝  ██╔══██╗
	   ██║   ██║╚██████╔╝███████╗██║  ██║
	   ╚═╝   ╚═╝ ╚═════╝ ╚══════╝╚═╝  ╚═╝ 
	

## 项目说明

###  TigerScan:
TigerScan是一款可扩展型风险扫描程序，通过自定义漏洞的检测风险规则逻辑，将其打包成so链接库的形式供tiger调用




### 所有参数:
```
Usage: TigerScan [-m] [-m redis] [-all] [-u] [-f filename] 
Options:
  -V	show version and exit
  -h	this help
  -all
      扫描检测plugin中的所有模块
  -list
    	可列出所有可使用的模块
  -p string
    	指定端口号
  -m string
    	指定扫描模块
  -u string
    	指定扫描地址
  -f string
  	选择批量扫描导入的txt文档
```

## 使用教程:
git clone 将该项目下载下来后，通过go build 生成二进制文件.
期间可能会需要下载plugin中使用到的第三方库，可通过go get的形式下载
```
cd TigerScan
go build 
./TigerScan -u 127.0.0.1 -p 22 -m ssh
```

## 插件扩展
### 注意事项
 - 插件主函数名称必须为Poc, 例子:  func Poc() {}
 - 插件主函数需要接入参数ip,port,username,password
 - 插件主函数Poc return返回数据必须包含下面的字段，且必须为string格式
```
type Filter  struct {
    Status	bool	`json:"status"`
    IP		string	`json:"ip"`
    PORT	string	`json:"port"`
    USERNAME	string	`json:"username"`
    PASSWORD	string	`json:"password"`
    Message	string	`json:"message"`
}
```
### 例子
redis.go
```
package main
import (
    "fmt"
)

func Poc(ip, port, username, password) string {
    ....
    #如果检测逻辑判断为存在漏洞风险则return true否则为false;这里true/false 为string格式
    return fmt.Sprintf("%s:%s:%s:%s:%s","true", ip, port, username, password, msg)
}

```
