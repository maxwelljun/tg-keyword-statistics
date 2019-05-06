package main

import (
	"github.com/robfig/cron"
)

func dbcron() {
	c := cron.New()
	//every day 00:00
	_ = c.AddFunc("@daily", func() {
		dbKeyword0("day")
	})
	//every month 00:00 no.1
	_ = c.AddFunc("@monthly", func() {
		dbKeyword0("month")
	})
	//every week 00:00 sun
	_ = c.AddFunc("@weekly", func() {
		dbKeyword0("week")
	})
	//every year
	_ = c.AddFunc("@yearly", func() {
		dbKeyword0("year")
	})
	c.Start()
}