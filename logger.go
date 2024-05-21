package logrus_starter

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)
import "github.com/archine/ioc"

type Logger struct {
}

type config struct {
	// Level log level
	// default: debug (supports: error、info、trace、warn、panic、fetal、debug)
	Level string `mapstructure:"level"`
}

func (l *Logger) Init() {
	var conf config
	v := ioc.GetBeanByName("viper.Viper").(*viper.Viper)
	v.SetDefault("log.level", "debug")
	if err := v.Sub("log").Unmarshal(&conf); err != nil {
		panic(fmt.Sprintf("Failed to read the log config, %s", err.Error()))
	}
	if level, err := log.ParseLevel(conf.Level); err != nil {
		panic(fmt.Sprintf("init log level failed, invalid log level string, %s", err.Error()))
	} else {
		log.SetLevel(level)
		log.SetFormatter(&LogFormat{})
	}
}

func (l *Logger) Info(msg string, args ...any) {
	log.Infof(msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	log.Warnf(msg, args...)
}

func (l *Logger) Debug(msg string, args ...any) {
	log.Debugf(msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	log.Errorf(msg, args...)
}

func (l *Logger) Fatal(format string, v ...any) {
	log.Fatalf(format, v...)
}
