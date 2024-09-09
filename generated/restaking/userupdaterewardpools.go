// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UserUpdateRewardPools is the `user_update_reward_pools` instruction.
type UserUpdateRewardPools struct {

	// [0] = [WRITE, SIGNER] user
	//
	// [1] = [] system_program
	//
	// [2] = [] receipt_token_mint
	//
	// [3] = [WRITE] reward_account
	//
	// [4] = [WRITE] user_reward_account
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUserUpdateRewardPoolsInstructionBuilder creates a new `UserUpdateRewardPools` instruction builder.
func NewUserUpdateRewardPoolsInstructionBuilder() *UserUpdateRewardPools {
	nd := &UserUpdateRewardPools{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	nd.AccountMetaSlice[1] = ag_solanago.Meta(Addresses["11111111111111111111111111111111"])
	nd.AccountMetaSlice[2] = ag_solanago.Meta(Addresses["FRAGSEthVFL7fdqM8hxfxkfCZzUvmg21cqPJVvC1qdbo"])
	return nd
}

// SetUserAccount sets the "user" account.
func (inst *UserUpdateRewardPools) SetUserAccount(user ag_solanago.PublicKey) *UserUpdateRewardPools {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(user).WRITE().SIGNER()
	return inst
}

// GetUserAccount gets the "user" account.
func (inst *UserUpdateRewardPools) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *UserUpdateRewardPools) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *UserUpdateRewardPools {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *UserUpdateRewardPools) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetReceiptTokenMintAccount sets the "receipt_token_mint" account.
func (inst *UserUpdateRewardPools) SetReceiptTokenMintAccount(receiptTokenMint ag_solanago.PublicKey) *UserUpdateRewardPools {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(receiptTokenMint)
	return inst
}

// GetReceiptTokenMintAccount gets the "receipt_token_mint" account.
func (inst *UserUpdateRewardPools) GetReceiptTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetRewardAccountAccount sets the "reward_account" account.
func (inst *UserUpdateRewardPools) SetRewardAccountAccount(rewardAccount ag_solanago.PublicKey) *UserUpdateRewardPools {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(rewardAccount).WRITE()
	return inst
}

func (inst *UserUpdateRewardPools) findFindRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: reward
	seeds = append(seeds, []byte{byte(0x72), byte(0x65), byte(0x77), byte(0x61), byte(0x72), byte(0x64)})
	// path: receiptTokenMint
	seeds = append(seeds, receiptTokenMint.Bytes())

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindRewardAccountAddressWithBumpSeed calculates RewardAccount account address with given seeds and a known bump seed.
func (inst *UserUpdateRewardPools) FindRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindRewardAccountAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *UserUpdateRewardPools) MustFindRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindRewardAccountAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindRewardAccountAddress finds RewardAccount account address with given seeds.
func (inst *UserUpdateRewardPools) FindRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindRewardAccountAddress(receiptTokenMint, 0)
	return
}

func (inst *UserUpdateRewardPools) MustFindRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindRewardAccountAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetRewardAccountAccount gets the "reward_account" account.
func (inst *UserUpdateRewardPools) GetRewardAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetUserRewardAccountAccount sets the "user_reward_account" account.
func (inst *UserUpdateRewardPools) SetUserRewardAccountAccount(userRewardAccount ag_solanago.PublicKey) *UserUpdateRewardPools {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(userRewardAccount).WRITE()
	return inst
}

func (inst *UserUpdateRewardPools) findFindUserRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: user_reward
	seeds = append(seeds, []byte{byte(0x75), byte(0x73), byte(0x65), byte(0x72), byte(0x5f), byte(0x72), byte(0x65), byte(0x77), byte(0x61), byte(0x72), byte(0x64)})
	// path: receiptTokenMint
	seeds = append(seeds, receiptTokenMint.Bytes())
	// path: user
	seeds = append(seeds, user.Bytes())

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindUserRewardAccountAddressWithBumpSeed calculates UserRewardAccount account address with given seeds and a known bump seed.
func (inst *UserUpdateRewardPools) FindUserRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindUserRewardAccountAddress(receiptTokenMint, user, bumpSeed)
	return
}

func (inst *UserUpdateRewardPools) MustFindUserRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserRewardAccountAddress(receiptTokenMint, user, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindUserRewardAccountAddress finds UserRewardAccount account address with given seeds.
func (inst *UserUpdateRewardPools) FindUserRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindUserRewardAccountAddress(receiptTokenMint, user, 0)
	return
}

func (inst *UserUpdateRewardPools) MustFindUserRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserRewardAccountAddress(receiptTokenMint, user, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetUserRewardAccountAccount gets the "user_reward_account" account.
func (inst *UserUpdateRewardPools) GetUserRewardAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst UserUpdateRewardPools) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UserUpdateRewardPools,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UserUpdateRewardPools) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UserUpdateRewardPools) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.User is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.ReceiptTokenMint is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.RewardAccount is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.UserRewardAccount is not set")
		}
	}
	return nil
}

func (inst *UserUpdateRewardPools) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UserUpdateRewardPools")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("              user", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("    system_program", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("receipt_token_mint", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("           reward_", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("      user_reward_", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj UserUpdateRewardPools) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *UserUpdateRewardPools) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewUserUpdateRewardPoolsInstruction declares a new UserUpdateRewardPools instruction with the provided parameters and accounts.
func NewUserUpdateRewardPoolsInstruction(
	// Accounts:
	user ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	receiptTokenMint ag_solanago.PublicKey,
	rewardAccount ag_solanago.PublicKey,
	userRewardAccount ag_solanago.PublicKey) *UserUpdateRewardPools {
	return NewUserUpdateRewardPoolsInstructionBuilder().
		SetUserAccount(user).
		SetSystemProgramAccount(systemProgram).
		SetReceiptTokenMintAccount(receiptTokenMint).
		SetRewardAccountAccount(rewardAccount).
		SetUserRewardAccountAccount(userRewardAccount)
}
