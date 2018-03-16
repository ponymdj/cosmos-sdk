package commands

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"

	crypto "github.com/tendermint/go-crypto"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/builder"
	"github.com/cosmos/cosmos-sdk/examples/basecoin/app"
	"github.com/cosmos/cosmos-sdk/wire"
	coin "github.com/cosmos/cosmos-sdk/x/bank" // XXX fix
	"github.com/cosmos/cosmos-sdk/x/stake"
)

// XXX remove dependancy
func PrefixedKey(app string, key []byte) []byte {
	prefix := append([]byte(app), byte(0))
	return append(prefix, key...)
}

//nolint
var (
	fsValAddr         = flag.NewFlagSet("", flag.ContinueOnError)
	fsDelAddr         = flag.NewFlagSet("", flag.ContinueOnError)
	FlagValidatorAddr = "address"
	FlagDelegatorAddr = "delegator-address"
)

func init() {
	//Add Flags
	fsValAddr.String(FlagValidatorAddr, "", "Address of the validator/candidate")
	fsDelAddr.String(FlagDelegatorAddr, "", "Delegator hex address")

}

// create command to query for all candidates
func GetCmdQueryCandidates(cdc *wire.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "candidates",
		Short: "Query for the set of validator-candidates pubkeys",
		RunE: func(cmd *cobra.Command, args []string) error {

			var pks []crypto.PubKey

			prove := !viper.GetBool(client.FlagTrustNode)
			key := PrefixedKey(stake.Name, stake.CandidatesAddrKey)

			res, err := builder.Query(key, "gaia-store-name") // XXX move gaia store name out of here
			if err != nil {
				return err
			}

			// parse out the candidates
			candidates := new(stake.Candidates)
			err = cdc.UnmarshalJSON(res, candidates)
			if err != nil {
				return err
			}
			output, err := json.MarshalIndent(candidates, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(output))
			return nil

			// TODO output with proofs / machine parseable etc.
		},
	}

	cmd.Flags().AddFlagSet(fsDelAddr)
	return cmd
}

// get the command to query a candidate
func GetCmdQueryCandidate(cdc *wire.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "candidate",
		Short: "Query a validator-candidate account",
		RunE: func(cmd *cobra.Command, args []string) error {

			addr, err := GetAddress(viper.GetString(FlagValidatorAddr))
			if err != nil {
				return err
			}

			prove := !viper.GetBool(client.FlagTrustNode)
			key := PrefixedKey(stake.Name, stake.GetCandidateKey(addr))

			// parse out the candidate
			candidate := new(stake.Candidate)
			err = cdc.UnmarshalBinary(res, candidate)
			if err != nil {
				return err
			}
			output, err := json.MarshalIndent(candidate, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(output))
			return nil

			// TODO output with proofs / machine parseable etc.
		},
	}

	cmd.Flags().AddFlagSet(fsValAddr)
	return cmd
}

// get the command to query a single delegator bond
func GetCmdQueryDelegatorBond(cdc *wire.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delegator-bond",
		Short: "Query a delegators bond based on address and candidate pubkey",
		RunE: func(cmd *cobra.Command, args []string) error {

			addr, err := GetAddress(viper.GetString(FlagValidatorAddr))
			if err != nil {
				return err
			}

			bz, err := hex.DecodeString(viper.GetString(FlagDelegatorAddr))
			if err != nil {
				return err
			}
			delegator := crypto.Address(bz)
			delegator = coin.ChainAddr(delegator)

			prove := !viper.GetBool(client.FlagTrustNode)
			key := PrefixedKey(stake.Name, stake.GetDelegatorBondKey(delegator, addr))

			// parse out the bond
			var bond stake.DelegatorBond
			cdc := app.MakeTxCodec() // XXX create custom Tx for Staking Module
			err = cdc.UnmarshalBinary(res, bond)
			if err != nil {
				return err
			}
			output, err := json.MarshalIndent(bond, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(output))
			return nil

			// TODO output with proofs / machine parseable etc.
		},
	}

	cmd.Flags().AddFlagSet(fsValAddr)
	cmd.Flags().AddFlagSet(fsDelAddr)
	return cmd
}

// get the command to query all the candidates bonded to a delegator
func GetCmdQueryDelegatorBond(cdc *wire.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delegator-candidates",
		Short: "Query all delegators candidates' pubkeys based on address",
		RunE: func(cmd *cobra.Command, args []string) error {

			bz, err := hex.DecodeString(viper.GetString(FlagDelegatorAddr))
			if err != nil {
				return err
			}
			delegator := crypto.Address(bz)
			delegator = coin.ChainAddr(delegator)

			prove := !viper.GetBool(client.FlagTrustNode)
			key := PrefixedKey(stake.Name, stake.GetDelegatorBondsKey(delegator))

			// parse out the candidates list
			var candidates []crypto.PubKey
			err = cdc.UnmarshalBinary(res, candidates)
			if err != nil {
				return err
			}
			output, err := json.MarshalIndent(candidates, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(output))
			return nil

			// TODO output with proofs / machine parseable etc.
		},
	}
	cmd.Flags().AddFlagSet(fsDelAddr)
	return cmd
}
