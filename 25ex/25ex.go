package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Print("are you ready?\n")
	betterSleep(1000000000)
	fmt.Print("lets go \n")
}

func betterSleep(seconds int) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(seconds))
	defer cancel()
	<-ctx.Done()
}
