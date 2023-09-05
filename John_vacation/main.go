package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print(Solution(2014, "April", "May", "Wednesday"))
}

// given Y = 2014, A = "April", B = "May" and W = "Wednesday"
// the function should return 7
// since John can leave his city on April 7th and come back on May 25th,
// so he will spend 7 weeks on Hawaii (the weeks beginning, respectively, on April 7th, 14th, 21st, 28th and May 5th, 12th, 19th).
func Solution(y int, a string, b string, w string) int {
	months := map[string]int{
		"January":   1,
		"February":  2,
		"March":     3,
		"April":     4,
		"May":       5,
		"June":      6,
		"July":      7,
		"August":    8,
		"September": 9,
		"October":   10,
		"November":  11,
		"December":  12,
	}

	days := map[string]int{
		"Sunday":    0,
		"Monday":    1,
		"Tuesday":   2,
		"Wednesday": 3,
		"Thursday":  4,
		"Friday":    5,
		"Saturday":  6,
	}

	// Tính số ngày trong tháng a
	daysInA := time.Date(y, time.Month(months[a]+1), 0, 0, 0, 0, 0, time.UTC).Day()

	// Tính ngày đầu tiên của tháng a
	firstDayOfA := time.Date(y, time.Month(months[a]), 1, 0, 0, 0, 0, time.UTC)

	// Tìm ngày đầu tiên trong tuần của tháng a
	firstWeekdayOfA := firstDayOfA.Weekday()

	// Tìm ngày đầu tiên có thể rời khỏi thành phố trong tháng a
	firstLeaveDay := days[w] - int(firstWeekdayOfA) + 1
	if firstLeaveDay <= 0 {
		firstLeaveDay += 7
	}

	// Tính ngày cuối cùng của tháng b
	lastDayOfB := time.Date(y, time.Month(months[b]+1), 0, 0, 0, 0, 0, time.UTC)

	// Tìm ngày cuối cùng có thể trở về thành phố trong tháng b
	lastReturnDay := lastDayOfB.Day() - (int(lastDayOfB.Weekday()) + 1)
	if lastReturnDay >= lastDayOfB.Day() {
		lastReturnDay -= 7
	}

	// Tính số ngày giữa hai ngày này
	daysBetween := lastReturnDay - firstLeaveDay + daysInA

	if months[a] < months[b] {
		// Nếu hai tháng khác nhau, cộng thêm số ngày của các tháng giữa chúng
		for i := months[a] + 1; i < months[b]; i++ {
			daysBetween += time.Date(y, time.Month(i+1), 0, 0, 0, 0, 0, time.UTC).Day()
		}
	}

	// Chia số ngày cho bảy để được số tuần
	weeks := daysBetween / 7

	// Trả về kết quả
	return weeks

}
