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

package common

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	valid "github.com/asaskevich/govalidator"
)

func ParseGeoLocation(geoLoc string) (float64, float64, error) {
	var lat float64
	var long float64
	if geoLoc == "" {
		return lat, long, fmt.Errorf("Missing geo location")
	}
	loc := strings.Split(geoLoc, ",")
	if len(loc) != 2 {
		return lat, long, fmt.Errorf("Invalid geo location %q. Valid format: <LatInDecimal,LongInDecimal>", geoLoc)
	}
	latStr, longStr := strings.TrimSpace(loc[0]), strings.TrimSpace(loc[1])
	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		return lat, long, fmt.Errorf("Invalid latitude %q, must be a valid decimal number", latStr)
	}
	if !valid.IsLatitude(latStr) {
		return lat, long, fmt.Errorf("Invalid latitude: %s", latStr)
	}

	long, err = strconv.ParseFloat(longStr, 64)
	if err != nil {
		return lat, long, fmt.Errorf("Invalid longitude %q, must be a valid decimal number", longStr)
	}
	if !valid.IsLongitude(longStr) {
		return lat, long, fmt.Errorf("Invalid longitude: %s", longStr)
	}

	return lat, long, nil
}

var nameMatch = regexp.MustCompile("^[a-zA-Z0-9][a-zA-Z0-9_-]*[a-zA-Z0-9]$")

func ValidateZoneId(zoneId string) error {
	if !nameMatch.MatchString(zoneId) {
		return fmt.Errorf("Invalid zone ID %q, can only contain alphanumeric, -, _ characters", zoneId)
	}
	if strings.HasPrefix(zoneId, "_") {
		return fmt.Errorf("Invalid zone ID %q, cannot start with _", zoneId)
	}
	if strings.HasPrefix(zoneId, "-") {
		return fmt.Errorf("Invalid zone ID %q, cannot start with -", zoneId)
	}
	if len(zoneId) < 8 || len(zoneId) > 32 {
		return fmt.Errorf("Invalid zone ID %q, valid length is 8 to 32 characters", zoneId)
	}
	return nil
}

func ValidateFederationName(name string) error {
	if !nameMatch.MatchString(name) {
		return fmt.Errorf("Invalid federation name %q, valid format is %s", name, nameMatch)
	}
	return nil
}

func ValidateCountryCode(countryCode string) error {
	if countryCode == "" {
		return fmt.Errorf("Missing country code")
	}
	if valid.IsISO3166Alpha2(countryCode) {
		return nil
	}
	return fmt.Errorf("Invalid country code %q. It must be a valid ISO 3166-1 Alpha-2 code for the country", countryCode)
}

func ValidateFederationId(fedId string) error {
	if !nameMatch.MatchString(fedId) {
		return fmt.Errorf("Invalid federation ID %q, can only contain alphanumeric, -, _ characters", fedId)
	}
	if strings.HasPrefix(fedId, "_") {
		return fmt.Errorf("Invalid federation ID %q, cannot start with _", fedId)
	}
	if strings.HasPrefix(fedId, "-") {
		return fmt.Errorf("Invalid federation ID %q, cannot start with -", fedId)
	}
	if len(fedId) < 8 || len(fedId) > 128 {
		return fmt.Errorf("Invalid federation ID %q, valid length is 8 to 128 characters", fedId)
	}
	return nil
}

func ValidateApiKey(apiKey string) error {
	if len(apiKey) < 8 || len(apiKey) > 128 {
		return fmt.Errorf("Invalid API key %q, valid length is 8 to 128 characters", apiKey)
	}
	return nil
}
