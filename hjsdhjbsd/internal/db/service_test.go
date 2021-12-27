package db

import (
	"log"
	"sync"
	"testing"
)

var initOnce sync.Once

func INIT() {
	initOnce.Do(func() {
		log.Println("A")
	})
}

func Test_Exec2(t *testing.T) {
	for i:=0; i < 50; i++{
		INIT()
		log.Println("B")
	}
}