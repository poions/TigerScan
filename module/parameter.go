package module

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

var (
	h	bool
	list	bool
	inner	bool
	m	string
	u	string
	p	string
	f	string
	r	string
)

const (
	redisType        string = "redis"
	checkAll         string = "all"
	defaultUsersDict string = "conf/user.txt"
	defaultPassDict  string = "conf/pwd.txt"
)

func Logo() {
	fmt.Fprintf(os.Stderr, `
	████████╗██╗ ██████╗ ███████╗██████╗ 
	╚══██╔══╝██║██╔════╝ ██╔════╝██╔══██╗
	   ██║   ██║██║  ███╗█████╗  ██████╔╝
	   ██║   ██║██║   ██║██╔══╝  ██╔══██╗
	   ██║   ██║╚██████╔╝███████╗██║  ██║
	   ╚═╝   ╚═╝ ╚═════╝ ╚══════╝╚═╝  ╚═╝ 
`)
}

func usage() {
	//Logo()
	fmt.Fprintf(os.Stderr, `
TigerScan version: TigerScan/1.10.0
Usage: Tiger [-upfhALL] [-all default] [-f filename] [-p port] [-u ipaddr] [-list --list] [-r --range]
Options:
	`)
	flag.PrintDefaults()
}

func init() {
	flag.BoolVar(&h, "h", false, "You can perform the scanning task by knowing the available modules")
	flag.BoolVar(&inner, "inner", false, "You can running host risk scan for inner network, the inner network address data from assets db")
	flag.BoolVar(&list, "list", false, "Lists all currently available modules")
	flag.StringVar(&m, "m", "", "Select the module you want to scan,the available modules can be viewed through --list example: ./scan -u 127.0.0.1 -p 6379 -m redis")
	flag.StringVar(&u, "u", "", "Specify an IP address for detection")
	flag.StringVar(&p, "p", "", "Specify a port service for instrumentation")
	flag.StringVar(&f, "f", "", "Select the list of targets to scan and recommend TXT format!")
	flag.StringVar(&r, "r", "", "扫描整个段如: 192.168.0.1/24, 192.168.0.1/16")

	flag.Usage = usage
}

func InputScan() {
	flag.Parse()
	/*
		if h {
			flag.Usage()
			os.Exit(0)
		}
	*/

	switch {
	case h == true:
		flag.Usage()
		os.Exit(0)
	case list == true:
		ListPluginsModuleLog()
	case len(f) != 0:
		//判断选择的是否是批量检测
		ChooseBatchScanAddr(f, p, m)

	case len(u) != 0:
		//判断选择的是否是单条记录检测
		ChooseSingleScanAddr(u, p, m)
	case len(r) != 0:
		//扫描段
		ChooseScanCIDRAddr(r, p, m)
	}
}

//judgePortInputValidity : 判断端口输入的合法性
func judgePortInputValidity(p string) bool {
	regStr := `^(\d+)$`
	match, _ := regexp.MatchString(regStr, p)
	if match == true {
		return true
	}
	log.Printf("Please enter the port number you want to scan! example: ./scan -f dict/ip.txt -p 6379")
	return false
}

//judgeModuleInputNone : 判断检测模块输入是否为空，默认为扫描全部all
func judgeModuleInputNone(m string) bool {
	if len(m) != 0 {
		return true
	} else {
		return false
	}
}

//judgeIPaddrValidity : 校验输入地址的合法性
func judgeIPaddrValidity(ipaddr string) bool {
	regStr := `^(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){2}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	match, _ := regexp.MatchString(regStr, ipaddr)
	return match
}

func judgePluginsNameValidity(moduleName string) bool {
	files, _ := ioutil.ReadDir("plugins")
	for _, f := range files {
		//fmt.Sprintln(f.Name(), "+++")
		regStr := fmt.Sprintf(`^(%s)$`, moduleName)
		match, _ := regexp.MatchString(regStr, f.Name())
		if match == true {
			return true
		}
	}
	return false
}

func judegeIpRangeValidity(ipaddr string) bool {
	matchs, _ := regexp.MatchString(`\d+\.\d+\.\d+\.\d+\/\d+`, ipaddr)
	return matchs
}
