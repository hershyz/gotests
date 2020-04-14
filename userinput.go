package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	term, _ := reader.ReadString('\n')
	term = strings.TrimSpace(term)
	commandLoop()
}
