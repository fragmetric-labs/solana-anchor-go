package main

const EventFileSnippet = `
type Event struct {
	Name string
	Data EventData
}

type EventData interface {
	UnmarshalWithDecoder(decoder *ag_binary.Decoder) error
	isEventData()
}

const eventLogPrefix = "Program data: "

func DecodeEvents(txData *ag_rpc.GetTransactionResult, targetProgramId ag_solanago.PublicKey, getAddressTables func(altAddresses []ag_solanago.PublicKey) (tables map[ag_solanago.PublicKey]ag_solanago.PublicKeySlice, err error)) (evts []*Event, err error) {
	var tx *ag_solanago.Transaction
	if tx, err = txData.Transaction.GetTransaction(); err != nil {
		return
	}

	altAddresses := make([]ag_solanago.PublicKey, len(tx.Message.AddressTableLookups))
	for i, alt := range tx.Message.AddressTableLookups {
		altAddresses[i] = alt.AccountKey
	}
	if len(altAddresses) > 0 {
		var tables map[ag_solanago.PublicKey]ag_solanago.PublicKeySlice
		if tables, err = getAddressTables(altAddresses); err != nil {
			return
		}
		tx.Message.SetAddressTables(tables)
		if err = tx.Message.ResolveLookups(); err != nil {
			return
		}
	}

	var base64Binaries [][]byte
	logMessageEventBinaries, err := decodeEventsFromLogMessage(txData.Meta.LogMessages)
	if err != nil {
		return
	}

	emitedCPIEventBinaries, err := decodeEventsFromEmitCPI(txData.Meta.InnerInstructions, tx.Message.AccountKeys, targetProgramId)
	if err != nil {
		return
	}

	base64Binaries = append(base64Binaries, logMessageEventBinaries...)
	base64Binaries = append(base64Binaries, emitedCPIEventBinaries...)
	evts, err = parseEvents(base64Binaries)
	return
}

func decodeEventsFromLogMessage(logMessages []string) (eventBinaries [][]byte, err error) {
	for _, log := range logMessages {
		if strings.HasPrefix(log, eventLogPrefix) {
			eventBase64 := log[len(eventLogPrefix):]

			var eventBinary []byte
			if eventBinary, err = base64.StdEncoding.DecodeString(eventBase64); err != nil {
				err = fmt.Errorf("failed to decode logMessage event: %s", eventBase64)
				return
			}
			eventBinaries = append(eventBinaries, eventBinary)
		}
	}
	return
}

func decodeEventsFromEmitCPI(InnerInstructions []ag_rpc.InnerInstruction, accountKeys ag_solanago.PublicKeySlice, targetProgramId ag_solanago.PublicKey) (eventBinaries [][]byte, err error) {
	for _, parsedIx := range InnerInstructions {
		for _, ix := range parsedIx.Instructions {
			if accountKeys[ix.ProgramIDIndex] != targetProgramId {
				continue
			}

			var ixData []byte
			if ixData, err = ag_base58.Decode(ix.Data.String()); err != nil {
				return
			}
			if len(ixData) < 8 {
				continue
			}

			eventBase64 := base64.StdEncoding.EncodeToString(ixData[8:])
			var eventBinary []byte
			if eventBinary, err = base64.StdEncoding.DecodeString(eventBase64); err != nil {
				return
			}
			eventBinaries = append(eventBinaries, eventBinary)
		}
	}
	return
}

func parseEvents(base64Binaries [][]byte) (evts []*Event, err error) {
	decoder := ag_binary.NewDecoderWithEncoding(nil, ag_binary.EncodingBorsh)

	for _, eventBinary := range base64Binaries {
		if len(eventBinary) < 8 {
			continue
		}
		eventDiscriminator := ag_binary.TypeID(eventBinary[:8])
		if eventType, ok := eventTypes[eventDiscriminator]; ok {
			eventData := reflect.New(eventType).Interface().(EventData)
			decoder.Reset(eventBinary)
			if err = eventData.UnmarshalWithDecoder(decoder); err != nil {
				err = fmt.Errorf("failed to unmarshal event %s: %w", eventType.String(), err)
				return
			}
			evts = append(evts, &Event{
				Name: eventNames[eventDiscriminator],
				Data: eventData,
			})
		}
	}
	return
}
`

const ErrorFileSnippet = `
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
`

const InstructionFileSnippet = `
func DecodeInstructions(message *ag_solanago.Message) (instructions []*Instruction, err error) {
	for _, ins := range message.Instructions {
		var programID ag_solanago.PublicKey
		if programID, err = message.Program(ins.ProgramIDIndex); err != nil {
			return
		}
		if !programID.Equals(ProgramID) {
			continue
		}
		var accounts []*ag_solanago.AccountMeta
		if accounts, err = ins.ResolveInstructionAccounts(message); err != nil {
			return
		}
		var insDecoded *Instruction
		if insDecoded, err = DecodeInstruction(accounts, ins.Data); err != nil {
			return
		}
		instructions = append(instructions, insDecoded)
	}
	return
}
`
