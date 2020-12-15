package util

import (
	"strconv"
	"time"
)

// 判断日期
func FormatToStr(str string) (str2 string) {
	layout := "2006-01-02 15:04:05"
	t, _ := time.Parse(layout, str)
	nTime := time.Now()
	yesTime := nTime.AddDate(0, 0, -1)
	if nTime.Format("20060102") == t.Format("20060102") {
		return "今天 " + t.Format("15:04:05")
	} else if yesTime.Format("20060102") == t.Format("20060102") {
		return "昨天 " + t.Format("15:04:05")
	} else if nTime.Format("2006") == t.Format("2006") {
		return t.Format("01-02 15:04:05")
	} else {
		return t.Format("2006-01-02 15:04:05")
	}
}

/**
 * 获取日期的描述
 * @param  str  {string} 时间字符串
 * @return str2 {string} 日期的描述
 */
func FormatToString(str string) (str2 string) {
	layout := "2006-01-02 15:04:05"
	loc := time.Local
	t, _ := time.ParseInLocation(layout, str, loc)
	nTime := time.Now()
	yesTime := nTime.AddDate(0, 0, -1)
	if nTime.Format("20060102") == t.Format("20060102") {
		mins := nTime.Sub(t)
		min := mins.Minutes()
		if min < 1 {
			return "刚刚"
		} else if min < 60 {
			return strconv.FormatFloat(min, 'f', 0, 64) + "分钟前"
		}
		return "今天 " + t.Format("15:04:05")
	} else if yesTime.Format("20060102") == t.Format("20060102") {
		return "昨天 " + t.Format("15:04:05")
	} else if nTime.Format("2006") == t.Format("2006") {
		return t.Format("01-02 15:04:05")
	} else {
		return t.Format("2006-01-02 15:04:05")
	}
}

// 获取日期的年月日
func GetDateItem(t *time.Time) (year int, month int, day int) {
	year = t.Year()
	month = int(t.Month())
	day = t.Day()
	return
}

//获取时间的时分秒
func GetTimeItem(t *time.Time) (hour int, min int, sec int) {
	hour = t.Hour()
	min = t.Minute()
	sec = t.Second()
	return
}

func GetNowDateFormat() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}
func GetNowDayFormat() string {
	t := time.Now()
	return t.Format("2006-01-02")
}
func GetNowMonthFormat() string {
	t := time.Now()
	return t.Format("2006-01")
}

func GetAddMonthFormat(month int) string {
	t := time.Now()
	getTime := t.AddDate(0, month, 0)
	// todo 减一个月，api有bug，额外处理下
	_, _, day := t.Date()
	if day == 31 {
		getTime = t.AddDate(0, month, -1)
	}
	return getTime.Format("2006-01")
}

func GetNowYearFormat() string {
	t := time.Now()
	return t.Format("2006")
}

//获取当前时间戳，单位毫秒
func GetNowTimestap() int64 {
	return time.Now().UnixNano() / 1e6
}

//获取当前时间戳，单位秒
func GetNowTsBySeconds() int64 {
	return time.Now().Unix()
}

func GetNowTimestapByString(dataStr string, format string) int64 {
	t, _ := time.Parse(format, dataStr)
	return t.Unix()
}

func GetNowDateTimeFormatByFormat(format string) string {
	t := time.Now()
	return t.Format(format)
}

func GetNowDateTimeFormat() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}

func GetNowDateDayFormat() string {
	t := time.Now()
	return t.Format("2006-01-02")
}

func GetTimeFromDefaultString(str string) time.Time {
	layout := "2006-01-02"
	local, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation(layout, str, local)
	return t
}

func GetDayLatestTime() time.Time {
	layout := "2006-01-02 15:04:05"
	local, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation(layout, GetNowDateDayFormat()+" 23:59:59", local)
	return t
}

func GetNowDateTimeFormatCustom(format string) string {
	t := time.Now()
	return t.Format(format)
}

func GetTimeFormat(ts int64) string {
	t := time.Unix(ts, 0)
	return t.Format("2006-01-02 15:04:05")
}

func GetTimeWithLayout(ts int64, layout string) string {
	t := time.Unix(ts, 0)
	return t.Format(layout)
}

func GetTimeFormatCustom(ts int64, format string) string {
	t := time.Unix(ts, 0)
	return t.Format(format)
}

func GetDateFormat(t time.Time) string {
	return t.Format("2006-01-02")
}

//判断是否为闰年
func IsLeapYear(year int) bool { //y == 2000, 2004
	//判断是否为闰年
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	}

	return false
}

//获取某月有多少天
func GetMonthDays(year int, month int) int {
	days := 0
	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			days = 30
		} else {
			days = 31
		}
	} else {
		if IsLeapYear(year) {
			days = 29
		} else {
			days = 28
		}
	}
	return days
}

//时间字符串
func GetNowTime() string {
	t := time.Now()
	year := strconv.Itoa(t.Year())
	month := strconv.Itoa(int(t.Month()))
	day := strconv.Itoa(t.Day())
	hour := strconv.Itoa(t.Hour())
	min := strconv.Itoa(t.Minute())
	sec := strconv.Itoa(t.Second())

	if len(month) == 1 {
		month = "0" + month
	}

	if len(day) == 1 {
		day = "0" + day
	}
	return year + month + day + hour + min + sec
}

