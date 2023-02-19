package cmd

import (
	"github.com/arrow2nd/anct/cmdutil"
	"github.com/arrow2nd/anct/gen"
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
	r.Flags().StringP("body", "b", "", "Body")

	return r
}

func (c *Command) reviewRun(cmd *cobra.Command, args []string) error {
	_, ID, err := cmdutil.SearchWorks(c.api, cmd, args)
	if err != nil {
		return err
	}

	rating := map[string]gen.RatingState{"overall": "", "movie": "", "character": "", "story": "", "music": ""}

	// 評価
	for name := range rating {
		r, err := cmdutil.ReceiveRating(cmd.Flags(), "rating-"+name)
		if err != nil {
			return err
		}

		rating[name] = r
	}

	// 本文
	body, err := cmdutil.ReceiveBody(cmd.Flags(), "body")
	if err != nil {
		return err
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

	// レビューを作成
	spinner := view.SpinnerStart(cmd.OutOrStdout(), "Creating work review")
	if err := c.api.CreateWorkReview(ID, body, rating["overall"], rating["movie"], rating["character"], rating["story"], rating["music"]); err != nil {
		return err
	}

	spinner.Stop()
	view.PrintDone(cmd.OutOrStdout(), "Reviewed!")

	return nil
}
