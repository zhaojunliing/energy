//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build !windows
// +build !windows

package cef

// 每一次拖拽区域改变都需要重新设置
func (m *TCEFChromiumBrowser) setDraggableRegions() {
}

// 非windows 没有 CompMsg 事件
func (m *TCEFChromiumBrowser) registerWindowsCompMsgEvent() {
}
