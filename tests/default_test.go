package test

import (
	_ "com.wangzhumo.iyouku/routers"
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestGet is a sample to run an endpoint test
func TestGet(t *testing.T) {
	goarch := runtime.GOOS
	fmt.Println(goarch)
	if goarch == "darwin" {
		fmt.Println("macOS")
	}
}

