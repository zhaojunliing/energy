package main

import (
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/examples/control-widget/src"
	_ "github.com/energye/energy/v2/examples/syso"
)

func main() {
	//开发环境中 MacOSX平台必须在"GlobalInit"之前设置CEF
	//设置使用CEF 和 CEF框架目录，生成开发执行应用程序包
	//环境变量 ENERGY_HOME="/app/cefframework" 配置框架所在目录
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	//创建应用
	app := cef.NewApplication()
	app.SetEnableGPU(true)
	app.SetUseMockKeyChain(true)
	//主进程窗口
	src.MainBrowserWindow(app)

	//运行应用
	cef.Run(app)
}
