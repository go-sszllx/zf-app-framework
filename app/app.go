package app

import (
	"github.com/astaxie/beego/logs"
)

var (
	// ZfApp is an application instance
	ZfApp *App
)

type App struct {
}

func init() {
	// create zf application
	ZfApp = NewApp()
}

func NewApp() *App {
	app := &App{}
	return app
}

func (a *App)RunMicro(params []string) {

}

func (a *App)RunMono(params []string) {

	for i := 0; i < len(params); i++ {
		switch params[i] {
		case "data":
			logs.Debug(params[i + 1])
			i++
		}
	}
}
