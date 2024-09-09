// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UserCancelWithdrawalRequest is the `user_cancel_withdrawal_request` instruction.
type UserCancelWithdrawalRequest struct {
	RequestId *uint64

	// [0] = [WRITE, SIGNER] user
	//
	// [1] = [] system_program
	//
	// [2] = [] associated_token_program
	//
	// [3] = [] receipt_token_program
	//
	// [4] = [WRITE] receipt_token_mint
	//
	// [5] = [] receipt_token_mint_authority
	//
	// [6] = [] receipt_token_lock_authority
	//
	// [7] = [WRITE] receipt_token_lock_account
	//
	// [8] = [WRITE] user_receipt_token_account
	//
	// [9] = [WRITE] fund_account
	//
	// [10] = [WRITE] user_fund_account
	//
	// [11] = [WRITE] reward_account
	//
	// [12] = [WRITE] user_reward_account
	//
	// [13] = [] instructions_sysvar
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUserCancelWithdrawalRequestInstructionBuilder creates a new `UserCancelWithdrawalRequest` instruction builder.
func NewUserCancelWithdrawalRequestInstructionBuilder() *UserCancelWithdrawalRequest {
	nd := &UserCancelWithdrawalRequest{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 14),
	}
	nd.AccountMetaSlice[1] = ag_solanago.Meta(Addresses["11111111111111111111111111111111"])
	nd.AccountMetaSlice[2] = ag_solanago.Meta(Addresses["ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"])
	nd.AccountMetaSlice[3] = ag_solanago.Meta(Addresses["TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb"])
	nd.AccountMetaSlice[4] = ag_solanago.Meta(Addresses["FRAGSEthVFL7fdqM8hxfxkfCZzUvmg21cqPJVvC1qdbo"]).WRITE()
	nd.AccountMetaSlice[13] = ag_solanago.Meta(Addresses["Sysvar1nstructions1111111111111111111111111"])
	return nd
}

// SetRequestId sets the "request_id" parameter.
func (inst *UserCancelWithdrawalRequest) SetRequestId(request_id uint64) *UserCancelWithdrawalRequest {
	inst.RequestId = &request_id
	return inst
}

// SetUserAccount sets the "user" account.
func (inst *UserCancelWithdrawalRequest) SetUserAccount(user ag_solanago.PublicKey) *UserCancelWithdrawalRequest {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(user).WRITE().SIGNER()
	return inst
}

// GetUserAccount gets the "user" account.
func (inst *UserCancelWithdrawalRequest) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *UserCancelWithdrawalRequest) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *UserCancelWithdrawalRequest {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *UserCancelWithdrawalRequest) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAssociatedTokenProgramAccount sets the "associated_token_program" account.
func (inst *UserCancelWithdrawalRequest) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *UserCancelWithdrawalRequest {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associated_token_program" account.
func (inst *UserCancelWithdrawalRequest) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetReceiptTokenProgramAccount sets the "receipt_token_program" account.
func (inst *UserCancelWithdrawalRequest) SetReceiptTokenProgramAccount(receiptTokenProgram ag_solanago.PublicKey) *UserCancelWithdrawalRequest {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(receiptTokenProgram)
	return inst
}

// GetReceiptTokenProgramAccount gets the "receipt_token_program" account.
func (inst *UserCancelWithdrawalRequest) GetReceiptTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetReceiptTokenMintAccount sets the "receipt_token_mint" account.
func (inst *UserCancelWithdrawalRequest) SetReceiptTokenMintAccount(receiptTokenMint ag_solanago.PublicKey) *UserCancelWithdrawalRequest {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(receiptTokenMint).WRITE()
	return inst
}

// GetReceiptTokenMintAccount gets the "receipt_token_mint" account.
func (inst *UserCancelWithdrawalRequest) GetReceiptTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetReceiptTokenMintAuthorityAccount sets the "receipt_token_mint_authority" account.
func (inst *UserCancelWithdrawalRequest) SetReceiptTokenMintAuthorityAccount(receiptTokenMintAuthority ag_solanago.PublicKey) *UserCancelWithdrawalRequest {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(receiptTokenMintAuthority)
	return inst
}

func (inst *UserCancelWithdrawalRequest) findFindReceiptTokenMintAuthorityAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: receipt_token_mint_authority
	seeds = append(seeds, []byte{byte(0x72), byte(0x65), byte(0x63), byte(0x65), byte(0x69), byte(0x70), byte(0x74), byte(0x5f), byte(0x74), byte(0x6f), byte(0x6b), byte(0x65), byte(0x6e), byte(0x5f), byte(0x6d), byte(0x69), byte(0x6e), byte(0x74), byte(0x5f), byte(0x61), byte(0x75), byte(0x74), byte(0x68), byte(0x6f), byte(0x72), byte(0x69), byte(0x74), byte(0x79)})
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

// FindReceiptTokenMintAuthorityAddressWithBumpSeed calculates ReceiptTokenMintAuthority account address with given seeds and a known bump seed.
func (inst *UserCancelWithdrawalRequest) FindReceiptTokenMintAuthorityAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindReceiptTokenMintAuthorityAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *UserCancelWithdrawalRequest) MustFindReceiptTokenMintAuthorityAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindReceiptTokenMintAuthorityAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindReceiptTokenMintAuthorityAddress finds ReceiptTokenMintAuthority account address with given seeds.
func (inst *UserCancelWithdrawalRequest) FindReceiptTokenMintAuthorityAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindReceiptTokenMintAuthorityAddress(receiptTokenMint, 0)
	return
}

