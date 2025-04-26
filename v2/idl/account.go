package idl

import (
	"encoding/json"
	"errors"
	"fmt"
)

type IdlInstructionAccountItem interface {
	isIdlInstructionAccountItem()
}

type IdlInstructionAccount struct {
	Name      string   `json:"name"`
	Docs      []string `json:"docs,omitempty"`
	Writable  bool     `json:"writable,omitempty"`
	Signer    bool     `json:"signer,omitempty"`
	Optional  bool     `json:"optional,omitempty"`
	Address   *string  `json:"address,omitempty"`
	Pda       *IdlPda  `json:"pda,omitempty"`
	Relations []string `json:"relations,omitempty"`
}

func (*IdlInstructionAccount) isIdlInstructionAccountItem() {}

type IdlInstructionAccounts struct {
	Name     string                      `json:"name"`
	Accounts []IdlInstructionAccountItem `json:"accounts"`
}

func (*IdlInstructionAccounts) isIdlInstructionAccountItem() {}

type IdlPda struct {
	Seeds   []IdlSeed `json:"seeds"`
	Program *IdlSeed  `json:"program,omitempty"`
}

type IdlSeed interface {
	isIdlSeed()
}

type IdlSeedConst struct {
	Kind  string `json:"kind"`
	Value []byte `json:"value"`
}

func (*IdlSeedConst) isIdlSeed() {}

type IdlSeedArg struct {
	Kind string `json:"kind"`
	Path string `json:"path"`
}

func (*IdlSeedArg) isIdlSeed() {}

type IdlSeedAccount struct {
	Kind    string  `json:"kind"`
	Path    string  `json:"path"`
	Account *string `json:"account,omitempty"`
}

func (*IdlSeedAccount) isIdlSeed() {}

type idlInstructionAccountItemWrapper struct {
	Item IdlInstructionAccountItem
}

type idlSeedWrapper struct {
	Seed IdlSeed
}

func (wrapper idlInstructionAccountItemWrapper) MarshalJSON() ([]byte, error) {
	switch item := wrapper.Item.(type) {
	case *IdlInstructionAccount:
		return json.Marshal(item)
	case *IdlInstructionAccounts:
		return json.Marshal(item)
	default:
		return nil, fmt.Errorf("unknown IdlInstructionAccountItem: %T", item)
	}
}

func (wrapper *idlInstructionAccountItemWrapper) UnmarshalJSON(data []byte) error {
	var objMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &objMap); err != nil {
		return err
	}

	if _, hasAccounts := objMap["accounts"]; hasAccounts {
		var accounts IdlInstructionAccounts
		if err := json.Unmarshal(data, &accounts); err != nil {
			return err
		}
		wrapper.Item = &accounts
	} else {
		var account IdlInstructionAccount
		if err := json.Unmarshal(data, &account); err != nil {
			return err
		}
		wrapper.Item = &account
	}

	return nil
}

func (wrapper idlSeedWrapper) MarshalJSON() ([]byte, error) {
	switch seed := wrapper.Seed.(type) {
	case *IdlSeedConst:
		return json.Marshal(seed)
	case *IdlSeedArg:
		return json.Marshal(seed)
	case *IdlSeedAccount:
		return json.Marshal(seed)
	default:
		return nil, fmt.Errorf("unknown IdlSeed: %T", seed)
	}
}

func (wrapper *idlSeedWrapper) UnmarshalJSON(data []byte) error {
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
		return errors.New("seed missing kind")
	}

	switch kind {
	case "const":
		var seed IdlSeedConst
		if err := json.Unmarshal(data, &seed); err != nil {
			return err
		}
		wrapper.Seed = &seed
	case "arg":
		var seed IdlSeedArg
		if err := json.Unmarshal(data, &seed); err != nil {
			return err
		}
		wrapper.Seed = &seed
	case "account":
		var seed IdlSeedAccount
		if err := json.Unmarshal(data, &seed); err != nil {
			return err
		}
		wrapper.Seed = &seed
	default:
		return fmt.Errorf("unknown seed kind: %s", kind)
	}

	return nil
}
