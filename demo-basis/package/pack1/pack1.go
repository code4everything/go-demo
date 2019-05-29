package pack1

import (
	"fmt"
	"time"
)

func PrintNow() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"))
}
