package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	lastP := os.Args[len(os.Args)-1]
	test := flag.Bool("t", false, "Test File Name")
	submit := flag.Bool("s", false, "Submit leetcode File Name")
	gen := flag.Bool("g", false, "Generate leetcode File Name")
	show := flag.Bool("sh", false, "Show leetcode File Name")

	flag.Parse()

	if _, err := os.Stat(lastP); err == nil {
		_, err := exec.Command("git", "add", lastP).Output()

		input, err := ioutil.ReadFile(lastP)

		if err != nil {
			log.Fatalln(err)
		}

		lines := strings.Split(string(input), "\n")

		for i, line := range lines {
			if strings.Contains(line, "package leetcode") {
				lines[i] = ""
			}
		}
		output := strings.Join(lines, "\n")
		err = ioutil.WriteFile(lastP, []byte(output), 0644)
		if err != nil {
			log.Fatalln(err)
		}
	}

	cmd := []string{"run"}
	if *submit {
		cmd = append(cmd, "s", lastP)
	} else if *test {
		cmd = append(cmd, "t", lastP)
	} else if *gen {
		cmd = append(cmd, "g", lastP)
	} else if *show {
		cmd = append(cmd, "sh", lastP)
	} else {
		cmd = append(cmd, "-h")
	}

	out, err := exec.Command("npm", cmd...).Output()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(out))

	out, err = exec.Command("git", "checkout", "--", lastP).Output()

	fmt.Println(string(out))
	os.Exit(0)
}
