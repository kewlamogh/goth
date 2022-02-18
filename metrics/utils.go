package metrics

import "fmt"

func checkError(err error) {
	if err != nil {
		fmt.Printf("err: %v", err)
	}
}
