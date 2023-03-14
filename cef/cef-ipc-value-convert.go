//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// ipcValueConvert -> v8ValueProcessMessageConvert
//
// IPC 和 ICefV8Value 数据序列化转换
package cef

import (
	"errors"
	"github.com/energye/energy/consts"
	"github.com/energye/energy/pkgs/json"
	jsoniter "github.com/json-iterator/go"
)

// ipcValueConvert
var ipcValueConvert v8ValueProcessMessageConvert

// v8ValueProcessMessageConvert ICefV8Value 和 ICefProcessMessage 转换
type v8ValueProcessMessageConvert uintptr

//ListValueToV8Value ICefListValue 转换 ICefV8Value
func (m *v8ValueProcessMessageConvert) ListValueToV8Value(list *ICefListValue) (*ICefV8Value, error) {
	if list == nil {
		return nil, errors.New("build v8 value error. Parameter null")
	}
	size := int(list.Size())
	result := V8ValueRef.NewArray(int32(size))
	for i := 0; i < size; i++ {
		value := list.GetValue(uint32(i))
		var newValue *ICefV8Value
		switch value.GetType() {
		case consts.VTYPE_NULL:
			newValue = V8ValueRef.NewNull()
		case consts.VTYPE_BOOL:
			newValue = V8ValueRef.NewBool(value.GetBool())
		case consts.VTYPE_INT:
			newValue = V8ValueRef.NewInt(value.GetInt())
		case consts.VTYPE_DOUBLE:
			newValue = V8ValueRef.NewDouble(value.GetDouble())
		case consts.VTYPE_STRING:
			newValue = V8ValueRef.NewString(value.GetString())
		case consts.VTYPE_BINARY: // []byte
			binaryValue := value.GetBinary()
			byteSize := binaryValue.GetSize()
			if byteSize > 0 {
				dataByte := make([]byte, binaryValue.GetSize())
				if c := binaryValue.GetData(dataByte, 0); c > 0 {
					newValue = V8ValueRef.NewArrayBuffer(dataByte, nil)
				}
			}
		case consts.VTYPE_DICTIONARY: // Object
			if v, err := m.DictionaryValueToV8Value(value.GetDictionary()); err == nil {
				newValue = v
			}
		case consts.VTYPE_LIST: // JSONArray
			if v, err := m.ListValueToV8Value(value.GetList()); err == nil {
				newValue = v
			}
		}
		if newValue == nil {
			newValue = V8ValueRef.NewNull()
		}
		result.SetValueByIndex(int32(i), newValue)
	}
	return result, nil
}

//DictionaryValueToV8Value ICefDictionaryValue 转换 ICefV8Value
func (m *v8ValueProcessMessageConvert) DictionaryValueToV8Value(dictionary *ICefDictionaryValue) (*ICefV8Value, error) {
	if dictionary == nil {
		return nil, errors.New("build v8 value error. Parameter null")
	}
	keys := dictionary.GetKeys()
	//bindSubObjectAccessor := V8AccessorRef.New()
	//bindSubObjectAccessor.Get(ipcRender.bindSubObjectGet)
	//bindSubObjectAccessor.Set(ipcRender.bindSubObjectSet)
	result := V8ValueRef.NewObject(nil)
	for i := 0; i < keys.Count(); i++ {
		key := keys.Get(i)
		value := dictionary.GetValue(key)
		var newValue *ICefV8Value
		switch value.GetType() {
		case consts.VTYPE_NULL:
			newValue = V8ValueRef.NewNull()
		case consts.VTYPE_BOOL:
			newValue = V8ValueRef.NewBool(value.GetBool())
		case consts.VTYPE_INT:
			newValue = V8ValueRef.NewInt(value.GetInt())
		case consts.VTYPE_DOUBLE:
			newValue = V8ValueRef.NewDouble(value.GetDouble())
		case consts.VTYPE_STRING:
			newValue = V8ValueRef.NewString(value.GetString())
		case consts.VTYPE_BINARY: // []byte
			binaryValue := value.GetBinary()
			byteSize := binaryValue.GetSize()
			if byteSize > 0 {
				dataByte := make([]byte, binaryValue.GetSize())
				if c := binaryValue.GetData(dataByte, 0); c > 0 {
					newValue = V8ValueRef.NewArrayBuffer(dataByte, nil)
				}
			}
		case consts.VTYPE_DICTIONARY: // Object
			if v, err := m.DictionaryValueToV8Value(value.GetDictionary()); err == nil {
				newValue = v
			}
		case consts.VTYPE_LIST: // JSONArray
			if v, err := m.ListValueToV8Value(value.GetList()); err == nil {
				newValue = v
			}
		}
		if newValue != nil {
			result.setValueByAccessor(key, consts.V8_ACCESS_CONTROL_DEFAULT, consts.V8_PROPERTY_ATTRIBUTE_NONE)
			result.setValueByKey(key, newValue, consts.V8_PROPERTY_ATTRIBUTE_NONE)
		}
	}
	return result, nil
}

