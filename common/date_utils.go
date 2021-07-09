package common

import "time"

// DateFormat 格式化时间
func DateFormat(date int64) string {
	unix := time.Unix(date, 0)
	return unix.Format("2006-01-02")
}
