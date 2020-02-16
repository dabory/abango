// Author : Eric Kim
// Build Date : 23 Jul 2018  Last Update 02 Aug 2018
// End-Agent for Passcon Multi OS go binding with Windows, MacOS, iOS, and Android
// All rights are reserved.

package etc

import (
	time "time"
)

func getNow() time.Time {
	loc, _ := time.LoadLocation("UTC")
	return time.Now().In(loc)
}

func GetNowUnix(sec ...int) int64 {
	var ret int64
	if sec == nil {
		ret = time.Now().UTC().Unix()
	} else {
		ret = time.Now().Add(time.Duration(sec[0]) * time.Second).UTC().Unix()
	}
	return ret
}
