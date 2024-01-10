package others

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestProduce(t *testing.T) {
	kk, _ := os.Create("k.txt")

	readFile, err := os.Open("tpi.log")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	i := 0
	for fileScanner.Scan() {
		i++
		//fmt.Println(i, ":", fileScanner.Text())
		temp := fileScanner.Text()
		split := strings.Split(temp, "->")
		temp = split[1]
		index := strings.Index(temp, "{")
		if index > 0 {
			temp = temp[index:]
		}
		index = strings.LastIndex(temp, "}")
		if index > 0 {
			temp = temp[:index+1]
		}
		//fmt.Println(temp)
		//if i == 10 {
		//	break
		//
		//}
		temp = "lpush QJCG_third_party_interface_yh " + "'" + temp + "'"

		kk.WriteString(temp + "\n")
	}

	readFile.Close()
}