// BytesToV8ArrayValue JSONArray 字节数组转换 TCefV8ValueArray
func (m *v8ValueProcessMessageConvert) BytesToV8ArrayValue(resultArgsBytes []byte) (*TCefV8ValueArray, error) {
	//只能是 JSONArray 对象类型
	jsonArray := json.NewJSONArray(resultArgsBytes)
	if jsonArray == nil {
		return nil, errors.New("parsing parameter failure")
	}
	resultArgs := V8ValueArrayRef.New()
	size := jsonArray.Size()
	for i := 0; i < size; i++ {
		value := jsonArray.GetByIndex(i)
		switch value.Type() {
		case consts.GO_VALUE_STRING:
			resultArgs.Add(V8ValueRef.NewString(value.String()))
		case consts.GO_VALUE_INT:
			resultArgs.Add(V8ValueRef.NewInt(int32(value.Int())))
		case consts.GO_VALUE_UINT:
			resultArgs.Add(V8ValueRef.NewUInt(uint32(value.UInt())))
		case consts.GO_VALUE_FLOAT64:
			resultArgs.Add(V8ValueRef.NewDouble(value.Float()))
		case consts.GO_VALUE_BOOL:
			resultArgs.Add(V8ValueRef.NewBool(value.Bool()))
		case consts.GO_VALUE_SLICE:
			if v := m.JSONArrayToV8Value(value.JSONArray()); v != nil {
				resultArgs.Add(v)
			} else {
				resultArgs.Add(V8ValueRef.NewNull())
			}
		case consts.GO_VALUE_MAP:
			if v := m.JSONObjectToV8Value(value.JSONObject()); v != nil {
				resultArgs.Add(v)
			} else {
				resultArgs.Add(V8ValueRef.NewNull())
			}
		}
	}
	jsonArray.Free()
	return resultArgs, nil
}

// JSONArrayToV8Value JSONArray 转 ICefV8Value
func (m *v8ValueProcessMessageConvert) JSONArrayToV8Value(array json.JSONArray) *ICefV8Value {
	if array == nil || !array.IsArray() {
		return nil
	}
	size := array.Size()
	result := V8ValueRef.NewArray(int32(size))
	for i := 0; i < size; i++ {
		value := array.GetByIndex(i)
		if value == nil {
			result.SetValueByIndex(int32(i), V8ValueRef.NewNull())
			continue
		}
		switch value.Type() {
		case consts.GO_VALUE_STRING:
			result.SetValueByIndex(int32(i), V8ValueRef.NewString(value.String()))
		case consts.GO_VALUE_INT:
			result.SetValueByIndex(int32(i), V8ValueRef.NewInt(int32(value.Int())))
		case consts.GO_VALUE_UINT:
			result.SetValueByIndex(int32(i), V8ValueRef.NewUInt(uint32(value.UInt())))
		case consts.GO_VALUE_FLOAT64:
			result.SetValueByIndex(int32(i), V8ValueRef.NewDouble(value.Float()))
		case consts.GO_VALUE_BOOL:
			result.SetValueByIndex(int32(i), V8ValueRef.NewBool(value.Bool()))
		case consts.GO_VALUE_SLICE:
			if v := m.JSONArrayToV8Value(value); v != nil {
				result.SetValueByIndex(int32(i), v)
			} else {
				result.SetValueByIndex(int32(i), V8ValueRef.NewNull())
			}
		case consts.GO_VALUE_MAP:
			if v := m.JSONObjectToV8Value(value.JSONObject()); v != nil {
				result.SetValueByIndex(int32(i), v)
			} else {
				result.SetValueByIndex(int32(i), V8ValueRef.NewNull())
			}
		default:
			result.SetValueByIndex(int32(i), V8ValueRef.NewNull())
		}
	}
	return result
}

