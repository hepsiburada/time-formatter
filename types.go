package time_formatter

import "time"

type LocaleType string
const (
	EN LocaleType = "en"
	TR LocaleType = "tr"
)

type FormatType string
const (
	M    FormatType = "$M$"    // 1 2 ... 11 12
	MM   FormatType = "$MM$"   // 01 01 ... 11 12
	MMM  FormatType = "$MMM$"  // Jan Feb ... Nov Dec
	MMMM FormatType = "$MMMM$" // January February ... November December

	D     FormatType = "$D$"     // 1 2 ... 30 31
	DD    FormatType = "$DD$"    // 01 02 ... 30 31
	DDD   FormatType = "$DDD$"   // 1 2 ... 364 365
	DDDD  FormatType = "$DDDD$"  // Mon, Tue ... Sat Sun
	DDDDD FormatType = "$DDDDD$" // Monday, Tuesday ... Saturday Sunday

	YY   FormatType = "$YY$"   // 70 71 ... 29 30
	YYYY FormatType = "$YYYY$" // 1970 1971 ... 2029 2030

	Q FormatType = "$Q$" // 1 2 3 4

	A FormatType = "$A$" // AM PM
	a FormatType = "$a$" // am pm

	H  FormatType = "$H$"  // 0 1 ... 22 23
	HH FormatType = "$HH$" // 00 01 ... 22 23
	h  FormatType = "$h$"  // 1 2 ... 11 12
	hh FormatType = "$hh$" // 01 02 ... 11 12

	m  FormatType = "$m$"  // 0 1 ... 58 59
	mm FormatType = "$mm$" // 00 01 ... 58 59

	s  FormatType = "$s$"  // 0 1 ... 58 59
	ss FormatType = "$ss$" // 00 01 ... 58 59

	Z  FormatType = "$Z$"  // -07:00 -06:00 ... +06:00 +07:00
	ZZ FormatType = "$ZZ$" // -0700 -0600 ... +0600 +0700

	X FormatType = "$X$" // 1360013296
)

type Formatter struct {
	currentLocale LocaleType
}

type IFormatter interface {
	ChangeLocale(localeType LocaleType)
	To(t time.Time, layout string) string
	AddOpts(opts LocaleTypeOptions) error
	CurrentLocaleType() LocaleType
}

type LocaleTypeOptions struct {
	LocaleType  LocaleType
	DayValues   []string
	MonthValues []string
}

type ToOpts struct {
	time   time.Time
	locale LocaleType
}