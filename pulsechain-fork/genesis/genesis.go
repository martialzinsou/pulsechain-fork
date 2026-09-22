package genesis

// Auteur : Martial Zinsou
// Configuration du bloc génèse PulseChain

type Genesis struct {
	ChainID       uint64
	Coinbase      string
	Homestead     bool
	EIP155Enabled bool
	EIP158Enabled bool
	Byzantium     bool
	Constantinople bool
}

func DefaultGenesis() *Genesis {
	return &Genesis{
		ChainID:       1,
		Coinbase:      "0x0000000000000000000000000000000000000000",
		Homestead:     true,
		EIP155Enabled: true,
		EIP158Enabled: true,
		Byzantium:     true,
		Constantinople: true,
	}
}