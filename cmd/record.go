package cmd

import (
	"github.com/arrow2nd/anct/cmdutil"
	"github.com/arrow2nd/anct/view"
	"github.com/spf13/cobra"
)

func (c *Command) newCmdRecord() *cobra.Command {
	r := &cobra.Command{
		Use:     "record [<query>]",
		Short:   "Record the watching of episode",
		Example: "  anct record お兄ちゃんはおしまい",
		RunE:    c.recordRun,
	}

	cmdutil.SetSearchFlags(r.Flags())
	r.Flags().StringP("rating", "R", "", "Episode rating: {great|good|average|bad}")
	r.Flags().StringP("comment", "C", "", "Comment")

	return r
}

func (c *Command) recordRun(cmd *cobra.Command, args []string) error {
	annictID, _, err := cmdutil.SearchWorks(c.api, cmd, args)
	if err != nil {
		return err
	}

	work, err := c.api.FetchWorkEpisodes(annictID)
	if err != nil {
		return err
	}

	// 記録するエピソードを選択
	episodeIDs, err := view.SelectEpisodes(work)
	if err != nil {
		return err
	}

	// 評価を取得
	rating, err := cmdutil.ReceiveRating(cmd.Flags())
	if err != nil {
		return err
	}

	// コメントを取得
	comment, _ := cmd.Flags().GetString("comment")
	if len(episodeIDs) > 1 {
		// 一括記録の場合コメントはつけない
		comment = ""
	} else if comment == "" {
		// 指定されていなければエディタを開く
		c, err := view.InputTextInEditor("Enter your comments")
		if err != nil {
			return err
		}
		comment = c
	}

	spinner := view.SpinnerStart(cmd.OutOrStdout(), "Creating episode record")

	if err := c.api.CreateEpisodeRecords(episodeIDs, rating, comment); err != nil {
		return err
	}

	spinner.Stop()

	view.PrintDone(cmd.OutOrStdout(), "Recorded!")
	return nil
}
