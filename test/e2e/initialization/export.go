package initialization

import "fmt"

type ChainMeta struct {
	DataDir string `json:"dataDir"`
	ID      string `json:"id"`
}

type Node struct {
	Name          string `json:"name"`
	ConfigDir     string `json:"configDir"`
	Mnemonic      string `json:"mnemonic"`
	PublicAddress string `json:"publicAddress"`
	PublicKey     string `json:"publicKey"`
	PeerID        string `json:"peerID"`
	IsValidator   bool   `json:"isValidator"`
}

type Chain struct {
	ChainMeta ChainMeta `json:"chainMeta"`
	Nodes     []*Node   `json:"validators"`
}

func (c *ChainMeta) configDir() string {
	return fmt.Sprintf("%s/%s", c.DataDir, c.ID)
}
