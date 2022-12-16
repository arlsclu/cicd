package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		<-time.Tick(3 * time.Second)
		fmt.Printf("hello world  , nowt time is : %s ", time.Now())
	}

}
