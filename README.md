# AeroReporting ⛈️🚨❗
Aero Reporting get weather reporting via the terminal. ⛈️🚨❗

## Install
```bash
sudo make install
```

## Uninstall
```bash
sudo make uninstall
```

## How-to-use
### Help!
```bash
aeroreporting --help
# or
aeroreporting -h
# or
aeroreporting
```

Result: 
```bash

```

### Get reports
There are multiple types of reporting you can get like [NOTAM](), [METAR](), [TAF](), etc.

For American pilots (it is also the default source):
```bash
aeroreporting --source=aviationweather
# or
aeroreporting
```

For Canadian pilots: 
```bash
aeroreporting --source=navcanada
```

> Note: Without indicating the `site` the only thing the command will be to check if the source (host) is up.

#### `site`
Takes a ICAO airport code: 
```bash
aeroreporting --site=CYUL, CYHU, CYYZ
```

There is a list of [ICAO airport code](https://en.wikipedia.org/wiki/ICAO_airport_code).

#### `metar-hours` - Historical Hours
```bash
aeroreporting --site CYUL --metar-hours
```

Options (only one can be chosen): 
- `0` - 0H (Current Data Only) (default)
- `1` - 1H
- `2` - 2H
- `3` - 3H
- `4` - 4H
- `5` - 5H
- `6` - 6H

#### `show-duplicates` - Show Duplicates
```bash
aeroreporting --site CYUL --show-duplicates
```
It is a boolean type, meaning if the flag is used then it is *true* that will make the duplicates display.

#### `display`
```bash
aeroreporting --site CYUL --display notam, taf, vfr_route
```

Options: 
- `sigmet` - SIGMET
- `airmet` - AIRMET
- `notam` - NOTAM
- `metar` - METAR
- `taf` - TAF
- `pirep` - PIREP
- `upperwind` - Upper Wind
- `vfr_route` - BC VFR Route Forecast

#### `analysis`
```bash
aeroreporting --site CYUL --analysis surface, 250
```

Options: 
- `surface` - Surface
- `250` - 250 hPa
- `500_thickness` - 500 hPa Thickness
- `500_vorticity` - 500 hPa Vorticity
- `700` - 700 hPa
- `850` - 850 hPa

#### `radar`

Options: 
- National: 
  - `echotop` - National ECHOTOP
  - `cappi_rain` - National CAPPI (RAIN)
  - `cappi_snow` - National CAPPI (SNOW)

```bash
aeroreporting --site CYUL --national-radar
```

- Regional: 
  - `echotop` - Regional ECHOTOP
  - `cappi_rain` - Regional CAPPI (RAIN)
  - `cappi_snow` - Regional CAPPI (SNOW)

```bash
aeroreporting --site CYUL --regional-radar
```

- Individual: 
  - `echotop` - Individual ECHOTOP
  - `cappi_rain` - Individual CAPPI (RAIN)
  - `cappi_snow` - Individual CAPPI (SNOW)

```bash
aeroreporting --site CYUL --individual-radar
```

#### `satellite`
```bash
aeroreporting --site CYUL --satellite ir, vis
```

Options: 
- `ir` - Infrared
- `vis` - Visible
- `3u` - Yukon and NWT 3u (only for **navcanada** source option)

#### `gfa` - Graphical Forecast
```bash
aeroreporting --site CYUL --gfa cldwx, turbc
```

Options: 
- `cldwx` - Clouds & Weather
- `turbc` - Icing, Turbulence & Freezing level
- `lgf` - Local (BC)

#### `significant-weather`
```bash
aeroreporting --site CYUL --significant-weather high, mid, surface_depiction
```

Options: 
- `high` - High Level
- `mid` - Mid Level
- `surface_depiction` - Surface Depiction

#### `turbulence`
```bash
aeroreporting --site CYUL --turbulence
```

It is a boolean type, meaning if the flag is used then it is *true* that the turbulence information needs to be displayed.

#### `upperwind`
```bash
aeroreporting --site CYUL --upperwind both
```

Options: 
- `both` (default)
- `high`
- `low`

#### `wind` 
```bash
aeroreporting --site CYUL --upperwind both --wind 3000, 6000, 12000
```

Options: 
- `3000`
- `6000`
- `9000`
- `12000`
- `FL180`
- `FL240`
- `FL340`
- `FL390`
- `FL450`

#### `report`
```bash
aeroreporting --report=notam, taf
```

Options: 
- `all` (default)
- `sigmet`
- `airmet`
- `notam`
- `metar`
- `taf`
- `pirep`
- `upperwind`
- `vfr_route`

#### `radius`
```bash
aeroreporting --radius 
```

Limit: `1` to `999`.

## What about my `X` country?!
Well just make a pull request 🤓!