package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret       string `init:"JWTSECRET"`
	PageSize        int    `ini:"PAGESIZE"`
	RuntimeRootPath string `ini:"RUNTIME_ROOTPATH"`

	PrefixUrl      string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string
	QrCodeSavePath string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

type Server struct {
	RunMode      string        `ini:"RUN_MODE"`
	HttpPort     int           `ini:"HTTP_PORT"`
	ReadTimeout  time.Duration `ini:"READ_TIMEOUT"`
	WriteTimeout time.Duration `ini:"WRITE_TIMEOUT"`
}

type Database struct {
	Type        string `ini:"TYPE"`
	User        string `ini:"USER"`
	Password    string `ini:"PASSWORD"`
	Host        string `ini:"HOST"`
	Name        string `ini:"NAME"`
	TablePrefix string `ini:"TABLE_PREFIX"`
}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var AppSetting = new(App)
var ServerSetting = new(Server)
var DatabaseSetting = new(Database)
var RedisSetting = new(Redis)

/*
	在 MapTo 中 typ.Kind() == reflect.Ptr 约束了必须使用指针，
	否则会返回 cannot map to non-pointer struct 的错误。这个是表面原因
*/

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)
	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
