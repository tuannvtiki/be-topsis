package main

import (
	"log"
	"os"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"topsis/client/notify"
	"topsis/config"
	"topsis/internal/domain/usecase"
	"topsis/internal/infrastructure/repository"
)

func main() {
	cfg, err := config.LoadConfig("build")
	if err != nil {
		logrus.Warnf("cannot load config: %v", err)
		return
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(cfg.DBSource), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal("failed to open database:", err)
		return
	}
	statisticalRepo := repository.NewStatisticalRepository(db)
	statisticalDomain := usecase.NewStatisticalDomain(statisticalRepo)

	notifyBot := notify.NewBotNotify(cfg, statisticalDomain)
	logrus.Info("Create new cron")
	c := cron.New()

	// Notify run at 17:30 every day
	notifyRun, err := c.AddFunc(cfg.CronNotifyRun, func() {
		err = notifyBot.ProcessNotifyRun()
		if err != nil {
			logrus.Warnf("send message for run is failed %v", err)
		}
	})

	// Notify summary 09:00 every day
	notifySummary, err := c.AddFunc(cfg.CronNotifySummary, func() {
		err = notifyBot.ProcessNotifySummary()
		if err != nil {
			logrus.Warnf("send message for summary is failed %v", err)
		}
	})

	// Notify statistical on day-of-month 1
	notifyStatistical, err := c.AddFunc(cfg.CronNotifyStatistical, func() {
		err = notifyBot.ProcessNotifyStatistical()
		if err != nil {
			logrus.Warnf("send message for statistical is failed %v", err)
		}
	})
	if err != nil {
		c.Remove(notifyRun)
		c.Remove(notifySummary)
		c.Remove(notifyStatistical)
		return
	}

	// Start cron with one scheduled job
	logrus.Info("Start cron")
	c.Run()
}
