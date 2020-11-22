# TimeFormatter

GoLang date time format - Helpful to convert normal date/time format into GoLang date/time format.

## Installation

First thing is to get your TimeFormatter package into your machine.

```go

go get "github.com/hepsiburada/time-formatter"

```

## Usage

```go

import (
	"fmt"
	tf "github.com/hepsiburada/time-formatter"
	"time"
)

func main() {
	formatter := tf.New()
	fmt.Println(formatter.To(time.Now(), fmt.Sprintf("In stock on %s %s!", tf.DD, tf.MMMM)))

	// or
	err := formatter.AddOpts(tf.LocaleTypeOptions{
		LocaleType:  "FR",
		DayValues:   []string{"Dimanche", "Lundi", "Mardi", "Mercredi", "Jeudi", "Vendredi", "Samedi"},
		MonthValues: []string{"Janvier", "Février", "Mars", "Avril", "Mai", "Juin", "Juillet", "Aout", "Septembre", "Octobre", "Novembre", "Décembre"},
	})
	if err != nil {
		panic(err)
	}

	formatter.ChangeLocale("FR")
	fmt.Println(formatter.To(time.Now(), fmt.Sprintf("En stock le %s %s!", tf.DD, tf.MMMM)))
}

```

## Importing packages

Import all necessary packages.("fmt" - Print, "time" - Getting time from machine, "testify" - A toolkit with common assertions) 

### Constants

|                | Token | Output                                 |
|----------------|-------|----------------------------------------|
| Month          | TIME_M     | 1 2 ... 11 12                          |
|                | TIME_MM    | 01 01 ... 11 12                        |
|                | TIME_MMM   | Jan Feb ... Nov Dec                    |
|                | TIME_MMMM  | January February ... November December |
| Day of Month   | TIME_D     | 1 2 ... 30 31                          |
|                | TIME_DD    | 01 02 ... 30 31                        |
|                | TIME_DDDD  | Mon, Tue ... Sat Sun                   |
|                | TIME_DDDDD | Monday, Tuesday ... Saturday Sunday    |
| Day of Year    | TIME_DDD   | 1 2 ... 364 365                        |
| Year           | TIME_YY    | 70 71 ... 29 30                        |
|                | TIME_YYYY  | 1970 1971 ... 2029 2030                |
| Quarter        | TIME_Q     | 1 2 3 4                                |
| AM/PM          | TIME_A     | AM PM                                  |
|                | TIME_a     | am pm                                  |
| Hour           | TIME_H     | 0 1 ... 22 23                          |
|                | TIME_HH    | 00 01 ... 22 23                        |
|                | TIME_h     | 1 2 ... 11 12                          |
|                | TIME_hh    | 01 02 ... 11 12                        |
| Minute         | TIME_m     | 0 1 ... 58 59                          |
|                | TIME_mm    | 00 01 ... 58 59                        |
| Second         | TIME_s     | 0 1 ... 58 59                          |
|                | TIME_ss    | 00 01 ... 58 59                        |
| Time Zone      | TIME_Z     | -07:00 -06:00 ... +06:00 +07:00        |
|                | TIME_ZZ    | -0700 -0600 ... +0600 +0700            |
| Unix Timestamp | TIME_X     | 1360013296                             |
