package global

import (
	"path/filepath"
	"runtime"
)
var (
	_,c,_,_=runtime.Caller(0)
	ROOT=filepath.Dir(filepath.Dir(filepath.Dir(c)))
)