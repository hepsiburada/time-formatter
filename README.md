# TimeFormatter

GoLang date time format - Helpful to convert normal date/time format into GoLang date/time format.

## Installation

First thing is to get your TimeFormatter package into your machine.

```go

go get "github.com/hepsiburada/time-formatter"

```

## Importing packages

Import all necessary packages.("fmt" - Print, "time" - Getting time from machine, "testify" - A toolkit with common assertions) 

### Constants

|                | Token | Output                                 |
|----------------|-------|----------------------------------------|
| Month          | M     | 1 2 ... 11 12                          |
|                | MM    | 01 01 ... 11 12                        |
|                | MMM   | Jan Feb ... Nov Dec                    |
|                | MMMM  | January February ... November December |
| Day of Month   | D     | 1 2 ... 30 31                          |
|                | DD    | 01 02 ... 30 31                        |
|                | DDDD  | Mon, Tue ... Sat Sun                   |
|                | DDDDD | Monday, Tuesday ... Saturday Sunday    |
| Day of Year    | DDD   | 1 2 ... 364 365                        |
| Year           | YY    | 70 71 ... 29 30                        |
|                | YYYY  | 1970 1971 ... 2029 2030                |
| Quarter        | Q     | 1 2 3 4                                |
| AM/PM          | A     | AM PM                                  |
|                | a     | am pm                                  |
| Hour           | H     | 0 1 ... 22 23                          |
|                | HH    | 00 01 ... 22 23                        |
|                | h     | 1 2 ... 11 12                          |
|                | hh    | 01 02 ... 11 12                        |
| Minute         | m     | 0 1 ... 58 59                          |
|                | mm    | 00 01 ... 58 59                        |
| Second         | s     | 0 1 ... 58 59                          |
|                | ss    | 00 01 ... 58 59                        |
| Time Zone      | Z     | -07:00 -06:00 ... +06:00 +07:00        |
|                | ZZ    | -0700 -0600 ... +0600 +0700            |
| Unix Timestamp | X     | 1360013296                             |

## Default Formatter

```go

DefaultFormatter.To(time.Now(), fmt.Sprintf("%s %s tarihinde stoklarda!", DD, MMMM))

```