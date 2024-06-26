//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build windows
// +build windows

// LCL窗口组件定义和实现-windows平台

package cef

import (
	"github.com/energye/energy/v2/cef/winapi"
	"github.com/energye/energy/v2/consts/messages"
	et "github.com/energye/energy/v2/types"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/win"
)

// 定义四角和边框范围
var (
	angleRange  int32 = 10 //四角
	borderRange int32 = 5  //四边框
)

// 组件消息类型
type compMessageType int8

const (
	cmtCEF compMessageType = iota
	cmtLCL
)

// ShowTitle 显示标题栏
func (m *LCLBrowserWindow) ShowTitle() {
	m.WindowProperty().EnableHideCaption = false
	//win.SetWindowLong(m.Handle(), win.GWL_STYLE, uintptr(win.GetWindowLong(m.Handle(), win.GWL_STYLE)|win.WS_CAPTION))
	//win.SetWindowPos(m.Handle(), m.Handle(), 0, 0, 0, 0, win.SWP_NOSIZE|win.SWP_NOMOVE|win.SWP_NOZORDER|win.SWP_NOACTIVATE|win.SWP_FRAMECHANGED)
	m.EnabledMaximize(m.WindowProperty().EnableMaximize)
	m.EnabledMinimize(m.WindowProperty().EnableMinimize)
	m.SetBorderStyle(types.BsSizeable)
}

// HideTitle 隐藏标题栏 无边框样式
func (m *LCLBrowserWindow) HideTitle() {
	m.WindowProperty().EnableHideCaption = true
	m.SetBorderStyle(types.BsNone)
	//m.Frameless()
}

// SetRoundRectRgn 窗口无边框时圆角设置
//
//	如果 rgn 值设置的过大同时开启GPU加速窗口会卡顿
func (m *LCLBrowserWindow) SetRoundRectRgn(rgn int) {
	if m.rgn == 0 && rgn > 0 {
		m.rgn = rgn
		m.SetOnPaint(func(sender lcl.IObject) {
			hnd := winapi.CreateRoundRectRgn(0, 0, et.LongInt(m.Width()), et.LongInt(m.Height()), et.LongInt(m.rgn), et.LongInt(m.rgn))
			winapi.SetWindowRgn(et.HWND(m.Handle()), hnd, true)
		})
	}
}

// FramelessForDefault 窗口四边框系统默认样式
//
//	TODO 窗口顶部有条线,
func (m *LCLBrowserWindow) FramelessForDefault() {
	gwlStyle := win.GetWindowLong(m.Handle(), win.GWL_STYLE)
	win.SetWindowLong(m.Handle(), win.GWL_STYLE, uintptr(gwlStyle&^win.WS_CAPTION&^win.WS_BORDER|win.WS_THICKFRAME))
	win.SetWindowPos(m.Handle(), 0, 0, 0, 0, 0, uint32(win.SWP_NOMOVE|win.SWP_NOSIZE|win.SWP_FRAMECHANGED))

	//winapi.SetClassLongPtr(this.Handle, GCL_STYLE, GetClassLong(this.Handle, GCL_STYLE)|CS_DropSHADOW)
	//gclStyle := winapi.GetClassLongPtr(et.HWND(m.Handle()), messages.GCL_STYLE)
	//winapi.WinSetClassLongPtr().SetClassLongPtr(et.HWND(m.Handle()), messages.GCL_STYLE, gclStyle|messages.CS_DROPSHADOW)
}

// FramelessForLine 窗口四边框是一条细线
func (m *LCLBrowserWindow) FramelessForLine() {
	gwlStyle := win.GetWindowLong(m.Handle(), win.GWL_STYLE)
	win.SetWindowLong(m.Handle(), win.GWL_STYLE, uintptr(gwlStyle&^win.WS_CAPTION&^win.WS_THICKFRAME|win.WS_BORDER))
	win.SetWindowPos(m.Handle(), 0, 0, 0, 0, 0, uint32(win.SWP_NOMOVE|win.SWP_NOSIZE|win.SWP_FRAMECHANGED))
}

