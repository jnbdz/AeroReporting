package navcanada

// https://plan.navcanada.ca/weather/api/alpha/
const (
	Tls     = true
	Url     = "plan.navcanada.ca"
	UrlPath = "/weather/api/alpha"
)

// Site
func Site(site []string) {
	// ?site=CYJN
	// &site=CYHU
	// &site=CYUL
}

// Display
/*func Display(what []string) {
	var options [8]string
	options[0] = "sigmet"
	options[1] = "airmet"
	options[2] = "notam"
	options[3] = "metar"
	options[4] = "taf"
	options[5] = "pirep"
	options[6] = "upperwind"
	options[7] = "vfr_route"

}*/

// Metar_hours
func Metar_hours() {

}

//
//  &image=SURFACE_ANALYSIS
//  &image=UPPER_ANALYSIS//250HPA
//  &image=UPPER_ANALYSIS//500HPA/THKNS&image=UPPER_ANALYSIS//500HPA/VORT&image=UPPER_ANALYSIS//700HPA
//  &image=UPPER_ANALYSIS//850HPA

//  &image=COMPOSITE_RADAR/ECHOTOP/NAT
//  &image=COMPOSITE_RADAR/LOW_RAIN~CAPPI/NAT
//  &image=COMPOSITE_RADAR/LOW_SNOW~CAPPI/NAT
//  &image=COMPOSITE_RADAR/ECHOTOP/PAC
//  &image=COMPOSITE_RADAR/ECHOTOP/WRN
//  &image=COMPOSITE_RADAR/ECHOTOP/ONT
//  &image=COMPOSITE_RADAR/ECHOTOP/QUE
//  &image=COMPOSITE_RADAR/ECHOTOP/ERN
//  &image=COMPOSITE_RADAR/LOW_RAIN~CAPPI/PAC
//  &image=COMPOSITE_RADAR/LOW_RAIN~CAPPI/WRN
//  &image=COMPOSITE_RADAR/LOW_RAIN~CAPPI/ONT
//  &image=COMPOSITE_RADAR/LOW_RAIN~CAPPI/QUE
//  &image=COMPOSITE_RADAR/LOW_RAIN~CAPPI/ERN
//  &image=COMPOSITE_RADAR/LOW_SNOW~CAPPI/PAC
//  &image=COMPOSITE_RADAR/LOW_SNOW~CAPPI/WRN
//  &image=COMPOSITE_RADAR/LOW_SNOW~CAPPI/ONT
//  &image=COMPOSITE_RADAR/LOW_SNOW~CAPPI/QUE
//  &image=COMPOSITE_RADAR/LOW_SNOW~CAPPI/ERN

//  &image=RADAR/ECHOTOP
//  &image=RADAR/CAPPI_RAIN
//  &image=RADAR/CAPPI_SNOW

//  &image=SATELLITE/IR
//  &image=SATELLITE/VIS
//  &image=SATELLITE/3u/*/*

//  &image=GFA/CLDWX
//  &image=GFA/TURBC
//  &image=LGF

//  &image=SIG_WX//HIGH_LEVEL/*
//  &image=SIG_WX//MID_LEVEL/*
//  &image=SIG_WX/DEPICTION/SURFACE/*

//  &image=TURBULENCE

//  &image=LOW_LEVEL_WIND/FL030
//  &image=LOW_LEVEL_WIND/FL060
//  &image=LOW_LEVEL_WIND/FL090
//  &image=LOW_LEVEL_WIND/FL120
//  &image=LOW_LEVEL_WIND/FL180
//  &image=HIGH_LEVEL_WIND/FL_240
//  &image=HIGH_LEVEL_WIND/FL_340
//  &image=HIGH_LEVEL_WIND/FL390
//  &image=HIGH_LEVEL_WIND/FL450

// &radius=10
// &notam_choice=default
// &upperwind_choice=both
// &_=1711059916244