func (inst *UserCancelWithdrawalRequest) MustFindReceiptTokenMintAuthorityAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindReceiptTokenMintAuthorityAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetReceiptTokenMintAuthorityAccount gets the "receipt_token_mint_authority" account.
func (inst *UserCancelWithdrawalRequest) GetReceiptTokenMintAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetReceiptTokenLockAuthorityAccount sets the "receipt_token_lock_authority" account.
func (inst *UserCancelWithdrawalRequest) SetReceiptTokenLockAuthorityAccount(receiptTokenLockAuthority ag_solanago.PublicKey) *UserCancelWithdrawalRequest {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(receiptTokenLockAuthority)
	return inst
}

func (inst *UserCancelWithdrawalRequest) findFindReceiptTokenLockAuthorityAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: receipt_token_lock_authority
	seeds = append(seeds, []byte{byte(0x72), byte(0x65), byte(0x63), byte(0x65), byte(0x69), byte(0x70), byte(0x74), byte(0x5f), byte(0x74), byte(0x6f), byte(0x6b), byte(0x65), byte(0x6e), byte(0x5f), byte(0x6c), byte(0x6f), byte(0x63), byte(0x6b), byte(0x5f), byte(0x61), byte(0x75), byte(0x74), byte(0x68), byte(0x6f), byte(0x72), byte(0x69), byte(0x74), byte(0x79)})
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

// FindReceiptTokenLockAuthorityAddressWithBumpSeed calculates ReceiptTokenLockAuthority account address with given seeds and a known bump seed.
func (inst *UserCancelWithdrawalRequest) FindReceiptTokenLockAuthorityAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindReceiptTokenLockAuthorityAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *UserCancelWithdrawalRequest) MustFindReceiptTokenLockAuthorityAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindReceiptTokenLockAuthorityAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindReceiptTokenLockAuthorityAddress finds ReceiptTokenLockAuthority account address with given seeds.
func (inst *UserCancelWithdrawalRequest) FindReceiptTokenLockAuthorityAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindReceiptTokenLockAuthorityAddress(receiptTokenMint, 0)
	return
}

func (inst *UserCancelWithdrawalRequest) MustFindReceiptTokenLockAuthorityAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindReceiptTokenLockAuthorityAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetReceiptTokenLockAuthorityAccount gets the "receipt_token_lock_authority" account.
func (inst *UserCancelWithdrawalRequest) GetReceiptTokenLockAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetReceiptTokenLockAccountAccount sets the "receipt_token_lock_account" account.
func (inst *UserCancelWithdrawalRequest) SetReceiptTokenLockAccountAccount(receiptTokenLockAccount ag_solanago.PublicKey) *UserCancelWithdrawalRequest {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(receiptTokenLockAccount).WRITE()
	return inst
}