// Frameless 无边框
func (m *LCLBrowserWindow) Frameless() {
	gwlStyle := win.GetWindowLong(m.Handle(), win.GWL_STYLE)
	win.SetWindowLong(m.Handle(), win.GWL_STYLE, uintptr(gwlStyle&^win.WS_CAPTION&^win.WS_THICKFRAME))
	win.SetWindowPos(m.Handle(), 0, 0, 0, 0, 0, uint32(win.SWP_NOMOVE|win.SWP_NOSIZE|win.SWP_FRAMECHANGED))
}

// windows无边框窗口任务栏处理
func (m *LCLBrowserWindow) taskMenu() {
	m.SetOnShow(func(sender lcl.IObject) bool {
		if m.WindowProperty().EnableHideCaption {
			gwlStyle := win.GetWindowLong(m.Handle(), win.GWL_STYLE)
			win.SetWindowLong(m.Handle(), win.GWL_STYLE, uintptr(gwlStyle|win.WS_SYSMENU|win.WS_MINIMIZEBOX))
		}
		return false
	})
}

// SetFocus
//
//	在窗口 (Visible = true) 显示之后设置窗口焦点
//	https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-showwindow
//	https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-setfocus
func (m *LCLBrowserWindow) SetFocus() {
	if m.TForm != nil {
		m.Visible()
		//窗口\激活在Z序中的下个顶层窗口
		m.Minimize()
		//激活窗口出现在前景
		m.Restore()
		//窗口设置焦点
		m.TForm.SetFocus()
	}
}

//func (m *LCLBrowserWindow) Frameless() {
//	var rect = &types.TRect{}
//	win.GetWindowRect(m.Handle(), rect)
//	win.SetWindowPos(m.Handle(), 0, rect.Left, rect.Top, rect.Right-rect.Left, rect.Bottom-rect.Top, win.SWP_FRAMECHANGED)
//}

