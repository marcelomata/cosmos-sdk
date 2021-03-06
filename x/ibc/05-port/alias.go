package port

// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/ibc/05-port/keeper
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/ibc/05-port/types

import (
	"github.com/cosmos/cosmos-sdk/x/ibc/05-port/keeper"
	"github.com/cosmos/cosmos-sdk/x/ibc/05-port/types"
)

const (
	SubModuleName = types.SubModuleName
	StoreKey      = types.StoreKey
	RouterKey     = types.RouterKey
	QuerierRoute  = types.QuerierRoute
)

var (
	// functions aliases
	NewKeeper       = keeper.NewKeeper
	NewRouter       = types.NewRouter
	ErrPortExists   = types.ErrPortExists
	ErrPortNotFound = types.ErrPortNotFound
	ErrInvalidPort  = types.ErrInvalidPort
	ErrInvalidRoute = types.ErrInvalidRoute
)

type (
	Keeper    = keeper.Keeper
	Router    = types.Router
	IBCModule = types.IBCModule
)
