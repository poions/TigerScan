package core
import (
	"strings"
	"os"
	"bufio"
	"io"
)

func ReadDict(fpath string) []string{
	var dict []string
	//fmt.Println(fpath)
	//content, err := ioutil.ReadFile(fpath)
	f, err := os.Open(fpath)
	if err != nil {
	    return nil
	}
	buf := bufio.NewReader(f)
	for {
	    line, err := buf.ReadString('\n')
	    line = strings.TrimSpace(line)
	    if err != nil {
		if err == io.EOF {
		    return dict
		}
		return nil
	    }
	    dict = append(dict, line)
	}
	return dict
}
