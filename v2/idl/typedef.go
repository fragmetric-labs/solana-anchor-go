package idl

import (
	"encoding/json"
	"errors"
	"fmt"
)

type IdlTypeDef struct {
	Name          string              `json:"name"`
	Docs          []string            `json:"docs,omitempty"`
	Serialization IdlSerialization    `json:"serialization,omitempty"`
	Repr          *IdlRepr            `json:"repr,omitempty"`
	Generics      []IdlTypeDefGeneric `json:"generics,omitempty"`
	Type          IdlTypeDefTy        `json:"type"`
}

type IdlSerialization string

const (
	IdlSerializationBorsh          IdlSerialization = "borsh"
	IdlSerializationBytemuck       IdlSerialization = "bytemuck"
	IdlSerializationBytemuckUnsafe IdlSerialization = "bytemuckunsafe"
	IdlSerializationCustom         IdlSerialization = "custom"
)

// ===== IdlRepr Types =====
type IdlRepr interface {
	isIdlRepr()
}

type IdlReprRust struct {
	Kind     string          `json:"kind"`
	Modifier IdlReprModifier `json:"modifier,omitempty"`
}

func (*IdlReprRust) isIdlRepr() {}

type IdlReprC struct {
	Kind     string          `json:"kind"`
	Modifier IdlReprModifier `json:"modifier,omitempty"`
}

func (*IdlReprC) isIdlRepr() {}

type IdlReprTransparent struct {
	Kind string `json:"kind"`
}

func (*IdlReprTransparent) isIdlRepr() {}

type IdlReprModifier struct {
	Packed bool  `json:"packed,omitempty"`
	Align  *uint `json:"align,omitempty"`
}

// ===== IdlTypeDefGeneric Types =====
type IdlTypeDefGeneric interface {
	isIdlTypeDefGeneric()
}

type IdlTypeDefGenericType struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
}

func (*IdlTypeDefGenericType) isIdlTypeDefGeneric() {}

type IdlTypeDefGenericConst struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func (*IdlTypeDefGenericConst) isIdlTypeDefGeneric() {}

// ===== IdlTypeDefTy Types =====
type IdlTypeDefTy interface {
	isIdlTypeDefTy()
}

type IdlTypeDefTyStruct struct {
	Kind   string            `json:"kind"`
	Fields *IdlDefinedFields `json:"fields,omitempty"`
}

func (*IdlTypeDefTyStruct) isIdlTypeDefTy() {}

type IdlTypeDefTyEnum struct {
	Kind     string           `json:"kind"`
	Variants []IdlEnumVariant `json:"variants"`
}

func (*IdlTypeDefTyEnum) isIdlTypeDefTy() {}

type IdlTypeDefTyType struct {
	Kind  string  `json:"kind"`
	Alias IdlType `json:"alias"`
}

func (*IdlTypeDefTyType) isIdlTypeDefTy() {}

type IdlEnumVariant struct {
	Name   string            `json:"name"`
	Fields *IdlDefinedFields `json:"fields,omitempty"`
}

// ===== IdlDefinedFields Types =====
type IdlDefinedFields interface {
	isIdlDefinedFields()
}

type IdlDefinedFieldsNamed struct {
	Fields []IdlField
}

func (*IdlDefinedFieldsNamed) isIdlDefinedFields() {}

type IdlDefinedFieldsTuple struct {
	Types []IdlType
}

func (*IdlDefinedFieldsTuple) isIdlDefinedFields() {}

// ===== Interface Wrappers =====
type idlReprWrapper struct {
	Repr IdlRepr
}

type idlTypeDefGenericWrapper struct {
	Generic IdlTypeDefGeneric
}

type idlTypeDefTyWrapper struct {
	Type IdlTypeDefTy
}

type idlDefinedFieldsWrapper struct {
	Fields IdlDefinedFields
}

func (wrapper idlReprWrapper) MarshalJSON() ([]byte, error) {
	switch repr := wrapper.Repr.(type) {
	case *IdlReprRust:
		return json.Marshal(repr)
	case *IdlReprC:
		return json.Marshal(repr)
	case *IdlReprTransparent:
		return json.Marshal(repr)
	default:
		return nil, fmt.Errorf("unknown IdlRepr: %T", repr)
	}
}

