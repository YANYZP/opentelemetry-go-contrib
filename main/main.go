package main

import (
	"context"
	"fmt"

	"go.opentelemetry.io/contrib/detect/gcp"
)

func main() {
	gcp := gcp.GCP{}
	res, err := gcp.Detect(context.Background())
	fmt.Println(res, err)

}
