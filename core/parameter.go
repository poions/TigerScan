package core
import (
	"flag"
	"os"
	"fmt"
)

var (
	m	string
	h	bool
	all	bool
	u	string
	p	string
	f	string
)

const (
	redisType		string = "redis"
	checkAll		string = "all"
	defaultUsersDict	string = "dict/user.txt"
	defaultPassDict		string = "dict/pwd.txt"
)


func init() {
	flag.BoolVar(&h, "h", false, "You can perform the scanning task by knowing the available modules")
	flag.BoolVar(&all, "all", true, "Start the full module scan of the input address!")
	flag.StringVar(&m, "m", "", "Select the module you want to scan,the available modules can be viewed through --list example: ./scan -u 127.0.0.1 -p 6379 -m redis")
	flag.StringVar(&u, "u", "", "Specify an IP address for detection")
	flag.StringVar(&p, "p", "", "Specify a port service for instrumentation")
	flag.StringVar(&f, "f", "", "Select the list of targets to scan and recommend TXT format!")

	flag.Usage = usage
}

func Input() {
	flag.Parse()
	if h {
	    flag.Usage()
	    os.Exit(0)
	}
	logo()
	switch {
	    case len(m) != 0:
	       ByhandleModule(m,u,p) 
	    case all:
		if len(f) != 0 {
		    if len(p) != 0{
		        targetList := ReadDict(f)
		        for _, target := range targetList {
		            CheckALLexec(target, p)
		        }
		    } else {
			    fmt.Println("Please enter the port number you want to scan! example: ./scan -f dict/ip.txt -p 6379")
		    }
		} else {
		    switch {
		      case len(u) == 0:
			  fmt.Println("Please enter the correct parameters, example: ./scan -u 127.0.0.1 -p 6379")
			  os.Exit(0)
		      case len(p) == 0:
			  fmt.Println("Please enter the port number you want to scan! example: ./scan -u 127.0.0.1 -p 6379")
			  os.Exit(0)
		      default:
			  CheckALLexec(u, p)
		      }
		}
	}
}


func logo() {
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
	logo()
	fmt.Fprintf(os.Stderr, `
TigerScan version: TigerScan/1.10.0
Usage: Tiger [-upfhALL] [-all default] [-f filename] [-p port] [-u ipaddr]
Options:
	`)
	flag.PrintDefaults()
}
