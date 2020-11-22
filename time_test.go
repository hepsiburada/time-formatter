package time_formatter

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAddOpts(t *testing.T) {
	t.Run("LocaleTypeValueWasNil", func(t *testing.T) {
		formatter := New()
		err := formatter.AddOpts(LocaleTypeOptions{
			LocaleType:  "",
			DayValues:   nil,
			MonthValues: nil,
		})

		assert.NotNil(t, err)
		assert.Error(t, err, "Locale type cannot be empty!")
	})

	t.Run("DayOrMonthValuesWasNil", func(t *testing.T) {
		formatter := New()
		err := formatter.AddOpts(LocaleTypeOptions{
			LocaleType:  "fb",
			DayValues:   nil,
			MonthValues: nil,
		})

		assert.NotNil(t, err)
		assert.Error(t, err, "Day or Month values cannot be empty!")
	})

	t.Run("DayOrMonthValuesWasNotNil", func(t *testing.T) {
		formatter := New()
		err := formatter.AddOpts(LocaleTypeOptions{
			LocaleType:  "fb",
			DayValues:   []string{"Dimanche", "Lundi", "Mardi", "Mercredi", "Jeudi", "Vendredi", "Samedi"},
			MonthValues: []string{"Janvier", "Février", "Mars", "Avril", "Mai", "Juin", "Juillet", "Aout", "Septembre", "Octobre", "Novembre", "Décembre"},
		})

		assert.NotNil(t, formatter)
		assert.Nil(t, err)
	})
}

func TestFormatter_ChangeLocale(t *testing.T) {
	formatter := New()
	assert.Equal(t, EN, formatter.CurrentLocaleType())

	formatter.ChangeLocale(TR)
	assert.Equal(t, TR, formatter.CurrentLocaleType())
}

