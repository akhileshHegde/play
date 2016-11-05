package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func Check(e error) {
        if e!=nil {
                fmt.Println(e)
		os.Exit(1)
        }
}

func runCmd(b string, args []string) {
	s := strings.Fields(b)
	j := 0
	if len(args) > 0 {
		for i, v := range s {
			if v == "<arg>" {
				s[i] = args[j]
				j++
			}
		}
	}
	if s[0] == "-" {
		s[0] = "run"
	}
	fmt.Printf("%s\n|\n|---->\n", s)
	switch s[0] {
	case "cd":
		os.Chdir(s[1])
	case "run":
		c := exec.Command(s[1])
		c.Args = s[1:]
		err := c.Run()
		Check(err)
	default:
		c := exec.Command(s[0])
		c.Args = s
		op, err := c.Output()
		Check(err)
		if op != nil {
			fmt.Printf("%s\n\n", op)
		} else {
			fmt.Printf("<nil>\n\n")
		}
	}
}

func list(dat *map[string][]string) {
	keys := make([]string, len(*dat))
	i := 0
	for k := range *dat {
	    keys[i] = k
	    i++
	}
	for _,v := range keys {
		fmt.Println(v, (*dat)[v])
	}
	os.Exit(0)
}

func main() {
	arg := os.Args[1]
	var dat map[string][]string
	file, err := ioutil.ReadFile(os.Getenv("GOPLAY"))
	Check(err)
	err = json.Unmarshal(file, &dat)
	if arg == "list" {
		list(&dat)
	}
	for _, v := range dat[arg] {
		runCmd(v, os.Args[2:])
	}

}
