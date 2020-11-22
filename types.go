package time_formatter

import "time"

type LocaleType string
const (
	EN LocaleType = "en"
	TR LocaleType = "tr"
)

type FormatType string
const (
	TIME_M    FormatType = "$M$"    // 1 2 ... 11 12
	TIME_MM   FormatType = "$MM$"   // 01 01 ... 11 12
	TIME_MMM  FormatType = "$MMM$"  // Jan Feb ... Nov Dec
	TIME_MMMM FormatType = "$MMMM$" // January February ... November December

	TIME_D     FormatType = "$D$"     // 1 2 ... 30 31
	TIME_DD    FormatType = "$DD$"    // 01 02 ... 30 31
	TIME_DDD   FormatType = "$DDD$"   // 1 2 ... 364 365
	TIME_DDDD  FormatType = "$DDDD$"  // Mon, Tue ... Sat Sun
	TIME_DDDDD FormatType = "$DDDDD$" // Monday, Tuesday ... Saturday Sunday

	TIME_YY   FormatType = "$YY$"   // 70 71 ... 29 30
	TIME_YYYY FormatType = "$YYYY$" // 1970 1971 ... 2029 2030

	TIME_Q FormatType = "$Q$" // 1 2 3 4

	TIME_A FormatType = "$A$" // AM PM
	TIME_a FormatType = "$a$" // am pm

	TIME_H  FormatType = "$H$"  // 0 1 ... 22 23
	TIME_HH FormatType = "$HH$" // 00 01 ... 22 23
	TIME_h  FormatType = "$h$"  // 1 2 ... 11 12
	TIME_hh FormatType = "$hh$" // 01 02 ... 11 12

	TIME_m  FormatType = "$m$"  // 0 1 ... 58 59
	TIME_mm FormatType = "$mm$" // 00 01 ... 58 59

	TIME_s  FormatType = "$s$"  // 0 1 ... 58 59
	TIME_ss FormatType = "$ss$" // 00 01 ... 58 59

	TIME_Z  FormatType = "$Z$"  // -07:00 -06:00 ... +06:00 +07:00
	TIME_ZZ FormatType = "$ZZ$" // -0700 -0600 ... +0600 +0700

	TIME_X FormatType = "$X$" // 1360013296
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