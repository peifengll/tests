package cc

import (
	"fmt"
	"testing"
)

func FormMetricsKey(ss ...string) string {
	keyLen := len(ss)
	if keyLen == 1 {
		return ss[0]
	}
	var tsl int
	for _, s := range ss {
		tsl += len(s) + 2
	}
	if keyLen&1 == 0 {
		tsl += 2
	}
	bs := make([]byte, tsl)
	bl := 0

	mainKey := ss[0]
	bl += copy(bs[bl:], []byte(mainKey+`{`))
	var s string
	var tempVale string
	for i := 1; i < keyLen; i += 2 {
		tempKey := ss[i]
		if i+1 == keyLen {
			tempVale = ""
		} else {
			tempVale = ss[i+1]
		}
		if i == 1 {
			s = tempKey + "=\"" + tempVale + "\""
		} else {
			s = "," + tempKey + "=\"" + tempVale + "\""
		}
		bl += copy(bs[bl:], []byte(s))
	}
	bl += copy(bs[bl:], []byte{'}'})
	return string(bs[:bl])
}

func TestFormMetricsKey(t *testing.T) {
	packetStatKey := FormMetricsKey(
		"_cron_job_retry_count",
		//"_3rd_party_mysql_query_count", "category", "CarCache",
	)
	fmt.Println(packetStatKey)

}
