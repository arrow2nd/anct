package cmd

import (
	"github.com/arrow2nd/anct/cmdutil"
	"github.com/arrow2nd/anct/gen"
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

	r.Flags().BoolP("unwatch", "u", false, "select from the unwatched episodes of the work you are watching")
	r.Flags().StringP("rating", "r", "", "episode rating: {great|good|average|bad}")
	r.Flags().StringP("comment", "c", "", "comment")
	r.Flags().BoolP("no-comment", "n", false, "skip comment input")

	return r
}

func (c *Command) recordRun(cmd *cobra.Command, args []string) error {
	var (
		work       *gen.WorkFragment = nil
		episodeIDs                   = []string{}
	)

	// 記録するエピソードを選択
	if unwatch, _ := cmd.Flags().GetBool("unwatch"); unwatch {
		id, err := c.recordSelectUnwatchEpisord()
		if err != nil {
			return err
		}

		episodeIDs = []string{id}
	} else {
		w, ids, err := c.recordSelectEpisodes(cmd, args)
		if err != nil {
			return err
		}

		work = w
		episodeIDs = ids
	}

	// 評価
	rating, err := cmdutil.ReceiveRating(cmd.Flags(), "rating")
	if err != nil {
		return err
	}

	comment := ""
	noComment, _ := cmd.Flags().GetBool("no-comment")

	// --no-comment が指定されていない & 記録するエピソードが1つの時のみコメントを受け取る
	if !noComment && len(episodeIDs) == 1 {
		if c, err := cmdutil.ReceiveBody(cmd.Flags(), "comment"); err != nil {
			return err
		} else {
			comment = c
		}
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

	// 初めての視聴でないならここで終了
	if work == nil || (work.ViewerStatusState != nil && *work.ViewerStatusState != gen.StatusStateNoState) {
		return nil
	}

	// 視聴ステータスの変更を確認
	changeStatus, err := view.Confirm("Change your viewing status to Watching?")
	if err != nil {
		return err
	}
	if !changeStatus {
		view.PrintCanceled(cmd.ErrOrStderr())
		return nil
	}

	spinner = view.SpinnerStart(cmd.OutOrStdout(), "Updating status")
	if err := c.api.UpdateWorkState(work.ID, gen.StatusStateWatching); err != nil {
		return err
	}

	spinner.Stop()
	view.PrintDone(cmd.OutOrStdout(), "Updated status!")

	return nil
}

// recordSelectEpisodes : 検索結果から記録するエピソードを選択
func (c *Command) recordSelectEpisodes(cmd *cobra.Command, args []string) (*gen.WorkFragment, []string, error) {
	work, _, err := cmdutil.SearchWorks(c.api, cmd, args)
	if err != nil {
		return nil, nil, err
	}

	episode, err := c.api.FetchWorkEpisodes(work.AnnictID)
	if err != nil {
		return nil, nil, err
	}

	ids, err := view.SelectEpisodes(episode)
	if err != nil {
		return nil, nil, err
	}

	return work, ids, nil
}

// recordSelectUnwatchEpisord : 未視聴のエピソードから選択
func (c *Command) recordSelectUnwatchEpisord() (string, error) {
	unwatchEpisodes, err := c.api.FetchUnwatchEpisodes()
	if err != nil {
		return "", err
	}

	return view.SelectUnwatchEpisode(unwatchEpisodes)
}
