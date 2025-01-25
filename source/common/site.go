package common

import "regexp"

// ValidateSite checks if the input string is a valid ICAO code
func ValidateSite(site string) bool {
	// Example ICAO code pattern: 4 letters, can be more specific to actual ICAO standards
	match, _ := regexp.MatchString("^[A-Z]{4}$", site)
	return match
}

// Site function to handle site parameters
func Site(sites []string) {
	// Validate each site code
	for _, site := range sites {
		if !ValidateSite(site) {
			// Handle invalid site code
		}
	}
	// Proceed with valid site codes
}
