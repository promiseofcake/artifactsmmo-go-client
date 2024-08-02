package cmd

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/promiseofcake/artifactsmmo-cli/internal/actions"
	"github.com/spf13/cobra"
)

var characterName string

// gatherCmd represents the gather command
var gatherCmd = &cobra.Command{
	Use:   "gather",
	Short: "Start a gather loop in your current location",
	RunE: func(cmd *cobra.Command, args []string) error {
		character, err := cmd.PersistentFlags().GetString("character")
		if err != nil {
			return fmt.Errorf("failed to get character: %w", err)
		}

		r := cmd.Context().Value(runnerKey).(*actions.Runner)
		for {
			slog.Info("about to gather")

			resp, err := r.Gather(cmd.Context(), character)
			if err != nil {
				slog.Error("failed to gather", "error", err.Error())
				return fmt.Errorf("failed to get gather: %w", err)
			}

			sec := resp.GetRemainingCooldown()
			slog.Info("gather results",
				"status", resp.StatusCode,
				"xp", resp.SkillInfo.Xp,
				"items", resp.SkillInfo.Items,
				"cooldown", sec,
			)
			time.Sleep(time.Duration(sec) * time.Second)
		}
		return nil
	},
}

func init() {
	gatherCmd.PersistentFlags().StringVar(&characterName, "character", "", "The name of your character")
	rootCmd.AddCommand(gatherCmd)
}
