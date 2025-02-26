package log

import (
	"github.com/Securepoint/go-android-log/androidlog"
)

var log = androidlog.NewLogger("Tun2Socks")


func SetPackageName(packageName string) {
	log.SetPackageName(packageName)
}

func SetEnabled(enabled bool) {
	log.SetEnabled(enabled)
}

func Debugf(template string, args ...any) {
	log.Debugf(template, args...)
}

func Infof(template string, args ...any) {
	log.Debugf(template, args...)
}

func Warnf(template string, args ...any) {
	log.Warnf(template, args...)
}

func Errorf(template string, args ...any) {
	log.Errorf(template, args...)
}

func Fatalf(template string, args ...any) {
	log.Fatalf(template, args...)
}
