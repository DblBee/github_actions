package cli_utils

import (
	"bufio"
	"fmt"
	"io"
)

func OutputLog(readCloser io.ReadCloser) {
	goTestScanner := bufio.NewScanner(readCloser)
	goTestScanner.Split(bufio.ScanLines)
	for goTestScanner.Scan() {
		m := goTestScanner.Text()
		fmt.Println(m)
	}
}
