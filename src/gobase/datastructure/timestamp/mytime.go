package main

import (
	"time"
	"fmt"
)

func main() {

	now := time.Now()
	fmt.Println("当前时间:", now)
	fmt.Println("当前时间2:", now.Format("2020-09-29 00:00:00"))
	fmt.Println("当前时间3:", now.Format(time.RFC3339))

	secs := now.Unix()
	fmt.Println(secs)

	nanos := now.UnixNano()
	fmt.Println(nanos)

	millis := nanos / 1000000
	fmt.Println(millis)

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))

	t1, e := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	fmt.Println("解析时间:", t1, e)
	fmt.Println(now.Format("3:04PM"))
	fmt.Println(now.Format("Mon Jan _2 15:04:05 2006"))
	fmt.Println(now.Format("2006-01-02T15:04:05.999999-07:00"))
	form:= "3 04 PM"
	t2,e:= time.Parse(form,"8 41 PM")
	fmt.Println(t2)

}
