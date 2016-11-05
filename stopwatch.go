package main

import "time"


type Stopwatch struct {
    start, stop time.Time
}

func (self *Stopwatch) Start() {
    self.start = time.Now()
}

func (self *Stopwatch) Stop() {
    self.stop = time.Now()
}

func (self *Stopwatch) Milliseconds() uint32 {
    return uint32(self.stop.Sub(self.start) / time.Millisecond)
}

func (self *Stopwatch) Nanoseconds() uint32 {
    return uint32(self.stop.Sub(self.start) / time.Nanosecond)
}

func StartStopwatch() *Stopwatch {
    stopwatch := Stopwatch{}
    stopwatch.Start()
    return &stopwatch
}


//func main() {
//    stopwatch := New()
//    stopwatch.Start()
//    stopwatch2 := NewStarted()
//    fmt.Println(stopwatch2.start)
//    stopwatch.Stop()
//    fmt.Println(stopwatch.Milliseconds())
//    fmt.Println(stopwatch.Nanoseconds())
//}

