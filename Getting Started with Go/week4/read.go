package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type names struct {
	fname string
	lname string
}

func main() {
	fmt.Print("enter file location and name: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	var name []names
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, _, err := rd.ReadLine()

		if err != nil || io.EOF == err {
			break
		}
		tname := strings.Split(string(line), " ")
		tpname := names{
			fixLongName(tname[0]),
			fixLongName(tname[1]),
		}
		name = append(name, tpname)
	}
	for i := 0; i < len(name); i++ {
		fmt.Println(i + 1)
		fmt.Println("First Name:" + name[i].fname)
		fmt.Println("Last Name:" + name[i].lname)
	}
}

// Cut the string if the name is too long
func fixLongName(buffer string) string {

	if len(buffer) > 20 {
		return string(buffer[0:20])
	} else {
		return buffer
	}
}
