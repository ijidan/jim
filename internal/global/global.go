package global

import (
	"jim/config"
	"jim/pkg"
	"os"
)

var (
	Root, _  = os.Getwd()
	Config   = config.GetConfigInstance(Root)
	Logger   = pkg.GetLoggerInstance(Config, Root)
	Db       = pkg.GetDbInstance(Config)
	Rd       = pkg.GetRdInstance(Config)
	Jaeger   = pkg.GetJaegerInstance(Config, "gim_api", "gim_api_root")
	Response = pkg.GetResponseInstance()
)

func Close() {
	_ = Rd.Close()
}
