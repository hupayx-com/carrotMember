package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	// ModuleName defines the module name
	ModuleName = "carrotmember"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_carrotmember"
)

var (
	// Keys for store prefixes
	// Last* values are constant during a block.
	MemberKey = []byte{0x11}
)

func GetRewardMemberKey(memAddr sdk.AccAddress) []byte {
	return append(MemberKey, address.MustLengthPrefix(memAddr)...)
}

func KeyPrefix(p string) []byte {
	return []byte(p)
}
