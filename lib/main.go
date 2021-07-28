package lib

func IsValidAddress(address string) bool {
	if address == "" || len(address) != 20 {
		return false
	}
	return true
}
