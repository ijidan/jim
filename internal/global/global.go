package global

import (
	"jim/config"
	"jim/pkg"
	"os"
)

var (
	Root, _   = os.Getwd()
	Config    = config.GetConfigInstance(Root)
	Logger    = pkg.GetLoggerInstance(Config, Root)
	Db        = pkg.GetDbInstance(Config)
	Rd        = pkg.GetRdInstance(Config)
	Response  = pkg.GetResponseInstance()
	RequestId = "X-Request-Id"
)

func Close() {
	_ = Rd.Close()
}
