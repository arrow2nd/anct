package cmd

import (
	"github.com/arrow2nd/anct/cmdutil"
	"github.com/arrow2nd/anct/view"
	"github.com/spf13/cobra"
)

func (c *Command) newCmdReview() *cobra.Command {
	r := &cobra.Command{
		Use:     "review [<query>]",
		Short:   "Review of the work",
		Example: "  anct review 戦姫絶唱シンフォギア",
		RunE:    c.reviewRun,
	}

	cmdutil.SetSearchFlags(r.Flags())

	r.Flags().StringP("rating-overall", "", "", "Overall rating: {great|good|average|bad}")
	r.Flags().StringP("rating-movie", "", "", "Move rating: {great|good|average|bad}")
	r.Flags().StringP("rating-character", "", "", "Character rating: {great|good|average|bad}")
	r.Flags().StringP("rating-story", "", "", "Story rating: {great|good|average|bad}")
	r.Flags().StringP("rating-music", "", "", "Music rating: {great|good|average|bad}")
	r.Flags().StringP("comment", "c", "", "Comment")

	return r
}

func (c *Command) reviewRun(cmd *cobra.Command, args []string) error {
	_, ID, err := cmdutil.SearchWorks(c.api, cmd, args)
	if err != nil {
		return nil, err
	}

	// 全体の評価

	// 映像の評価

	// キャラクターの評価

	// 音楽の評価

	// コメント
	comment, err := cmdutil.ReceiveComment(cmd.Flags())
	if err != nil {
		return err
	}

	view.PrintDone(cmd.OutOrStdout(), "Reviewed!")

	return nil
}
