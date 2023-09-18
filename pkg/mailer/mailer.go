package gomail

import (
	"bufio"
	"fmt"
	"os"
)

func ReadNewFile(file string) []string {

	var result []string

	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {

		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return result

}
