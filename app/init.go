package app

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type InitType struct {
	Port     int
	Location *time.Location
}

func Init() (*InitType, error) {
	// โหลดไฟล์ .env ก่อน แล้วเอาไฟล์ .env.local มา merge ทับอีกที
	if IsDebug() {
		_ = godotenv.Load(".env.local")
		_ = godotenv.Load(".env")
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Time Zone
	loc, ex := time.LoadLocation("Asia/Bangkok")
	if ex != nil {
		log.Fatalf("set time zone, %v", ex.Error())
	}
	time.Local = loc

	return &InitType{
		Location: loc,
		Port:     8080,
	}, nil

}

func IsDebug() bool {
	return strings.Contains(os.Args[0], `__debug_bin`)
}
