package cli

import (
    "strconv"
	
	"github.com/spf13/cobra"
    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/hupayx-com/carrotMember/x/carrotmember/types"
)

var _ = strconv.Itoa(0)

func CmdRewardPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reward-pool",
		Short: "Query rewardPool",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryRewardPoolRequest{
				
            }

            

			res, err := queryClient.RewardPool(cmd.Context(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}