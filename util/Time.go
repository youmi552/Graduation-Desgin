package util

import "time"

func GetTimeStamp() int {
	time := time.Now().Unix()
	return int(time)
}
