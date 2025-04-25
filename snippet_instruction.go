package main

import (
	. "github.com/dave/jennifer/jen"
)

func generateInstructionSnippet(file *File) {
	generateDecodeInstructionsFunc(file)
}

func generateDecodeInstructionsFunc(file *File) {
	file.Func().Id("DecodeInstructions").Params(
		Id("message").Op("*").Qual("github.com/gagliardetto/solana-go", "Message"),
	).Params(
		Id("instructions").Index().Op("*").Id("Instruction"),
		Err().Error(),
	).Block(
		For(List(Id("_"), Id("ins")).Op(":=").Range().Id("message").Dot("Instructions")).Block(
			Var().Id("programID").Qual("github.com/gagliardetto/solana-go", "PublicKey"),
			If(
				List(Id("programID"), Err()).Op("=").Id("message").Dot("Program").Call(Id("ins").Dot("ProgramIDIndex")),
				Err().Op("!=").Nil(),
			).Block(
				Return(),
			),
			If(Op("!").Id("programID").Dot("Equals").Call(Id("ProgramID"))).Block(
				Continue(),
			),
			Var().Id("accounts").Index().Op("*").Qual("github.com/gagliardetto/solana-go", "AccountMeta"),
			If(
				List(Id("accounts"), Err()).Op("=").Id("ins").Dot("ResolveInstructionAccounts").Call(Id("message")),
				Err().Op("!=").Nil(),
			).Block(
				Return(),
			),
			Var().Id("insDecoded").Op("*").Id("Instruction"),
			If(
				List(Id("insDecoded"), Err()).Op("=").Id("DecodeInstruction").Call(Id("accounts"), Id("ins").Dot("Data")),
				Err().Op("!=").Nil(),
			).Block(
				Return(),
			),
			Id("instructions").Op("=").Append(Id("instructions"), Id("insDecoded")),
		),
		Return(),
	)
}
