package main

import (
	"context"
)

func main() {
	ctx1 := context.TODO()
	ctx2 := context.Background()

	_ = ctx1 == ctx2
}
