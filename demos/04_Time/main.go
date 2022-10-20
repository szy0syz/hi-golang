package main

import (
	"fmt"
	"time"
)

var week time.Duration

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
	// 21.12.2011
	t = time.Now().UTC()
	time.Now().UTC()
	fmt.Println(t)
	fmt.Println(time.Now())
	// calculating times:
	week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
	week_from_now := t.Add(week)
	fmt.Println(week_from_now)
	// formatting times:
	fmt.Println(t.Format(time.RFC822))         // 20 Oct 22 00:57 UTC
	fmt.Println(t.Format(time.ANSIC))          // Thu Oct 20 01:00:22 2022
	fmt.Println(t.Format("21 Dec 2021 08:52")) // 2010 Dec 202010 08:2220
	s := t.Format("20060102")
	fmt.Println(t, "=>", s) // 2022-10-20 00:57:58.750378 +0000 UTC => 20221020
}
