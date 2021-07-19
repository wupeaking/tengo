package stdlib

import (
	"time"

	"github.com/d5/tengo/v2/common"
)

var timesModule = map[string]common.Object{
	"format_ansic":        &common.String{Value: time.ANSIC},
	"format_unix_date":    &common.String{Value: time.UnixDate},
	"format_ruby_date":    &common.String{Value: time.RubyDate},
	"format_rfc822":       &common.String{Value: time.RFC822},
	"format_rfc822z":      &common.String{Value: time.RFC822Z},
	"format_rfc850":       &common.String{Value: time.RFC850},
	"format_rfc1123":      &common.String{Value: time.RFC1123},
	"format_rfc1123z":     &common.String{Value: time.RFC1123Z},
	"format_rfc3339":      &common.String{Value: time.RFC3339},
	"format_rfc3339_nano": &common.String{Value: time.RFC3339Nano},
	"format_kitchen":      &common.String{Value: time.Kitchen},
	"format_stamp":        &common.String{Value: time.Stamp},
	"format_stamp_milli":  &common.String{Value: time.StampMilli},
	"format_stamp_micro":  &common.String{Value: time.StampMicro},
	"format_stamp_nano":   &common.String{Value: time.StampNano},
	"nanosecond":          &common.Int{Value: int64(time.Nanosecond)},
	"microsecond":         &common.Int{Value: int64(time.Microsecond)},
	"millisecond":         &common.Int{Value: int64(time.Millisecond)},
	"second":              &common.Int{Value: int64(time.Second)},
	"minute":              &common.Int{Value: int64(time.Minute)},
	"hour":                &common.Int{Value: int64(time.Hour)},
	"january":             &common.Int{Value: int64(time.January)},
	"february":            &common.Int{Value: int64(time.February)},
	"march":               &common.Int{Value: int64(time.March)},
	"april":               &common.Int{Value: int64(time.April)},
	"may":                 &common.Int{Value: int64(time.May)},
	"june":                &common.Int{Value: int64(time.June)},
	"july":                &common.Int{Value: int64(time.July)},
	"august":              &common.Int{Value: int64(time.August)},
	"september":           &common.Int{Value: int64(time.September)},
	"october":             &common.Int{Value: int64(time.October)},
	"november":            &common.Int{Value: int64(time.November)},
	"december":            &common.Int{Value: int64(time.December)},
	"sleep": &common.UserFunction{
		Name:  "sleep",
		Value: timesSleep,
	}, // sleep(int)
	"parse_duration": &common.UserFunction{
		Name:  "parse_duration",
		Value: timesParseDuration,
	}, // parse_duration(str) => int
	"since": &common.UserFunction{
		Name:  "since",
		Value: timesSince,
	}, // since(time) => int
	"until": &common.UserFunction{
		Name:  "until",
		Value: timesUntil,
	}, // until(time) => int
	"duration_hours": &common.UserFunction{
		Name:  "duration_hours",
		Value: timesDurationHours,
	}, // duration_hours(int) => float
	"duration_minutes": &common.UserFunction{
		Name:  "duration_minutes",
		Value: timesDurationMinutes,
	}, // duration_minutes(int) => float
	"duration_nanoseconds": &common.UserFunction{
		Name:  "duration_nanoseconds",
		Value: timesDurationNanoseconds,
	}, // duration_nanoseconds(int) => int
	"duration_seconds": &common.UserFunction{
		Name:  "duration_seconds",
		Value: timesDurationSeconds,
	}, // duration_seconds(int) => float
	"duration_string": &common.UserFunction{
		Name:  "duration_string",
		Value: timesDurationString,
	}, // duration_string(int) => string
	"month_string": &common.UserFunction{
		Name:  "month_string",
		Value: timesMonthString,
	}, // month_string(int) => string
	"date": &common.UserFunction{
		Name:  "date",
		Value: timesDate,
	}, // date(year, month, day, hour, min, sec, nsec) => time
	"now": &common.UserFunction{
		Name:  "now",
		Value: timesNow,
	}, // now() => time
	"parse": &common.UserFunction{
		Name:  "parse",
		Value: timesParse,
	}, // parse(format, str) => time
	"unix": &common.UserFunction{
		Name:  "unix",
		Value: timesUnix,
	}, // unix(sec, nsec) => time
	"add": &common.UserFunction{
		Name:  "add",
		Value: timesAdd,
	}, // add(time, int) => time
	"add_date": &common.UserFunction{
		Name:  "add_date",
		Value: timesAddDate,
	}, // add_date(time, years, months, days) => time
	"sub": &common.UserFunction{
		Name:  "sub",
		Value: timesSub,
	}, // sub(t time, u time) => int
	"after": &common.UserFunction{
		Name:  "after",
		Value: timesAfter,
	}, // after(t time, u time) => bool
	"before": &common.UserFunction{
		Name:  "before",
		Value: timesBefore,
	}, // before(t time, u time) => bool
	"time_year": &common.UserFunction{
		Name:  "time_year",
		Value: timesTimeYear,
	}, // time_year(time) => int
	"time_month": &common.UserFunction{
		Name:  "time_month",
		Value: timesTimeMonth,
	}, // time_month(time) => int
	"time_day": &common.UserFunction{
		Name:  "time_day",
		Value: timesTimeDay,
	}, // time_day(time) => int
	"time_weekday": &common.UserFunction{
		Name:  "time_weekday",
		Value: timesTimeWeekday,
	}, // time_weekday(time) => int
	"time_hour": &common.UserFunction{
		Name:  "time_hour",
		Value: timesTimeHour,
	}, // time_hour(time) => int
	"time_minute": &common.UserFunction{
		Name:  "time_minute",
		Value: timesTimeMinute,
	}, // time_minute(time) => int
	"time_second": &common.UserFunction{
		Name:  "time_second",
		Value: timesTimeSecond,
	}, // time_second(time) => int
	"time_nanosecond": &common.UserFunction{
		Name:  "time_nanosecond",
		Value: timesTimeNanosecond,
	}, // time_nanosecond(time) => int
	"time_unix": &common.UserFunction{
		Name:  "time_unix",
		Value: timesTimeUnix,
	}, // time_unix(time) => int
	"time_unix_nano": &common.UserFunction{
		Name:  "time_unix_nano",
		Value: timesTimeUnixNano,
	}, // time_unix_nano(time) => int
	"time_format": &common.UserFunction{
		Name:  "time_format",
		Value: timesTimeFormat,
	}, // time_format(time, format) => string
	"time_location": &common.UserFunction{
		Name:  "time_location",
		Value: timesTimeLocation,
	}, // time_location(time) => string
	"time_string": &common.UserFunction{
		Name:  "time_string",
		Value: timesTimeString,
	}, // time_string(time) => string
	"is_zero": &common.UserFunction{
		Name:  "is_zero",
		Value: timesIsZero,
	}, // is_zero(time) => bool
	"to_local": &common.UserFunction{
		Name:  "to_local",
		Value: timesToLocal,
	}, // to_local(time) => time
	"to_utc": &common.UserFunction{
		Name:  "to_utc",
		Value: timesToUTC,
	}, // to_utc(time) => time
}

