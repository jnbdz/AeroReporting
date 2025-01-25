package navcanada

// Define the allowed options as a type and use iota for simplicity
type DisplayOption int

const (
	Sigmet DisplayOption = iota
	Airmet
	Notam
	Metar
	Taf
	Pirep
	Upperwind
	VfrRoute
)

// Map DisplayOption to its string representation
var displayOptions = map[DisplayOption]string{
	Sigmet:    "sigmet",
	Airmet:    "airmet",
	Notam:     "notam",
	Metar:     "metar",
	Taf:       "taf",
	Pirep:     "pirep",
	Upperwind: "upperwind",
	VfrRoute:  "vfr_route",
}

// ValidateDisplay checks if the input string matches one of the allowed display options
func ValidateDisplay(input string) bool {
	for _, option := range displayOptions {
		if input == option {
			return true
		}
	}
	return false
}

// Display function to handle what parameters
func Display(whats []string) {
	for _, what := range whats {
		if !ValidateDisplay(what) {
			// Handle invalid display option
		}
	}
	// Proceed with valid display options
}
