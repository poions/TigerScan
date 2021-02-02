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

## 目前检测类型有:
| 规则ID | 规则描述 | 备注 |\n
| `1`  | redis未授权访问 | |
| `2`  | redis弱口令 | |
| `3`  | ssh弱口令 | |
