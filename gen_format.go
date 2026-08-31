package main

import (
	. "github.com/dave/jennifer/jen"
)

// genFormatFile emits a "format.go" file into the generated package, replacing the
// github.com/gagliardetto/solana-go/text and .../text/format packages that fluxrpc/solana-go
// has no equivalent for. It preserves the generated code's public API (EncodableToTree,
// Encoder, Option) so EncodeToTree/TextEncode keep compiling and behaving informatively;
// nothing in this codebase's own tests exercises these programmatically, so exact cosmetic
// fidelity with the old SDK's output isn't required.
func genFormatFile(idl IDL) (*FileWrapper, error) {
	file := NewGoFile(idl.Metadata.Name, true)

	// EncodableToTree replaces github.com/gagliardetto/solana-go/text.EncodableToTree.
	file.Add(
		Type().Id("EncodableToTree").Interface(
			Id("EncodeToTree").Params(Id("parent").Qual(PkgTreeout, "Branches")),
		).Line(),
	)

	// Option/Encoder replace github.com/gagliardetto/solana-go/text.{Option,Encoder}.
	file.Add(Type().Id("Option").Struct().Line())
	file.Add(
		Type().Id("Encoder").Struct(
			Id("w").Qual("io", "Writer"),
		).Line(),
	)
	file.Add(
		Func().Id("NewEncoder").Params(Id("w").Qual("io", "Writer")).Op("*").Id("Encoder").Block(
			Return(Op("&").Id("Encoder").Values(Dict{Id("w"): Id("w")})),
		).Line(),
	)
	file.Add(
		Func().Params(Id("e").Op("*").Id("Encoder")).Id("Encode").
			Params(Id("v").Interface(), Id("opt").Op("*").Id("Option")).
			Error().
			BlockFunc(func(body *Group) {
				body.List(Id("_"), Id("err")).Op(":=").Qual("fmt", "Fprintf").Call(Id("e").Dot("w"), Lit("%+v\n"), Id("v"))
				body.Return(Id("err"))
			}).Line(),
	)

	// formatProgram/formatInstruction/formatParam/formatMeta replace the
	// github.com/gagliardetto/solana-go/text/format helpers used inside each
	// instruction's EncodeToTree method body.
	file.Add(
		Func().Id("formatProgram").
			Params(Id("name").String(), Id("programID").Qual(PkgSolanaGo, "PublicKey")).
			String().
			Block(
				Return(Qual("fmt", "Sprintf").Call(Lit("Program: %s: %s"), Id("name"), Id("programID").Dot("String").Call())),
			).Line(),
	)
	file.Add(
		Func().Id("formatInstruction").
			Params(Id("name").String()).
			String().
			Block(
				Return(Qual("fmt", "Sprintf").Call(Lit("Instruction: %s"), Id("name"))),
			).Line(),
	)
	file.Add(
		Func().Id("formatParam").
			Params(Id("label").String(), Id("value").Interface()).
			String().
			Block(
				Return(Qual("fmt", "Sprintf").Call(Lit("%s: %v"), Id("label"), Id("value"))),
			).Line(),
	)
	file.Add(
		Func().Id("formatMeta").
			Params(Id("label").String(), Id("meta").Op("*").Qual(PkgSolanaGo, "AccountMeta")).
			String().
			BlockFunc(func(body *Group) {
				body.If(Id("meta").Op("==").Nil()).Block(
					Return(Qual("fmt", "Sprintf").Call(Lit("%s: <nil>"), Id("label"))),
				)
				body.Var().Id("flags").Index().String()
				body.If(Id("meta").Dot("IsWritable")).Block(
					Id("flags").Op("=").Append(Id("flags"), Lit("WRITE")),
				)
				body.If(Id("meta").Dot("IsSigner")).Block(
					Id("flags").Op("=").Append(Id("flags"), Lit("SIGNER")),
				)
				body.If(Len(Id("flags")).Op("==").Lit(0)).Block(
					Return(Qual("fmt", "Sprintf").Call(Lit("%s: %s"), Id("label"), Id("meta").Dot("PublicKey").Dot("String").Call())),
				)
				body.Return(
					Qual("fmt", "Sprintf").Call(
						Lit("%s: %s [%s]"),
						Id("label"),
						Id("meta").Dot("PublicKey").Dot("String").Call(),
						Qual("strings", "Join").Call(Id("flags"), Lit(", ")),
					),
				)
			}).Line(),
	)

	return &FileWrapper{
		Name: "format",
		File: file,
	}, nil
}
