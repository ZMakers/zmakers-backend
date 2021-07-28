package eth

type callObj struct {
	To string `json:"to"`
	Value int64 `json:"value"`
	Data []byte `json:"data"`
	PrivKey string `json:"privKey"`
}

