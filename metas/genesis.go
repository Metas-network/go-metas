package metas

import (
	"math/big"

	"github.com/Metas-network/lachesis-base/hash"
	"github.com/Metas-network/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Metas-network/go-metas/inter"
	"github.com/Metas-network/go-metas/metas/genesis"
	"github.com/Metas-network/go-metas/metas/genesis/gpos"
)

type Genesis struct {
	Accounts    genesis.Accounts
	Storage     genesis.Storage
	Delegations genesis.Delegations
	Blocks      genesis.Blocks
	RawEvmItems genesis.RawEvmItems
	Validators  gpos.Validators

	FirstEpoch    idx.Epoch
	PrevEpochTime inter.Timestamp
	Time          inter.Timestamp
	ExtraData     []byte

	TotalSupply *big.Int

	DriverOwner common.Address

	Rules Rules

	Hash func() hash.Hash
}