func timesSleep(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 1 {
		err = common.ErrWrongNumArguments
		return
	}

	i1, ok := common.ToInt64(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	time.Sleep(time.Duration(i1))
	ret = common.UndefinedValue

	return
}

func timesParseDuration(args ...common.Object) (
	ret common.Object,
	err error,
) {
	if len(args) != 1 {
		err = common.ErrWrongNumArguments
		return
	}

	s1, ok := common.ToString(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	dur, err := time.ParseDuration(s1)
	if err != nil {
		ret = wrapError(err)
		return
	}

	ret = &common.Int{Value: int64(dur)}

	return
}

func timesSince(args ...common.Object) (
	ret common.Object,
	err error,
) {
	if len(args) != 1 {
		err = common.ErrWrongNumArguments
		return
	}

	t1, ok := common.ToTime(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &common.Int{Value: int64(time.Since(t1))}

	return
}

func timesUntil(args ...common.Object) (
	ret common.Object,
	err error,
) {
	if len(args) != 1 {
		err = common.ErrWrongNumArguments
		return
	}

	t1, ok := common.ToTime(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &common.Int{Value: int64(time.Until(t1))}

	return
}

func timesDurationHours(args ...common.Object) (
	ret common.Object,
	err error,
) {
	if len(args) != 1 {
		err = common.ErrWrongNumArguments
		return
	}

	i1, ok := common.ToInt64(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &common.Float{Value: time.Duration(i1).Hours()}

	return
}

func timesDurationMinutes(args ...common.Object) (
	ret common.Object,
	err error,
) {
	if len(args) != 1 {
		err = common.ErrWrongNumArguments
		return
	}

	i1, ok := common.ToInt64(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &common.Float{Value: time.Duration(i1).Minutes()}

	return
}

func timesDurationNanoseconds(args ...common.Object) (
	ret common.Object,
	err error,
) {
	if len(args) != 1 {
		err = common.ErrWrongNumArguments
		return
	}

	i1, ok := common.ToInt64(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &common.Int{Value: time.Duration(i1).Nanoseconds()}

	return
}

func timesDurationSeconds(args ...common.Object) (
	ret common.Object,
	err error,
) {
	if len(args) != 1 {
		err = common.ErrWrongNumArguments
		return
	}

	i1, ok := common.ToInt64(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &common.Float{Value: time.Duration(i1).Seconds()}

	return
}

func timesDurationString(args ...common.Object) (
	ret common.Object,
	err error,
) {
	if len(args) != 1 {
		err = common.ErrWrongNumArguments
		return
	}

	i1, ok := common.ToInt64(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &common.String{Value: time.Duration(i1).String()}

	return
}

func timesMonthString(args ...common.Object) (
	ret common.Object,
	err error,
) {
	if len(args) != 1 {
		err = common.ErrWrongNumArguments
		return
	}

	i1, ok := common.ToInt64(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &common.String{Value: time.Month(i1).String()}

	return
}

func timesDate(args ...common.Object) (
	ret common.Object,
	err error,
) {
	if len(args) != 7 {
		err = common.ErrWrongNumArguments
		return
	}

	i1, ok := common.ToInt(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}
	i2, ok := common.ToInt(args[1])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}
	i3, ok := common.ToInt(args[2])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}
	i4, ok := common.ToInt(args[3])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "int(compatible)",
			Found:    args[3].TypeName(),
		}
		return
	}
	i5, ok := common.ToInt(args[4])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "fifth",
			Expected: "int(compatible)",
			Found:    args[4].TypeName(),
		}
		return
	}
	i6, ok := common.ToInt(args[5])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "sixth",
			Expected: "int(compatible)",
			Found:    args[5].TypeName(),
		}
		return
	}
	i7, ok := common.ToInt(args[6])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "seventh",
			Expected: "int(compatible)",
			Found:    args[6].TypeName(),
		}
		return
	}

	ret = &common.Time{
		Value: time.Date(i1,
			time.Month(i2), i3, i4, i5, i6, i7, time.Now().Location()),
	}

	return
}

