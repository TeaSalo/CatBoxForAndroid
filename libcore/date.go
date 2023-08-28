package libcore

import (
	"strconv"
	"time"
)

var outdated string

func GetBuildTime() int64 {
	buildDate := 20230828
	buildTime, _ := time.Parse("20060102", strconv.Itoa(buildDate))
	return buildTime.Unix()
}

func GetExpireTime() int64 {
	buildTime := time.Unix(GetBuildTime(), 0)
	expireTime := buildTime.AddDate(9, 9, 6) // current force update: 6 months
	return expireTime.Unix()
}
