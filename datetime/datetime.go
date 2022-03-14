package datetime

import (
    "strconv"
    "time"
)

const (
    ZeroDateTime = "1970-01-01 00:00:00"
)

var timeFormat = map[string]string{
    "yyyy-mm-dd hh:mm:ss": "2006-01-02 15:04:05",
    "yyyy-mm-dd hh:mm":    "2006-01-02 15:04",
    "yyyy-mm-dd hh":       "2006-01-02 15",
    "yyyy-mm-dd":          "2006-01-02",
    "yyyy-mm":             "2006-01",
    "mm-dd":               "01-02",
    "dd-mm-yy hh:mm:ss":   "02-01-06 15:04:05",
    "yyyy/mm/dd hh:mm:ss": "2006/01/02 15:04:05",
    "yyyy/mm/dd hh:mm":    "2006/01/02 15:04",
    "yyyy/mm/dd hh":       "2006/01/02 15",
    "yyyy/mm/dd":          "2006/01/02",
    "yyyy/mm":             "2006/01",
    "mm/dd":               "01/02",
    "dd/mm/yy hh:mm:ss":   "02/01/06 15:04:05",
    "yyyy":                "2006",
    "mm":                  "01",
    "hh:mm:ss":            "15:04:05",
    "mm:ss":               "04:05",
}

// NowUnix 秒时间戳
func NowUnix() int64 {
    return time.Now().Unix()
}

// FromUnix 秒时间戳转时间
func FromUnix(unix int64) time.Time {
    return time.Unix(unix, 0)
}

// NowTimestamp 当前毫秒时间戳
func NowTimestamp() int64 {
    return Timestamp(time.Now())
}

// Timestamp 毫秒时间戳
func Timestamp(t time.Time) int64 {
    return t.UnixNano() / 1e6
}

// FromTimestamp 毫秒时间戳转时间
func FromTimestamp(timestamp int64) time.Time {
    return time.Unix(0, timestamp*int64(time.Millisecond))
}

// FormatTimeToStr 时间格式化
func FormatTimeToStr(time time.Time, format string) string {
    layout := timeFormat[format]
    if len(layout) == 0 {
        layout = "2006-01-02 15:04:05"
    }
    return time.Format(layout)
}

// FormatStrToTime 字符串时间转时间类型
func FormatStrToTime(timeStr, format string) (time.Time, error) {
    layout := timeFormat[format]
    if len(layout) == 0 {
        layout = "2006-01-02 15:04:05"
    }
    return time.Parse(layout, timeStr)
}

// GetNowDate 获取现在的日期 return format yyyy-mm-dd of current date
func GetNowDate() string {
    return time.Now().Format("2006-01-02")
}

// GetNowTime 获取现在的时间 return format hh-mm-ss of current time
func GetNowTime() string {
    return time.Now().Format("15:04:05")
}

// WithTimeAsStartOfDay 返回指定时间当天的开始时间
func WithTimeAsStartOfDay(t time.Time) time.Time {
    return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// IsZero 1970-01-01 00:00:00
func IsZero(theTime time.Time) bool {
    return FormatTimeToStr(theTime, "2006-01-02 15:04:05") == ZeroDateTime
}

// PrettyTime 多用于消息或者邮件显示
//
//  将时间格式换成 xx秒前，xx分钟前...
//  规则：
//  59秒--->刚刚
//  1-59分钟--->x分钟前（23分钟前）
//  1-24小时--->x小时前（5小时前）
//  昨天--->昨天 hh:mm（昨天 16:15）
//  前天--->前天 hh:mm（前天 16:15）
//  前天以后--->mm-dd（2月18日）
func PrettyTime(milliseconds int64) string {
    t := FromTimestamp(milliseconds)
    duration := (NowTimestamp() - milliseconds) / 1000
    if duration < 60 {
        return "刚刚"
    } else if duration < 3600 {
        return strconv.FormatInt(duration/60, 10) + "分钟前"
    } else if duration < 86400 {
        return strconv.FormatInt(duration/3600, 10) + "小时前"
    } else if Timestamp(WithTimeAsStartOfDay(time.Now().Add(-time.Hour*24))) <= milliseconds {
        return "昨天 " + FormatTimeToStr(t, timeFormat["hh:mm:ss"])
    } else if Timestamp(WithTimeAsStartOfDay(time.Now().Add(-time.Hour*24*2))) <= milliseconds {
        return "前天 " + FormatTimeToStr(t, timeFormat["hh:mm:ss"])
    } else {
        return FormatTimeToStr(t, timeFormat["yyyy-mm-dd"])
    }
}

// AddMinute add or sub minute to the time
func AddMinute(t time.Time, minute int64) time.Time {
    return t.Add(time.Minute * time.Duration(minute))
}

// AddHour add or sub hour to the time
func AddHour(t time.Time, hour int64) time.Time {
    return t.Add(time.Hour * time.Duration(hour))
}

// AddDay add or sub day to the time
func AddDay(t time.Time, day int64) time.Time {
    return t.Add(24 * time.Hour * time.Duration(day))
}
