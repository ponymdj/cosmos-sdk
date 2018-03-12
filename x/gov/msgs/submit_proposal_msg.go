package gov

import (
	"encoding/json"
	"fmt"

	crypto "github.com/tendermint/go-crypto"
)

//----------------------------------------
// SubmitProposalMsg

type SubmitProposalMsg struct {
	Title          string         //  Title of the proposal
	Description    string         //  Description of the proposal
	ProposalType   string         //  Type of proposal. Initial set {PlainTextProposal, SoftwareUpgradeProposal}
	Proposer       crypto.address //  Address of the proposer
	InitialDeposit sdk.Coins      //  Initial deposit paid by sender. Must be strictly positive.
}

func NewSubmitProposalMsg(title string, description string, proposalType string, initialDeposit int64) SendMsg {
	return SubmitProposalMsg{
		Title:          title,
		Description:    description,
		ProposalType:   proposalType,
		InitialDeposit: initialDeposit,
	}
}

// Implements Msg.
func (msg SubmitProposalMsg) Type() string { return "gov" }

// Implements Msg.
func (msg SubmitProposalMsg) ValidateBasic() sdk.Error {

	if len(msg.Title) == 0 {
		return ErrInvalidAttribute(msg.Title) // TODO: Proper Error
	}
	if len(msg.Description) == 0 {
		return ErrInvalidAttribute(msg.Description) // TODO: Proper Error
	}
	if len(msg.ProposalType) == 0 {
		return ErrInvalidAttribute(msg.ProposalType) // TODO: Proper Error
	}
	if len(msg.Proposer) == 0 {
		return ErrInvalidAddress(msg.Proposer.String())
	}
	if !msg.Amount.IsValid() {
		return ErrInvalidCoins(msg.Amount.String())
	}
	if !msg.Amount.IsPositive() {
		return ErrInvalidCoins(msg.Amount.String())
	}
	return nil
}

func (msg SubmitProposalMsg) String() string {
	return fmt.Sprintf("SubmitProposalMsg{%v, %v, %v, %v}", msg.Title, msg.Description, msg.ProposalType, msg.InitialDeposit)
}

// Implements Msg.
func (msg SubmitProposalMsg) Get(key interface{}) (value interface{}) {
	return nil
}

// Implements Msg.
func (msg SubmitProposalMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg) // XXX: ensure some canonical form
	if err != nil {
		panic(err)
	}
	return b
}

// Implements Msg.
func (msg SendMsg) GetSigners() []crypto.Address {
	return []crypto.Address{msg.Proposer}
}