func timesNow(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 0 {
		err = common.ErrWrongNumArguments
		return
	}

	ret = &common.Time{Value: time.Now()}

	return
}

func timesParse(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 2 {
		err = common.ErrWrongNumArguments
		return
	}

	s1, ok := common.ToString(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := common.ToString(args[1])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	parsed, err := time.Parse(s1, s2)
	if err != nil {
		ret = wrapError(err)
		return
	}

	ret = &common.Time{Value: parsed}

	return
}

func timesUnix(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 2 {
		err = common.ErrWrongNumArguments
		return
	}

	i1, ok := common.ToInt64(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := common.ToInt64(args[1])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &common.Time{Value: time.Unix(i1, i2)}

	return
}

func timesAdd(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 2 {
		err = common.ErrWrongNumArguments
		return
	}

	t1, ok := common.ToTime(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := common.ToInt64(args[1])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &common.Time{Value: t1.Add(time.Duration(i2))}

	return
}

func timesSub(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 2 {
		err = common.ErrWrongNumArguments
		return
	}

	t1, ok := common.ToTime(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	t2, ok := common.ToTime(args[1])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "time(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &common.Int{Value: int64(t1.Sub(t2))}

	return
}

func timesAddDate(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 4 {
		err = common.ErrWrongNumArguments
		return
	}

	t1, ok := common.ToTime(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := common.ToInt(args[1])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	i3, ok := common.ToInt(args[2])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}

	i4, ok := common.ToInt(args[3])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "int(compatible)",
			Found:    args[3].TypeName(),
		}
		return
	}

	ret = &common.Time{Value: t1.AddDate(i2, i3, i4)}

	return
}

func timesAfter(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 2 {
		err = common.ErrWrongNumArguments
		return
	}

	t1, ok := common.ToTime(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	t2, ok := common.ToTime(args[1])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "time(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	if t1.After(t2) {
		ret = common.TrueValue
	} else {
		ret = common.FalseValue
	}

	return
}

func timesBefore(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 2 {
		err = common.ErrWrongNumArguments
		return
	}

	t1, ok := common.ToTime(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	t2, ok := common.ToTime(args[1])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	if t1.Before(t2) {
		ret = common.TrueValue
	} else {
		ret = common.FalseValue
	}

	return
}

func timesTimeYear(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 1 {
		err = common.ErrWrongNumArguments
		return
	}

	t1, ok := common.ToTime(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &common.Int{Value: int64(t1.Year())}

	return
}

func timesTimeMonth(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 1 {
		err = common.ErrWrongNumArguments
		return
	}

	t1, ok := common.ToTime(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &common.Int{Value: int64(t1.Month())}

	return
}

func timesTimeDay(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 1 {
		err = common.ErrWrongNumArguments
		return
	}

	t1, ok := common.ToTime(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &common.Int{Value: int64(t1.Day())}

	return
}

func timesTimeWeekday(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 1 {
		err = common.ErrWrongNumArguments
		return
	}

	t1, ok := common.ToTime(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &common.Int{Value: int64(t1.Weekday())}

	return
}

func timesTimeHour(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 1 {
		err = common.ErrWrongNumArguments
		return
	}

	t1, ok := common.ToTime(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &common.Int{Value: int64(t1.Hour())}

	return
}

func timesTimeMinute(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 1 {
		err = common.ErrWrongNumArguments
		return
	}

	t1, ok := common.ToTime(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &common.Int{Value: int64(t1.Minute())}

	return
}

func timesTimeSecond(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 1 {
		err = common.ErrWrongNumArguments
		return
	}

	t1, ok := common.ToTime(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &common.Int{Value: int64(t1.Second())}

	return
}

func timesTimeNanosecond(args ...common.Object) (
	ret common.Object,
	err error,
) {
	if len(args) != 1 {
		err = common.ErrWrongNumArguments
		return
	}

	t1, ok := common.ToTime(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &common.Int{Value: int64(t1.Nanosecond())}

	return
}

func timesTimeUnix(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 1 {
		err = common.ErrWrongNumArguments
		return
	}

	t1, ok := common.ToTime(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &common.Int{Value: t1.Unix()}

	return
}

func timesTimeUnixNano(args ...common.Object) (
	ret common.Object,
	err error,
) {
	if len(args) != 1 {
		err = common.ErrWrongNumArguments
		return
	}

	t1, ok := common.ToTime(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &common.Int{Value: t1.UnixNano()}

	return
}

func timesTimeFormat(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 2 {
		err = common.ErrWrongNumArguments
		return
	}

	t1, ok := common.ToTime(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := common.ToString(args[1])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	s := t1.Format(s2)
	if len(s) > common.MaxStringLen {

		return nil, common.ErrStringLimit
	}

	ret = &common.String{Value: s}

	return
}

func timesIsZero(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 1 {
		err = common.ErrWrongNumArguments
		return
	}

	t1, ok := common.ToTime(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	if t1.IsZero() {
		ret = common.TrueValue
	} else {
		ret = common.FalseValue
	}

	return
}

func timesToLocal(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 1 {
		err = common.ErrWrongNumArguments
		return
	}

	t1, ok := common.ToTime(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &common.Time{Value: t1.Local()}

	return
}

func timesToUTC(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 1 {
		err = common.ErrWrongNumArguments
		return
	}

	t1, ok := common.ToTime(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &common.Time{Value: t1.UTC()}

	return
}

func timesTimeLocation(args ...common.Object) (
	ret common.Object,
	err error,
) {
	if len(args) != 1 {
		err = common.ErrWrongNumArguments
		return
	}

	t1, ok := common.ToTime(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &common.String{Value: t1.Location().String()}

	return
}

func timesTimeString(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 1 {
		err = common.ErrWrongNumArguments
		return
	}

	t1, ok := common.ToTime(args[0])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &common.String{Value: t1.String()}

	return
}
