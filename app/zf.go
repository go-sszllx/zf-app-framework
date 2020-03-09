package app

// Hook function to run
type hookfunc func() error

var (
	hooks = make([]hookfunc, 0) //hook function slice to store the hookfunc
)

func AddAPPStartHook(hf ...hookfunc) {
	hooks = append(hooks, hf...)
}

func Run(params ...string) {
	initBeforeRun()

	if len(params) > 0 {
		switch params[0] {
		case "micro":
			ZfApp.RunMicro(params)
		default:
			ZfApp.RunMono(params)
		}
	} else {
		ZfApp.RunMono(params)
	}


}

func initBeforeRun() {

	for _, hk := range hooks {
		if err := hk(); err != nil {
			panic(err)
		}
	}
}
