package main

import (
	"fmt"

	"github.com/langbox/timef"
)

func main() {
	var timestamp int64 = 1640682392
	str := timef.GetNormalStringByTimestamp(timestamp)
	fmt.Println(str)
}
