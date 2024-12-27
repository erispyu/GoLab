package main

func AmountCheckNAPAS(amount float64) ErrorCode {
	if amount >= MaxAmountNAPAS {
		return ErrorTxnAmountTooHigh
	} else if amount >= MinAmountNAPAS && amount < MaxAmountNAPAS {
		return Success
	} else {
		return ErrorTxnAmountTooLow
	}
}

func AmountCheckCITAD(amount float64, isATMCard bool) ErrorCode {
	if amount > MaxAmountCITAD {
		return ErrorTxnAmountTooHigh
	} else if amount >= MinAmountCITAD && amount <= MaxAmountCITAD {
		if !isATMCard {
			return Success
		} else {
			return ErrorInvalidBankAccount
		}
	} else {
		return ErrorTxnAmountTooLow
	}
}

func CheckRoute(amount float64, isATMCard bool) (ErrorCode, Route) {
	amountCheckResultNAPAS := AmountCheckNAPAS(amount)
	if amountCheckResultNAPAS == Success {
		return Success, NAPAS
	}

	amountCheckResultCITAD := AmountCheckCITAD(amount, isATMCard)
	if amountCheckResultCITAD == Success {
		return Success, CITAD
	}

	return amountCheckResultCITAD, ErrorType
}