func (m *LCLBrowserWindow) doOnRenderCompMsg(chromiumBrowser ICEFChromiumBrowser, messageType compMessageType, message *types.TMessage, lResult *types.LRESULT, aHandled *bool) {
	switch message.Msg {
	case messages.WM_NCLBUTTONDBLCLK: // 163 NC left dclick
		//标题栏拖拽区域 双击最大化和还原
		if m.cwcap.canCaption && m.WindowProperty().EnableWebkitAppRegionDClk && m.WindowProperty().EnableMaximize {
			*lResult = messages.HTCAPTION
			*aHandled = true
			if win.ReleaseCapture() {
				if m.WindowState() == types.WsNormal {
					win.PostMessage(m.Handle(), messages.WM_SYSCOMMAND, messages.SC_MAXIMIZE, 0)
				} else {
					win.PostMessage(m.Handle(), messages.WM_SYSCOMMAND, messages.SC_RESTORE, 0)
				}
				win.SendMessage(m.Handle(), messages.WM_NCLBUTTONUP, messages.HTCAPTION, 0)
			}
		}
	case messages.WM_NCMOUSEMOVE: // 160 nc mouse move
		m.cwcap.onNCMouseMove(m.Handle(), message, lResult, aHandled)
	case messages.WM_NCLBUTTONDOWN: // 161 nc left down
		// 标题栏和边框处理
		m.cwcap.onNCLButtonDown(m.Handle(), message, lResult, aHandled)
	case messages.WM_NCLBUTTONUP: // 162 nc l up
		if m.cwcap.canCaption {
			*lResult = messages.HTCAPTION
			*aHandled = true
		}
	case messages.WM_SETCURSOR: // 32 设置鼠标图标样式
		m.cwcap.onSetCursor(message, lResult, aHandled)
	case messages.WM_NCHITTEST: // 132 NCHITTEST
		if m.cwcap.borderMD { //TODO 测试windows7, 161消息之后再次处理132消息导致消息错误
			m.cwcap.borderMD = false
			return
		}
		var (
			x, y    int32 // 鼠标在当前窗口的坐标
			caption bool  // CEF HTML 自定义标题栏
		)
		if messageType == cmtCEF {
			//鼠标坐标是否在标题区域
			x, y, caption = m.cwcap.isCaption(chromiumBrowser, et.HWND(m.Handle()), message)
		} else if messageType == cmtLCL {
			x, y = m.cwcap.toPoint(message)
			p := &et.Point{
				X: x,
				Y: y,
			}
			winapi.ScreenToClient(et.HWND(m.Handle()), p)
			x, y = p.X, p.Y
		}
		if caption { //窗口标题栏
			*lResult = messages.HTCAPTION
			*aHandled = true
		} else if m.WindowProperty().EnableHideCaption && m.WindowProperty().EnableResize && m.WindowState() == types.WsNormal { //1.窗口隐藏标题栏 2.启用了调整窗口大小 3.非最大化、最小化、全屏状态
			//全屏时不能调整窗口大小
			if m.WindowProperty().current.ws == types.WsFullScreen {
				return
			}
			var rect types.TRect
			// 当前类型的消息取出计算的宽高
			if messageType == cmtCEF {
				rect = chromiumBrowser.WindowParent().BoundsRect()
			} else if messageType == cmtLCL {
				rect = m.BoundsRect()
			}
			// 判断当前鼠标是否在边框范围
			// 窗口边框和CEF组件边框

			handled := m.cwcap.onCanBorder(chromiumBrowser, x, y, &rect)
			if handled {
				// 鼠标在边框范围
				// 当是CEF组件消息，判断一次组件四边距离窗口四边间距，如果大于边框范围则取消操作
				// TODO 暂时不使用
				//if messageType == cmtCEF {
				//	windowRect := m.BoundsRect()
				//	switch m.cwcap.borderHT {
				//	case messages.HTTOP: // 上
				//		if rect.Top > borderRange {
				//			return
				//		}
				//	case messages.HTBOTTOM: // 下
				//		if (windowRect.Height() - rect.Bottom) > borderRange {
				//			return
				//		}
				//	case messages.HTLEFT: // 左
				//		if rect.Left > borderRange {
				//			return
				//		}
				//	case messages.HTRIGHT: // 右
				//		if (windowRect.Width() - rect.Right) > borderRange {
				//			return
				//		}
				//	case messages.HTTOPRIGHT: // 右上
				//		if (windowRect.Width()-rect.Right) > borderRange || rect.Top > borderRange {
				//			return
				//		}
				//	case messages.HTBOTTOMRIGHT: // 右下
				//		if (windowRect.Width()-rect.Right) > borderRange || (windowRect.Height()-rect.Bottom) > borderRange {
				//			return
				//		}
				//	case messages.HTTOPLEFT: // 左上
				//		if rect.Left > borderRange || rect.Top > borderRange {
				//			return
				//		}
				//	case messages.HTBOTTOMLEFT: // 左下
				//		if rect.Left > borderRange || (windowRect.Height()-rect.Bottom) > borderRange {
				//			return
				//		}
				//	}
				//}

				// 鼠标在边框位置
				*lResult = types.LRESULT(m.cwcap.borderHT)
				*aHandled = true
			}
		}
	}
}

// Restore Windows平台，窗口还原
func (m *LCLBrowserWindow) Restore() {
	if m.TForm == nil {
		return
	}
	RunOnMainThread(func() {
		if win.ReleaseCapture() {
			win.SendMessage(m.Handle(), messages.WM_SYSCOMMAND, messages.SC_RESTORE, 0)
		}
	})
}

