package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	s, _ := in.ReadString('\n')
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	if strings.HasPrefix(s, "i") &&
		strings.HasSuffix(s, "n") &&
		strings.Contains(s, "a") {
		fmt.Println("Found!")
		return
	}
	fmt.Println("Not Found!")
}
