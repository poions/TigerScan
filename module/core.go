package module

import (
	"fmt"
	"log"
	"time"
	"os"
	"plugin"
	"strings"

	"TigerScan/event"
)

func ChooseSingleScanAddr(addr, port, m string) {
	if judgeModuleInputNone(m) == true {
		//ReadLocalDict(f)
		fName := fmt.Sprintf("%s.so", m)
		execRunningMain(addr, port, fName)
	} else {
		os.Exit(0)
	}

}

//ChooseBatchScanAddr : 批量检测入口
func ChooseBatchScanAddr(filepath, port, m string) {
	//判断端口输入是否为空
	ipaddr := event.ReadLocalDict(filepath)
	for i := 0; i < len(ipaddr); i++ {
		if judgeModuleInputNone(m) == true {
			fName := fmt.Sprintf("%s.so", m)
			execRunningMain(ipaddr[i], port, fName)
		} else {
			os.Exit(0)
		}
	}
}

func ChooseScanCIDRAddr(r, port, m string) {
	for {
		if judegeIpRangeValidity(r) != true {
			os.Exit(0)
		}
		addrs, _ := event.Hosts(r)
		fName := fmt.Sprintf("%s.so", m)
		if judgeModuleInputNone(m) == true {
			for i := 0; i < len(addrs); i++ {
				execRunningMain(addrs[i], port, fName)
			}
		}else {
			os.Exit(0)
		}
	}
}


func execRunningMain(ip, port, m string) {
	switch {
	case judgeIPaddrValidity(ip) != true:
		log.Println("the input ip is illegality!")
		os.Exit(0)
	case judgePortInputValidity(port) != true:
		log.Println("the input port is illegality!")
		os.Exit(0)
	case judgePluginsNameValidity(m) != true:
		log.Println("the plugins module is illegality!")
		os.Exit(0)
	}
	users := event.ReadLocalDict(event.LoadLocalConf().LocalUserPath)
	passwd := event.ReadLocalDict(event.LoadLocalConf().LocalPwdPath)
	callPlugins(m, ip, port, users, passwd)
	//test(m, ip, port, users, passwd)
}

//ListPluginsModuleLog : 列出可用模块
func ListPluginsModuleLog() {
	plugins, _ := event.CheckPluginsList()
	fmt.Println("The modules currently available are：")
	for i := 0; i < len(plugins); i++ {
		fmt.Println("[", i, "]", plugins[i])
	}
}



func callPlugins(fName, ip, port string, users, passwd []string) {

	go func(ip, port string){
		for _, username := range users {
			for _, password := range passwd {
				pocPath := fmt.Sprintf("plugins/%s", fName)
				p, err := plugin.Open(pocPath)
				if err != nil {
					log.Println("Could not find this module(" + fName + "), please check to see if the module name exists!")
					os.Exit(0)
				}
				call, err := p.Lookup("Poc")
				if err != nil {
					panic(err)
				}
				run := call.(func(string, string, string, string) string)(ip, port, username, password)
				str := strings.Split(run, "#")
				var context string

				switch str[0] {
				case "0":
					context = "successed"
				case "1":
					context = "error, username or password err!"
				case "2":
					context = "error, server port err!"
					return
				}
				for r := 0; r < len(str); r++ {
					if str[0] == "0"{
						fmt.Println(str[1],str[2],str[3],str[4],context,"**************")
						return
					}
				}
			}
		}
	}(ip,port)
	time.Sleep(50e6)
}

