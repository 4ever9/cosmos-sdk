package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/tendermint/tendermint/crypto"
)

// ModuleAccount defines an account type for pools that hold tokens in an escrow
type ModuleAccount interface {
	auth.Account

	Name() string
}

//-----------------------------------------------------------------------------
// Module Holder Account

var _ ModuleAccount = (*ModuleHolderAccount)(nil)

// ModuleHolderAccount defines an account for modules that holds coins on a pool
type ModuleHolderAccount struct {
	*auth.BaseAccount

	PoolName string `json:"name"` // name of the pool
}

// NewModuleHolderAccount creates a new ModuleHolderAccount instance
func NewModuleHolderAccount(name string) *ModuleHolderAccount {
	moduleAddress := sdk.AccAddress(crypto.AddressHash([]byte(name)))

	baseAcc := auth.NewBaseAccountWithAddress(moduleAddress)
	return &ModuleHolderAccount{
		BaseAccount: &baseAcc,
		PoolName:    name,
	}
}

// Name returns the the name of the holder's module
func (pha ModuleHolderAccount) Name() string {
	return pha.PoolName
}

// SetPubKey - Implements Account
func (pha *ModuleHolderAccount) SetPubKey(pubKey crypto.PubKey) error {
	return fmt.Errorf("not supported for pool accounts")
}

// SetSequence - Implements Account
func (pha *ModuleHolderAccount) SetSequence(seq uint64) error {
	return fmt.Errorf("not supported for pool accounts")
}

// String follows stringer interface
func (pha ModuleHolderAccount) String() string {
	// we ignore the other fields as they will always be empty
	return fmt.Sprintf(`Pool Holder Account:
Address:  %s
Coins:    %s
Name:     %s`,
		pha.Address, pha.Coins, pha.PoolName)
}

//-----------------------------------------------------------------------------
// Module Minter Account

var _ ModuleAccount = (*ModuleMinterAccount)(nil)

// ModuleMinterAccount defines an account for modules that holds coins on a pool
type ModuleMinterAccount struct {
	*ModuleHolderAccount
}

// NewModuleMinterAccount creates a new  ModuleMinterAccount instance
func NewModuleMinterAccount(name string) *ModuleMinterAccount {
	moduleHolderAcc := NewModuleHolderAccount(name)

	return &ModuleMinterAccount{ModuleHolderAccount: moduleHolderAcc}
}

// String follows stringer interface
func (pma ModuleMinterAccount) String() string {
	// we ignore the other fields as they will always be empty
	return fmt.Sprintf(`Pool Minter Account:
Address: %s
Coins:   %s
Name:    %s`,
		pma.Address, pma.Coins, pma.PoolName)
}
