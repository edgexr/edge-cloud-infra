package main

import (
	"context"

	"github.com/mobiledgex/edge-cloud/cloud-resource-manager/platform/pc"
	"github.com/mobiledgex/edge-cloud/edgeproto"
	"github.com/mobiledgex/edge-cloud/log"
)

// K8s Cluster
type K8sClusterStats struct {
	key      edgeproto.ClusterInstKey
	promAddr string
	client   pc.PlatformClient
	ClusterMetrics
}

func (c *K8sClusterStats) GetClusterStats(ctx context.Context) *ClusterMetrics {
	if c.promAddr == "" {
		return nil
	}
	if err := collectClusterPrometheusMetrics(ctx, c); err != nil {
		log.SpanLog(ctx, log.DebugLevelMetrics, "Could not collect cluster metrics", "K8s Cluster", c)
		return nil
	}
	return &c.ClusterMetrics
}

// Currently we are collecting stats for all apps in the cluster in one shot
// Implementing  EDGECLOUD-1183 would allow us to query by label and we can have each app be an individual metric
func (c *K8sClusterStats) GetAppStats(ctx context.Context) map[MetricAppInstKey]*AppMetrics {
	if c.promAddr == "" {
		return nil
	}
	metrics := collectAppPrometheusMetrics(ctx, c)
	if metrics == nil {
		log.SpanLog(ctx, log.DebugLevelMetrics, "Could not collect app metrics", "K8s Cluster", c)
	}
	return metrics
}