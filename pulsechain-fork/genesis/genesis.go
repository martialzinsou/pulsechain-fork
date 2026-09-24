package genesis

// Auteur : Martial Zinsou
// Configuration et spécification du bloc génèse pour le Fork de PulseChain

// Genesis définit les paramètres de lancement du réseau PulseChain Fork
type Genesis struct {
	ChainID        uint64 `json:"chainId"`
	Coinbase       string `json:"coinbase"`
	Homestead      bool   `json:"homestead"`
	EIP155Enabled  bool   `json:"eip155"`
	EIP158Enabled  bool   `json:"eip158"`
	Byzantium      bool   `json:"byzantium"`
	Constantinople bool   `json:"constantinople"`
	Petersburg     bool   `json:"petersburg"`
	Istanbul       bool   `json:"istanbul"`
	Berlin         bool   `json:"berlin"`
	London         bool   `json:"london"`
	Shanghai       bool   `json:"shanghai"`
	Period         uint64 `json:"period"`
	Epoch          uint64 `json:"epoch"`
}

// DefaultGenesis initialise la configuration par défaut du fork PulseChain (ChainID 369)
func DefaultGenesis() *Genesis {
	return &Genesis{
		ChainID:        369, // PulseChain Mainnet Chain ID
		Coinbase:       "0x2b5AD5c4795c026514f8317c7a215E218DcCD6cF", // Adresse PulseChain validateur
		Homestead:      true,
		EIP155Enabled:  true,
		EIP158Enabled:  true,
		Byzantium:      true,
		Constantinople: true,
		Petersburg:     true,
		Istanbul:       true,
		Berlin:         true,
		London:         true,
		Shanghai:       true,
		Period:         3,     // 3 secondes par bloc
		Epoch:          30000,
	}
}

// TestnetGenesis initialise la configuration pour un réseau de test PulseChain Fork (ChainID 943)
func TestnetGenesis() *Genesis {
	gen := DefaultGenesis()
	gen.ChainID = 943 // PulseChain Testnet-v4 ID
	return gen
}
