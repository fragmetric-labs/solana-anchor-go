package main

import (
	. "github.com/dave/jennifer/jen"
)

func generateErrorSnippet(file *File) {
	file.Type().Id("CustomError").Interface(
		Id("Code").Params().Int(),
		Id("Name").Params().String(),
		Id("Error").Params().String(),
	).Line()

	file.Type().Id("customErrorDef").Struct(
		Id("code").Int(),
		Id("name").String(),
		Id("msg").String(),
	)

	file.Func().Params(Id("e").Op("*").Id("customErrorDef")).Id("Code").Params().Int().Block(
		Return(Id("e").Dot("code")),
	).Line()

	file.Func().Params(Id("e").Op("*").Id("customErrorDef")).Id("Name").Params().String().Block(
		Return(Id("e").Dot("name")),
	).Line()

	file.Func().Params(Id("e").Op("*").Id("customErrorDef")).Id("Error").Params().String().Block(
		Return(
			Qual("fmt", "Sprintf").Call(
				Lit("%s(%d): %s"),
				Id("e").Dot("name"),
				Id("e").Dot("code"),
				Id("e").Dot("msg"),
			),
		),
	).Line()

	generateDecodeCustomErrorFunc(file)
	generateDecodeErrorCodeFunc(file)
}

func generateDecodeCustomErrorFunc(file *File) {
	file.Func().Id("DecodeCustomError").Params(
		Id("rpcErr").Error(),
	).Params(
		Err().Error(),
		Id("ok").Bool(),
	).Block(
		If(
			List(Id("errCode"), Id("o")).Op(":=").Id("decodeErrorCode").Call(Id("rpcErr")),
			Id("o"),
		).Block(
			If(
				List(Id("customErr"), Id("o")).Op(":=").Id("Errors").Index(Id("errCode")),
				Id("o"),
			).Block(
				Err().Op("=").Id("customErr"),
				Id("ok").Op("=").Lit(true),
				Return(),
			),
		),
		Return(),
	).Line()
}

func generateDecodeErrorCodeFunc(file *File) {
	file.Func().Id("decodeErrorCode").Params(
		Id("rpcErr").Error(),
	).Params(
		Id("errorCode").Int(),
		Id("ok").Bool(),
	).Block(
		Var().Id("jErr").Op("*").Qual("github.com/gagliardetto/solana-go/rpc/jsonrpc", "RPCError"),
		If(
			Qual("errors", "As").Call(Id("rpcErr"), Op("&").Id("jErr")).
				Op("&&").Id("jErr").Dot("Data").Op("!=").Nil(),
		).Block(
			If(
				List(Id("root"), Id("o")).Op(":=").Id("jErr").Dot("Data").Assert(Map(String()).Interface()),
				Id("o"),
			).Block(
				If(
					List(Id("rootErr"), Id("o")).Op(":=").Id("root").Index(Lit("err")).Assert(Map(String()).Interface()),
					Id("o"),
				).Block(
					If(
						List(Id("rootErrInstructionError"), Id("o")).Op(":=").Id("rootErr").Index(Lit("InstructionError")),
						Id("o"),
					).Block(
						If(
							List(Id("rootErrInstructionErrorItems"), Id("o")).Op(":=").Id("rootErrInstructionError").Assert(Index().Interface()),
							Id("o"),
						).Block(
							If(Len(Id("rootErrInstructionErrorItems")).Op("==").Lit(2)).Block(
								If(
									List(Id("v"), Id("o")).Op(":=").Id("rootErrInstructionErrorItems").Index(Lit(1)).Assert(Map(String()).Interface()),
									Id("o"),
								).Block(
									If(
										List(Id("v2"), Id("o")).Op(":=").Id("v").Index(Lit("Custom")).Assert(Qual("encoding/json", "Number")),
										Id("o"),
									).Block(
										If(
											List(Id("code"), Err()).Op(":=").Id("v2").Dot("Int64").Call(),
											Err().Op("==").Nil(),
										).Block(
											Id("ok").Op("=").Lit(true),
											Id("errorCode").Op("=").Int().Call(Id("code")),
										),
									).Else().If(
										List(Id("v2"), Id("o")).Op(":=").Id("v").Index(Lit("Custom")).Assert(Float64()),
										Id("o"),
									).Block(
										Id("ok").Op("=").Lit(true),
										Id("errorCode").Op("=").Int().Call(Id("v2")),
									),
								),
							),
						),
					),
				),
			),
		),
		Return(),
	).Line()
}