// JSONObjectToV8Value JSONObject 转 ICefV8Value
func (m *v8ValueProcessMessageConvert) JSONObjectToV8Value(object json.JSONObject) *ICefV8Value {
	if object == nil || !object.IsObject() {
		return nil
	}
	size := object.Size()
	result := V8ValueRef.NewObject(nil)
	keys := object.Keys()
	for i := 0; i < size; i++ {
		key := keys[i]
		value := object.GetByKey(key)
		if value == nil {
			result.setValueByKey(key, V8ValueRef.NewNull(), consts.V8_PROPERTY_ATTRIBUTE_NONE)
			continue
		}
		switch value.Type() {
		case consts.GO_VALUE_STRING:
			result.setValueByKey(key, V8ValueRef.NewString(value.String()), consts.V8_PROPERTY_ATTRIBUTE_NONE)
		case consts.GO_VALUE_INT:
			result.setValueByKey(key, V8ValueRef.NewInt(int32(value.Int())), consts.V8_PROPERTY_ATTRIBUTE_NONE)
		case consts.GO_VALUE_UINT:
			result.setValueByKey(key, V8ValueRef.NewUInt(uint32(value.UInt())), consts.V8_PROPERTY_ATTRIBUTE_NONE)
		case consts.GO_VALUE_FLOAT64:
			result.setValueByKey(key, V8ValueRef.NewDouble(value.Float()), consts.V8_PROPERTY_ATTRIBUTE_NONE)
		case consts.GO_VALUE_BOOL:
			result.setValueByKey(key, V8ValueRef.NewBool(value.Bool()), consts.V8_PROPERTY_ATTRIBUTE_NONE)
		case consts.GO_VALUE_SLICE:
			if v := m.JSONArrayToV8Value(value.JSONArray()); v != nil {
				result.setValueByKey(key, v, consts.V8_PROPERTY_ATTRIBUTE_NONE)
			} else {
				result.setValueByKey(key, V8ValueRef.NewNull(), consts.V8_PROPERTY_ATTRIBUTE_NONE)
			}
		case consts.GO_VALUE_MAP:
			if v := m.JSONObjectToV8Value(value); v != nil {
				result.setValueByKey(key, v, consts.V8_PROPERTY_ATTRIBUTE_NONE)
			} else {
				result.setValueByKey(key, V8ValueRef.NewNull(), consts.V8_PROPERTY_ATTRIBUTE_NONE)
			}
		default:
			result.setValueByKey(key, V8ValueRef.NewNull(), consts.V8_PROPERTY_ATTRIBUTE_NONE)
		}
	}
	return result
}

// V8ValueToProcessMessageBytes ICefV8Value 转换 []byte 进程消息
func (m *v8ValueProcessMessageConvert) V8ValueToProcessMessageBytes(v8value *ICefV8Value) []byte {
	if result, err := m.V8valueArrayToSlice(v8value); err == nil {
		if v, err := jsoniter.Marshal(result); err == nil {
			return v
		}
	}
	return nil
}

