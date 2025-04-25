package main

import (
	. "github.com/dave/jennifer/jen"
)

func generateEventSnippet(file *File) {
	file.Type().Id("Event").Struct(
		Id("Name").String(),
		Id("Data").Id("EventData"),
	)

	file.Type().Id("EventData").Interface(
		Id("UnmarshalWithDecoder").Params(
			Id("decoder").Op("*").Qual("github.com/gagliardetto/binary", "Decoder"),
		).Error(),
		Id("isEventData").Params(),
	)

	file.Const().Id("eventLogPrefix").Op("=").Lit("Program data: ")

	generateDecodeEventsFunc(file)
	generateDecodeEventsFromLogMessageFunc(file)
	generateDecodeEventsFromEmitCPIFunc(file)
	generateParseEventsFunc(file)
}

func generateDecodeEventsFunc(file *File) {
	file.Func().Id("DecodeEvents").Params(
		Id("txData").Op("*").Qual("github.com/gagliardetto/solana-go/rpc", "GetTransactionResult"),
		Id("targetProgramId").Qual("github.com/gagliardetto/solana-go", "PublicKey"),
		Id("getAddressTables").Func().Params(
			Id("altAddresses").Index().Qual("github.com/gagliardetto/solana-go", "PublicKey"),
		).Params(
			Id("tables").Map(Qual("github.com/gagliardetto/solana-go", "PublicKey")).Qual("github.com/gagliardetto/solana-go", "PublicKeySlice"),
			Err().Error(),
		),
	).Params(
		Id("evts").Index().Op("*").Id("Event"),
		Err().Error(),
	).Block(
		Var().Id("tx").Op("*").Qual("github.com/gagliardetto/solana-go", "Transaction"),
		If(
			List(Id("tx"), Err()).Op("=").Id("txData").Dot("Transaction").Dot("GetTransaction").Call(),
			Err().Op("!=").Nil(),
		).Block(
			Return(),
		),

		Id("altAddresses").Op(":=").Make(Index().Qual("github.com/gagliardetto/solana-go", "PublicKey"), Len(Id("tx").Dot("Message").Dot("AddressTableLookups"))),
		For(
			List(Id("i"), Id("alt")).Op(":=").Range().Id("tx").Dot("Message").Dot("AddressTableLookups"),
		).Block(
			Id("altAddresses").Index(Id("i")).Op("=").Id("alt").Dot("AccountKey"),
		),

		If(Len(Id("altAddresses")).Op(">").Lit(0)).Block(
			Var().Id("tables").Map(Qual("github.com/gagliardetto/solana-go", "PublicKey")).Qual("github.com/gagliardetto/solana-go", "PublicKeySlice"),
			If(
				List(Id("tables"), Err()).Op("=").Id("getAddressTables").Call(Id("altAddresses")),
				Err().Op("!=").Nil(),
			).Block(
				Return(),
			),
			Id("tx").Dot("Message").Dot("SetAddressTables").Call(Id("tables")),
			If(
				Err().Op("=").Id("tx").Dot("Message").Dot("ResolveLookups").Call(),
				Err().Op("!=").Nil(),
			).Block(
				Return(),
			),
		),

		Var().Id("base64Binaries").Index().Index().Byte(),
		List(Id("logMessageEventBinaries"), Err()).Op(":=").Id("decodeEventsFromLogMessage").Call(Id("txData").Dot("Meta").Dot("LogMessages")),
		If(Err().Op("!=").Nil()).Block(
			Return(),
		),

		List(Id("emitedCPIEventBinaries"), Err()).Op(":=").Id("decodeEventsFromEmitCPI").Call(
			Id("txData").Dot("Meta").Dot("InnerInstructions"),
			Id("tx").Dot("Message").Dot("AccountKeys"),
			Id("targetProgramId"),
		),
		If(Err().Op("!=").Nil()).Block(
			Return(),
		),

		Id("base64Binaries").Op("=").Append(Id("base64Binaries"), Id("logMessageEventBinaries").Op("...")),
		Id("base64Binaries").Op("=").Append(Id("base64Binaries"), Id("emitedCPIEventBinaries").Op("...")),
		List(Id("evts"), Err()).Op("=").Id("parseEvents").Call(Id("base64Binaries")),
		Return(),
	).Line()
}

