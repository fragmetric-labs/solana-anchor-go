package idl

// Ref: https://github.com/solana-foundation/anchor/blob/v0.31.1/idl/spec/src/lib.rs
const IDL_SPEC = "0.31.1"

type Idl struct {
	Address      string           `json:"address"`
	Metadata     IdlMetadata      `json:"metadata"`
	Docs         []string         `json:"docs,omitempty"`
	Instructions []IdlInstruction `json:"instructions"`
	Accounts     []IdlAccount     `json:"accounts,omitempty"`
	Events       []IdlEvent       `json:"events,omitempty"`
	Errors       []IdlErrorCode   `json:"errors,omitempty"`
	Types        []IdlTypeDef     `json:"types,omitempty"`
	Constants    []IdlConst       `json:"constants,omitempty"`
}

type IdlMetadata struct {
	Name         string          `json:"name"`
	Version      string          `json:"version"`
	Spec         string          `json:"spec"`
	Description  *string         `json:"description,omitempty"`
	Repository   *string         `json:"repository,omitempty"`
	Dependencies []IdlDependency `json:"dependencies,omitempty"`
	Contact      *string         `json:"contact,omitempty"`
	Deployments  *IdlDeployments `json:"deployments,omitempty"`
}

type IdlDependency struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type IdlDeployments struct {
	Mainnet  *string `json:"mainnet,omitempty"`
	Testnet  *string `json:"testnet,omitempty"`
	Devnet   *string `json:"devnet,omitempty"`
	Localnet *string `json:"localnet,omitempty"`
}

type IdlInstruction struct {
	Name          string                      `json:"name"`
	Docs          []string                    `json:"docs,omitempty"`
	Discriminator IdlDiscriminator            `json:"discriminator"`
	Accounts      []IdlInstructionAccountItem `json:"accounts"`
	Args          []IdlField                  `json:"args"`
	Returns       IdlType                     `json:"returns,omitempty"`
}

type IdlAccount struct {
	Name          string           `json:"name"`
	Discriminator IdlDiscriminator `json:"discriminator"`
}

type IdlEvent struct {
	Name          string           `json:"name"`
	Discriminator IdlDiscriminator `json:"discriminator"`
}

type IdlConst struct {
	Name  string   `json:"name"`
	Docs  []string `json:"docs,omitempty"`
	Type  IdlType  `json:"type"`
	Value string   `json:"value"`
}

type IdlErrorCode struct {
	Code uint32  `json:"code"`
	Name string  `json:"name"`
	Msg  *string `json:"msg,omitempty"`
}

type IdlField struct {
	Name string   `json:"name"`
	Docs []string `json:"docs,omitempty"`
	Type IdlType  `json:"type"`
}

type IdlDiscriminator []byte
