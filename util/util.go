package util

import "time"
import "log"

func CheckAlive(updateTime string, days int) bool {
	layout := "Mon Jan 02 15:04:05 MST 2006"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, err := time.ParseInLocation(layout, updateTime, loc)
	t = t.AddDate(0, 0, days)
	if err != nil {
		log.Fatalf("time parse error", err)
	}
	log.Println(t)
	return t.Compare(time.Now()) > 0
}
