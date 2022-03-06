package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hupayx-com/carrotMember/x/carrotmember/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAddMember() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-member [member-address] [total-cnt] [amount]",
		Short: "Broadcast message add-member",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argMemberAddress := args[0]
			argTotalCnt, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argAmount, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddMember(
				clientCtx.GetFromAddress().String(),
				argMemberAddress,
				argTotalCnt,
				argAmount,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
