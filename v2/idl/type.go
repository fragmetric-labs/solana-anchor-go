package idl

import (
	"encoding/json"
	"errors"
	"fmt"
)

type IdlType interface {
	isIdlType()
}

type IdlTypeSimple struct {
	// Bool,
	// U8,
	// I8,
	// U16,
	// I16,
	// U32,
	// I32,
	// F32,
	// U64,
	// I64,
	// F64,
	// U128,
	// I128,
	// U256,
	// I256,
	// Bytes,
	// String,
	// Pubkey,
	Kind string
}

func (*IdlTypeSimple) isIdlType() {}

type IdlTypeOption struct {
	Inner IdlType
}

func (*IdlTypeOption) isIdlType() {}

type IdlTypeVec struct {
	Inner IdlType
}

func (*IdlTypeVec) isIdlType() {}

type IdlTypeArray struct {
	Inner IdlType
	Len   IdlArrayLen
}

func (*IdlTypeArray) isIdlType() {}

type IdlTypeDefined struct {
	Name     string
	Generics []IdlGenericArg
}

func (*IdlTypeDefined) isIdlType() {}

type IdlTypeGeneric struct {
	Name string
}

func (*IdlTypeGeneric) isIdlType() {}

type IdlArrayLen interface {
	isIdlArrayLen()
}

type IdlArrayLenGeneric struct {
	Value string
}

func (*IdlArrayLenGeneric) isIdlArrayLen() {}

type IdlArrayLenValue struct {
	Value uint
}

func (*IdlArrayLenValue) isIdlArrayLen() {}

type IdlGenericArg interface {
	isIdlGenericArg()
}

type IdlGenericArgType struct {
	Kind string  `json:"kind"`
	Type IdlType `json:"type"`
}

func (*IdlGenericArgType) isIdlGenericArg() {}

type IdlGenericArgConst struct {
	Kind  string `json:"kind"`
	Value string `json:"value"`
}

func (*IdlGenericArgConst) isIdlGenericArg() {}

type idlTypeWrapper struct {
	Type IdlType
}

func (wrapper idlTypeWrapper) MarshalJSON() ([]byte, error) {
	switch t := wrapper.Type.(type) {
	case *IdlTypeSimple:
		return json.Marshal(t.Kind)
	case *IdlTypeOption:
		return json.Marshal(map[string]any{
			"option": idlTypeWrapper{Type: t.Inner},
		})
	case *IdlTypeVec:
		return json.Marshal(map[string]any{
			"vec": idlTypeWrapper{Type: t.Inner},
		})
	case *IdlTypeArray:
		return json.Marshal(map[string]any{
			"array": idlTypeWrapper{Type: t.Inner},
			"len":   idlArrayLenWrapper{Len: t.Len},
		})
	case *IdlTypeDefined:
		result := map[string]any{
			"defined": true,
			"name":    t.Name,
		}
		if len(t.Generics) > 0 {
			generics := make([]any, len(t.Generics))
			for i, g := range t.Generics {
				generics[i] = idlGenericArgWrapper{Arg: g}
			}
			result["generics"] = generics
		}
		return json.Marshal(result)
	case *IdlTypeGeneric:
		return json.Marshal(t.Name)
	default:
		return nil, fmt.Errorf("unknown IdlType: %T", t)
	}
}

