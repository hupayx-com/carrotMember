package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/hupayx-com/carrotMember/x/carrotmember/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdNextRewardTime() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "next-reward-time",
		Short: "Query nextRewardTime",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryNextRewardTimeRequest{}

			res, err := queryClient.NextRewardTime(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