func generateDecodeEventsFromLogMessageFunc(file *File) {
	file.Func().Id("decodeEventsFromLogMessage").Params(
		Id("logMessages").Index().String(),
	).Params(
		Id("eventBinaries").Index().Index().Byte(),
		Err().Error(),
	).Block(
		For(List(Id("_"), Id("log")).Op(":=").Range().Id("logMessages")).Block(
			If(Qual("strings", "HasPrefix").Call(Id("log"), Id("eventLogPrefix"))).Block(
				Id("eventBase64").Op(":=").Id("log").Index(Len(Id("eventLogPrefix")).Op(":")),

				Var().Id("eventBinary").Index().Byte(),
				If(
					List(Id("eventBinary"), Err()).Op("=").Qual("encoding/base64", "StdEncoding").Dot("DecodeString").Call(Id("eventBase64")),
					Err().Op("!=").Nil(),
				).Block(
					Err().Op("=").Qual("fmt", "Errorf").Call(Lit("failed to decode logMessage event: %s"), Id("eventBase64")),
					Return(),
				),
				Id("eventBinaries").Op("=").Append(Id("eventBinaries"), Id("eventBinary")),
			),
		),
		Return(),
	).Line()
}

func generateDecodeEventsFromEmitCPIFunc(file *File) {
	file.Func().Id("decodeEventsFromEmitCPI").Params(
		Id("InnerInstructions").Index().Qual("github.com/gagliardetto/solana-go/rpc", "InnerInstruction"),
		Id("accountKeys").Qual("github.com/gagliardetto/solana-go", "PublicKeySlice"),
		Id("targetProgramId").Qual("github.com/gagliardetto/solana-go", "PublicKey"),
	).Params(
		Id("eventBinaries").Index().Index().Byte(),
		Err().Error(),
	).Block(
		For(List(Id("_"), Id("parsedIx")).Op(":=").Range().Id("InnerInstructions")).Block(
			For(List(Id("_"), Id("ix")).Op(":=").Range().Id("parsedIx").Dot("Instructions")).Block(
				If(Id("accountKeys").Index(Id("ix").Dot("ProgramIDIndex")).Op("!=").Id("targetProgramId")).Block(
					Continue(),
				),

				Var().Id("ixData").Index().Byte(),
				If(
					List(Id("ixData"), Err()).Op("=").Qual("github.com/mr-tron/base58", "Decode").Call(Id("ix").Dot("Data").Dot("String").Call()),
					Err().Op("!=").Nil(),
				).Block(
					Return(),
				),
				If(Len(Id("ixData")).Op("<").Lit(8)).Block(
					Continue(),
				),

				Id("eventBase64").Op(":=").Qual("encoding/base64", "StdEncoding").Dot("EncodeToString").Call(Id("ixData").Index(Lit(8), Empty())),
				Var().Id("eventBinary").Index().Byte(),
				If(
					List(Id("eventBinary"), Err()).Op("=").Qual("encoding/base64", "StdEncoding").Dot("DecodeString").Call(Id("eventBase64")),
					Err().Op("!=").Nil(),
				).Block(
					Return(),
				),
				Id("eventBinaries").Op("=").Append(Id("eventBinaries"), Id("eventBinary")),
			),
		),
		Return(),
	).Line()
}

func generateParseEventsFunc(file *File) {
	file.Func().Id("parseEvents").Params(
		Id("base64Binaries").Index().Index().Byte(),
	).Params(
		Id("evts").Index().Op("*").Id("Event"),
		Err().Error(),
	).Block(
		Id("decoder").Op(":=").Qual("github.com/gagliardetto/binary", "NewDecoderWithEncoding").Call(
			Nil(),
			Qual("github.com/gagliardetto/binary", "EncodingBorsh"),
		),

		For(List(Id("_"), Id("eventBinary")).Op(":=").Range().Id("base64Binaries")).Block(
			If(Len(Id("eventBinary")).Op("<").Lit(8)).Block(
				Continue(),
			),
			Id("eventDiscriminator").Op(":=").Qual("github.com/gagliardetto/binary", "TypeID").Call(Id("eventBinary").Index(Lit(0), Lit(8))),
			If(
				List(Id("eventType"), Id("ok")).Op(":=").Id("eventTypes").Index(Id("eventDiscriminator")),
				Id("ok"),
			).Block(
				Id("eventData").Op(":=").Qual("reflect", "New").Call(Id("eventType")).Dot("Interface").Call().Assert(Id("EventData")),
				Id("decoder").Dot("Reset").Call(Id("eventBinary")),
				If(
					Err().Op("=").Id("eventData").Dot("UnmarshalWithDecoder").Call(Id("decoder")),
					Err().Op("!=").Nil(),
				).Block(
					Err().Op("=").Qual("fmt", "Errorf").Call(Lit("failed to unmarshal event %s: %w"), Id("eventType").Dot("String").Call(), Err()),
					Return(),
				),
				Id("evts").Op("=").Append(Id("evts"), Op("&").Id("Event").Values(Dict{
					Id("Name"): Id("eventNames").Index(Id("eventDiscriminator")),
					Id("Data"): Id("eventData"),
				})),
			),
		),
		Return(),
	).Line()
}