// V8valueArrayToSlice ICefV8Value 转换 Slice
func (m *v8ValueProcessMessageConvert) V8valueArrayToSlice(v8value *ICefV8Value) ([]any, error) {
	if !v8value.IsArray() {
		return nil, errors.New("convert list value error. Please pass in the array type")
	}
	argsLen := v8value.GetArrayLength()
	result := make([]any, argsLen)
	for i := 0; i < argsLen; i++ {
		args := v8value.GetValueByIndex(i)
		if args.IsString() {
			result[i] = args.GetStringValue()
		} else if args.IsInt() {
			result[i] = int(args.GetIntValue())
		} else if args.IsUInt() {
			result[i] = uint(args.GetUIntValue())
		} else if args.IsDouble() {
			result[i] = args.GetDoubleValue()
		} else if args.IsBool() {
			result[i] = args.GetBoolValue()
		} else if args.IsNull() {
			result[i] = "null"
		} else if args.IsUndefined() {
			result[i] = "undefined"
		} else if args.IsArray() {
			if v, err := m.V8valueArrayToSlice(args); err == nil {
				result[i] = v
			} else {
				result[i] = nil
			}
		} else if args.IsObject() {
			if v, err := m.V8valueObjectToMap(args); err == nil {
				result[i] = v
			} else {
				result[i] = nil
			}
		} else if args.IsArrayBuffer() {
			result[i] = ""
		} else {
			result[i] = ""
		}
		args.Free()
	}
	return result, nil
}

// V8valueObjectToMap ICefV8Value 转换 Maps
func (m *v8ValueProcessMessageConvert) V8valueObjectToMap(v8value *ICefV8Value) (map[string]any, error) {
	if !v8value.IsObject() {
		return nil, errors.New("convert dictionary value error. Please pass in the object type")
	}
	keys := v8value.GetKeys()
	result := make(map[string]any, keys.Count())
	for i := 0; i < keys.Count(); i++ {
		key := keys.Get(i)
		args := v8value.getValueByKey(key)
		if args.IsString() {
			result[key] = args.GetStringValue()
		} else if args.IsInt() {
			result[key] = int(args.GetIntValue())
		} else if args.IsUInt() {
			result[key] = uint(args.GetUIntValue())
		} else if args.IsDouble() {
			result[key] = args.GetDoubleValue()
		} else if args.IsBool() {
			result[key] = args.GetBoolValue()
		} else if args.IsNull() {
			result[key] = "null"
		} else if args.IsUndefined() {
			result[key] = "undefined"
		} else if args.IsArray() {
			if v, err := m.V8valueArrayToSlice(args); err == nil {
				result[key] = v
			} else {
				result[key] = nil
			}
		} else if args.IsObject() {
			if v, err := m.V8valueObjectToMap(args); err == nil {
				result[key] = v
			} else {
				result[key] = nil
			}
		} else if args.IsArrayBuffer() {
			//arrayValue.SetBinary()
			result[key] = ""
		} else {
			result[key] = ""
		}
		args.Free()
	}
	keys.Free()
	return result, nil
}

// V8ValueToProcessMessage ICefV8Value 转换 进程消息
func (m *v8ValueProcessMessageConvert) V8ValueToProcessMessage(v8value *ICefV8Value) (*ICefListValue, error) {
	if v8value == nil {
		return nil, errors.New("build process value error. Parameter null")
	}
	if v8value.IsArray() {
		return m.V8valueArrayToListValue(v8value)
	} else if v8value.IsObject() {
		if v, err := m.V8valueObjectToDictionaryValue(v8value); err == nil {
			arrayValue := ListValueRef.New()
			arrayValue.SetDictionary(uint32(0), v)
			return arrayValue, nil
		} else {
			return nil, err
		}
	} else {
		arrayValue := ListValueRef.New()
		if v8value.IsString() {
			arrayValue.SetString(uint32(0), v8value.GetStringValue())
		} else if v8value.IsInt() {
			arrayValue.SetInt(uint32(0), v8value.GetIntValue())
		} else if v8value.IsUInt() {
			arrayValue.SetInt(uint32(0), int32(v8value.GetUIntValue()))
		} else if v8value.IsDouble() {
			arrayValue.SetDouble(uint32(0), v8value.GetDoubleValue())
		} else if v8value.IsBool() {
			arrayValue.SetBool(uint32(0), v8value.GetBoolValue())
		} else if v8value.IsArrayBuffer() {
			//arrayValue.SetBinary()
		} else {
			arrayValue.SetNull(uint32(0))
		}
		return arrayValue, nil
	}
}

