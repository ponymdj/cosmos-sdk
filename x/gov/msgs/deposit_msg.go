package gov

import (
	"encoding/json"
	"fmt"
)

type DepositMsg struct {
	ProposalId int64          `json:"proposal_id"` // ID of the proposal
	Depositer  crypto.address `json:"depositer"`   // Address of the depositer
	Amount     sdk.Coins      `json:"amount"`      // Coins to add to the proposal's deposit
}

func NewDepositMsgMsg(proposalId int64, depositer crypto.address, amount sdk.Coins) DepositMsg {
	return DepositMsg{
		ProposalId: proposalId,
		Depositer:  depositer,
		Amount:     Amount,
	}
}

// Implements Msg.
func (msg DepositMsg) Type() string { return "gov" }

// Implements Msg.
func (msg DepositMsg) ValidateBasic() sdk.Error {
	if len(msg.Depositer) == 0 {
		return ErrInvalidAddress(msg.Depositer.String())
	}
	if !msg.Amount.IsValid() {
		return ErrInvalidCoins(msg.Amount.String())
	}
	if !msg.Amount.IsPositive() {
		return ErrInvalidCoins(msg.Amount.String())
	}
	return nil
}

func (msg DepositMsg) String() string {
	return fmt.Sprintf("DepositMsg{%v=>%v: %v}", msg.Depositer, msg.ProposerId, msg.Amount)
}

// Implements Msg.
func (msg DepositMsg) Get(key interface{}) (value interface{}) {
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
func (msg SubmitProposalMsg) GetSigners() []crypto.Address {
	return []crypto.Address{msg.Depositer}
}
