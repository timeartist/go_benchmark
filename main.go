package main

import "fmt"
import "os"
import "os/signal"
import "syscall"
import "time"

func performGetTest(testFixture fixture, latencies chan<- uint32, done <-chan bool) {   
    for {
        select {
            case <- done:
                break
            default:
                stopwatch := StartStopwatch()
                testFixture.Get("foo")
                stopwatch.Stop()
                latency := stopwatch.Nanoseconds()
                latencies <- latency
        }
    }
}

func performSetTest(testFixture fixture, latencies chan<- uint32, done <-chan bool) {   
    for {
        select {
            case <- done:
                break
            default:
                stopwatch := StartStopwatch()
                testFixture.Set("foo", "bar")
                stopwatch.Stop()
                latency := stopwatch.Nanoseconds()
                latencies <- latency
        }
    }
}

func main() {
    signals := make(chan os.Signal, 1)
    done := make(chan bool, 1)
    latencies := make(chan uint32, 2147483647)
    
    signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
    
    go func() {
        signal := <- signals
        fmt.Println(signal)
        done <- true
    }()
    
    R1 := createRedis("redis://localhost:6379/0")
    defer R1.Close()
    
    R2 := createRedis("redis://localhost:6379/0")
    defer R2.Close()

    go performGetTest(&R1, latencies, done)
    go performSetTest(&R2, latencies, done)
    
    tick := time.Tick(time.Second)
    go func() {
        for {
            <-tick
            var sum uint32 = 0
            select {
                case <- done:
                    break
                default:
                    count := uint32(len(latencies)) 
                    for j := uint32(0); j < count; j++ {
                        sum += <-latencies
                    }
                    
                    fmt.Println("sum")
                    fmt.Println(sum)
                    fmt.Println("count")
                    fmt.Println(count)
                    fmt.Println("average")
                    fmt.Println(sum/count)
                    
                    //items <- i
                    //i++
            }
        }
    }()
    
    fmt.Println("awaiting signal")
    <-done
    fmt.Println("done")
    
}