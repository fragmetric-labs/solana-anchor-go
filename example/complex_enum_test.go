package example

import (
	"bytes"
	"encoding/hex"
	"github.com/davecgh/go-spew/spew"
	"github.com/fragmetric-labs/solana-anchor-go/generated/restaking"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestComplexEnum(t *testing.T) {
	var res restaking.OperationCommandResult = &restaking.OperationCommandResultProcessWithdrawalBatchTuple{
		Elem0: restaking.ProcessWithdrawalBatchCommandResult{
			RequestedReceiptTokenAmount:   1,
			ProcessedReceiptTokenAmount:   2,
			AssetTokenMint:                nil,
			RequiredAssetAmount:           3,
			ReservedAssetUserAmount:       4,
			DeductedAssetFeeAmount:        5,
			OffsettedAssetReceivables:     nil,
			TransferredAssetRevenueAmount: 6,
			WithdrawalFeeRateBps:          7,
		},
	}
	src := restaking.OperatorRanFundCommand{
		ReceiptTokenMint: ag_solanago.MustPublicKeyFromBase58("GPKjBDTNexAsis6zqGnioAzhauvHzs6UzGXKxx37HdkA"),
		FundAccount:      ag_solanago.MustPublicKeyFromBase58("HdZM8mzEH7JAcswjJNgCC8Zmbu97LCzYo4WCSvFkfWKx"),
		NextSequence:     123,
		NumOperated:      456,
		Command: restaking.OperationCommandProcessWithdrawalBatchTuple{
			Elem0: restaking.ProcessWithdrawalBatchCommand{
				State: &restaking.ProcessWithdrawalBatchCommandStateExecuteTuple{
					AssetTokenMint:       nil,
					NumProcessingBatches: 1,
					ReceiptTokenAmount:   12345,
				},
				Forced: true,
			},
		},
		Result: &res,
	}

	// encoding
	buf := new(bytes.Buffer)
	require.NoError(t, ag_binary.NewBorshEncoder(buf).Encode(src), "1")

	dst := restaking.OperatorRanFundCommand{}
	require.NoError(t, ag_binary.NewBorshDecoder(buf.Bytes()).Decode(&dst), "2")

	printer := spew.ConfigState{
		Indent:                  " ",
		DisablePointerAddresses: true,
		DisableCapacities:       true,
	}
	require.Equal(t, printer.Sdump(src), printer.Sdump(dst), "3")
	printer.Dump(src)
}

func TestComplexEnum_DecodeRawEvent(t *testing.T) {
	eventData, err := hex.DecodeString("e445a52e51cb9a1d0a001dcc807de395d634089bb695733914837de852fbd24cff284e2722a63480697643ca75f76c92349707ef78937aa1fdee335943b8d8e3eabc7a07f9384299867f65d1a69d98080000f0050000000000000702000100e1f5050000000000010700e1f5050000000000000000000000000000e13c0600000000000000000000000000000000000000000000000000000000000000001400")
	require.NoError(t, err, "failed to decode bytes")
	buf := bytes.NewBuffer(eventData[8:])

	dst := restaking.OperatorRanFundCommandEventData{}
	require.NoError(t, ag_binary.NewBorshDecoder(buf.Bytes()).Decode(&dst), "2")

	printer := spew.ConfigState{
		Indent:                  " ",
		DisablePointerAddresses: true,
		DisableCapacities:       true,
	}
	printer.Dump(dst)
	require.NotEmpty(t, dst.Result, "result is empty .. see: https://explorer.solana.com/tx/48CL5KorgSLNmSSpS2SgZaNPsjKNkC9qoax9bTBPcZ3td57eDdiwDvQ2m1UPzxG6JcjNBqmUMJ8mEqE5eyqmDg7P?cluster=devnet")
}
