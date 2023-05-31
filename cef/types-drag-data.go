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
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// Instance 实例
func (m *ICefDragData) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefDragData) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefDragData) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefDragData) Clone() *ICefDragData {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(internale_CefDragData_Clone).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDragData{instance: getInstance(result)}
	}
	return nil
}

func (m *ICefDragData) IsReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_CefDragData_IsReadOnly).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefDragData) IsLink() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_CefDragData_IsLink).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefDragData) IsFragment() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_CefDragData_IsFragment).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefDragData) IsFile() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_CefDragData_IsFile).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefDragData) GetLinkUrl() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_CefDragData_GetLinkUrl).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefDragData) GetLinkTitle() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_CefDragData_GetLinkTitle).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefDragData) GetLinkMetadata() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_CefDragData_GetLinkMetadata).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefDragData) GetFragmentText() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_CefDragData_GetFragmentText).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefDragData) GetFragmentHtml() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_CefDragData_GetFragmentHtml).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefDragData) GetFragmentBaseUrl() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_CefDragData_GetFragmentBaseUrl).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefDragData) GetFileName() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_CefDragData_GetFileName).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefDragData) GetFileContents(writer *ICefStreamWriter) uint32 {
	if !m.IsValid() || !writer.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(internale_CefDragData_GetFileContents).Call(m.Instance(), writer.Instance())
	return uint32(r1)
}

func (m *ICefDragData) GetFileNames() ([]string, int32) {
	if !m.IsValid() {
		return nil, 0
	}
	var result uintptr
	r1, _, _ := imports.Proc(internale_CefDragData_GetFileNames).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 && r1 > 0 {
		fileNamesList := lcl.AsStrings(result)
		if fileNamesList.IsValid() {
			count := int(fileNamesList.Count())
			fileNames := make([]string, count, count)
			for i := 0; i < count; i++ {
				fileNames[i] = fileNamesList.Strings(int32(i))
			}
			fileNamesList.Free()
			return fileNames, int32(count)
		}
	}
	return nil, 0
}

func (m *ICefDragData) SetLinkUrl(url string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefDragData_SetLinkUrl).Call(m.Instance(), api.PascalStr(url))
}

func (m *ICefDragData) SetLinkTitle(title string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefDragData_SetLinkTitle).Call(m.Instance(), api.PascalStr(title))
}

func (m *ICefDragData) SetLinkMetadata(data string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefDragData_SetLinkMetadata).Call(m.Instance(), api.PascalStr(data))
}

func (m *ICefDragData) SetFragmentText(text string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefDragData_SetFragmentText).Call(m.Instance(), api.PascalStr(text))
}

func (m *ICefDragData) SetFragmentHtml(html string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefDragData_SetFragmentHtml).Call(m.Instance(), api.PascalStr(html))
}

func (m *ICefDragData) SetFragmentBaseUrl(baseUrl string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefDragData_SetFragmentBaseUrl).Call(m.Instance(), api.PascalStr(baseUrl))
}

func (m *ICefDragData) ResetFileContents() {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefDragData_ResetFileContents).Call(m.Instance())
}

func (m *ICefDragData) AddFile(path, displayName string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefDragData_AddFile).Call(m.Instance(), api.PascalStr(path), api.PascalStr(displayName))
}

func (m *ICefDragData) ClearFilenames() {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefDragData_ClearFilenames).Call(m.Instance())
}

func (m *ICefDragData) GetImage() *ICefImage {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(internale_CefDragData_GetImage).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefImage{instance: getInstance(result)}
	}
	return nil
}

func (m *ICefDragData) GetImageHotspot() *TCefPoint {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(internale_CefDragData_GetImageHotspot).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return (*TCefPoint)(unsafe.Pointer(result))
	}
	return nil
}

func (m *ICefDragData) HasImage() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_CefDragData_HasImage).Call(m.Instance())
	return api.GoBool(r1)
}
