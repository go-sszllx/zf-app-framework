package app

import "framework/scheduler"

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
	// TODO:
	//     1. 通过入参决定是否创建普通service还是web service
	//     2. 暂使用beego的config和orm模块

	/*
	普通micro
		ip := beego.AppConfig.DefaultString("nsq::ip", cm.BrokerAddress)
		port := beego.AppConfig.DefaultString("nsq::port", cm.BrokerPort)
		var BrokerHosts = []string{ // specified nsq service host
			0: ip + ":" + port,
		}

		ip = beego.AppConfig.DefaultString("etcd::ip", cm.RegistryAddress)
		port = beego.AppConfig.DefaultString("etcd::port", cm.RegistryPort)
		var RegistryHosts = []string{ // specified etcd service host
			0: ip + ":" + port,
		}

		name := beego.AppConfig.DefaultString("micro::serviceName", cm.MicroServiceName)
		service := micro.NewService(
			// This name must match the package name given in your protobuf definition
			micro.Name(name),
			micro.Version("0.1"),
			micro.Broker(nsq.NewBroker(func(o *broker.Options) { // specified using nsq broker
				o.Addrs = BrokerHosts
			})),
			micro.Registry(etcd.NewRegistry(func(o *registry.Options) { // specified using etcd registry
				o.Addrs = RegistryHosts
			})),
		)

		// Init will parse the command line flags.
		service.Init()
	*/

	/*
	web micro
		// specified web service host
		ip := beego.AppConfig.DefaultString("micro::web.ip", cm.MicroWebAddress)
		port := beego.AppConfig.DefaultString("micro::web.port", cm.MicroWebPort)
		webHost := ip + ":" + port

		name := beego.AppConfig.DefaultString("micro::web.serviceName", cm.MicroWebServiceName)
		service := web.NewService(
			web.Name(name),
			web.Address(webHost),
		)

		// add all routes
		for key, value := range instance.router.RouterMap {
			service.HandleFunc(key, value)
		}

		// Init will parse the command line flags.
		service.Init()
	*/

	initInternal()
}

func (a *App)RunMono(params []string) {
	initInternal()
}

func initInternal() {
	scheduler.Init()
}
