package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// ANSI color codes
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Yellow = "\033[33m"
	Green  = "\033[32m"
	Cyan   = "\033[36m"
	Blue   = "\033[34m"
)

func timestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func Info(msg string, args ...interface{}) {
	fmt.Printf("%s%s [INFO] %s%s\n", Cyan, timestamp(), fmt.Sprintf(msg, args...), Reset)
}

func Warn(msg string, args ...interface{}) {
	fmt.Printf("%s%s [WARN] %s%s\n", Yellow, timestamp(), fmt.Sprintf(msg, args...), Reset)
}

func Error(msg string, args ...interface{}) {
	fmt.Printf("%s%s [ERROR] %s%s\n", Red, timestamp(), fmt.Sprintf(msg, args...), Reset)
}

func Debug(msg string, args ...interface{}) {
	fmt.Printf("%s%s [DEBUG] %s%s\n", Blue, timestamp(), fmt.Sprintf(msg, args...), Reset)
}

func Success(msg string, args ...interface{}) {
	fmt.Printf("%s%s [OK] %s%s\n", Green, timestamp(), fmt.Sprintf(msg, args...), Reset)
}

func Fatal(msg string, args ...interface{}) {
	fmt.Printf("%s%s [FATAL] %s%s\n", Red, timestamp(), fmt.Sprintf(msg, args...), Reset)
	os.Exit(1)
}

func Panic(msg string, args ...interface{}) {
	panic(fmt.Sprintf("[%s] [PANIC] %s", timestamp(), fmt.Sprintf(msg, args...)))
}
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()
		latency := endTime.Sub(startTime)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path

		// Log based on status code
		switch {
		case statusCode >= 500:
			Error("%s - [%s] \"%s %s\" %d %s", clientIP, endTime.Format(time.RFC1123), method, path, statusCode, latency)
		case statusCode >= 400:
			Warn("%s - [%s] \"%s %s\" %d %s", clientIP, endTime.Format(time.RFC1123), method, path, statusCode, latency)
		default:
			Info("%s - [%s] \"%s %s\" %d %s", clientIP, endTime.Format(time.RFC1123), method, path, statusCode, latency)
		}
	}
}
