// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"encoding/json"
	"errors"
	"fmt"
	ag_jsonrpc "github.com/gagliardetto/solana-go/rpc/jsonrpc"
)

var (
	_ *json.Encoder        = nil
	_ *ag_jsonrpc.RPCError = nil
	_ fmt.Formatter        = nil
	_                      = errors.ErrUnsupported
)
var (
	ErrInvalidSignatureError = &customErrorDef{
		code: 6000,
		msg:  "signature verification failed",
		name: "InvalidSignatureError",
	}
	ErrTokenNotTransferableError = &customErrorDef{
		code: 6001,
		msg:  "token is not transferable currently",
		name: "TokenNotTransferableError",
	}
	ErrTokenNotTransferringException = &customErrorDef{
		code: 6002,
		msg:  "token is not transferring currently",
		name: "TokenNotTransferringException",
	}
	ErrCalculationArithmeticException = &customErrorDef{
		code: 6003,
		msg:  "calculation arithmetic exception",
		name: "CalculationArithmeticException",
	}
	ErrDecodeInvalidUtf8FormatException = &customErrorDef{
		code: 6004,
		msg:  "decode invalid utf-8 format exception",
		name: "DecodeInvalidUtf8FormatException",
	}
	ErrFundInvalidUpdateError = &customErrorDef{
		code: 6005,
		msg:  "fund: cannot apply invalid update",
		name: "FundInvalidUpdateError",
	}
	ErrFundSOLTransferFailedException = &customErrorDef{
		code: 6006,
		msg:  "fund: sol transfer failed",
		name: "FundSOLTransferFailedException",
	}
	ErrFundTokenTransferFailedException = &customErrorDef{
		code: 6007,
		msg:  "fund: token transfer failed",
		name: "FundTokenTransferFailedException",
	}
	ErrFundAlreadySupportedTokenError = &customErrorDef{
		code: 6008,
		msg:  "fund: already supported token",
		name: "FundAlreadySupportedTokenError",
	}
	ErrFundNotSupportedTokenError = &customErrorDef{
		code: 6009,
		msg:  "fund: not supported the token",
		name: "FundNotSupportedTokenError",
	}
	ErrFundExceededSOLCapacityAmountError = &customErrorDef{
		code: 6010,
		msg:  "fund: exceeded sol capacity amount",
		name: "FundExceededSOLCapacityAmountError",
	}
	ErrFundExceededTokenCapacityAmountError = &customErrorDef{
		code: 6011,
		msg:  "fund: exceeded token capacity amount",
		name: "FundExceededTokenCapacityAmountError",
	}
	ErrFundExceededMaxWithdrawalRequestError = &customErrorDef{
		code: 6012,
		msg:  "fund: exceeded max withdrawal request per user",
		name: "FundExceededMaxWithdrawalRequestError",
	}
	ErrFundOperationReservedSOLExhaustedException = &customErrorDef{
		code: 6013,
		msg:  "fund: operation reserved sol is exhausted",
		name: "FundOperationReservedSOLExhaustedException",
	}
	ErrFundWithdrawalRequestNotFoundError = &customErrorDef{
		code: 6014,
		msg:  "fund: withdrawal request not found",
		name: "FundWithdrawalRequestNotFoundError",
	}
	ErrFundPendingWithdrawalRequestError = &customErrorDef{
		code: 6015,
		msg:  "fund: withdrawal request not completed yet",
		name: "FundPendingWithdrawalRequestError",
	}
	ErrFundWithdrawalReservedSOLExhaustedException = &customErrorDef{
		code: 6016,
		msg:  "fund: withdrawal reserved sol is exhausted",
		name: "FundWithdrawalReservedSOLExhaustedException",
	}
	ErrFundWithdrawalDisabledError = &customErrorDef{
		code: 6017,
		msg:  "fund: withdrawal is currently disabled",
		name: "FundWithdrawalDisabledError",
	}
	ErrFundProcessingWithdrawalRequestError = &customErrorDef{
		code: 6018,
		msg:  "fund: withdrawal request is already in progress",
		name: "FundProcessingWithdrawalRequestError",
	}
	ErrFundTokenPricingSourceNotFoundException = &customErrorDef{
		code: 6019,
		msg:  "fund: token pricing source is not found",
		name: "FundTokenPricingSourceNotFoundException",
	}
	ErrOperatorJobUnmetThresholdError = &customErrorDef{
		code: 6020,
		msg:  "operator: job unmet threshold",
		name: "OperatorJobUnmetThresholdError",
	}
	ErrRewardInvalidTransferArgsException = &customErrorDef{
		code: 6021,
		msg:  "reward: invalid token transfer args",
		name: "RewardInvalidTransferArgsException",
	}
	ErrRewardInvalidMetadataNameLengthError = &customErrorDef{
		code: 6022,
		msg:  "reward: invalid metadata name length",
		name: "RewardInvalidMetadataNameLengthError",
	}
	ErrRewardInvalidMetadataDescriptionLengthError = &customErrorDef{
		code: 6023,
		msg:  "reward: invalid metadata description length",
		name: "RewardInvalidMetadataDescriptionLengthError",
	}
	ErrRewardInvalidRewardType = &customErrorDef{
		code: 6024,
		msg:  "reward: invalid reward type",
		name: "RewardInvalidRewardType",
	}
	ErrRewardAlreadyExistingHolderError = &customErrorDef{
		code: 6025,
		msg:  "reward: already existing holder",
		name: "RewardAlreadyExistingHolderError",
	}
	ErrRewardAlreadyExistingRewardError = &customErrorDef{
		code: 6026,
		msg:  "reward: already existing reward",
		name: "RewardAlreadyExistingRewardError",
	}
	ErrRewardAlreadyExistingPoolError = &customErrorDef{
		code: 6027,
		msg:  "reward: already existing pool",
		name: "RewardAlreadyExistingPoolError",
	}
	ErrRewardHolderNotFoundError = &customErrorDef{
		code: 6028,
		msg:  "reward: holder not found",
		name: "RewardHolderNotFoundError",
	}
	ErrRewardNotFoundError = &customErrorDef{
		code: 6029,
		msg:  "reward: reward not found",
		name: "RewardNotFoundError",
	}
	ErrRewardPoolNotFoundError = &customErrorDef{
		code: 6030,
		msg:  "reward: pool not found",
		name: "RewardPoolNotFoundError",
	}
	ErrRewardUserPoolNotFoundError = &customErrorDef{
		code: 6031,
		msg:  "reward: user pool not found",
		name: "RewardUserPoolNotFoundError",
	}
	ErrRewardPoolClosedError = &customErrorDef{
		code: 6032,
		msg:  "reward: pool is closed",
		name: "RewardPoolClosedError",
	}
	ErrRewardInvalidPoolConfigurationException = &customErrorDef{
		code: 6033,
		msg:  "reward: invalid pool configuration",
		name: "RewardInvalidPoolConfigurationException",
	}
	ErrRewardInvalidPoolAccessException = &customErrorDef{
		code: 6034,
		msg:  "reward: invalid reward pool access",
		name: "RewardInvalidPoolAccessException",
	}
	ErrRewardUnmetAccountReallocError = &customErrorDef{
		code: 6035,
		msg:  "reward: unmet account size reallocation",
		name: "RewardUnmetAccountReallocError",
	}
	ErrRewardInvalidAccountingException = &customErrorDef{
		code: 6036,
		msg:  "reward: incorrect accounting exception",
		name: "RewardInvalidAccountingException",
	}
	ErrRewardInvalidAllocatedAmountDeltaException = &customErrorDef{
		code: 6037,
		msg:  "reward: invalid amount or contribution accrual rate",
		name: "RewardInvalidAllocatedAmountDeltaException",
	}
	ErrRewardExceededMaxHoldersException = &customErrorDef{
		code: 6038,
		msg:  "reward: exceeded max holders",
		name: "RewardExceededMaxHoldersException",
	}
	ErrRewardExceededMaxRewardsException = &customErrorDef{
		code: 6039,
		msg:  "reward: exceeded max rewards",
		name: "RewardExceededMaxRewardsException",
	}
	ErrRewardExceededMaxRewardPoolsException = &customErrorDef{
		code: 6040,
		msg:  "reward: exceeded max reward pools",
		name: "RewardExceededMaxRewardPoolsException",
	}
	ErrRewardExceededMaxUserRewardPoolsException = &customErrorDef{
		code: 6041,
		msg:  "reward: exceeded max user reward pools",
		name: "RewardExceededMaxUserRewardPoolsException",
	}
	ErrRewardExceededMaxHolderPubkeysException = &customErrorDef{
		code: 6042,
		msg:  "reward: exceeded max pubkeys per holder",
		name: "RewardExceededMaxHolderPubkeysException",
	}
	ErrRewardExceededMaxTokenAllocatedAmountRecordException = &customErrorDef{
		code: 6043,
		msg:  "reward: exceeded max token allocated amount record",
		name: "RewardExceededMaxTokenAllocatedAmountRecordException",
	}
	ErrRewardExceededMaxRewardSettlementException = &customErrorDef{
		code: 6044,
		msg:  "reward: exceeded max reward settlements per pool",
		name: "RewardExceededMaxRewardSettlementException",
	}
	ErrRewardStaleSettlementBlockNotExistError = &customErrorDef{
		code: 6045,
		msg:  "reward: stale settlement block not exist",
		name: "RewardStaleSettlementBlockNotExistError",
	}
	ErrRewardInvalidSettlementBlockHeightException = &customErrorDef{
		code: 6046,
		msg:  "reward: invalid settlement block height",
		name: "RewardInvalidSettlementBlockHeightException",
	}
	ErrRewardInvalidSettlementBlockContributionException = &customErrorDef{
		code: 6047,
		msg:  "reward: invalid settlement block contribution",
		name: "RewardInvalidSettlementBlockContributionException",
	}
	ErrRewardInvalidTotalUserSettledAmountException = &customErrorDef{
		code: 6048,
		msg:  "reward: sum of user settled amount cannot exceed total amount",
		name: "RewardInvalidTotalUserSettledAmountException",
	}
	ErrRewardInvalidTotalUserSettledContributionException = &customErrorDef{
		code: 6049,
		msg:  "reward: sum of user settled contribution cannot exceed total contribution",
		name: "RewardInvalidTotalUserSettledContributionException",
	}
	Errors = map[int]CustomError{
		6000: ErrInvalidSignatureError,
		6001: ErrTokenNotTransferableError,
		6002: ErrTokenNotTransferringException,
		6003: ErrCalculationArithmeticException,
		6004: ErrDecodeInvalidUtf8FormatException,
		6005: ErrFundInvalidUpdateError,
		6006: ErrFundSOLTransferFailedException,
		6007: ErrFundTokenTransferFailedException,
		6008: ErrFundAlreadySupportedTokenError,
		6009: ErrFundNotSupportedTokenError,
		6010: ErrFundExceededSOLCapacityAmountError,
		6011: ErrFundExceededTokenCapacityAmountError,
		6012: ErrFundExceededMaxWithdrawalRequestError,
		6013: ErrFundOperationReservedSOLExhaustedException,
		6014: ErrFundWithdrawalRequestNotFoundError,
		6015: ErrFundPendingWithdrawalRequestError,
		6016: ErrFundWithdrawalReservedSOLExhaustedException,
		6017: ErrFundWithdrawalDisabledError,
		6018: ErrFundProcessingWithdrawalRequestError,
		6019: ErrFundTokenPricingSourceNotFoundException,
		6020: ErrOperatorJobUnmetThresholdError,
		6021: ErrRewardInvalidTransferArgsException,
		6022: ErrRewardInvalidMetadataNameLengthError,
		6023: ErrRewardInvalidMetadataDescriptionLengthError,
		6024: ErrRewardInvalidRewardType,
		6025: ErrRewardAlreadyExistingHolderError,
		6026: ErrRewardAlreadyExistingRewardError,
		6027: ErrRewardAlreadyExistingPoolError,
		6028: ErrRewardHolderNotFoundError,
		6029: ErrRewardNotFoundError,
		6030: ErrRewardPoolNotFoundError,
		6031: ErrRewardUserPoolNotFoundError,
		6032: ErrRewardPoolClosedError,
		6033: ErrRewardInvalidPoolConfigurationException,
		6034: ErrRewardInvalidPoolAccessException,
		6035: ErrRewardUnmetAccountReallocError,
		6036: ErrRewardInvalidAccountingException,
		6037: ErrRewardInvalidAllocatedAmountDeltaException,
		6038: ErrRewardExceededMaxHoldersException,
		6039: ErrRewardExceededMaxRewardsException,
		6040: ErrRewardExceededMaxRewardPoolsException,
		6041: ErrRewardExceededMaxUserRewardPoolsException,
		6042: ErrRewardExceededMaxHolderPubkeysException,
		6043: ErrRewardExceededMaxTokenAllocatedAmountRecordException,
		6044: ErrRewardExceededMaxRewardSettlementException,
		6045: ErrRewardStaleSettlementBlockNotExistError,
		6046: ErrRewardInvalidSettlementBlockHeightException,
		6047: ErrRewardInvalidSettlementBlockContributionException,
		6048: ErrRewardInvalidTotalUserSettledAmountException,
		6049: ErrRewardInvalidTotalUserSettledContributionException,
	}
)

