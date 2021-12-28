package timef

import "time"

// 时间操作

const (
	//SHORTTIMESTRING 短时间
	SHORTTIMESTRING = "20060102"
	//TIMEFORMAT 时间格式化
	TIMEFORMAT = "20060102150405"
	//NORMALTIMEFORMAT 普通格式化
	NORMALTIMEFORMAT    = "2006-01-02 15:04:05"
	ISONORMALTIMEFORMAT = "2006-01-02T15:04:05"
)

//GetTime 当前时间
//GetTime current time
func GetTime() time.Time {
	return time.Now()
}

//GetTimeByInt 时间戳转时间
//GetTimeByInt convert second unix timestamp to time
func GetTimeByInt(t1 int64) time.Time {
	return time.Unix(t1, 0)
}

//GetTimeByString 字符串转时间 eg: 20060102150405
//GetTimeByString convert string to time
func GetTimeByString(timestring string) (time.Time, error) {
	if timestring == "" {
		return time.Time{}, nil
	}
	return time.ParseInLocation(TIMEFORMAT, timestring, time.Local)
}

//GetTimeByNormalString 标准字符串 转 时间 eg: 2006-01-02 15:04:05
//GetTimeByNormalString convert standard string to time
func GetTimeByNormalString(timestring string) (time.Time, error) {
	if timestring == "" {
		return time.Time{}, nil
	}
	return time.ParseInLocation(NORMALTIMEFORMAT, timestring, time.Local)
}

//GetTimeByNormalString 标准字符串 转 时间 eg: 2006-01-02T15:04:05
//GetTimeByNormalString convert iso string to time
func GetTimeByISONormalString(timestring string) (time.Time, error) {
	if timestring == "" {
		return time.Time{}, nil
	}
	return time.ParseInLocation(ISONORMALTIMEFORMAT, timestring, time.Local)
}

//GetTimeString 格式化为： 20060102150405
//GetTimeString format time info to 20060102150405
func GetTimeString(t time.Time) string {
	return t.Format(TIMEFORMAT)
}

//GetNormalTimeString 格式化为：2006-01-02 15:04:05
//GetNormalTimeString format time info to 2006-01-02 15:04:05
func GetNormalTimeString(t time.Time) string {
	return t.Format(NORMALTIMEFORMAT)
}

//GetTimeUnix 转为时间戳：秒数
//GetTimeUnix convert time to second unix timestamp
func GetTimeUnix(t time.Time) int64 {
	return t.Unix()
}

//GetTimeMills 转为时间戳: 毫秒数
//GetTimeMills convert time to milliseconds unix timestamp
func GetTimeMills(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// GetNormalStringByTimestamp 时间戳 转 字符串 eg: 1640681215 to 2021-12-28 16:46:55
// GetNormalStringByTimestamp convert timestamp to string
func GetNormalStringByTimestamp(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format(NORMALTIMEFORMAT)
}

//CompareTime 比较两个时间大小
//CompareTime compare t1 is before  t2
func CompareTime(t1, t2 time.Time) bool {
	return t1.Before(t2)
}

//GetNextHourTime n小时候后的时间字符串
//GetNextHourTime get string time after n hour
func GetNextHourTime(s string, n int64) string {
	t2, _ := time.ParseInLocation(TIMEFORMAT, s, time.Local)
	t1 := t2.Add(time.Hour * time.Duration(n))
	return GetTimeString(t1)
}

//GetHourDiffer 计算两个时间差多少小时
//GetHourDiffer calculate how hours diff between startTime and endTime
func GetHourDiffer(startTime, endTime string) float32 {
	var hour float32
	t1, err := time.ParseInLocation(TIMEFORMAT, startTime, time.Local)
	t2, err := time.ParseInLocation(TIMEFORMAT, endTime, time.Local)
	if err == nil && CompareTime(t1, t2) {
		diff := GetTimeUnix(t2) - GetTimeUnix(t1)
		hour = float32(diff) / 3600
		return hour
	}
	return hour
}

//Checkhours 判断当前时间是否整点
//Checkhours check Whether now is Hour
func Checkhours() bool {
	_, m, s := GetTime().Clock()
	if m == s && m == 0 && s == 0 {
		return true
	}
	return false
}

//StringToNormalString 时间字符串转为标准字符串 eg: 20060102150  to 2006-01-02 15:04:05
//StringToNormalString convert string to normal string
func StringToNormalString(t string) string {
	if !(len(TIMEFORMAT) == len(t) || len(SHORTTIMESTRING) == len(t)) {
		return t
	}
	if len(SHORTTIMESTRING) == len(t) {
		t += "000000"
	}
	if len(TIMEFORMAT) == len(t) {
		t1, err := GetTimeByString(t)
		if err != nil {
			return t
		}
		t = GetTimeString(t1)
	}
	return t
}
