package main

import (
    "time"
    "context"
    "golang.org/x/time/rate"
)

func main() {
    RateLimit := 10
    BucketSize := 10
    ctx := context.Background()
    e := rate.Every(time.Second/RateLimit)
    l := rate.NewLimiter(e, BucketSize)

    for _, task := range tasks {
        err := l.Wait(ctx)
        if err != nil {
            panic(err)
        }
        // do something
    }
}