func TestFormatter_To(t *testing.T) {
	t.Run("DefaultFormatterChangeLocaleWithSupportedLocale", func(t *testing.T) {
		formatter := New()

		timeNow := time.Now()
		currentLocale := formatter.CurrentLocaleType()

		assert.Equal(t, languageDaysMap[currentLocale][timeNow.Weekday()], formatter.To(timeNow, fmt.Sprintf("%s", TIME_DDDDD)))
		assert.Equal(t, languageMonthsMap[currentLocale][timeNow.Month()-1], formatter.To(timeNow, fmt.Sprintf("%s", TIME_MMMM)))
	})

	t.Run("CreateFormatterWithNewLocale", func(t *testing.T) {
		formatter := New()
		err := formatter.AddOpts(LocaleTypeOptions{
			LocaleType:  "FR",
			DayValues:   []string{"Dimanche", "Lundi", "Mardi", "Mercredi", "Jeudi", "Vendredi", "Samedi"},
			MonthValues: []string{"Janvier", "Février", "Mars", "Avril", "Mai", "Juin", "Juillet", "Aout", "Septembre", "Octobre", "Novembre", "Décembre"},
		})
		assert.Nil(t, err)

		timeNow := time.Now()
		currentLocale := formatter.CurrentLocaleType()

		assert.Nil(t, err)
		assert.NotNil(t, formatter)
		assert.Equal(t, languageDaysMap[currentLocale][timeNow.Weekday()], formatter.To(timeNow, fmt.Sprintf("%s", TIME_DDDDD)))
		assert.Equal(t, languageMonthsMap[currentLocale][timeNow.Month()-1], formatter.To(timeNow, fmt.Sprintf("%s", TIME_MMMM)))
	})

	t.Run("DateFormatWithDefaultLocale", func(t *testing.T) {
		formatter := New()
		currentLocale := formatter.CurrentLocaleType()

		type fields struct {
			options LocaleTypeOptions
		}
		type args struct {
			t      time.Time
			layout string
		}
		tests := []struct {
			name   string
			fields fields
			args   args
			want   func(timeNow time.Time) string
		}{
			{
				// 1 2 ... 30 31
				name:   "OnlyDayNumber",
				fields: fields{},
				args: args{
					t:      time.Now(),
					layout: fmt.Sprintf("%s", TIME_D),
				},
				want: func(timeNow time.Time) string {
					return fmt.Sprintf("%d", timeNow.Day())
				},
			},
			{
				// 01 02 ... 30 31
				name:   "OnlyDayNumberWithTwoLength",
				fields: fields{},
				args: args{
					t:      time.Now(),
					layout: fmt.Sprintf("%s", TIME_DD),
				},
				want: func(timeNow time.Time) string {
					return fmt.Sprintf("%02d", timeNow.Day())
				},
			},
			{
				// 1 2 ... 364 365
				name:   "OnlyDayNumberOfYear",
				fields: fields{},
				args: args{
					t:      time.Now(),
					layout: fmt.Sprintf("%s", TIME_DDD),
				},
				want: func(timeNow time.Time) string {
					return fmt.Sprintf("%d", timeNow.YearDay())
				},
			},
			{
				// Mon, Tue ... Sat Sun
				name:   "OnlyShortDayText",
				fields: fields{},
				args: args{
					t:      time.Now(),
					layout: fmt.Sprintf("%s", TIME_DDDD),
				},
				want: func(timeNow time.Time) string {
					return languageDaysMap[currentLocale][timeNow.Weekday()][:3]
				},
			},
			{
				// Mon, Tue ... Sat Sun
				name:   "OnlyLongDayText",
				fields: fields{},
				args: args{
					t:      time.Now(),
					layout: fmt.Sprintf("%s", TIME_DDDDD),
				},
				want: func(timeNow time.Time) string {
					return languageDaysMap[currentLocale][timeNow.Weekday()]
				},
			},
			{
				// 1 2 ... 11 12
				name:   "OnlyMonthNumber",
				fields: fields{},
				args: args{
					t:      time.Now(),
					layout: fmt.Sprintf("%s", TIME_M),
				},
				want: func(timeNow time.Time) string {
					return fmt.Sprintf("%d", timeNow.Month())
				},
			},
			{
				// 01 01 ... 11 12
				name:   "OnlyMonthNumberWithTwoLength",
				fields: fields{},
				args: args{
					t:      time.Now(),
					layout: fmt.Sprintf("%s", TIME_MM),
				},
				want: func(timeNow time.Time) string {
					return fmt.Sprintf("%02d", timeNow.Month())
				},
			},
			{
				// Mon, Tue ... Sat Sun
				name:   "OnlyShortMonthText",
				fields: fields{},
				args: args{
					t:      time.Now(),
					layout: fmt.Sprintf("%s", TIME_MMM),
				},
				want: func(timeNow time.Time) string {
					return languageMonthsMap[currentLocale][timeNow.Month()-1][:3]
				},
			},
			{
				// Mon, Tue ... Sat Sun
				name:   "OnlyLongMonthText",
				fields: fields{},
				args: args{
					t:      time.Now(),
					layout: fmt.Sprintf("%s", TIME_MMMM),
				},
				want: func(timeNow time.Time) string {
					return languageMonthsMap[currentLocale][timeNow.Month()-1]
				},
			},
			{
				// 70 71 ... 29 30
				name:   "OnlyShortYear",
				fields: fields{},
				args: args{
					t:      time.Now(),
					layout: fmt.Sprintf("%s", TIME_YY),
				},
				want: func(timeNow time.Time) string {
					return fmt.Sprintf("%d", timeNow.Year())[2:]
				},
			},
			{
				// 1970 1971 ... 2029 2030
				name:   "OnlyLongYear",
				fields: fields{},
				args: args{
					t:      time.Now(),
					layout: fmt.Sprintf("%s", TIME_YYYY),
				},
				want: func(timeNow time.Time) string {
					return fmt.Sprintf("%d", timeNow.Year())
				},
			},
			{
				// 1 2 3 4
				name:   "OnlyQuarter",
				fields: fields{},
				args: args{
					t:      time.Now(),
					layout: fmt.Sprintf("%s", TIME_Q),
				},
				want: func(timeNow time.Time) string {
					return fmt.Sprintf("%d", (timeNow.Month()/4)+1)
				},
			},
			{
				// AM PM
				name:   "Only AM/PM-1",
				fields: fields{},
				args: args{
					t:      time.Now(),
					layout: fmt.Sprintf("%s", TIME_A),
				},
				want: func(timeNow time.Time) string {
					if timeNow.Hour() >= 12 {
						return "PM"
					} else {
						return "AM"
					}
				},
			},
			{
				// AM PM
				name:   "Only AM/PM-2",
				fields: fields{},
				args: args{
					t:      time.Now().Add(12 * time.Hour),
					layout: fmt.Sprintf("%s", TIME_A),
				},
				want: func(timeNow time.Time) string {
					if timeNow.Hour() >= 12 {
						return "PM"
					} else {
						return "AM"
					}
				},
			},
			{
				// am pm
				name:   "Only am/pm-1",
				fields: fields{},
				args: args{
					t:      time.Now(),
					layout: fmt.Sprintf("%s", TIME_a),
				},
				want: func(timeNow time.Time) string {
					if timeNow.Hour() >= 12 {
						return "pm"
					} else {
						return "am"
					}
				},
			},
			{
				// am pm
				name:   "Only am/pm-2",
				fields: fields{},
				args: args{
					t:      time.Now().Add(12 * time.Hour),
					layout: fmt.Sprintf("%s", TIME_a),
				},
				want: func(timeNow time.Time) string {
					if timeNow.Hour() >= 12 {
						return "pm"
					} else {
						return "am"
					}
				},
			},
			{
				// 0 1 ... 22 23
				name:   "OnlyHourNumberWith24Hour",
				fields: fields{},
				args: args{
					t:      time.Now(),
					layout: fmt.Sprintf("%s", TIME_H),
				},
				want: func(timeNow time.Time) string {
					return fmt.Sprintf("%d", timeNow.Hour())
				},
			},
			{
				// 00 01 ... 22 23
				name:   "OnlyHourNumberTwoLengthWith24Hour",
				fields: fields{},
				args: args{
					t:      time.Now(),
					layout: fmt.Sprintf("%s", TIME_HH),
				},
				want: func(timeNow time.Time) string {
					return fmt.Sprintf("%02d", timeNow.Hour())
				},
			},
			{
				// 1 2 ... 11 12
				name:   "OnlyHourNumberWith12Hour-1",
				fields: fields{},
				args: args{
					t:      time.Now(),
					layout: fmt.Sprintf("%s", TIME_h),
				},
				want: func(timeNow time.Time) string {
					if timeNow.Hour() > 12 {
						return fmt.Sprintf("%d", timeNow.Hour()-12)
					} else {
						return fmt.Sprintf("%d", timeNow.Hour())
					}
				},
			},
			{
				// 1 2 ... 11 12
				name:   "OnlyHourNumberWith12Hour-2",
				fields: fields{},
				args: args{
					t:      time.Now().Add(12 * time.Hour),
					layout: fmt.Sprintf("%s", TIME_h),
				},
				want: func(timeNow time.Time) string {
					if timeNow.Hour() > 12 {
						return fmt.Sprintf("%d", timeNow.Hour()-12)
					} else {
						return fmt.Sprintf("%d", timeNow.Hour())
					}
				},
			},
			{
				// 01 02 ... 11 12
				name:   "OnlyHourNumberTwoLengthWith12Hour-1",
				fields: fields{},
				args: args{
					t:      time.Now(),
					layout: fmt.Sprintf("%s", TIME_hh),
				},
				want: func(timeNow time.Time) string {
					if timeNow.Hour() > 12 {
						return fmt.Sprintf("%02d", timeNow.Hour()-12)
					} else {
						return fmt.Sprintf("%02d", timeNow.Hour())
					}
				},
			},
			{
				// 01 02 ... 11 12
				name:   "OnlyHourNumberTwoLengthWith12Hour-2",
				fields: fields{},
				args: args{
					t:      time.Now().Add(12 * time.Hour),
					layout: fmt.Sprintf("%s", TIME_hh),
				},
				want: func(timeNow time.Time) string {
					if timeNow.Hour() > 12 {
						return fmt.Sprintf("%02d", timeNow.Hour()-12)
					} else {
						return fmt.Sprintf("%02d", timeNow.Hour())
					}
				},
			},
			{
				// 0 1 ... 58 59
				name:   "OnlyMinutesNumberTwoLengthWith12Hour",
				fields: fields{},
				args: args{
					t:      time.Now(),
					layout: fmt.Sprintf("%s", TIME_m),
				},
				want: func(timeNow time.Time) string {
					return fmt.Sprintf("%d", timeNow.Minute())
				},
			},
			{
				// 00 01 ... 58 59
				name:   "OnlyMinutesNumberTwoLength",
				fields: fields{},
				args: args{
					t:      time.Now(),
					layout: fmt.Sprintf("%s", TIME_mm),
				},
				want: func(timeNow time.Time) string {
					return fmt.Sprintf("%02d", timeNow.Minute())
				},
			},
			{
				// 0 1 ... 58 59
				name:   "OnlySecondNumberTwoLength",
				fields: fields{},
				args: args{
					t:      time.Now(),
					layout: fmt.Sprintf("%s", TIME_s),
				},
				want: func(timeNow time.Time) string {
					return fmt.Sprintf("%d", timeNow.Second())
				},
			},
			{
				// 00 01 ... 58 59
				name:   "OnlySecondNumberTwoLength",
				fields: fields{},
				args: args{
					t:      time.Now(),
					layout: fmt.Sprintf("%s", TIME_ss),
				},
				want: func(timeNow time.Time) string {
					return fmt.Sprintf("%02d", timeNow.Second())
				},
			},
			{
				// -07:00 -06:00 ... +06:00 +07:00
				name:   "OnlyTimeZoneSeparateWithColon",
				fields: fields{},
				args: args{
					t:      time.Now(),
					layout: fmt.Sprintf("%s", TIME_Z),
				},
				want: func(timeNow time.Time) string {
					name, _ := timeNow.Zone()
					return fmt.Sprintf("%s:00", name)
				},
			},
			{
				// -0700 -0600 ... +0600 +0700
				name:   "OnlyTimeZone",
				fields: fields{},
				args: args{
					t:      time.Now(),
					layout: fmt.Sprintf("%s", TIME_ZZ),
				},
				want: func(timeNow time.Time) string {
					name, _ := timeNow.Zone()
					return fmt.Sprintf("%s00", name)
				},
			},
			{
				// 1360013296
				name:   "OnlyUnixTimestamp",
				fields: fields{},
				args: args{
					t:      time.Now(),
					layout: fmt.Sprintf("%s", TIME_X),
				},
				want: func(timeNow time.Time) string {
					return fmt.Sprintf("%d", timeNow.Unix())
				},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f := New()
				if got := f.To(tt.args.t, tt.args.layout); got != tt.want(tt.args.t) {
					t.Errorf("To() = %v, want %v", got, tt.want(tt.args.t))
				}
			})
		}

	})

	t.Run("MultipleSameTokenInLayout", func(t *testing.T) {
		formatter := New()
		currentLocale := formatter.CurrentLocaleType()

		timeNow := time.Now()

		value := formatter.To(timeNow, fmt.Sprintf("%s:%s:%s, %s", TIME_D, TIME_MMMM, TIME_YYYY, TIME_D))

		assert.Equal(t, fmt.Sprintf("%s:%s:%s, %s",
			fmt.Sprintf("%d", timeNow.Day()),
			languageMonthsMap[currentLocale][timeNow.Month()-1],
			fmt.Sprintf("%d", timeNow.Year()),
			fmt.Sprintf("%d", timeNow.Day()),
		), value)
	})
}
