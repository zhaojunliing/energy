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

import "unsafe"

// TODO no impl

// ICefNavigationEntry
//
//	/include/capi/cef_navigation_entry_capi.h (cef_navigation_entry_t)
type ICefNavigationEntry struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}