func (inst *UserCancelWithdrawalRequest) findFindReceiptTokenLockAccountAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: receipt_token_lock
	seeds = append(seeds, []byte{byte(0x72), byte(0x65), byte(0x63), byte(0x65), byte(0x69), byte(0x70), byte(0x74), byte(0x5f), byte(0x74), byte(0x6f), byte(0x6b), byte(0x65), byte(0x6e), byte(0x5f), byte(0x6c), byte(0x6f), byte(0x63), byte(0x6b)})
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

// FindReceiptTokenLockAccountAddressWithBumpSeed calculates ReceiptTokenLockAccount account address with given seeds and a known bump seed.
func (inst *UserCancelWithdrawalRequest) FindReceiptTokenLockAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindReceiptTokenLockAccountAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *UserCancelWithdrawalRequest) MustFindReceiptTokenLockAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindReceiptTokenLockAccountAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindReceiptTokenLockAccountAddress finds ReceiptTokenLockAccount account address with given seeds.
func (inst *UserCancelWithdrawalRequest) FindReceiptTokenLockAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindReceiptTokenLockAccountAddress(receiptTokenMint, 0)
	return
}

func (inst *UserCancelWithdrawalRequest) MustFindReceiptTokenLockAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindReceiptTokenLockAccountAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetReceiptTokenLockAccountAccount gets the "receipt_token_lock_account" account.
func (inst *UserCancelWithdrawalRequest) GetReceiptTokenLockAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetUserReceiptTokenAccountAccount sets the "user_receipt_token_account" account.
func (inst *UserCancelWithdrawalRequest) SetUserReceiptTokenAccountAccount(userReceiptTokenAccount ag_solanago.PublicKey) *UserCancelWithdrawalRequest {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(userReceiptTokenAccount).WRITE()
	return inst
}

func (inst *UserCancelWithdrawalRequest) findFindUserReceiptTokenAccountAddress(user ag_solanago.PublicKey, receiptTokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// path: user
	seeds = append(seeds, user.Bytes())
	// path: receiptTokenProgram
	seeds = append(seeds, receiptTokenProgram.Bytes())
	// path: receiptTokenMint
	seeds = append(seeds, receiptTokenMint.Bytes())

	programID := Addresses["ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"]

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, programID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, programID)
	}
	return
}

// FindUserReceiptTokenAccountAddressWithBumpSeed calculates UserReceiptTokenAccount account address with given seeds and a known bump seed.
func (inst *UserCancelWithdrawalRequest) FindUserReceiptTokenAccountAddressWithBumpSeed(user ag_solanago.PublicKey, receiptTokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindUserReceiptTokenAccountAddress(user, receiptTokenProgram, receiptTokenMint, bumpSeed)
	return
}

