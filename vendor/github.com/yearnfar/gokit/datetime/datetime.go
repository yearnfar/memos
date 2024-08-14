package datetime

import (
	"strings"
	"time"
)

// 参考：https://docs.python.org/3/library/datetime.html

// DateFormat pattern rules.
var layoutPatterns = []string{
	// 年
	"%Y", "2006", // 以世纪为十进制数的年份。例如: 1999 或 2003
	"%-Y", "06", // 没有世纪的年份并以零填充的十进制数字。例如，99 或 03

	// 月
	"%m", "01", // 月份为零填充的十进制数字。例如：01, 02, …, 12
	"%-m", "1", // 月份，不用0填充。例如：1, 2, …, 12
	"%b", "Jan", // 月份缩写名称。例如：Jan, Feb, …, Dec
	"%B", "January", // 月份作为区域设置的全名。例如：January, February, …, December

	// 日
	"%d", "02", // 使用0填充。例如：01, 02, …, 31
	"%-d", "2", // 不用0填充。例如：1, 2, …, 31

	// 时
	"%H", "15", // 小时（24小时时钟）作为零填充的十进制数字。例如：00, 01, …, 23
	"%I", "03", // 使用0填充的2位十进制数字。例如：01, 02, …, 12
	"%-I", "3", // 不使用0天子的十进制数字。例如: 1, 2, …, 12

	// 分
	"%M", "04", // 使用0填充的2位十进制数字。例如: 00, 01, …, 59
	"%-M", "4", // 不使用0填充的十进制数字。例如: 1, 2, …, 59

	// 秒
	"%S", "05", // 使用0填充的2位十进制数字。例如: 00, 01, …, 59
	"%-S", "5", // 不使用0填充的十进制数字。例如: 1, 2, …, 59

	// 周
	"%a", "Mon", // 星期作为区域设置的缩写名称。例如：Mon, Tue, …, Sun
	"%A", "Monday", // 星期作为区域设置的全名。例如：Monday, Tuesday, …, Sunday

	// 上午、下午
	"%p", "PM",

	// 时区
	"%Z", "MST",
	"%z", "-0700",
}

var replacer = strings.NewReplacer(layoutPatterns...)

// Strptime 解析时间
func Strptime(v string, layout string) (time.Time, error) {
	return time.ParseInLocation(replacer.Replace(layout), v, time.Local)
}

// Strftime 格式化时间
func Strftime(t time.Time, layout string) string {
	return t.Format(replacer.Replace(layout))
}
