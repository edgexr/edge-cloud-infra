package cliwrapper

import (
	"github.com/mobiledgex/edge-cloud-infra/mc/ormapi"
)

func (s *Client) ShowAppEvents(uri, token string, query *ormapi.RegionAppInstMetrics) (*ormapi.AllMetrics, int, error) {
	args := []string{"events", "app"}
	metrics := ormapi.AllMetrics{}
	st, err := s.runObjs(uri, token, args, query, &metrics)
	return &metrics, st, err
}
func (s *Client) ShowClusterEvents(uri, token string, query *ormapi.RegionClusterInstMetrics) (*ormapi.AllMetrics, int, error) {
	args := []string{"events", "cluster"}
	metrics := ormapi.AllMetrics{}
	st, err := s.runObjs(uri, token, args, query, &metrics)
	return &metrics, st, err
}

func (s *Client) ShowCloudletEvents(uri, token string, query *ormapi.RegionCloudletMetrics) (*ormapi.AllMetrics, int, error) {
	args := []string{"events", "cloudlet"}
	metrics := ormapi.AllMetrics{}
	st, err := s.runObjs(uri, token, args, query, &metrics)
	return &metrics, st, err
}