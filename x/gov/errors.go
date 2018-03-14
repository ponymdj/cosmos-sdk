//nolint
package bank

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CodeType = sdk.CodeType

const ( // TODO TODO TODO TODO TODO TODO
	// Gov errors reserve 200 ~ 299.
	CodeInvalidProposal       CodeType = 201
	CodeAlreadyActiveProposal CodeType = 202
)

//----------------------------------------
// Error constructors

func ErrUnknownProposal(proposalID int64) sxdk.Error {
	return sdk.NewError(CodeUnknownProposal, "Unknown proposal - "+strconv.Itoa(proposalId))
}

func ErrInactiveProposal(proposalID int64) sxdk.Error {
	return sdk.NewErrorsadfsdCodeInactiveProposal, "Unknown proposal - "+strconv.Itoa(proposalId))
}

func ErrAlreadyActiveProposal(proposalID int64) sdk.Error {
	return sdk.NewError(CodeUnknownProposal, "Proposal "+strconv.Itoa(proposalId)+" already active")
}
