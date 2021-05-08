# Tiger Scanner 🐯
## 描述

	████████╗██╗ ██████╗ ███████╗██████╗ 
	╚══██╔══╝██║██╔════╝ ██╔════╝██╔══██╗
	   ██║   ██║██║  ███╗█████╗  ██████╔╝
	   ██║   ██║██║   ██║██╔══╝  ██╔══██╗
	   ██║   ██║╚██████╔╝███████╗██║  ██║
	   ╚═╝   ╚═╝ ╚═════╝ ╚══════╝╚═╝  ╚═╝ 
	

## 项目说明

###  TigerScan🐯:
TigerScan是一款可扩展型风险扫描程序，通过自定义漏洞的检测风险规则逻辑，将其打包成so链接库的形式供tiger调用.




### 所有参数:
```
TigerScan version: TigerScan/3.0.1
Usage: Tiger [-upfmhALL] [-all default] [-f filename] [-p port] [-u ipaddr] [-list list_all_module] [-m choose_module]
Options:
  -V	show version and exit
  -h	this help
  -all
      扫描检测plugin中的所有模块(默认： 不带入-m方法的话默认进入-all模式)
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
  -r string
	选择扫描C段或者B地址范围
```

## 使用教程:
git clone 将该项目下载下来后，通过go build 生成二进制文件.
期间可能会需要下载plugin中使用到的第三方库，可通过go get的形式下载
```bash
cd TigerScan
go build 
./TigerScan -u 127.0.0.1 -p 22 -m ssh
```

## 插件扩展
### 说明
 - 支持新增插件
 - 不支持更新插件
### 注意事项
 - 插件主函数名称必须为Poc, 例子:  func Poc() string {}
 - 插件主函数需要接入参数ip,port,username,password
 - 插件主函数Poc return返回数据必须包含下面的字段，且必须为string格式
```golang
type Filter  struct {
    Status	string	`json:"status"`
    IP		string	`json:"ip"`
    PORT	string	`json:"port"`
    USERNAME	string	`json:"username"`
    PASSWORD	string	`json:"password"`
    Message	string	`json:"message"`
}
```
### 例子
redis.go
```golang
package main
import (
    "fmt"
)

func Poc(ip, port, username, password) string {
    ....
    #如果检测逻辑判断为存在漏洞风险则return 0; 否则为1或2
    switch {
    case 检测失败:
    	return fmt.Sprintf("%s:%s:%s:%s:%s","1", ip, port, username, password, msg)
    case 检测成功:
    	return fmt.Sprintf("%s:%s:%s:%s:%s","0", ip, port, username, password, msg)
}
```

### Poc 返回状态码响应说明
状态码 | 格式 | 备注 |
:-----:|:-----:|:-----:|
 0 | string | 扫描成功,发现漏洞 
 1 | string | 扫描失败,账号或密码错误 
 2 | string | 扫描失败,端口访问出错 
 3 | string | 扫描失败,漏洞利用失败 

### 编写完成规则后，需要将检测模块源码打包成链接库形式
```golang
go build --buildmode=plugin -o plugins/redis.so redis.go
```
