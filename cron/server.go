package cron

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/go-kratos/kratos/v2/transport"
	"math/rand"
	"sync"
	"time"
)

var _ transport.Server = (*Cron)(nil)

var cron *Cron
var single sync.Once

type Cron struct {
	*gocron.Scheduler
}

func Get() *Cron {
	single.Do(func() {
		cron = &Cron{
			gocron.NewScheduler(time.Local),
		}
	})
	return cron
}

func (c *Cron) Endpoint() (string, error) {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("private-cron://%s-%v-%v:0", "cron", time.Now().Unix(), rand.Int()), nil
}

func (c *Cron) Start() error {
	c.StartAsync()
	return nil
}

func (c *Cron) Stop() error {
	c.Scheduler.Stop()
	return nil
}
