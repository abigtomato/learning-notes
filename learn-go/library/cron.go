package main

import (
	"log"
	"github.com/robfig/cron"
)

// 定时任务
func main() {
	i := 0
	
	c := cron.New()
	c.AddFunc("*/3 * * * * ?", func() {
		i++
		log.Println("cron running: ", i)
	})
	c.Start()

	select{}
}
