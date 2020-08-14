package time

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type (
	LocaleType string

	Formatter struct {
		options LocaleTypeOptions
	}

	LocaleTypeOptions struct {
		LocaleType  LocaleType
		DayValues   []string
		MonthValues []string
	}
)

const (
	EN LocaleType = "en"
	TR LocaleType = "tr"
)

var currentLocale = TR

var DefaultFormatter = &Formatter{
	LocaleTypeOptions{
		LocaleType:  EN,
		DayValues:   []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
		MonthValues: []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	},
}

var languageDaysMap = map[LocaleType][]string{
	EN: {
		"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday",
	},
	TR: {
		"Pazar", "Pazartesi", "Salı", "Çarşamba", "Perşembe", "Cuma", "Cumartesi",
	},
}

var languageMonthsMap = map[LocaleType][]string{
	EN: {
		"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December",
	},
	TR: {
		"Ocak", "Şubat", "Mart", "Nisan", "Mayıs", "Haziran", "Temmuz", "Ağustos", "Eylül", "Ekim", "Kasım", "Aralık",
	},
}

const (
	M    = "$M$"    // 1 2 ... 11 12
	MM   = "$MM$"   // 01 01 ... 11 12
	MMM  = "$MMM$"  // Jan Feb ... Nov Dec
	MMMM = "$MMMM$" // January February ... November December

	D     = "$D$"     // 1 2 ... 30 31
	DD    = "$DD$"    // 01 02 ... 30 31
	DDD   = "$DDD$"   // 1 2 ... 364 365
	DDDD  = "$DDDD$"  // Mon, Tue ... Sat Sun
	DDDDD = "$DDDDD$" // Monday, Tuesday ... Saturday Sunday

	YY   = "$YY$"   // 70 71 ... 29 30
	YYYY = "$YYYY$" // 1970 1971 ... 2029 2030

	Q = "$Q$" // 1 2 3 4

	A = "$A$" // AM PM
	a = "$a$" // am pm

	H  = "$H$"  // 0 1 ... 22 23
	HH = "$HH$" // 00 01 ... 22 23
	h  = "$h$"  // 1 2 ... 11 12
	hh = "$hh$" // 01 02 ... 11 12

	m  = "$m$"  // 0 1 ... 58 59
	mm = "$mm$" // 00 01 ... 58 59

	s  = "$s$"  // 0 1 ... 58 59
	ss = "$ss$" // 00 01 ... 58 59

	Z  = "$Z$"  // -07:00 -06:00 ... +06:00 +07:00
	ZZ = "$ZZ$" // -0700 -0600 ... +0600 +0700

	X = "$X$" // 1360013296
)

var operationMap = map[string]func(t time.Time) string{
	D: func(t time.Time) string {
		return fmt.Sprintf("%d", t.Day())
	},
	DD: func(t time.Time) string {
		return fmt.Sprintf("%02d", t.Day())
	},
	DDD: func(t time.Time) string {
		return fmt.Sprintf("%d", t.YearDay())
	},
	DDDD: func(t time.Time) string {
		return languageDaysMap[currentLocale][t.Weekday()][:3]
	},
	DDDDD: func(t time.Time) string {
		return languageDaysMap[currentLocale][t.Weekday()]
	},
	M: func(t time.Time) string {
		return fmt.Sprintf("%d", t.Month())
	},
	MM: func(t time.Time) string {
		return fmt.Sprintf("%02d", t.Month())
	},
	MMM: func(t time.Time) string {
		return languageMonthsMap[currentLocale][t.Month()-1][:3]
	},
	MMMM: func(t time.Time) string {
		return languageMonthsMap[currentLocale][t.Month()-1]
	},
	YY: func(t time.Time) string {
		return fmt.Sprintf("%d", t.Year())[2:]
	},
	YYYY: func(t time.Time) string {
		return fmt.Sprintf("%d", t.Year())
	},
	Q: func(t time.Time) string {
		return fmt.Sprintf("%d", (t.Month()/4)+1)
	},
	A: func(t time.Time) string {
		if t.Hour() >= 12 {
			return "PM"
		} else {
			return "AM"
		}
	},
	a: func(t time.Time) string {
		if t.Hour() >= 12 {
			return "pm"
		} else {
			return "am"
		}
	},
	H: func(t time.Time) string {
		return fmt.Sprintf("%d", t.Hour())
	},
	HH: func(t time.Time) string {
		return fmt.Sprintf("%02d", t.Hour())
	},
	h: func(t time.Time) string {
		if t.Hour() > 12 {
			return fmt.Sprintf("%d", t.Hour()-12)
		} else {
			return fmt.Sprintf("%d", t.Hour())
		}
	},
	hh: func(t time.Time) string {
		if t.Hour() > 12 {
			return fmt.Sprintf("%02d", t.Hour()-12)
		} else {
			return fmt.Sprintf("%02d", t.Hour())
		}
	},
	m: func(t time.Time) string {
		return fmt.Sprintf("%d", t.Minute())
	},
	mm: func(t time.Time) string {
		return fmt.Sprintf("%02d", t.Minute())
	},
	s: func(t time.Time) string {
		return fmt.Sprintf("%d", t.Second())
	},
	ss: func(t time.Time) string {
		return fmt.Sprintf("%02d", t.Second())
	},
	Z: func(time time.Time) string {
		name, _ := time.Zone()
		return fmt.Sprintf("%s:00", name)
	},
	ZZ: func(time time.Time) string {
		name, _ := time.Zone()
		return fmt.Sprintf("%s00", name)
	},
	X: func(t time.Time) string {
		return fmt.Sprintf("%d", t.Unix())
	},
}

func (f *Formatter) ChangeLocale(localeType LocaleType) {
	currentLocale = localeType
}

func (f *Formatter) To(t time.Time, layout string) string {
	tokenSplitPattern := regexp.MustCompile(`\$(.*?\S)\$`)

	fields := tokenSplitPattern.FindAllString(layout, -1)

	keys := make(map[string]bool)

	for _, field := range fields {
		if _, ok := keys[field]; !ok {
			if operationFunc, ok := operationMap[field]; ok {
				layout = strings.ReplaceAll(layout, field, operationFunc(t))
			}
			keys[field] = true
		}
	}
	return layout
}

func New(options LocaleTypeOptions) (*Formatter, error) {
	if options.DayValues == nil || options.MonthValues == nil {
		return nil, fmt.Errorf("%s or %s was null", "DayValues", "MonthValues")
	}

	currentLocale = options.LocaleType
	languageDaysMap[options.LocaleType] = options.DayValues
	languageMonthsMap[options.LocaleType] = options.MonthValues

	return &Formatter{options: options}, nil
}