func (wrapper *idlTypeWrapper) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		switch s {
		case "bool", "u8", "i8", "u16", "i16", "u32", "i32", "f32", "u64", "i64",
			"f64", "u128", "i128", "u256", "i256", "bytes", "string", "pubkey":
			wrapper.Type = &IdlTypeSimple{Kind: s}
			return nil
		default:
			wrapper.Type = &IdlTypeGeneric{Name: s}
			return nil
		}
	}

	var objMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &objMap); err != nil {
		return err
	}

	if opt, ok := objMap["option"]; ok {
		var inner idlTypeWrapper
		if err := json.Unmarshal(opt, &inner); err != nil {
			return err
		}
		wrapper.Type = &IdlTypeOption{Inner: inner.Type}
		return nil
	}

	if vec, ok := objMap["vec"]; ok {
		var inner idlTypeWrapper
		if err := json.Unmarshal(vec, &inner); err != nil {
			return err
		}
		wrapper.Type = &IdlTypeVec{Inner: inner.Type}
		return nil
	}

	if array, ok := objMap["array"]; ok {
		var inner idlTypeWrapper
		if err := json.Unmarshal(array, &inner); err != nil {
			return err
		}

		var arrayLen idlArrayLenWrapper
		if lenData, ok := objMap["len"]; ok {
			if err := json.Unmarshal(lenData, &arrayLen); err != nil {
				return err
			}
		} else {
			return errors.New("array type missing length")
		}

		wrapper.Type = &IdlTypeArray{
			Inner: inner.Type,
			Len:   arrayLen.Len,
		}
		return nil
	}

	if _, ok := objMap["defined"]; ok || objMap["name"] != nil {
		var name string
		if nameData, ok := objMap["name"]; ok {
			if err := json.Unmarshal(nameData, &name); err != nil {
				return err
			}
		} else {
			return errors.New("defined type missing name")
		}

		var generics []IdlGenericArg
		if genericsData, ok := objMap["generics"]; ok {
			var genericWrappers []idlGenericArgWrapper
			if err := json.Unmarshal(genericsData, &genericWrappers); err != nil {
				return err
			}
			generics = make([]IdlGenericArg, len(genericWrappers))
			for i, w := range genericWrappers {
				generics[i] = w.Arg
			}
		}

		wrapper.Type = &IdlTypeDefined{
			Name:     name,
			Generics: generics,
		}
		return nil
	}

	return errors.New("unable to unmarshal IdlType")
}

type idlArrayLenWrapper struct {
	Len IdlArrayLen
}

func (wrapper idlArrayLenWrapper) MarshalJSON() ([]byte, error) {
	switch l := wrapper.Len.(type) {
	case *IdlArrayLenValue:
		return json.Marshal(l.Value)
	case *IdlArrayLenGeneric:
		return json.Marshal(map[string]any{
			"generic": l.Value,
		})
	default:
		return nil, fmt.Errorf("unknown IdlArrayLen: %T", l)
	}
}

func (wrapper *idlArrayLenWrapper) UnmarshalJSON(data []byte) error {
	var val uint
	if err := json.Unmarshal(data, &val); err == nil {
		wrapper.Len = &IdlArrayLenValue{Value: val}
		return nil
	}

	var objMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &objMap); err != nil {
		return err
	}

	if generic, ok := objMap["generic"]; ok {
		var g string
		if err := json.Unmarshal(generic, &g); err != nil {
			return err
		}
		wrapper.Len = &IdlArrayLenGeneric{Value: g}
		return nil
	}

	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		wrapper.Len = &IdlArrayLenGeneric{Value: s}
		return nil
	}

	return errors.New("unable to unmarshal IdlArrayLen")
}

type idlGenericArgWrapper struct {
	Arg IdlGenericArg
}

func (wrapper idlGenericArgWrapper) MarshalJSON() ([]byte, error) {
	switch arg := wrapper.Arg.(type) {
	case *IdlGenericArgType:
		return json.Marshal(map[string]any{
			"kind": "type",
			"type": idlTypeWrapper{Type: arg.Type},
		})
	case *IdlGenericArgConst:
		return json.Marshal(map[string]any{
			"kind":  "const",
			"value": arg.Value,
		})
	default:
		return nil, fmt.Errorf("unknown IdlGenericArg: %T", arg)
	}
}

func (wrapper *idlGenericArgWrapper) UnmarshalJSON(data []byte) error {
	var objMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &objMap); err != nil {
		return err
	}

	var kind string
	if kindData, ok := objMap["kind"]; ok {
		if err := json.Unmarshal(kindData, &kind); err != nil {
			return err
		}
	} else {
		return errors.New("generic arg missing kind")
	}

	switch kind {
	case "type":
		var typeWrapper idlTypeWrapper
		if typeData, ok := objMap["type"]; ok {
			if err := json.Unmarshal(typeData, &typeWrapper); err != nil {
				return err
			}
		} else {
			return errors.New("type generic arg missing type")
		}

		wrapper.Arg = &IdlGenericArgType{
			Kind: "type",
			Type: typeWrapper.Type,
		}
		return nil

	case "const":
		var value string
		if valueData, ok := objMap["value"]; ok {
			if err := json.Unmarshal(valueData, &value); err != nil {
				return err
			}
		} else {
			return errors.New("const generic arg missing value")
		}

		wrapper.Arg = &IdlGenericArgConst{
			Kind:  "const",
			Value: value,
		}
		return nil

	default:
		return fmt.Errorf("unknown generic arg kind: %s", kind)
	}
}
