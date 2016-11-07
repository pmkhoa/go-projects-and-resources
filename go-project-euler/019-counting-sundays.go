// Counting Sundays
// You are given the following information, but you may prefer to do some research for yourself.
//
//     1 Jan 1900 was a Monday.
//     Thirty days has September,
//     April, June and November.
//     All the rest have thirty-one,
//     Saving February alone,
//     Which has twenty-eight, rain or shine.
//     And on leap years, twenty-nine.
//     A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
//
// How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?

package main

import (
    "fmt"
)

const NumDaysInRegularYear = 365

type Month struct {
    Name    string
    numDays int
}

func isLeapYear(year int) bool {
    return year % 4 == 0 && ( year % 100 != 0 && year % 400 == 0 )
}

var months = []Month{
    {"January", 31},
    {"February", 28},
    {"March", 31},
    {"April", 30},
    {"May", 31},
    {"June", 30},
    {"July", 31},
    {"August", 31},
    {"September", 30},
    {"October", 31},
    {"November", 30},
    {"December", 31},
}

func main() {
    curDay := (1 + NumDaysInRegularYear) % 7
    count := 0

    fmt.Println(curDay)
    for i := 1901; i <= 1902; i ++ {
        for _, month := range months {
            curDay += month.numDays
            if month.Name == "February" && isLeapYear(i) {
                curDay++
            }
            fmt.Println(curDay)
            curDay = curDay % 7
            if curDay == 0 {
                count++
            }
        }
    }
    fmt.Println(count)
}
