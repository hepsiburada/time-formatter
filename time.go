package time_formatter

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

var DefaultFormatter = New(EN)

func New(locale LocaleType) IFormatter {
	return &Formatter{currentLocale: locale}
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

var formatFuncsMap = map[FormatType]func(to ToOpts) string{
	TIME_D: func(to ToOpts) string {
		return fmt.Sprintf("%d", to.time.Day())
	},
	TIME_DD: func(to ToOpts) string {
		return fmt.Sprintf("%02d", to.time.Day())
	},
	TIME_DDD: func(to ToOpts) string {
		return fmt.Sprintf("%d", to.time.YearDay())
	},
	TIME_DDDD: func(to ToOpts) string {
		return languageDaysMap[to.locale][to.time.Weekday()][:3]
	},
	TIME_DDDDD: func(to ToOpts) string {
		return languageDaysMap[to.locale][to.time.Weekday()]
	},
	TIME_M: func(to ToOpts) string {
		return fmt.Sprintf("%d", to.time.Month())
	},
	TIME_MM: func(to ToOpts) string {
		return fmt.Sprintf("%02d", to.time.Month())
	},
	TIME_MMM: func(to ToOpts) string {
		return languageMonthsMap[to.locale][to.time.Month()-1][:3]
	},
	TIME_MMMM: func(to ToOpts) string {
		return languageMonthsMap[to.locale][to.time.Month()-1]
	},
	TIME_YY: func(to ToOpts) string {
		return fmt.Sprintf("%d", to.time.Year())[2:]
	},
	TIME_YYYY: func(to ToOpts) string {
		return fmt.Sprintf("%d", to.time.Year())
	},
	TIME_Q: func(to ToOpts) string {
		return fmt.Sprintf("%d", (to.time.Month()/4)+1)
	},
	TIME_A: func(to ToOpts) string {
		if to.time.Hour() >= 12 {
			return "PM"
		} else {
			return "AM"
		}
	},
	TIME_a: func(to ToOpts) string {
		if to.time.Hour() >= 12 {
			return "pm"
		} else {
			return "am"
		}
	},
	TIME_H: func(to ToOpts) string {
		return fmt.Sprintf("%d", to.time.Hour())
	},
	TIME_HH: func(to ToOpts) string {
		return fmt.Sprintf("%02d", to.time.Hour())
	},
	TIME_h: func(to ToOpts) string {
		if to.time.Hour() > 12 {
			return fmt.Sprintf("%d", to.time.Hour()-12)
		} else {
			return fmt.Sprintf("%d", to.time.Hour())
		}
	},
	TIME_hh: func(to ToOpts) string {
		if to.time.Hour() > 12 {
			return fmt.Sprintf("%02d", to.time.Hour()-12)
		} else {
			return fmt.Sprintf("%02d", to.time.Hour())
		}
	},
	TIME_m: func(to ToOpts) string {
		return fmt.Sprintf("%d", to.time.Minute())
	},
	TIME_mm: func(to ToOpts) string {
		return fmt.Sprintf("%02d", to.time.Minute())
	},
	TIME_s: func(to ToOpts) string {
		return fmt.Sprintf("%d", to.time.Second())
	},
	TIME_ss: func(to ToOpts) string {
		return fmt.Sprintf("%02d", to.time.Second())
	},
	TIME_Z: func(to ToOpts) string {
		name, _ := to.time.Zone()
		return fmt.Sprintf("%s:00", name)
	},
	TIME_ZZ: func(to ToOpts) string {
		name, _ := to.time.Zone()
		return fmt.Sprintf("%s00", name)
	},
	TIME_X: func(to ToOpts) string {
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
				layout = strings.ReplaceAll(layout, field, operationFunc(ToOpts{time: t, locale: f.currentLocale}))
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
