package time

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	t.Run("DayOrMonthValuesWasNil", func(t *testing.T) {
		formatter, err := New(LocaleTypeOptions{
			LocaleType:  TR,
			DayValues:   nil,
			MonthValues: nil,
		})

		assert.Nil(t, formatter)
		assert.NotNil(t, err)
		assert.Error(t, err, "DayValues or MonthValues was null")
	})

	t.Run("DayOrMonthValuesWasNotNil", func(t *testing.T) {
		formatter, err := New(LocaleTypeOptions{
			LocaleType:  "FR",
			DayValues:   []string{"Dimanche", "Lundi", "Mardi", "Mercredi", "Jeudi", "Vendredi", "Samedi"},
			MonthValues: []string{"Janvier", "Février", "Mars", "Avril", "Mai", "Juin", "Juillet", "Aout", "Septembre", "Octobre", "Novembre", "Décembre"},
		})

		assert.NotNil(t, formatter)
		assert.Nil(t, err)
	})
}

func TestFormatter_ChangeLocale(t *testing.T) {
	DefaultFormatter.ChangeLocale(EN)

	assert.Equal(t, EN, currentLocale)

	DefaultFormatter.ChangeLocale(TR)

	assert.Equal(t, TR, currentLocale)
}

func TestFormatter_To(t *testing.T) {
	t.Run("DefaultFormatterChangeLocaleWithSupportedLocale", func(t *testing.T) {
		DefaultFormatter.ChangeLocale(EN)

		timeNow := time.Now()

		assert.Equal(t, languageDaysMap[currentLocale][timeNow.Weekday()], DefaultFormatter.To(timeNow, fmt.Sprintf("%s", DDDDD)))
		assert.Equal(t, languageMonthsMap[currentLocale][timeNow.Month()-1], DefaultFormatter.To(timeNow, fmt.Sprintf("%s", MMMM)))
	})

	t.Run("CreateFormatterWithNewLocale", func(t *testing.T) {
		formatter, err := New(LocaleTypeOptions{
			LocaleType:  "FR",
			DayValues:   []string{"Dimanche", "Lundi", "Mardi", "Mercredi", "Jeudi", "Vendredi", "Samedi"},
			MonthValues: []string{"Janvier", "Février", "Mars", "Avril", "Mai", "Juin", "Juillet", "Aout", "Septembre", "Octobre", "Novembre", "Décembre"},
		})

		timeNow := time.Now()

		assert.Nil(t, err)
		assert.NotNil(t, formatter)
		assert.Equal(t, languageDaysMap[currentLocale][timeNow.Weekday()], formatter.To(timeNow, fmt.Sprintf("%s", DDDDD)))
		assert.Equal(t, languageMonthsMap[currentLocale][timeNow.Month()-1], formatter.To(timeNow, fmt.Sprintf("%s", MMMM)))
	})

	t.Run("DateFormatWithDefaultLocale", func(t *testing.T) {
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
					layout: fmt.Sprintf("%s", D),
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
					layout: fmt.Sprintf("%s", DD),
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
					layout: fmt.Sprintf("%s", DDD),
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
					layout: fmt.Sprintf("%s", DDDD),
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
					layout: fmt.Sprintf("%s", DDDDD),
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
					layout: fmt.Sprintf("%s", M),
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
					layout: fmt.Sprintf("%s", MM),
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
					layout: fmt.Sprintf("%s", MMM),
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
					layout: fmt.Sprintf("%s", MMMM),
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
					layout: fmt.Sprintf("%s", YY),
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
					layout: fmt.Sprintf("%s", YYYY),
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
					layout: fmt.Sprintf("%s", Q),
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
					layout: fmt.Sprintf("%s", A),
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
					layout: fmt.Sprintf("%s", A),
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
					layout: fmt.Sprintf("%s", a),
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
					layout: fmt.Sprintf("%s", a),
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
					layout: fmt.Sprintf("%s", H),
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
					layout: fmt.Sprintf("%s", HH),
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
					layout: fmt.Sprintf("%s", h),
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
					layout: fmt.Sprintf("%s", h),
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
					layout: fmt.Sprintf("%s", hh),
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
					layout: fmt.Sprintf("%s", hh),
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
					layout: fmt.Sprintf("%s", m),
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
					layout: fmt.Sprintf("%s", mm),
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
					layout: fmt.Sprintf("%s", s),
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
					layout: fmt.Sprintf("%s", ss),
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
					layout: fmt.Sprintf("%s", Z),
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
					layout: fmt.Sprintf("%s", ZZ),
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
					layout: fmt.Sprintf("%s", X),
				},
				want: func(timeNow time.Time) string {
					return fmt.Sprintf("%d", timeNow.Unix())
				},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				f := &Formatter{
					options: tt.fields.options,
				}
				if got := f.To(tt.args.t, tt.args.layout); got != tt.want(tt.args.t) {
					t.Errorf("To() = %v, want %v", got, tt.want(tt.args.t))
				}
			})
		}

	})

	t.Run("MultipleSameTokenInLayout", func(t *testing.T) {
		timeNow := time.Now()

		value := DefaultFormatter.To(timeNow, fmt.Sprintf("%s:%s:%s, %s", D, MMMM, YYYY, D))

		assert.Equal(t, fmt.Sprintf("%s:%s:%s, %s",
			fmt.Sprintf("%d", timeNow.Day()),
			languageMonthsMap[currentLocale][timeNow.Month()-1],
			fmt.Sprintf("%d", timeNow.Year()),
			fmt.Sprintf("%d", timeNow.Day()),
		), value)
	})
}
