// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gke

import (
	"context"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/compute/metadata"

	"go.opentelemetry.io/otel/api/kv"
	"go.opentelemetry.io/otel/api/standard"
	"go.opentelemetry.io/otel/sdk/resource"

	"go.opentelemetry.io/contrib/detect/gcp"
)

// Detect detects associated resources when running in GKE environment.
func Detect(ctx context.Context) (*resource.Resource, error) {
	if os.Getenv("KUBERNETES_SERVICE_HOST") == "" {
		return nil, nil
	}

	labels := []kv.KeyValue{
		standard.K8SNamespaceNameKey.String(os.Getenv("NAMESPACE")),
		standard.K8SPodNameKey.String(os.Getenv("HOSTNAME")),
		standard.ContainerNameKey.String(os.Getenv("CONTAINER_NAME")),
	}

	clusterName, err := metadata.InstanceAttributeValue("cluster-name")
	logError(err)

	if clusterName != "" {
		labels = append(labels, standard.K8SClusterNameKey.String(clusterName))
	}

	k8sLabelRes := resource.New(labels...)

	gcpDetecor := gcp.GCP{}

	gceLablRes, err := gcpDetecor.Detect(ctx)

	if err != nil {
		return nil, err
	}

	return resource.Merge(gceLablRes, k8sLabelRes), nil
}

// logError logs error only if the error is present and it is not 'not defined'
func logError(err error) {
	if err != nil {
		if !strings.Contains(err.Error(), "not defined") {
			log.Printf("Error retrieving gcp metadata: %v", err)
		}
	}
}
