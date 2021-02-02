package core
import (
    "fmt"
    "log"
    "plugin"
    "io/ioutil"
    "time"
    "strings"
    "unsafe"
    "reflect"
    "os"
)



func BytesToMyStruct(b []byte) *Filter {
    return (*Filter)(unsafe.Pointer(
	    (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
    ))
}


func CheckALLexec(ip, port string) {
	files, _ := ioutil.ReadDir("plugins")
	for _, f := range files {
	    callPlugins(f.Name(),ip, port)
	}
}

func ByhandleModule(m, ip, port string){
	fName := fmt.Sprintf("%s.so", m)
	callPlugins(fName, ip, port)
}


func callPlugins(fName, ip, port string) {
	users := ReadDict(defaultUsersDict)
	passwd := ReadDict(defaultPassDict)

	for _, username := range users {
	    for _, password := range passwd {
	        go func(username, password string) {
		    fmt.Fprintf(os.Stdout, "Type: %s ip:%s port:%s user:%s passwd:%s\r",fName,ip,port,username,password)
	            pocPath := fmt.Sprintf("plugins/%s", fName)
	            p, err := plugin.Open(pocPath)
	            if err != nil {
	                log.Println("Could not find this module("+fName+"), please check to see if the module name exists!")
			os.Exit(0)
	            }
	            call, err := p.Lookup("Poc")
	            if err != nil {
	                panic(err)
	            }
	            run := call.(func(string, string, string, string) string)(ip, port, username, password)
		    str := strings.Split(run, ":")
		    if str[0] == "true"{
		        Outputres(run)
			return
		    }
		}(username, password)
		time.Sleep(10e7)
	    }
	}
}

func Outputres(r string) {
	log.Println(r)
}