// 报表时间: 年月日
func GetTimeString() (y, m, d string) {
	t := time.Now()
	y = t.Format("2006")
	m = t.Format("2006-01")
	d = t.Format("2006-01-02")
	return
}

//时间字符串
func GetNowTime4Day() string {
	t := time.Now()
	day := t.Format("20060102")
	return day
}

// 报表事件入库参数
func GetEventTime(str string) (param1, param2, param3 string, res bool) {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, str)
	if err != nil {
		return "", "", "", false
	}
	day := t.Format("20060102")
	month := t.Format("200601")
	year := t.Format("2006")

	return year, month, day, true
}

// 获取报表查询时间参数
func GetStringTime4Day(str string) (param1 string, res bool) {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, str)
	if err != nil {
		return "", false
	}
	day := t.Format("20060102")
	return day, true
}

// 获取
func GetTimeArray(startTime, endTime string) []string {
	layout := "2006-01-02 15:04:05"
	start, err := time.Parse(layout, startTime)
	resArr := make([]string, 0)
	if err != nil {
		return resArr
	}
	str1 := start.Format("20060102")
	resArr = append(resArr, str1)
	end, err := time.Parse(layout, endTime)
	if err != nil {
		return resArr
	}
	timeBuild(start, end, &resArr)
	str2 := end.Format("20060102")

	if str2 != str1 {
		resArr = append(resArr, str2)
	}
	return resArr
}

func timeBuild(start time.Time, end time.Time, between *[]string) {
	d, err := time.ParseDuration("24h")
	if err != nil {
		return
	}
	nextDay := start.Add(d)
	if end.After(nextDay) {
		dayStr := nextDay.Format("20060102")
		*between = append(*between, dayStr)
		timeBuild(nextDay, end, between)
	}
}

// 获取某时刻多少天前的时间
func GetTime7Day(start string, num int) (string, string) {
	t := time.Now()
	if start != "" {
		layout := "2006-01-02 15:04:05"
		startTime, err := time.Parse(layout, start)
		if err != nil {
			t = startTime
		}
	} else {
		start = t.Format("2006-01-02 15:04:05")
	}

	allHour := num * 24
	allHourStr := "-" + strconv.Itoa(allHour) + "h"

	d, err := time.ParseDuration(allHourStr)
	if err != nil {
		return "", start
	}
	nextDay := t.Add(d)
	day := nextDay.Format("2006-01-02 15:04:05")
	return day, start
}

// 获取某时刻多少天前的时间
func GetTimeBeforeDay(start string, num int) (string, string) {
	t := time.Now()
	if start != "" {
		layout := "2006-01-02 15:04:05"
		startTime, err := time.Parse(layout, start)
		if err != nil {
			t = startTime
		}
	} else {
		start = t.Format("2006-01-02 15:04:05")
	}

	allHour := num * 24
	allHourStr := "-" + strconv.Itoa(allHour) + "h"

	d, err := time.ParseDuration(allHourStr)
	if err != nil {
		return "", start
	}
	nextDay := t.Add(d)
	day := nextDay.Format("2006-01-02")
	return day, start
}

//时间字符串
func GetNowYearMoth() []string {
	t := time.Now()
	day := t.Format("2006")
	resArr := make([]string, 0)
	fMonth, err := strconv.Atoi(day + "01")
	if err != nil {
		return resArr
	}
	for i := 0; i < 12; i++ {
		mStr := strconv.Itoa(fMonth + i)
		resArr = append(resArr, mStr)
	}
	return resArr
}

//时间运算
func GetNextTime(duration string) string {
	now := time.Now()
	s, _ := time.ParseDuration(duration)
	nowAfter15Second := now.Add(s)
	t := nowAfter15Second.Format("2006-01-02 15:04:05")
	return t
}

/**
 * @Description : 计算两个日期的差值
 * @Date        : 2019-04-12 12:22
 * @Modify      : return day hour minute second
 */
func GetDataSubValue(newData string) (day, hour, minute, second float64) {
	nowData := GetNowDateTimeFormat()
	if len(newData) == 10 {
		newData += " 00:00:00"
		nowData = GetNowDayFormat() + " 00:00:00"
	}

	layout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	newTime, _ := time.ParseInLocation(layout, newData, loc) //使用模板在对应时区转化为time.time类型
	nowTime, _ := time.ParseInLocation(layout, nowData, loc)

	dtn := newTime.Sub(nowTime)
	hour = dtn.Hours()
	minute = dtn.Minutes()
	second = dtn.Seconds()
	day = hour / 24
	return
}


func UtcToLocal(utcTime string) string {
	layout := "2006-01-02T15:04:05Z"
	utc, _ := time.LoadLocation("UTC")
	newTime, _ := time.ParseInLocation(layout, utcTime, utc)

	return newTime.Local().Format("2006-01-02 15:04:05")
}


func LocalToUtc(localTime string) string {
	layout := "2006-01-02 15:04:05"
	local, _ := time.LoadLocation("Local")
	newTime, _ := time.ParseInLocation(layout, localTime, local)

	return newTime.UTC().Format("2006-01-02T15:04:05Z")
}

//比较时间大小 true表示time2>time1
func CompareTimeString(time1,time2 string) bool {


	//先把时间字符串格式化成相同的时间类型
	t1, err := time.Parse("2006-01-02 15:04:05", time1)
	t2, err := time.Parse("2006-01-02 15:04:05", time2)
	if err == nil && t1.Before(t2) {
		//处理逻辑
		return true
	}
	return false
}
