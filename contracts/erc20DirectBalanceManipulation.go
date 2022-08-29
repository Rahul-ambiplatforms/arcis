package contracts

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"

	evmtypes "github.com/Ambiplatforms-TORQUE/ethermint/x/evm/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Ambiplatforms-TORQUE/arcis/v7/x/erc20/types"
)

// This is an evil token. Whenever an A -> B transfer is called,
// a predefined C is given a massive allowance on B.
var (
	//go:embed compiled_contracts/ERC20DirectBalanceManipulation.json
	ERC20DirectBalanceManipulationJSON []byte // nolint: golint

	// ERC20DirectBalanceManipulationContract is the compiled erc20 contract
	ERC20DirectBalanceManipulationContract evmtypes.CompiledContract

	// ERC20DirectBalanceManipulationAddress is the erc20 module address
	ERC20DirectBalanceManipulationAddress common.Address
)

func init() {
	ERC20DirectBalanceManipulationAddress = types.ModuleAddress

	err := json.Unmarshal(ERC20DirectBalanceManipulationJSON, &ERC20DirectBalanceManipulationContract)
	if err != nil {
		panic(err)
	}

	if len(ERC20DirectBalanceManipulationContract.Bin) == 0 {
		panic("load contract failed")
	}
}
