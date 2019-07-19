package setting

import (
	"fmt"
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
	if currentPath, err = os.Getwd(); err != nil {
		log.Fatal("get setting current path err %v:", err)
	}

	currentPathLen := len(currentPath)

	if sysType := runtime.GOOS; sysType == "windows" {
		prixfLen := len("pkg\\setting")
		configAppPath = currentPath[0:currentPathLen-prixfLen] + "config\\app.ini"
	} else {
		prixfLen := len("pkg/setting")
		configAppPath = currentPath[0:currentPathLen-prixfLen] + "config/app.ini"
	}

	if Cfg, err = ini.Load(configAppPath); err != nil {
		log.Fatal("Fail to pare 'config/app.ini':%v", err)
	}
}

func GetDefuultDatabase() string {
	fmt.Println(configAppPath)
	DefaultDatabaseName = Cfg.Section("default_database").Key("NAME").String()
	return DefaultDatabaseName
}
