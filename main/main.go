package main

import (
	"context"
	"fmt"
)

func main() {
	gcp := GCE{}
	res, err := gcp.Detect(context.Background())
	fmt.Println(res, err)
}
