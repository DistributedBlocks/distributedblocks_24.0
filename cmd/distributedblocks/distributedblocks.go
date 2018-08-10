package main

import (
	_ "net/http/pprof"

	"github.com/skycoin/skycoin/src/skycoin"
	"github.com/skycoin/skycoin/src/util/logging"
	"github.com/skycoin/skycoin/src/visor"
)

var (
	// Version of the node. Can be set by -ldflags
	Version = "0.24.1"
	// Commit ID. Can be set by -ldflags
	Commit = ""
	// Branch name. Can be set by -ldflags
	Branch = ""
	// ConfigMode (possible values are "", "STANDALONE_CLIENT").
	// This is used to change the default configuration.
	// Can be set by -ldflags
	ConfigMode = ""

	logger = logging.MustGetLogger("main")

	// GenesisSignatureStr hex string of genesis signature
	GenesisSignatureStr = "b248bf0606c85c87d86fe61f0645b98a30ccf213dbe221004b1a1898423215664cf0b094e18cb2a7f4c6ce113d4544891c2c876f8a6d69de98a9ae92b49a6c6501"
	// GenesisAddressStr genesis address string
	GenesisAddressStr = "2FiHFKWKqhnDCio47z9Q2o28qWF3yePUCg7"
	// BlockchainPubkeyStr pubic key string
	BlockchainPubkeyStr = "02ae5c7701dcd71e7448116e5766130f2eacecc585f3e320b1d85f2951382a424a"
	// BlockchainSeckeyStr empty private key string
	BlockchainSeckeyStr = ""

	// GenesisTimestamp genesis block create unix time
	GenesisTimestamp uint64 = 1530619056
	// GenesisCoinVolume represents the coin capacity
	GenesisCoinVolume uint64 = 100e12

	// DefaultConnections the default trust node addresses
	DefaultConnections = []string{
	"139.162.161.41:9610",
"139.162.7.215:9610",
"139.162.7.215:9611",
"139.162.7.215:9612",
"18.218.142.16:9600",
"178.79.170.194:9600",
"13.58.196.172:9600",
	}
)

func main() {
	// get node config
	nodeConfig := skycoin.NewNodeConfig(ConfigMode, skycoin.NodeParameters{
		GenesisSignatureStr: GenesisSignatureStr,
		GenesisAddressStr:   GenesisAddressStr,
		GenesisCoinVolume:   GenesisCoinVolume,
		GenesisTimestamp:    GenesisTimestamp,
		BlockchainPubkeyStr: BlockchainPubkeyStr,
		BlockchainSeckeyStr: BlockchainSeckeyStr,
		DefaultConnections:  DefaultConnections,
		PeerListURL:         "https://distributedblocks.com/peers.txt",
		Port:                9600,
		WebInterfacePort:    9620,
		DataDirectory:       "$HOME/.distributedblocks",
		ProfileCPUFile:      "skycoin.prof",
	})

	// create a new fiber coin instance
	coin := skycoin.NewCoin(
		skycoin.Config{
			Node: *nodeConfig,
			Build: visor.BuildInfo{
				Version: Version,
				Commit:  Commit,
				Branch:  Branch,
			},
		},
		logger,
	)

	// parse config values
	coin.ParseConfig()

	// run fiber coin node
	coin.Run()
}
