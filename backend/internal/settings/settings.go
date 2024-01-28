package settings

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/ini.v1"
)

type App struct {
	JwtSecret string

	RuntimeRootPath string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string

	EnableCors bool
}

var AppSettings = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSettings = &Server{}

type Database struct {
	Database string
	Password string
	Username string
	Port     string
	Host     string
}

var DatabaseSettings = &Database{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("config.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'config/config.ini': %v", err)
	}

	mapTo("app", AppSettings)
	mapTo("server", ServerSettings)
	mapTo("database", DatabaseSettings)

	AppSettings.LogSavePath = fmt.Sprintf("%s%s", AppSettings.RuntimeRootPath, AppSettings.LogSavePath)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
