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
	f := flag.String("f", "1.two-sum.go", "Leetcode File Name")
	test := flag.Bool("t", false, "Test File Name")
	submit := flag.Bool("s", false, "Submit leetcode File Name")

	flag.Parse()

	out, err := exec.Command("git", "add", *f).Output()

	input, err := ioutil.ReadFile(*f)

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
	err = ioutil.WriteFile(*f, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}

	cmd := []string{"run"}
	if *submit {
		cmd = append(cmd, "s", *f)
	} else if *test {
		cmd = append(cmd, "t", *f)
	} else {
		cmd = append(cmd, "-h")
	}

	out, err = exec.Command("npm", cmd...).Output()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(out))

	out, err = exec.Command("git", "checkout", "--", *f).Output()

	fmt.Println(string(out))
	os.Exit(0)
}
