package gov

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	abci "github.com/tendermint/abci/types"
)

func checkProposal(ctx sdk.Context, gm governanceMapper) {
	proposal := gm.ProposalQueuePeek(ctx)
	if proposal == nil {
		return
	}

	// Don't want to do urgent for now

	// // Urgent proposal accepted
	// if proposal.Votes.YesVotes/proposal.InitTotalVotingPower >= 2/3 {
	// 	gm.PopProposalQueue(ctx)
	// 	refund(ctx, gm, proposalId, proposal)
	// 	return checkProposal()
	// }

	// Proposal reached the end of the voting period
	if ctx.BlockHeight() == proposal.VotingStartBlock+proposal.Procedure.VotingPeriod {
		gm.ProposalQueuePop(ctx)

		// Slash validators if not voted
		for _, validatorGovInfo := range proposal.ValidatorGovInfos {
			if validatorOption.LastVoteWeight < 0 {
				// TODO: SLASH MWAHAHAHAHAHA
			}
		}

		//Proposal was accepted
		nonAbstainTotal := proposal.Votes.YesVotes + proposal.Votes.NoVotes + proposal.Votes.NoWithVetoVotes
		if proposal.YesVotes/nonAbstainTotal > 0.5 && proposal.NoWithVetoVotes/nonAbstainTotal < 1/3 { // TODO: Deal with decimals

			//  TODO:  Act upon accepting of proposal

			// Refund deposits
			for _, deposit := range proposal.Deposits {
				gm.ck.AddCoins(ctx, deposit.Depositer, deposit.Amount)
				if err != nil {
					panic("should not happen")
				}
			}

			// check next proposal recursively
			return checkProposal()
		}

		//  TODO: Prune proposal
	}
}

func NewBeginBlocker(gm governanceMapper) sdk.BeginBlocker {
	return func(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
		checkProposal(ctx, gm)
		return abci.ResponseBeginBlock{}
	}
}
