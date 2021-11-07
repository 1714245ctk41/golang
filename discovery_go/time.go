package main

import (
    "fmt"
    "time"
)

func timeHandle() {
    now := time.Now()

    fmt.Println("now:", now)

    nowNanosecond := time.Now().UnixNano()
    then := now.Add(10 * time.Minute)

    // if we had fix number of units to subtract, we can use following line instead fo above 2 lines. It does type convertion automatically.
    thenNanosecond := then.UnixNano() 
    time := thenNanosecond - nowNanosecond

    fmt.Println("10 second now:", nowNanosecond)
    
    fmt.Println("10 nanosecond ago:", then)
    fmt.Println("10 thenNanosecond ago:", thenNanosecond)
    fmt.Println("between: ", time)
}
func makeTimestamp() int64 {
    return time.Now().UnixNano() / int64(time.Millisecond)
}
func nowAsUnixMilliseconds() int64{
    return time.Now().Round(time.Millisecond).UnixNano() / 1e6
}
func findMinAndMax(a [5]int) (min int, max int) {
    min = a[0]
    max = a[0]
    for _, value := range a {
        if value < min {
            min = value
        }
        if value > max {
            max = value
        }
    }
    return min, max
}