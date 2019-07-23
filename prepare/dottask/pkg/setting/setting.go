package setting

import (
	"github.com/go-ini/ini"
	"github.com/gpmgo/gopm/modules/log"
	"os"
	"runtime"
)

var (
	Cfg                 *ini.File
	DefaultDatabaseName string
	currentPath         string
	configAppPath       string
)

func init() {
	var err error

	currentPath := os.Getenv("DOTTASK_WORKER_DIR")

	if sysType := runtime.GOOS; sysType == "windows" {
		configAppPath = currentPath + "config\\app.ini"
	} else {
		configAppPath = currentPath + "config/app.ini"
	}

	if Cfg, err = ini.Load(configAppPath); err != nil {
		log.Fatal("Fail to pare 'config/app.ini':%v", err)
	}
}

func GetDefuultDatabase() string {
	//fmt.Println(configAppPath)
	DefaultDatabaseName = Cfg.Section("default_database").Key("NAME").String()
	return DefaultDatabaseName
}
