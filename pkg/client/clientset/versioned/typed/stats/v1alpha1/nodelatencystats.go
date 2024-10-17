// Copyright 2024 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	v1alpha1 "antrea.io/antrea/pkg/apis/stats/v1alpha1"
	scheme "antrea.io/antrea/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gentype "k8s.io/client-go/gentype"
)

// NodeLatencyStatsGetter has a method to return a NodeLatencyStatsInterface.
// A group's client should implement this interface.
type NodeLatencyStatsGetter interface {
	NodeLatencyStats() NodeLatencyStatsInterface
}

// NodeLatencyStatsInterface has methods to work with NodeLatencyStats resources.
type NodeLatencyStatsInterface interface {
	Create(ctx context.Context, nodeLatencyStats *v1alpha1.NodeLatencyStats, opts v1.CreateOptions) (*v1alpha1.NodeLatencyStats, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.NodeLatencyStats, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.NodeLatencyStatsList, error)
	NodeLatencyStatsExpansion
}

// nodeLatencyStats implements NodeLatencyStatsInterface
type nodeLatencyStats struct {
	*gentype.ClientWithList[*v1alpha1.NodeLatencyStats, *v1alpha1.NodeLatencyStatsList]
}

// newNodeLatencyStats returns a NodeLatencyStats
func newNodeLatencyStats(c *StatsV1alpha1Client) *nodeLatencyStats {
	return &nodeLatencyStats{
		gentype.NewClientWithList[*v1alpha1.NodeLatencyStats, *v1alpha1.NodeLatencyStatsList](
			"nodelatencystats",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1alpha1.NodeLatencyStats { return &v1alpha1.NodeLatencyStats{} },
			func() *v1alpha1.NodeLatencyStatsList { return &v1alpha1.NodeLatencyStatsList{} }),
	}
}