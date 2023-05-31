//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// V8 JSValue JSDouble 类型实现

package bind

type JSDouble interface {
	JSValue
	AsDouble() JSDouble
	DoubleValue() float64
}

type jsDouble struct {
	V8Value
}

func (m *jsDouble) AsDouble() JSDouble {
	if m.IsDouble() {
		return m
	}
	return nil
}

func (m *jsDouble) DoubleValue() float64 {
	return m.JsonData.Float()
}
