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
		Example: "  anct record お兄ちゃんはおしまい！",
		RunE:    c.recordRun,
	}

	cmdutil.SetSearchFlags(r.Flags())

	r.Flags().BoolP("unwatch", "u", false, "Select from the unwatched episodes of the work you are watching")
	r.Flags().StringP("rating", "r", "", "Episode rating: {great|good|average|bad}")
	r.Flags().StringP("comment", "c", "", "Comment")

	return r
}

func (c *Command) recordRun(cmd *cobra.Command, args []string) error {
	episodeIDs := []string{}

	// 記録するエピソードを選択
	if unwatch, _ := cmd.Flags().GetBool("unwatch"); unwatch {
		id, err := c.recordSelectUnwatchEpisord()
		if err != nil {
			return err
		}

		episodeIDs = []string{id}
	} else {
		ids, err := c.recordSelectEpisodes(cmd, args)
		if err != nil {
			return err
		}

		episodeIDs = ids
	}

	// 評価
	rating, err := cmdutil.ReceiveRating(cmd.Flags(), "rating")
	if err != nil {
		return err
	}

	// コメント
	comment := ""
	if len(episodeIDs) == 1 {
		// 記録するエピソードが1つの時のみコメントを受け取る
		c, err := cmdutil.ReceiveBody(cmd.Flags(), "comment")
		if err != nil {
			return err
		}

		comment = c
	}

	// 確認
	submit, err := view.Confirm("Submit?")
	if err != nil {
		return err
	}
	if !submit {
		view.PrintCanceled(cmd.ErrOrStderr())
		return nil
	}

	spinner := view.SpinnerStart(cmd.OutOrStdout(), "Creating episode record")
	if err := c.api.CreateEpisodeRecords(episodeIDs, rating, comment); err != nil {
		return err
	}

	spinner.Stop()
	view.PrintDone(cmd.OutOrStdout(), "Recorded!")

	return nil
}

// recordSelectEpisodes : 検索結果から記録するエピソードを選択
func (c *Command) recordSelectEpisodes(cmd *cobra.Command, args []string) ([]string, error) {
	annictID, _, err := cmdutil.SearchWorks(c.api, cmd, args)
	if err != nil {
		return nil, err
	}

	work, err := c.api.FetchWorkEpisodes(annictID)
	if err != nil {
		return nil, err
	}

	return view.SelectEpisodes(work)
}

// recordSelectUnwatchEpisord : 未視聴のエピソードから選択
func (c *Command) recordSelectUnwatchEpisord() (string, error) {
	unwatchEpisodes, err := c.api.FetchUnwatchEpisodes()
	if err != nil {
		return "", err
	}

	return view.SelectUnwatchEpisode(unwatchEpisodes)
}
