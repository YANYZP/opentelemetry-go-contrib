package main

import (
	"context"
	"fmt"

	"contrib.go.opencensus.io/exporter/stackdriver/monitoredresource/gcp"
)

func main() {
	gcp := gcp.GCP{}
	res, err := gcp.Autodetect(context.Background())
	fmt.Println(res)

}