func (inst *UserCancelWithdrawalRequest) MustFindUserReceiptTokenAccountAddressWithBumpSeed(user ag_solanago.PublicKey, receiptTokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserReceiptTokenAccountAddress(user, receiptTokenProgram, receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindUserReceiptTokenAccountAddress finds UserReceiptTokenAccount account address with given seeds.
func (inst *UserCancelWithdrawalRequest) FindUserReceiptTokenAccountAddress(user ag_solanago.PublicKey, receiptTokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindUserReceiptTokenAccountAddress(user, receiptTokenProgram, receiptTokenMint, 0)
	return
}

func (inst *UserCancelWithdrawalRequest) MustFindUserReceiptTokenAccountAddress(user ag_solanago.PublicKey, receiptTokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserReceiptTokenAccountAddress(user, receiptTokenProgram, receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetUserReceiptTokenAccountAccount gets the "user_receipt_token_account" account.
func (inst *UserCancelWithdrawalRequest) GetUserReceiptTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetFundAccountAccount sets the "fund_account" account.
func (inst *UserCancelWithdrawalRequest) SetFundAccountAccount(fundAccount ag_solanago.PublicKey) *UserCancelWithdrawalRequest {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(fundAccount).WRITE()
	return inst
}

func (inst *UserCancelWithdrawalRequest) findFindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: fund
	seeds = append(seeds, []byte{byte(0x66), byte(0x75), byte(0x6e), byte(0x64)})
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

// FindFundAccountAddressWithBumpSeed calculates FundAccount account address with given seeds and a known bump seed.
func (inst *UserCancelWithdrawalRequest) FindFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindFundAccountAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *UserCancelWithdrawalRequest) MustFindFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAccountAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindFundAccountAddress finds FundAccount account address with given seeds.
func (inst *UserCancelWithdrawalRequest) FindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindFundAccountAddress(receiptTokenMint, 0)
	return
}

func (inst *UserCancelWithdrawalRequest) MustFindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAccountAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetFundAccountAccount gets the "fund_account" account.
func (inst *UserCancelWithdrawalRequest) GetFundAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetUserFundAccountAccount sets the "user_fund_account" account.
func (inst *UserCancelWithdrawalRequest) SetUserFundAccountAccount(userFundAccount ag_solanago.PublicKey) *UserCancelWithdrawalRequest {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(userFundAccount).WRITE()
	return inst
}

func (inst *UserCancelWithdrawalRequest) findFindUserFundAccountAddress(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: user_fund
	seeds = append(seeds, []byte{byte(0x75), byte(0x73), byte(0x65), byte(0x72), byte(0x5f), byte(0x66), byte(0x75), byte(0x6e), byte(0x64)})
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

// FindUserFundAccountAddressWithBumpSeed calculates UserFundAccount account address with given seeds and a known bump seed.
func (inst *UserCancelWithdrawalRequest) FindUserFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindUserFundAccountAddress(receiptTokenMint, user, bumpSeed)
	return
}

func (inst *UserCancelWithdrawalRequest) MustFindUserFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserFundAccountAddress(receiptTokenMint, user, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindUserFundAccountAddress finds UserFundAccount account address with given seeds.
func (inst *UserCancelWithdrawalRequest) FindUserFundAccountAddress(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindUserFundAccountAddress(receiptTokenMint, user, 0)
	return
}

func (inst *UserCancelWithdrawalRequest) MustFindUserFundAccountAddress(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserFundAccountAddress(receiptTokenMint, user, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetUserFundAccountAccount gets the "user_fund_account" account.
func (inst *UserCancelWithdrawalRequest) GetUserFundAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetRewardAccountAccount sets the "reward_account" account.
func (inst *UserCancelWithdrawalRequest) SetRewardAccountAccount(rewardAccount ag_solanago.PublicKey) *UserCancelWithdrawalRequest {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(rewardAccount).WRITE()
	return inst
}

func (inst *UserCancelWithdrawalRequest) findFindRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *UserCancelWithdrawalRequest) FindRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindRewardAccountAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *UserCancelWithdrawalRequest) MustFindRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindRewardAccountAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindRewardAccountAddress finds RewardAccount account address with given seeds.
func (inst *UserCancelWithdrawalRequest) FindRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindRewardAccountAddress(receiptTokenMint, 0)
	return
}

func (inst *UserCancelWithdrawalRequest) MustFindRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindRewardAccountAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetRewardAccountAccount gets the "reward_account" account.
func (inst *UserCancelWithdrawalRequest) GetRewardAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetUserRewardAccountAccount sets the "user_reward_account" account.
func (inst *UserCancelWithdrawalRequest) SetUserRewardAccountAccount(userRewardAccount ag_solanago.PublicKey) *UserCancelWithdrawalRequest {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(userRewardAccount).WRITE()
	return inst
}

func (inst *UserCancelWithdrawalRequest) findFindUserRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *UserCancelWithdrawalRequest) FindUserRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindUserRewardAccountAddress(receiptTokenMint, user, bumpSeed)
	return
}

func (inst *UserCancelWithdrawalRequest) MustFindUserRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserRewardAccountAddress(receiptTokenMint, user, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindUserRewardAccountAddress finds UserRewardAccount account address with given seeds.
func (inst *UserCancelWithdrawalRequest) FindUserRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindUserRewardAccountAddress(receiptTokenMint, user, 0)
	return
}

