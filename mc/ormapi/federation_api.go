// Copyright 2022 MobiledgeX, Inc
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

package ormapi

import (
	"github.com/lib/pq"
)

type Federator struct {
	// Globally unique string used to indentify a federation with partner federation
	FederationId string `gorm:"primary_key" json:"federationid"`
	// Globally unique string to identify an operator platform
	OperatorId string `gorm:"type:citext" json:"operatorid"`
	// ISO 3166-1 Alpha-2 code for the country where operator platform is located
	CountryCode string `json:"countrycode"`
	// Federation access point address
	FederationAddr string `json:"federationaddr"`
	// Region to which this federator is associated with
	Region string `json:"region"`
	// Mobile country code of operator sending the request
	MCC string `json:"mcc"`
	// List of mobile network codes of operator sending the request
	MNC pq.StringArray `gorm:"type:text[]" json:"mnc"`
	// IP and Port of discovery service URL of operator platform
	LocatorEndPoint string `json:"locatorendpoint"`
	// Revision ID to track object changes. We use jaeger traceID for easy debugging
	// but this can differ with what partner federator uses
	// read only: true
	Revision string `json:"revision"`
	// API Key used for authentication (stored in secure storage)
	ApiKey string `json:"apikey"`
	// read only: true
	ApiKeyHash string `gorm:"not null"`
	// read only: true
	Salt string `gorm:"not null"`
	// read only: true
	Iter int `gorm:"not null"`
}

type Federation struct {
	// Name to uniquely identify a federation
	Name string `gorm:"unique; not null" json:"name"`
	// Self federation ID
	SelfFederationId string `gorm:"primary_key; unique" json:"selffederationid"`
	// Self operator ID
	SelfOperatorId string `json:"selfoperatorid"`
	// Partner Federator
	Federator `json:",inline"`
	// Partner shares its zones with self federator as part of federation
	// read only: true
	PartnerRoleShareZonesWithSelf bool
	// Partner is allowed access to self federator zones as part of federation
	// read only: true
	PartnerRoleAccessToSelfZones bool
}

// Details of zone owned by a federator. MC defines a zone as a group of cloudlets,
// but currently it is restricted to one cloudlet
type FederatorZone struct {
	// Globally unique string used to authenticate operations over federation interface
	ZoneId string `gorm:"primary_key" json:"zoneid"`
	// Globally unique string to identify an operator platform
	OperatorId string `gorm:"type:citext" json:"operatorid"`
	// ISO 3166-1 Alpha-2 code for the country where operator platform is located
	CountryCode string `json:"countrycode"`
	// GPS co-ordinates associated with the zone (in decimal format)
	GeoLocation string `json:"geolocation"`
	// Comma seperated list of cities under this zone
	City string `json:"city"`
	// Comma seperated list of states under this zone
	State string `json:"state"`
	// Type of locality eg rural, urban etc.
	Locality string `json:"locality"`
	// Region in which cloudlets reside
	Region string `json:"region"`
	// List of cloudlets part of this zone
	Cloudlets pq.StringArray `gorm:"type:text[]" json:"cloudlets"`
	// Revision ID to track object changes. We use jaeger traceID for easy debugging
	// but this can differ with what partner federator uses
	// read only: true
	Revision string `json:"revision"`
}

// Information of the partner federator with whom the self federator zone is shared
type FederatedSelfZone struct {
	// Globally unique identifier of the federator zone
	ZoneId string `gorm:"primary_key;type:text REFERENCES federator_zones(zone_id)" json:"zoneid"`
	// Self operator ID
	SelfOperatorId string `json:"selfoperatorid"`
	// Name of the Federation
	FederationName string `gorm:"primary_key" json:"federationname"`
	// Zone registered by partner federator
	// read only: true
	Registered bool
	// Revision ID to track object changes. We use jaeger traceID for easy debugging
	// but this can differ with what partner federator uses
	// read only: true
	Revision string `json:"revision"`
}

// Zones shared as part of federation with partner federator
type FederatedPartnerZone struct {
	// Self operator ID
	SelfOperatorId string `json:"selfoperatorid"`
	// Name of the Federation
	FederationName string `gorm:"primary_key" json:"federationname"`
	// Partner federator zone
	FederatorZone `json:",inline"`
	// Zone registered by self federator
	// read only: true
	Registered bool
}

// Register/Deregister partner zones shared as part of federation
type FederatedZoneRegRequest struct {
	// Self operator ID
	SelfOperatorId string `json:"selfoperatorid"`
	// Name of the Federation
	FederationName string `json:"federationname"`
	// Partner federator zones to be registered/deregistered
	Zones []string `json:"zones"`
}

func (f *Federator) GetTags() map[string]string {
	tags := make(map[string]string)
	tags["org"] = f.OperatorId
	tags["region"] = f.Region
	return tags
}

func (f *Federation) GetTags() map[string]string {
	tags := make(map[string]string)
	tags["org"] = f.SelfOperatorId
	tags["federatedorg"] = f.OperatorId
	tags["federationname"] = f.Name
	return tags
}

func (f *FederatorZone) GetTags() map[string]string {
	tags := make(map[string]string)
	tags["org"] = f.OperatorId
	tags["region"] = f.Region
	tags["zoneid"] = f.ZoneId
	return tags
}

func (f *FederatedSelfZone) GetTags() map[string]string {
	tags := make(map[string]string)
	tags["org"] = f.SelfOperatorId
	tags["federationname"] = f.FederationName
	tags["zoneid"] = f.ZoneId
	return tags
}

func (f *FederatedPartnerZone) GetTags() map[string]string {
	tags := make(map[string]string)
	tags["org"] = f.SelfOperatorId
	tags["federatedorg"] = f.OperatorId
	tags["federationname"] = f.FederationName
	tags["zoneid"] = f.ZoneId
	return tags
}

func (f *FederatedZoneRegRequest) GetTags() map[string]string {
	tags := make(map[string]string)
	tags["org"] = f.SelfOperatorId
	tags["federationname"] = f.FederationName
	return tags
}
