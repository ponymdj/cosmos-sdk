package gov

import (
	"github.com/cosmos/cosmos-sdk/x/bank"
	wire "github.com/tendermint/go-wire"
)

type governanceMapper struct {
	// The reference to the CoinKeeper to modify balances
	ck bank.CoinKeeper

	// The (unexposed) keys used to access the stores from the Context.
	proposalStoreKey sdk.StoreKey

	// The wire codec for binary encoding/decoding of accounts.
	cdc *wire.Codec
}

// NewGovernanceMapper returns a mapper that uses go-wire to (binary) encode and decode gov types.
func NewGovernanceMapper(key sdk.StoreKey, ck bank.CoinKeeper) accountMapper {
	cdc := wire.NewCodec()
	return governanceMapper{
		proposalStoreKey: key,
		ck:               ck,
		cdc:              cdc,
	}
}

// Returns the go-wire codec.  You may need to register interfaces
// and concrete types here, if your app's sdk.Account
// implementation includes interface fields.
// NOTE: It is not secure to expose the codec, so check out
// .Seal().
func (gm governanceMapper) WireCodec() *wire.Codec {
	return gm.cdc
}

func (gm governanceMapper) GetProposal(ctx sdk.Context, proposalID int64) Proposal {
	store := ctx.KVStore(gm.proposalStoreKey)
	bz := store.Get(proposalID)
	if bz == nil {
		return nil
	}

	proposal := &Proposal{}
	err := gm.cdc.UnmarshalBinary(bz, proposal)
	if err != nil {
		panic(err)
	}

	return proposal
}

// Implements sdk.AccountMapper.
func (gm governanceMapper) SetProposal(ctx sdk.Context, proposal Proposal) {
	store := ctx.KVStore(gm.proposalStoreKey)

	bz, err := gm.cdc.MarshalBinary(proposal)
	if err != nil {
		panic(err)
	}

	store.Set(proposalID, bz)
}

func (gm governanceMapper) getNewProposalID(ctx sdk.Context) int64 {
	store := ctx.KVStore(gm.proposalStoreKey)
	bz := store.Get([]byte("newProposalID"))
	if bz == nil {
		return nil
	}

	proposalID = new(int64)
	err := gm.cdc.UnmarshalBinaryBare(bz, proposalID)
	if err != nil {
		panic("should not happen")
	}

	bz, err := gm.cdc.MarshalBinaryBare(proposalID + 1)
	if err != nil {
		panic("should not happen")
	}

	store.Set([]byte("newProposalID"), bz)

	return proposalID
}

func (gm governanceMapper) getProposalQueue(ctx sdk.Context) ProposalQueue {
	store := ctx.KVStore(gm.proposalStoreKey)
	bz := store.Get([]byte("proposalQueue"))
	if bz == nil {
		return nil
	}

	proposalQueue := &ProposalQueue{}
	err := gm.cdc.UnmarshalBinaryBare(bz, proposalQueue)
	if err != nil {
		panic(err)
	}

	return proposalQueue
}

func (gm governanceMapper) setProposalQueue(ctx sdk.Context, proposalQueue ProposalQueue) {
	store := ctx.KVStore(gm.proposalStoreKey)

	bz, err := gm.cdc.MarshalBinaryBare(proposalQueue)
	if err != nil {
		panic(err)
	}

	store.Set([]byte("proposalQueue"), bz)
}

func (gm governanceMapper) ProposalQueuePeek(ctx sdk.Context) Proposal {
	proposalQueue := gm.getProposaljueue(ctx)
	if len(proposalQueue) == 0 {
		return nil
	}
	return gm.GetProposal(ctx, proposalQueue[0])
}

func (gm governanceMapper) ProposalQueuePop(ctx sdk.Context) Proposal {
	proposalQueue := gm.getProposalQueue(ctx)
	if len(proposalQueue) == 0 {
		return nil
	}
	frontElement, proposalQueue = proposalQueue[0], proposalQueue[1:]
	gm.setProposalQueue(ctx, proposalQueue)
	return gm.GetProposal(ctx, frontElement)
}

func (gm governanceMapper) ProposalQueuePush(ctx sdk.Context, proposal Proposal) {
	proposalQueue := append(gm.getProposalQueue(ctx), proposal.ProposalID)
	bz, err := gm.cdc.MarshalBinary(proposalQueue)
	if err != nil {
		panic(err)
	}
	store.Set([]byte("proposalQueue"), bz)
}
