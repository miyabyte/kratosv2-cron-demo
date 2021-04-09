package test

import (
	"context"
	"github.com/go-co-op/gocron"
	"github.com/miyabyte/kratosv2-cron-demo/cron"
	"log"
	"testing"
	"time"
)

/*
package server

import (
	"...YOUR-PATH.../server"
)

func NewCronServer() (*server.Cron, error) {
	s := server.Get()
	_, e := s.Every(1).Seconds().Do(func() { log.Println(".") })
	_, e := s.Every(2).Seconds().Do(func() { log.Println("..") })
	return s,nil
}
*/

func TestRegister(t *testing.T) {
	s := cron.Get()
	_, e := s.Every(1).Seconds().Do(func() { log.Println("...") })
	if e != nil {
		log.Fatal(e)
	}
	s.StartAsync()
	time.Sleep(5 * time.Second)
	log.Printf("stop-err:%v", s.Stop())
}

func TestGocronEverySecondsCron(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)

	s := gocron.NewScheduler(time.Local)
	_, e := s.Every(1).Seconds().Do(func() { log.Println("...") })
	if e != nil {
		log.Fatal(e)
	}

	go func() {
		<-ctx.Done()
		s.Stop()
		return
	}()
	s.StartAsync()

	select {}
}
