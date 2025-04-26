package idl

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

func ParseIdlType(s string) (IdlType, error) {
	s = strings.TrimSpace(s)
	switch s {
	case "bool":
		return &IdlTypeSimple{Kind: "bool"}, nil
	case "u8":
		return &IdlTypeSimple{Kind: "u8"}, nil
	case "i8":
		return &IdlTypeSimple{Kind: "i8"}, nil
	case "u16":
		return &IdlTypeSimple{Kind: "u16"}, nil
	case "i16":
		return &IdlTypeSimple{Kind: "i16"}, nil
	case "u32":
		return &IdlTypeSimple{Kind: "u32"}, nil
	case "i32":
		return &IdlTypeSimple{Kind: "i32"}, nil
	case "f32":
		return &IdlTypeSimple{Kind: "f32"}, nil
	case "u64":
		return &IdlTypeSimple{Kind: "u64"}, nil
	case "i64":
		return &IdlTypeSimple{Kind: "i64"}, nil
	case "f64":
		return &IdlTypeSimple{Kind: "f64"}, nil
	case "u128":
		return &IdlTypeSimple{Kind: "u128"}, nil
	case "i128":
		return &IdlTypeSimple{Kind: "i128"}, nil
	case "u256":
		return &IdlTypeSimple{Kind: "u256"}, nil
	case "i256":
		return &IdlTypeSimple{Kind: "i256"}, nil
	case "Vec<u8>":
		return &IdlTypeSimple{Kind: "bytes"}, nil
	case "String", "&str", "&'staticstr":
		return &IdlTypeSimple{Kind: "string"}, nil
	case "Pubkey":
		return &IdlTypeSimple{Kind: "pubkey"}, nil
	}

	// Option<T>
	if strings.HasPrefix(s, "Option<") && strings.HasSuffix(s, ">") {
		inner, err := ParseIdlType(s[7 : len(s)-1])
		if err != nil {
			return nil, errors.New("invalid Option type: " + err.Error())
		}
		return &IdlTypeOption{Inner: inner}, nil
	}

	// Vec<T>
	if strings.HasPrefix(s, "Vec<") && strings.HasSuffix(s, ">") {
		inner, err := ParseIdlType(s[4 : len(s)-1])
		if err != nil {
			return nil, errors.New("invalid Vec type: " + err.Error())
		}
		return &IdlTypeVec{Inner: inner}, nil
	}

	// Array [T; len]
	if strings.HasPrefix(s, "[") {
		var parseArrayFn func(string) (IdlType, error)
		parseArrayFn = func(inner string) (IdlType, error) {
			if strings.HasSuffix(inner, "]") {
				nestedInner := inner[1 : len(inner)-1]
				return parseArrayFn(nestedInner)
			} else {
				parts := strings.Split(inner, ";")
				if len(parts) != 2 {
					return nil, errors.New("invalid array format")
				}

				elemType, err := ParseIdlType(parts[0])
				if err != nil {
					return nil, err
				}

				lenStr := strings.ReplaceAll(parts[1], "_", "")
				var arrayLen IdlArrayLen

				if val, err := strconv.ParseUint(lenStr, 10, 64); err == nil {
					arrayLen = &IdlArrayLenValue{Value: uint(val)}
				} else {
					arrayLen = &IdlArrayLenGeneric{Value: lenStr}
				}

				return &IdlTypeArray{
					Inner: elemType,
					Len:   arrayLen,
				}, nil
			}
		}

		if len(s) < 2 || !strings.HasSuffix(s, "]") {
			return nil, errors.New("invalid array format")
		}
		return parseArrayFn(s[1 : len(s)-1])
	}

	if idx := strings.Index(s, "<"); idx != -1 {
		name := s[:idx]
		genericsStr := s[idx+1 : len(s)-1]

		genericParts := strings.Split(genericsStr, ",")
		generics := make([]IdlGenericArg, 0, len(genericParts))

		for _, g := range genericParts {
			g = strings.TrimSpace(g)

			_, errBool := strconv.ParseBool(g)
			_, errUint := strconv.ParseUint(g, 10, 128)
			_, errInt := strconv.ParseInt(g, 10, 128)

			if errBool == nil || errUint == nil || errInt == nil {
				generics = append(generics, &IdlGenericArgConst{
					Kind:  "const",
					Value: g,
				})
				continue
			}

			gType, err := ParseIdlType(g)
			if err != nil {
				return nil, err
			}

			generics = append(generics, &IdlGenericArgType{
				Kind: "type",
				Type: gType,
			})
		}

		return &IdlTypeDefined{
			Name:     name,
			Generics: generics,
		}, nil
	}

	return &IdlTypeDefined{
		Name:     s,
		Generics: []IdlGenericArg{},
	}, nil
}

func MarshalIdlType(t IdlType) ([]byte, error) {
	return json.Marshal(idlTypeWrapper{Type: t})
}

func UnmarshalIdlType(data []byte) (IdlType, error) {
	var wrapper idlTypeWrapper
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, err
	}
	return wrapper.Type, nil
}

func MarshalInstructionAccountItem(item IdlInstructionAccountItem) ([]byte, error) {
	return json.Marshal(idlInstructionAccountItemWrapper{Item: item})
}

func UnmarshalInstructionAccountItem(data []byte) (IdlInstructionAccountItem, error) {
	var wrapper idlInstructionAccountItemWrapper
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, err
	}
	return wrapper.Item, nil
}
