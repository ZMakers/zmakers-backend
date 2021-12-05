package lib

func IsValidAddress(address string) bool {
	if address == "" || len(address) != 42 {
		return false
	}
	return true
}

func IsValidTxHash(txHash string) bool {
	if txHash == "" || len(txHash) != 66 {
		return false
	}
	return true
}
