package main

import (
    "fmt"
    "unsafe"
)

func main() {
    fmt.Println(unsafe.Sizeof("11111111111111111111111111111111"))
    fmt.Println(unsafe.Sizeof("1"))
    //events2 := make(chan int, 100)
    //go func() {
    //    for{
    //        v := <- events1
    //        fmt.Println(v)
    //    }
    //
    //}()

    //go func() {
    //    time.Sleep(2*time.Second)
    //    events1 <- 1
    //    events2 <- 2
    //    time.Sleep(2*time.Second)
    //    close(events1)
    //    close(events2)
    //    events1 <- 1
    //    events2 <- 2
    //}()
    //
    //select {
    //
    //}
}