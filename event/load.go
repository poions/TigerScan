package event

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"net"
)



//ReadLocalDict : 读取本地字典数据
func ReadLocalDict(dictNamePath string) []string {
	file, err := os.Open(dictNamePath)
	if err != nil {
		//
		log.Printf("%s,not find files!", err)
		os.Exit(0)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("%s,get dict data error !", err)
		os.Exit(0)
	}
	string_slice := strings.Split(string(content), "\n")
	return string_slice
}

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

//LoadLocalConf : 获取配置文件数据内容
func LoadLocalConf() SettingConf {
	JsonParse := NewJsonStruct()
	v := SettingConf{}
	JsonParse.Load(SettingPath, &v)

	return v
}

func (jst *JsonStruct) Load(filename string, v interface{}) {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("%s,get setting data error !", err)
		return
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		log.Printf("%s,get dict data error !", err)
		return
	}
}

//CheckPluginsList : 检查可用插件
func CheckPluginsList() ([]string, []string) {
	listPluginsName := []string{}
	listModuleName := []string{}
	files, _ := ioutil.ReadDir("plugins")
	for _, f := range files {
		//fmt.Println(f.Name())
		pluginName := strings.Split(f.Name(), ".")[0]
		listPluginsName = append(listPluginsName, pluginName)
		listModuleName = append(listModuleName, f.Name())
	}
	return listPluginsName, listModuleName
}


func Hosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}
	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	return ips[1 : len(ips)-1], nil
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}


