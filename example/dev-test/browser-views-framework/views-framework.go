//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/cef/process"
	"github.com/energye/golcl/lcl"
)

//go:embed resources
var resources embed.FS

func main() {
	cef.GlobalInit(nil, &resources)
	fmt.Println("main", process.Args.ProcessType())

	application := cef.NewApplication()
	application.SetMultiThreadedMessageLoop(false)
	application.SetExternalMessagePump(false)
	application.SetOnContextInitialized(func() {
		fmt.Println("OnContextInitialized()")
		component := lcl.NewComponent(nil)
		chromiumConfig := cef.NewChromiumConfig()
		chromium := cef.NewChromium(component, chromiumConfig)
		browserViewComponent := cef.NewBrowserViewComponent(component)
		windowComponent := cef.NewWindowComponent(component)

		chromium.SetOnBeforePopup(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo, client *cef.ICefClient, noJavascriptAccess *bool) bool {
			fmt.Println("OnBeforePopup TargetUrl:", beforePopupInfo.TargetUrl)

			return false
		})
		windowComponent.SetOnWindowCreated(func(sender lcl.IObject, window *cef.ICefWindow) {
			fmt.Println("OnWindowCreated")
			b := chromium.CreateBrowserByBrowserViewComponent("https://www.baidu.com", browserViewComponent)
			fmt.Println("\tCreateBrowserByBrowserViewComponent", b)
			windowComponent.AddChildView(browserViewComponent)
			display := windowComponent.Display()
			fmt.Println("\tdisplay", display, "ClientAreaBoundsInScreen", windowComponent.ClientAreaBoundsInScreen(), display.ID(), display.DeviceScaleFactor())
			fmt.Println("\t", display.Bounds(), display.WorkArea())
			window.CenterWindow(cef.NewCefSize(1024, 768))
			browserViewComponent.RequestFocus()
			window.SetWindowAppIconFS(1, "resources/golang.jpeg")
			windowComponent.SetAlwaysOnTop(true)
			//window.SetFullscreen(true)
			window.Show()
			fmt.Println("SetOnWindowCreated end")
		})
		windowComponent.SetOnCanClose(func(sender lcl.IObject, window *cef.ICefWindow, aResult *bool) {
			fmt.Println("OnCanClose")
			application.QuitMessageLoop()
			*aResult = true
		})

		windowComponent.CreateTopLevelWindow()
	})
	//application.SetOnGetDefaultClient(func(client *cef.ICefClient) {
	//	fmt.Println("OnGetDefaultClient")
	//})
	var processType bool
	if !process.Args.IsMain() {
		processType = application.StartSubProcess()
	} else {
		processType = application.StartMainProcess()
		application.RunMessageLoop()
	}
	fmt.Println("application.StartMainProcess()", processType)
	if processType {
		fmt.Println("application.RunMessageLoop()")
	}
}