func (wrapper *idlReprWrapper) UnmarshalJSON(data []byte) error {
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
		return errors.New("repr missing kind")
	}

	switch kind {
	case "rust":
		var repr IdlReprRust
		if err := json.Unmarshal(data, &repr); err != nil {
			return err
		}
		wrapper.Repr = &repr
	case "c":
		var repr IdlReprC
		if err := json.Unmarshal(data, &repr); err != nil {
			return err
		}
		wrapper.Repr = &repr
	case "transparent":
		var repr IdlReprTransparent
		if err := json.Unmarshal(data, &repr); err != nil {
			return err
		}
		wrapper.Repr = &repr
	default:
		return fmt.Errorf("unknown repr kind: %s", kind)
	}

	return nil
}

func (wrapper idlTypeDefGenericWrapper) MarshalJSON() ([]byte, error) {
	switch g := wrapper.Generic.(type) {
	case *IdlTypeDefGenericType:
		return json.Marshal(map[string]any{
			"kind": "type",
			"name": g.Name,
		})
	case *IdlTypeDefGenericConst:
		return json.Marshal(map[string]any{
			"kind": "const",
			"name": g.Name,
			"type": g.Type,
		})
	default:
		return nil, fmt.Errorf("unknown IdlTypeDefGeneric: %T", g)
	}
}

func (wrapper *idlTypeDefGenericWrapper) UnmarshalJSON(data []byte) error {
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
		return errors.New("typedef generic missing kind")
	}

	switch kind {
	case "type":
		var name string
		if nameData, ok := objMap["name"]; ok {
			if err := json.Unmarshal(nameData, &name); err != nil {
				return err
			}
		} else {
			return errors.New("type generic missing name")
		}

		wrapper.Generic = &IdlTypeDefGenericType{
			Kind: "type",
			Name: name,
		}
		return nil

	case "const":
		var name string
		if nameData, ok := objMap["name"]; ok {
			if err := json.Unmarshal(nameData, &name); err != nil {
				return err
			}
		} else {
			return errors.New("const generic missing name")
		}

		var typeStr string
		if typeData, ok := objMap["type"]; ok {
			if err := json.Unmarshal(typeData, &typeStr); err != nil {
				return err
			}
		} else {
			return errors.New("const generic missing type")
		}

		wrapper.Generic = &IdlTypeDefGenericConst{
			Kind: "const",
			Name: name,
			Type: typeStr,
		}
		return nil

	default:
		return fmt.Errorf("unknown typedef generic kind: %s", kind)
	}
}

func (wrapper idlTypeDefTyWrapper) MarshalJSON() ([]byte, error) {
	switch ty := wrapper.Type.(type) {
	case *IdlTypeDefTyStruct:
		return json.Marshal(ty)
	case *IdlTypeDefTyEnum:
		return json.Marshal(ty)
	case *IdlTypeDefTyType:
		return json.Marshal(ty)
	default:
		return nil, fmt.Errorf("unknown IdlTypeDefTy: %T", ty)
	}
}

func (wrapper *idlTypeDefTyWrapper) UnmarshalJSON(data []byte) error {
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
		return errors.New("typedef type missing kind")
	}

	switch kind {
	case "struct":
		var ty IdlTypeDefTyStruct
		if err := json.Unmarshal(data, &ty); err != nil {
			return err
		}
		wrapper.Type = &ty
	case "enum":
		var ty IdlTypeDefTyEnum
		if err := json.Unmarshal(data, &ty); err != nil {
			return err
		}
		wrapper.Type = &ty
	case "type":
		var ty IdlTypeDefTyType
		if err := json.Unmarshal(data, &ty); err != nil {
			return err
		}
		wrapper.Type = &ty
	default:
		return fmt.Errorf("unknown typedef type kind: %s", kind)
	}

	return nil
}

func (wrapper idlDefinedFieldsWrapper) MarshalJSON() ([]byte, error) {
	switch fields := wrapper.Fields.(type) {
	case *IdlDefinedFieldsNamed:
		return json.Marshal(fields.Fields)
	case *IdlDefinedFieldsTuple:
		return json.Marshal(fields.Types)
	default:
		return nil, fmt.Errorf("unknown IdlDefinedFields: %T", fields)
	}
}

func (wrapper *idlDefinedFieldsWrapper) UnmarshalJSON(data []byte) error {
	var named []IdlField
	if err := json.Unmarshal(data, &named); err == nil && len(named) > 0 && named[0].Name != "" {
		wrapper.Fields = &IdlDefinedFieldsNamed{Fields: named}
		return nil
	}

	var tupleWrappers []idlTypeWrapper
	if err := json.Unmarshal(data, &tupleWrappers); err == nil {
		types := make([]IdlType, len(tupleWrappers))
		for i, w := range tupleWrappers {
			types[i] = w.Type
		}
		wrapper.Fields = &IdlDefinedFieldsTuple{Types: types}
		return nil
	}

	return errors.New("unable to unmarshal defined fields")
}
