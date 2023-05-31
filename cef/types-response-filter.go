//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ************************** creates ************************** //

// ResponseFilterRef -> ICefResourceHandler
var ResponseFilterRef responseFilter

type responseFilter uintptr

func (*responseFilter) New() *ICefResponseFilter {
	var result uintptr
	imports.Proc(internale_CefResponseFilterRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefResponseFilter{instance: unsafe.Pointer(result)}
	}
	return nil
}

// ************************** impl ************************** //

// Instance 实例
func (m *ICefResponseFilter) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefResponseFilter) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefResponseFilter) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefResponseFilter) InitFilter(fn responseFilterInitFilter) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefResponseFilter_InitFilter).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefResponseFilter) Filter(fn responseFilterFilter) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefResponseFilter_Filter).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ************************** events ************************** //

type responseFilterInitFilter func() bool                                                                                                                                                   //InitFilter
type responseFilterFilter func(dataIn uintptr, dataInSize uint32, dataInRead *uint32, dataOut uintptr, dataOutSize uint32, dataOutWritten *uint32) (status consts.TCefResponseFilterStatus) //Filter

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case responseFilterInitFilter:
			resultPtr := (*bool)(getPtr(0))
			*resultPtr = fn.(responseFilterInitFilter)()
		case responseFilterFilter:
			dataIn := getVal(0)
			dataInSize := uint32(getVal(1))
			dataInRead := (*uint32)(getPtr(2))
			dataOut := getVal(3)
			dataOutSize := uint32(getVal(4))
			dataOutWritten := (*uint32)(getPtr(5))
			statusPtr := (*int32)(getPtr(6))
			status := fn.(responseFilterFilter)(dataIn, dataInSize, dataInRead, dataOut, dataOutSize, dataOutWritten)
			*statusPtr = int32(status)
		default:
			return false
		}
		return true
	})
}
