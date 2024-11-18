package utils

import (
	"fmt"
	"os"
)

func Debug(a ...any) {
	if os.Getenv("DEBUG") != "true" {
		return
	}
	fmt.Println(a...)
}
