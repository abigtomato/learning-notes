package main

import (
	"fmt"
	"time"
)

// 时间日期库
func timeConst() {
	i := 0
	for {
		i++
		fmt.Println(i)
		// Nanosecond	纳秒
		// Microsecond	微秒 = 1000 * Nanosecond
		// Millisecond  毫秒 = 1000 * Microsecond
		// Second		秒 = 1000 * Millisecond
		// Minute		分 = 60 * Second
		// Hour			时 = 60 * Minute
		time.Sleep(time.Millisecond * 100)
		if i == 100 {
			break
		}
	}
}

func format(now time.Time) {
	str := fmt.Sprintf("%v/%v/%v %v:%v:%v\n", 
		now.Year(),
		int(now.Month()),
		now.Day(),
		now.Hour(),
		now.Minute(),
		now.Second())
	fmt.Printf("Time: %v", str)

	// 格式化时间日期，2006-01-02 15/04/05是格式模板
	fmt.Println(now.Format("2006-01-02 15/04/05"))
}

func main() {
	now := time.Now()
	fmt.Printf("now=%v, type=%T\n", now, now)

	fmt.Printf("%v年%v月%v日%v时%v分%v秒\n",
		now.Year(),
		int(now.Month()),
		now.Day(),
		now.Hour(),
		now.Minute(),
		now.Second())

	// 时间日期格式化
	format(now)

	// 时间戳
	fmt.Printf("Unix时间戳=%v, UnixNano时间戳=%v\n", now.Unix(), now.UnixNano())
}
