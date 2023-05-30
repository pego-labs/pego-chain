package systemcontracts

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

type UpgradeConfig struct {
	BeforeUpgrade upgradeHook
	AfterUpgrade  upgradeHook
	ContractAddr  common.Address
	CommitUrl     string
	Code          string
}

type Upgrade struct {
	UpgradeName string
	Configs     []*UpgradeConfig
}

type upgradeHook func(blockNumber *big.Int, contractAddr common.Address, statedb *state.StateDB) error

const (
	mainNet    = "Mainnet"
	chapelNet  = "Chapel"
	rialtoNet  = "Rialto"
	defaultNet = "Default"
)

var (
	GenesisHash common.Hash
	//upgrade config
	ramanujanUpgrade = make(map[string]*Upgrade)

	nielsUpgrade = make(map[string]*Upgrade)

	mirrorUpgrade = make(map[string]*Upgrade)

	brunoUpgrade = make(map[string]*Upgrade)

	eulerUpgrade = make(map[string]*Upgrade)

	gibbsUpgrade = make(map[string]*Upgrade)

	moranUpgrade = make(map[string]*Upgrade)
)

func init() {

}

func UpgradeBuildInSystemContract(config *params.ChainConfig, blockNumber *big.Int, statedb *state.StateDB) {
	if config == nil || blockNumber == nil || statedb == nil {
		return
	}
	var network string
	switch GenesisHash {
	/* Add mainnet genesis hash */
	case params.BSCGenesisHash:
		network = mainNet
	case params.ChapelGenesisHash:
		network = chapelNet
	case params.RialtoGenesisHash:
		network = rialtoNet
	default:
		network = defaultNet
	}

	logger := log.New("system-contract-upgrade", network)
	if config.IsOnRamanujan(blockNumber) {
		applySystemContractUpgrade(ramanujanUpgrade[network], blockNumber, statedb, logger)
	}

	if config.IsOnNiels(blockNumber) {
		applySystemContractUpgrade(nielsUpgrade[network], blockNumber, statedb, logger)
	}

	if config.IsOnMirrorSync(blockNumber) {
		applySystemContractUpgrade(mirrorUpgrade[network], blockNumber, statedb, logger)
	}

	if config.IsOnBruno(blockNumber) {
		applySystemContractUpgrade(brunoUpgrade[network], blockNumber, statedb, logger)
	}

	if config.IsOnEuler(blockNumber) {
		applySystemContractUpgrade(eulerUpgrade[network], blockNumber, statedb, logger)
	}

	if config.IsOnGibbs(blockNumber) {
		applySystemContractUpgrade(gibbsUpgrade[network], blockNumber, statedb, logger)
	}

	if config.IsOnMoran(blockNumber) {
		applySystemContractUpgrade(moranUpgrade[network], blockNumber, statedb, logger)
	}

	/*
		apply other upgrades
	*/
}

func applySystemContractUpgrade(upgrade *Upgrade, blockNumber *big.Int, statedb *state.StateDB, logger log.Logger) {
	if upgrade == nil {
		logger.Info("Empty upgrade config", "height", blockNumber.String())
		return
	}

	logger.Info(fmt.Sprintf("Apply upgrade %s at height %d", upgrade.UpgradeName, blockNumber.Int64()))
	for _, cfg := range upgrade.Configs {
		logger.Info(fmt.Sprintf("Upgrade contract %s to commit %s", cfg.ContractAddr.String(), cfg.CommitUrl))

		if cfg.BeforeUpgrade != nil {
			err := cfg.BeforeUpgrade(blockNumber, cfg.ContractAddr, statedb)
			if err != nil {
				panic(fmt.Errorf("contract address: %s, execute beforeUpgrade error: %s", cfg.ContractAddr.String(), err.Error()))
			}
		}

		newContractCode, err := hex.DecodeString(cfg.Code)
		if err != nil {
			panic(fmt.Errorf("failed to decode new contract code: %s", err.Error()))
		}
		statedb.SetCode(cfg.ContractAddr, newContractCode)

		if cfg.AfterUpgrade != nil {
			err := cfg.AfterUpgrade(blockNumber, cfg.ContractAddr, statedb)
			if err != nil {
				panic(fmt.Errorf("contract address: %s, execute afterUpgrade error: %s", cfg.ContractAddr.String(), err.Error()))
			}
		}
	}
}
