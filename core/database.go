package core

import (
	"context"
	"dance/conf"
	"errors"
	"time"

	logrus "github.com/sirupsen/logrus"
	driver "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

var (
	db *gorm.DB
)

func GetDB() *gorm.DB {
	return db
}

type sqlLogger struct {
	*logrus.Entry
	logger.Config
}

// LogMode log mode
func (l *sqlLogger) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

// Info print info
func (l sqlLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		info := append([]interface{}{msg}, data...)
		l.Entry.WithFields(logrus.Fields{"location": utils.FileWithLineNum()}).Info(info...)
	}
}

// Warn print warn messages
func (l sqlLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		info := append([]interface{}{msg}, data...)
		l.Entry.WithFields(logrus.Fields{"location": utils.FileWithLineNum()}).Warn(info...)
	}
}

// Error print error messages
func (l sqlLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		info := append([]interface{}{msg}, data...)
		l.Entry.WithFields(logrus.Fields{"location": utils.FileWithLineNum()}).Error(info...)
	}
}

// Trace print sql message
func (l sqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= logger.Error && (!errors.Is(err, gorm.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		fields := logrus.Fields{"error": err, "delay": elapsed.Milliseconds(), "rows": rows, "sql": sql}
		fields["location"] = utils.FileWithLineNum()
		l.WithFields(fields).Error("database error")
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
		sql, rows := fc()
		fields := logrus.Fields{"error": err, "delay": elapsed.Milliseconds(), "rows": rows, "sql": sql}
		fields["location"] = utils.FileWithLineNum()
		l.WithFields(fields).Warn("database warn")
	case l.LogLevel == logger.Info:
		sql, rows := fc()
		fields := logrus.Fields{"error": err, "delay": elapsed.Milliseconds(), "rows": rows, "sql": sql}
		fields["location"] = utils.FileWithLineNum()
		l.WithFields(fields).Info("database info")
	}
}

func InitDB() {
	var log = &sqlLogger{
		logrus.WithFields(logrus.Fields{"signature": "sqlmonitor"}),
		logger.Config{
			SlowThreshold:             10 * time.Millisecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	}

	var err error
	var dialect = driver.Open(conf.Config.Database)
	var configs = &gorm.Config{
		SkipDefaultTransaction:   true,
		DisableNestedTransaction: true,
		Logger:                   log,
	}
	db, err = gorm.Open(dialect, configs)
	if err != nil {
		panic(err)
	}
	db, err := db.DB()
	// 数据库连接池相关设置
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(time.Hour)
	db.SetConnMaxLifetime(time.Hour)
}