// V8valueArrayToListValue ICefV8Value 转换 ICefListValue
func (m *v8ValueProcessMessageConvert) V8valueArrayToListValue(v8value *ICefV8Value) (*ICefListValue, error) {
	if !v8value.IsArray() {
		return nil, errors.New("convert list value error. Please pass in the array type")
	}
	arrayValue := ListValueRef.New()
	argsLen := v8value.GetArrayLength()
	for i := 0; i < argsLen; i++ {
		args := v8value.GetValueByIndex(i)
		if args.IsString() {
			arrayValue.SetString(uint32(i), args.GetStringValue())
		} else if args.IsInt() {
			arrayValue.SetInt(uint32(i), args.GetIntValue())
		} else if args.IsUInt() {
			arrayValue.SetInt(uint32(i), int32(args.GetUIntValue()))
		} else if args.IsDouble() {
			arrayValue.SetDouble(uint32(i), args.GetDoubleValue())
		} else if args.IsBool() {
			arrayValue.SetBool(uint32(i), args.GetBoolValue())
		} else if args.IsNull() || args.IsUndefined() {
			arrayValue.SetNull(uint32(i))
		} else if args.IsArray() {
			if v, err := m.V8valueArrayToListValue(args); err == nil {
				arrayValue.SetList(uint32(i), v)
			}
		} else if args.IsObject() {
			if v, err := m.V8valueObjectToDictionaryValue(args); err == nil {
				arrayValue.SetDictionary(uint32(i), v)
			}
		} else if args.IsArrayBuffer() {
			//arrayValue.SetBinary()
		} else {
			arrayValue.SetNull(uint32(i))
		}
		args.Free()
	}
	return arrayValue, nil
}

// V8valueObjectToDictionaryValue ICefV8Value 转换 ICefDictionaryValue
func (m *v8ValueProcessMessageConvert) V8valueObjectToDictionaryValue(v8value *ICefV8Value) (*ICefDictionaryValue, error) {
	if !v8value.IsObject() {
		return nil, errors.New("convert dictionary value error. Please pass in the object type")
	}
	dictionaryValue := DictionaryValueRef.New()
	keys := v8value.GetKeys()
	for i := 0; i < keys.Count(); i++ {
		key := keys.Get(i)
		args := v8value.getValueByKey(key)
		if args.IsString() {
			dictionaryValue.SetString(key, args.GetStringValue())
		} else if args.IsInt() {
			dictionaryValue.SetInt(key, args.GetIntValue())
		} else if args.IsUInt() {
			dictionaryValue.SetInt(key, int32(args.GetUIntValue()))
		} else if args.IsDouble() {
			dictionaryValue.SetDouble(key, args.GetDoubleValue())
		} else if args.IsBool() {
			dictionaryValue.SetBool(key, args.GetBoolValue())
		} else if args.IsNull() || args.IsUndefined() {
			dictionaryValue.SetNull(key)
		} else if args.IsArray() {
			if v, err := m.V8valueArrayToListValue(args); err == nil {
				dictionaryValue.SetList(key, v)
			}
		} else if args.IsObject() {
			if v, err := m.V8valueObjectToDictionaryValue(args); err == nil {
				dictionaryValue.SetDictionary(key, v)
			}
		} else {
			dictionaryValue.SetNull(key)
		}
		args.Free()
	}
	keys.Free()
	return dictionaryValue, nil
}
