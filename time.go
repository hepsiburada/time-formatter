package time_formatter

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

func New() IFormatter {
	return &Formatter{currentLocale: EN}
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

var formatFuncsMap = map[FormatType]func(to ToOptions) string{
	D: func(to ToOptions) string {
		return fmt.Sprintf("%d", to.time.Day())
	},
	DD: func(to ToOptions) string {
		return fmt.Sprintf("%02d", to.time.Day())
	},
	DDD: func(to ToOptions) string {
		return fmt.Sprintf("%d", to.time.YearDay())
	},
	DDDD: func(to ToOptions) string {
		return languageDaysMap[to.locale][to.time.Weekday()][:3]
	},
	DDDDD: func(to ToOptions) string {
		return languageDaysMap[to.locale][to.time.Weekday()]
	},
	M: func(to ToOptions) string {
		return fmt.Sprintf("%d", to.time.Month())
	},
	MM: func(to ToOptions) string {
		return fmt.Sprintf("%02d", to.time.Month())
	},
	MMM: func(to ToOptions) string {
		return languageMonthsMap[to.locale][to.time.Month()-1][:3]
	},
	MMMM: func(to ToOptions) string {
		return languageMonthsMap[to.locale][to.time.Month()-1]
	},
	YY: func(to ToOptions) string {
		return fmt.Sprintf("%d", to.time.Year())[2:]
	},
	YYYY: func(to ToOptions) string {
		return fmt.Sprintf("%d", to.time.Year())
	},
	Q: func(to ToOptions) string {
		return fmt.Sprintf("%d", (to.time.Month()/4)+1)
	},
	A: func(to ToOptions) string {
		if to.time.Hour() >= 12 {
			return "PM"
		} else {
			return "AM"
		}
	},
	a: func(to ToOptions) string {
		if to.time.Hour() >= 12 {
			return "pm"
		} else {
			return "am"
		}
	},
	H: func(to ToOptions) string {
		return fmt.Sprintf("%d", to.time.Hour())
	},
	HH: func(to ToOptions) string {
		return fmt.Sprintf("%02d", to.time.Hour())
	},
	h: func(to ToOptions) string {
		if to.time.Hour() > 12 {
			return fmt.Sprintf("%d", to.time.Hour()-12)
		} else {
			return fmt.Sprintf("%d", to.time.Hour())
		}
	},
	hh: func(to ToOptions) string {
		if to.time.Hour() > 12 {
			return fmt.Sprintf("%02d", to.time.Hour()-12)
		} else {
			return fmt.Sprintf("%02d", to.time.Hour())
		}
	},
	m: func(to ToOptions) string {
		return fmt.Sprintf("%d", to.time.Minute())
	},
	mm: func(to ToOptions) string {
		return fmt.Sprintf("%02d", to.time.Minute())
	},
	s: func(to ToOptions) string {
		return fmt.Sprintf("%d", to.time.Second())
	},
	ss: func(to ToOptions) string {
		return fmt.Sprintf("%02d", to.time.Second())
	},
	Z: func(to ToOptions) string {
		name, _ := to.time.Zone()
		return fmt.Sprintf("%s:00", name)
	},
	ZZ: func(to ToOptions) string {
		name, _ := to.time.Zone()
		return fmt.Sprintf("%s00", name)
	},
	X: func(to ToOptions) string {
		return fmt.Sprintf("%d", to.time.Unix())
	},
}

func (f *Formatter) ChangeLocale(localeType LocaleType) {
	f.currentLocale = localeType
}

func (f *Formatter) To(t time.Time, layout string) string {
	tokenSplitPattern := regexp.MustCompile(`\$(.*?\S)\$`)
	fields := tokenSplitPattern.FindAllString(layout, -1)

	keys := make(map[string]bool)
	for _, field := range fields {
		if _, ok := keys[field]; !ok {
			if operationFunc, ok := formatFuncsMap[FormatType(field)]; ok {
				layout = strings.ReplaceAll(layout, field, operationFunc(ToOptions{time: t, locale: f.currentLocale}))
			}
			keys[field] = true
		}
	}
	return layout
}

func (f *Formatter) AddOpts(opts LocaleTypeOptions) error {
	if opts.LocaleType == "" {
		return errors.New("Locale type cannot be empty!")
	} else if len(opts.DayValues) == 0 || len(opts.MonthValues) == 0 {
		return errors.New("Day or Month values cannot be empty!")
	}

	languageDaysMap[opts.LocaleType] = opts.DayValues
	languageMonthsMap[opts.LocaleType] = opts.MonthValues
	return nil
}

func (f *Formatter) CurrentLocaleType() LocaleType {
	return f.currentLocale
}
