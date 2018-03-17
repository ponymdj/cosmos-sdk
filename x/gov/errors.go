//nolint
package gov

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CodeType = sdk.CodeType

const ( // TODO TODO TODO TODO TODO TODO
	// Gov errors reserve 200 ~ 299.
	CodeUnknownProposal       CodeType = 201
	CodeInactiveProposal      CodeType = 202
	CodeAlreadyActiveProposal CodeType = 203
)

//----------------------------------------
// Error constructors

func ErrUnknownProposal(proposalID int64) sdk.Error {
	return sdk.NewError(CodeUnknownProposal, "Unknown proposal - "+strconv.FormatInt(proposalID, 10))
}

func ErrInactiveProposal(proposalID int64) sdk.Error {
	return sdk.NewError(CodeInactiveProposal, "Unknown proposal - "+strconv.FormatInt(proposalID, 10))
}

func ErrAlreadyActiveProposal(proposalID int64) sdk.Error {
	return sdk.NewError(CodeAlreadyActiveProposal, "Proposal "+strconv.FormatInt(proposalID, 10)+" already active")
}