type CustomError interface {
	Code() int
	Name() string
	Error() string
}

type customErrorDef struct {
	code int
	name string
	msg  string
}

func (e *customErrorDef) Code() int {
	return e.code
}

func (e *customErrorDef) Name() string {
	return e.name
}

func (e *customErrorDef) Error() string {
	return fmt.Sprintf("%s(%d): %s", e.name, e.code, e.msg)
}

func DecodeCustomError(rpcErr error) (err error, ok bool) {
	if errCode, o := decodeErrorCode(rpcErr); o {
		if customErr, o := Errors[errCode]; o {
			err = customErr
			ok = true
			return
		}
	}
	return
}

func decodeErrorCode(rpcErr error) (errorCode int, ok bool) {
	var jErr *ag_jsonrpc.RPCError
	if errors.As(rpcErr, &jErr) && jErr.Data != nil {
		if root, o := jErr.Data.(map[string]interface{}); o {
			if rootErr, o := root["err"].(map[string]interface{}); o {
				if rootErrInstructionError, o := rootErr["InstructionError"]; o {
					if rootErrInstructionErrorItems, o := rootErrInstructionError.([]interface{}); o {
						if len(rootErrInstructionErrorItems) == 2 {
							if v, o := rootErrInstructionErrorItems[1].(map[string]interface{}); o {
								if v2, o := v["Custom"].(json.Number); o {
									if code, err := v2.Int64(); err == nil {
										ok = true
										errorCode = int(code)
									}
								} else if v2, o := v["Custom"].(float64); o {
									ok = true
									errorCode = int(v2)
								}
							}
						}
					}
				}
			}
		}
	}
	return
}
