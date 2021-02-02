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
    	List all detection types
  -p string
    	Specify an ports for risk scanning
  -m string
    	Specifies the type of risk scan:
  -u string
    	Specify an address for risk scanning
```

## 使用教程:
git clone 将该项目下载下来后，通过go build 生成二进制文件.
期间可能会需要下载plugin中使用到的第三方库，可通过go get的形式下载
'''
cd TigerScan
go build 
./TigerScan -u 127.0.0.1 -p 22 -m ssh
'''
