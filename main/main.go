package main

import (
	"context"
	"fmt"
)

func main() {
	gcp := GCE{}
	res, _ := gcp.Detect(context.Background())
	for _, ele := range res.Attributes() {
		fmt.Println(ele)
	}
}
