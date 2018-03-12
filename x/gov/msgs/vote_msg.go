package gov

import (
	"encoding/json"
	"fmt"

	crypto "github.com/tendermint/go-crypto"
)

type VoteMsg struct {
	Voter      crypto.address //  address of the voter
	ProposalID int64          //  proposalID of the proposal
	Option     string         //  option from OptionSet chosen by the voter
}

func NewVoteMsg(voter crypto.address, proposalId int64, option string) SendMsg {
	return VoteMsg{
		Voter:      voter,
		ProposalID: proposalID,
		Option:     option,
	}
}

// Implements Msg.
func (msg VoteMsg) Type() string { return "gov" }

// Implements Msg.
func (msg VoteMsg) ValidateBasic() sdk.Error {

	if len(msg.Voter) == 0 {
		return ErrInvalidAddress(msg.Voter.String())
	}
	if msg.Option != "Yes" || msg.Option != "No" || msg.Option != "NoWithVeto" || msg.Option != "Abstain" {
		return ErrInvalidAttribute(msg.Option) // TODO: Proper Error
	}
	return nil
}

func (msg VoteMsg) String() string {
	return fmt.Sprintf("VoteMsg{%v - %v}", msg.ProposalId, msg.Option)
}

// Implements Msg.
func (msg VoteMsg) Get(key interface{}) (value interface{}) {
	return nil
}

// Implements Msg.
func (msg VoteMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg) // XXX: ensure some canonical form
	if err != nil {
		panic(err)
	}
	return b
}

// Implements Msg.
func (msg VoteMsg) GetSigners() []crypto.Address {
	return []crypto.Address{msg.Voter}
}
