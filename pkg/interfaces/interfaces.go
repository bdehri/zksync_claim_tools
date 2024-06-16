package interfaces

type Config struct {
	Wallets       []Wallet `yaml:"wallets"`
	TargetAddress string   `yaml:"target_address"`
	RpcUrl        string   `yaml:"rpc_url"`
	GatherAll     bool     `yaml:"gather_all"`
	Approve       bool     `yaml:"approve"`
	OnlyApprove   bool     `yaml:"only_approve"`
}

type Wallet struct {
	PrivateKey   string   `yaml:"private_key"`
	MerkleIndex  string   `yaml:"merkle_index"`
	MerkleProofs []string `yaml:"merkle_proofs"`
	TokenAmount  string   `yaml:"token_amount"`
}
