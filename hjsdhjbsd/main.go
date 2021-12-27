package main

import (
	"context"
	"log"
	"time"
)

func main() {
	// ctx := context.Background()
	// ctx2 := context.TODO()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//ctx2 := context.WithTimeout(ctx, 1*time.Second)
	ch := make(chan int, 10)
	ch2 := make(chan struct{}, 10)
	go func() {
		for i := 0; i < 150; i++ {
			ch <- i
			if i == 50 {
				close(ch)
				return
			}
		}
	}()

	go func(ctx2 context.Context) {
		for {
			select {
			case <-ctx2.Done():
				return
			default:
				select {
				case v := <-ch:
					log.Println(v)
				case v := <-ch2:
					log.Println(v)
				// default:
				// 	log.Println("Не прочитал")
				default:
				}
			}
		}
	}(ctx)

	time.Sleep(time.Second * 10)
}