// Minimize Windows平台，窗口最小化
func (m *LCLBrowserWindow) Minimize() {
	if m.TForm == nil {
		return
	}
	RunOnMainThread(func() {
		if win.ReleaseCapture() {
			win.PostMessage(m.Handle(), messages.WM_SYSCOMMAND, messages.SC_MINIMIZE, 0)
		}
	})
}

// Maximize Windows平台，窗口最大化/还原
func (m *LCLBrowserWindow) Maximize() {
	if m.TForm == nil || m.IsFullScreen() {
		return
	}
	RunOnMainThread(func() {
		if win.ReleaseCapture() {
			if m.WindowState() == types.WsNormal {
				win.PostMessage(m.Handle(), messages.WM_SYSCOMMAND, messages.SC_MAXIMIZE, 0)
			} else {
				win.SendMessage(m.Handle(), messages.WM_SYSCOMMAND, messages.SC_RESTORE, 0)
			}
		}
	})
}

// FullScreen 窗口全屏
func (m *LCLBrowserWindow) FullScreen() {
	if m.WindowProperty().EnableHideCaption {
		RunOnMainThread(func() {
			if m.WindowState() == types.WsMinimized || m.WindowState() == types.WsMaximized {
				if win.ReleaseCapture() {
					win.SendMessage(m.Handle(), messages.WM_SYSCOMMAND, messages.SC_RESTORE, 0)
				}
			}
			m.WindowProperty().current.ws = types.WsFullScreen
			m.setCurrentProperty()
			m.SetBoundsRect(m.Monitor().BoundsRect())
		})
	}
}

// ExitFullScreen 窗口退出全屏
func (m *LCLBrowserWindow) ExitFullScreen() {
	wp := m.WindowProperty()
	if wp.EnableHideCaption && wp.current.ws == types.WsFullScreen {
		RunOnMainThread(func() {
			wp.current.ws = types.WsNormal
			m.SetWindowState(types.WsNormal)
			m.SetBounds(wp.current.x, wp.current.y, wp.current.w, wp.current.h)
		})
	}
}

// IsFullScreen 是否全屏
func (m *LCLBrowserWindow) IsFullScreen() bool {
	return m.WindowProperty().current.ws == types.WsFullScreen
}

// 窗口透明
//func (m *LCLBrowserWindow) SetTransparentColor() {
//	m.SetColor(colors.ClNavy)
//	Exstyle := win.GetWindowLong(m.Handle(), win.GWL_EXSTYLE)
//	Exstyle = Exstyle | win.WS_EX_LAYERED&^win.WS_EX_TRANSPARENT
//	win.SetWindowLong(m.Handle(), win.GWL_EXSTYLE, uintptr(Exstyle))
//	win.SetLayeredWindowAttributes(m.Handle(),
//		colors.ClNavy, //crKey 指定需要透明的背景颜色值
//		255,           //bAlpha 设置透明度,0完全透明，255不透明
//		//LWA_ALPHA: crKey无效,bAlpha有效
//		//LWA_COLORKEY: 窗体中的所有颜色为crKey的地方全透明，bAlpha无效
//		//LWA_ALPHA | LWA_COLORKEY: crKey的地方全透明，其它地方根据bAlpha确定透明度
//		win.LWA_ALPHA|win.LWA_COLORKEY)
//}

func (m *LCLBrowserWindow) doDrag() {
	// Windows Drag Window
	// m.drag != nil 时，这里处理的是 up 事件, 给标题栏标记为false
	if m.drag != nil {
		m.drag.drag()
	} else {
		// 全屏时不能拖拽窗口
		if m.WindowProperty().current.ws == types.WsFullScreen {
			return
		}
		// 此时是 down 事件, 拖拽窗口
		if win.ReleaseCapture() {
			win.PostMessage(m.Handle(), messages.WM_NCLBUTTONDOWN, messages.HTCAPTION, 0)
			m.cwcap.canCaption = true
		}
	}
}
