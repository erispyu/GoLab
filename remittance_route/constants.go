package main

const (
	MaxAmountCITAD = 2e9 * 100 // 2 billion in cent
	MaxAmountNAPAS = 5e6 * 100 // 5 million in cent
	MinAmountCITAD = 1         // one cent
	MinAmountNAPAS = 2e3 * 100 // 2 thousand in cent
)

type ErrorCode string

const (
	Success                 ErrorCode = "Success"
	ErrorTxnAmountTooHigh   ErrorCode = "ErrorTxnAmountTooHigh"
	ErrorTxnAmountTooLow    ErrorCode = "ErrorTxnAmountTooLow"
	ErrorInvalidBankAccount ErrorCode = "ErrorInvalidBankAccount"
)

type Route string

const (
	ErrorType Route = "ErrorType"
	NAPAS     Route = "NAPAS"
	CITAD     Route = "CITAD"
)