func (inst *UserCancelWithdrawalRequest) MustFindUserRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserRewardAccountAddress(receiptTokenMint, user, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetUserRewardAccountAccount gets the "user_reward_account" account.
func (inst *UserCancelWithdrawalRequest) GetUserRewardAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetInstructionsSysvarAccount sets the "instructions_sysvar" account.
func (inst *UserCancelWithdrawalRequest) SetInstructionsSysvarAccount(instructionsSysvar ag_solanago.PublicKey) *UserCancelWithdrawalRequest {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(instructionsSysvar)
	return inst
}

// GetInstructionsSysvarAccount gets the "instructions_sysvar" account.
func (inst *UserCancelWithdrawalRequest) GetInstructionsSysvarAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

func (inst UserCancelWithdrawalRequest) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UserCancelWithdrawalRequest,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UserCancelWithdrawalRequest) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UserCancelWithdrawalRequest) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.RequestId == nil {
			return errors.New("RequestId parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.User is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.ReceiptTokenProgram is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.ReceiptTokenMint is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.ReceiptTokenMintAuthority is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.ReceiptTokenLockAuthority is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.ReceiptTokenLockAccount is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.UserReceiptTokenAccount is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.FundAccount is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.UserFundAccount is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.RewardAccount is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.UserRewardAccount is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.InstructionsSysvar is not set")
		}
	}
	return nil
}

func (inst *UserCancelWithdrawalRequest) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UserCancelWithdrawalRequest")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param(" RequestId", *inst.RequestId))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=14]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                        user", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("              system_program", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("    associated_token_program", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("       receipt_token_program", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("          receipt_token_mint", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("receipt_token_mint_authority", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("receipt_token_lock_authority", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("         receipt_token_lock_", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("         user_receipt_token_", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("                       fund_", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("                  user_fund_", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("                     reward_", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("                user_reward_", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("         instructions_sysvar", inst.AccountMetaSlice.Get(13)))
					})
				})
		})
}

func (obj UserCancelWithdrawalRequest) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `RequestId` param:
	err = encoder.Encode(obj.RequestId)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UserCancelWithdrawalRequest) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `RequestId`:
	err = decoder.Decode(&obj.RequestId)
	if err != nil {
		return err
	}
	return nil
}

// NewUserCancelWithdrawalRequestInstruction declares a new UserCancelWithdrawalRequest instruction with the provided parameters and accounts.
func NewUserCancelWithdrawalRequestInstruction(
	// Parameters:
	request_id uint64,
	// Accounts:
	user ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey,
	receiptTokenProgram ag_solanago.PublicKey,
	receiptTokenMint ag_solanago.PublicKey,
	receiptTokenMintAuthority ag_solanago.PublicKey,
	receiptTokenLockAuthority ag_solanago.PublicKey,
	receiptTokenLockAccount ag_solanago.PublicKey,
	userReceiptTokenAccount ag_solanago.PublicKey,
	fundAccount ag_solanago.PublicKey,
	userFundAccount ag_solanago.PublicKey,
	rewardAccount ag_solanago.PublicKey,
	userRewardAccount ag_solanago.PublicKey,
	instructionsSysvar ag_solanago.PublicKey) *UserCancelWithdrawalRequest {
	return NewUserCancelWithdrawalRequestInstructionBuilder().
		SetRequestId(request_id).
		SetUserAccount(user).
		SetSystemProgramAccount(systemProgram).
		SetAssociatedTokenProgramAccount(associatedTokenProgram).
		SetReceiptTokenProgramAccount(receiptTokenProgram).
		SetReceiptTokenMintAccount(receiptTokenMint).
		SetReceiptTokenMintAuthorityAccount(receiptTokenMintAuthority).
		SetReceiptTokenLockAuthorityAccount(receiptTokenLockAuthority).
		SetReceiptTokenLockAccountAccount(receiptTokenLockAccount).
		SetUserReceiptTokenAccountAccount(userReceiptTokenAccount).
		SetFundAccountAccount(fundAccount).
		SetUserFundAccountAccount(userFundAccount).
		SetRewardAccountAccount(rewardAccount).
		SetUserRewardAccountAccount(userRewardAccount).
		SetInstructionsSysvarAccount(instructionsSysvar)
}
