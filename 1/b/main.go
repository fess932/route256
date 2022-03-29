package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		n := s.Text()
		if len(n) != 11 {
			fmt.Printf("")
			return
		}

		// opcode, err := strconv.Atoi(n[1:4])
		// if err != nil {
		// 	log.Fatal(err)
		// }

		fmt.Printf(n[1:4])
	}

}
