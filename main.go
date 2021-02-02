package main
import (
	"TigerScan/core"
	"time"
	"fmt"
)


func main(){
	//core.CheckALLexec()

	beforeTime := time.Now()
	core.Input()
	str := time.Since(beforeTime)
	fmt.Printf("\n扫描检测消耗的时间:%s\n", str)
}
