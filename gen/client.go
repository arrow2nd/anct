// Code generated by github.com/Yamashou/gqlgenc, DO NOT EDIT.

package gen

import (
	"context"
	"net/http"

	"github.com/Yamashou/gqlgenc/clientv2"
)

type Client struct {
	Client *clientv2.Client
}

func NewClient(cli *http.Client, baseURL string, options *clientv2.Options, interceptors ...clientv2.RequestInterceptor) *Client {
	return &Client{Client: clientv2.NewClient(cli, baseURL, options, interceptors...)}
}

type Query struct {
	Node                Node                    "json:\"node,omitempty\" graphql:\"node\""
	Nodes               []Node                  "json:\"nodes\" graphql:\"nodes\""
	SearchCharacters    *CharacterConnection    "json:\"searchCharacters,omitempty\" graphql:\"searchCharacters\""
	SearchEpisodes      *EpisodeConnection      "json:\"searchEpisodes,omitempty\" graphql:\"searchEpisodes\""
	SearchOrganizations *OrganizationConnection "json:\"searchOrganizations,omitempty\" graphql:\"searchOrganizations\""
	SearchPeople        *PersonConnection       "json:\"searchPeople,omitempty\" graphql:\"searchPeople\""
	SearchWorks         *WorkConnection         "json:\"searchWorks,omitempty\" graphql:\"searchWorks\""
	User                *User                   "json:\"user,omitempty\" graphql:\"user\""
	Viewer              *User                   "json:\"viewer,omitempty\" graphql:\"viewer\""
}
type Mutation struct {
	CreateRecord *CreateRecordPayload "json:\"createRecord,omitempty\" graphql:\"createRecord\""
	CreateReview *CreateReviewPayload "json:\"createReview,omitempty\" graphql:\"createReview\""
	DeleteRecord *DeleteRecordPayload "json:\"deleteRecord,omitempty\" graphql:\"deleteRecord\""
	DeleteReview *DeleteReviewPayload "json:\"deleteReview,omitempty\" graphql:\"deleteReview\""
	UpdateRecord *UpdateRecordPayload "json:\"updateRecord,omitempty\" graphql:\"updateRecord\""
	UpdateReview *UpdateReviewPayload "json:\"updateReview,omitempty\" graphql:\"updateReview\""
	UpdateStatus *UpdateStatusPayload "json:\"updateStatus,omitempty\" graphql:\"updateStatus\""
}
type WorkFragment struct {
	AnnictID          int64        "json:\"annictId\" graphql:\"annictId\""
	ID                string       "json:\"id\" graphql:\"id\""
	Title             string       "json:\"title\" graphql:\"title\""
	Media             Media        "json:\"media\" graphql:\"media\""
	SeasonName        *SeasonName  "json:\"seasonName,omitempty\" graphql:\"seasonName\""
	SeasonYear        *int64       "json:\"seasonYear,omitempty\" graphql:\"seasonYear\""
	ViewerStatusState *StatusState "json:\"viewerStatusState,omitempty\" graphql:\"viewerStatusState\""
}
type EpisodeFragment struct {
	ID                 string  "json:\"id\" graphql:\"id\""
	Number             *int64  "json:\"number,omitempty\" graphql:\"number\""
	NumberText         *string "json:\"numberText,omitempty\" graphql:\"numberText\""
	Title              *string "json:\"title,omitempty\" graphql:\"title\""
	ViewerRecordsCount int64   "json:\"viewerRecordsCount\" graphql:\"viewerRecordsCount\""
}
type WorkEpisodesFragment struct {
	Episodes   *WorkEpisodesFragment_Episodes "json:\"episodes,omitempty\" graphql:\"episodes\""
	NoEpisodes bool                           "json:\"noEpisodes\" graphql:\"noEpisodes\""
}
type UnwatchLibraryEntryFragment struct {
	Work        UnwatchLibraryEntryFragment_Work "json:\"work\" graphql:\"work\""
	NextEpisode *EpisodeFragment                 "json:\"nextEpisode,omitempty\" graphql:\"nextEpisode\""
}
type WorkInfoFragment struct {
	AnnictID          int64                                           "json:\"annictId\" graphql:\"annictId\""
	ID                string                                          "json:\"id\" graphql:\"id\""
	Title             string                                          "json:\"title\" graphql:\"title\""
	Media             Media                                           "json:\"media\" graphql:\"media\""
	SeasonName        *SeasonName                                     "json:\"seasonName,omitempty\" graphql:\"seasonName\""
	SeasonYear        *int64                                          "json:\"seasonYear,omitempty\" graphql:\"seasonYear\""
	ViewerStatusState *StatusState                                    "json:\"viewerStatusState,omitempty\" graphql:\"viewerStatusState\""
	Episodes          *WorkInfoFragment_WorkEpisodesFragment_Episodes "json:\"episodes,omitempty\" graphql:\"episodes\""
	NoEpisodes        bool                                            "json:\"noEpisodes\" graphql:\"noEpisodes\""
	Image             *WorkInfoFragment_Image                         "json:\"image,omitempty\" graphql:\"image\""
	OfficialSiteURL   *string                                         "json:\"officialSiteUrl,omitempty\" graphql:\"officialSiteUrl\""
	WatchersCount     int64                                           "json:\"watchersCount\" graphql:\"watchersCount\""
}
type WorkEpisodesFragment_Episodes struct {
	Nodes []*EpisodeFragment "json:\"nodes,omitempty\" graphql:\"nodes\""
}
type UnwatchLibraryEntryFragment_Work struct {
	Title string "json:\"title\" graphql:\"title\""
}
type WorkInfoFragment_WorkEpisodesFragment_Episodes struct {
	Nodes []*EpisodeFragment "json:\"nodes,omitempty\" graphql:\"nodes\""
}
type WorkInfoFragment_Image struct {
	Copyright           *string "json:\"copyright,omitempty\" graphql:\"copyright\""
	RecommendedImageURL *string "json:\"recommendedImageUrl,omitempty\" graphql:\"recommendedImageUrl\""
	FacebookOgImageURL  *string "json:\"facebookOgImageUrl,omitempty\" graphql:\"facebookOgImageUrl\""
}
type UpdateWorkState_UpdateStatus struct {
	ClientMutationID *string "json:\"clientMutationId,omitempty\" graphql:\"clientMutationId\""
}
type CreateEpisodeRecord_CreateRecord struct {
	ClientMutationID *string "json:\"clientMutationId,omitempty\" graphql:\"clientMutationId\""
}
type CreateWorkReview_CreateReview struct {
	ClientMutationID *string "json:\"clientMutationId,omitempty\" graphql:\"clientMutationId\""
}
type SearchWorksByKeyword_SearchWorks struct {
	Nodes []*WorkFragment "json:\"nodes,omitempty\" graphql:\"nodes\""
}
type FetchWorkInfo_SearchWorks_Nodes_WorkInfoFragment_WorkEpisodesFragment_Episodes struct {
	Nodes []*EpisodeFragment "json:\"nodes,omitempty\" graphql:\"nodes\""
}
type FetchWorkInfo_SearchWorks_Nodes_WorkInfoFragment_Image struct {
	Copyright           *string "json:\"copyright,omitempty\" graphql:\"copyright\""
	RecommendedImageURL *string "json:\"recommendedImageUrl,omitempty\" graphql:\"recommendedImageUrl\""
	FacebookOgImageURL  *string "json:\"facebookOgImageUrl,omitempty\" graphql:\"facebookOgImageUrl\""
}
type FetchWorkInfo_SearchWorks struct {
	Nodes []*WorkInfoFragment "json:\"nodes,omitempty\" graphql:\"nodes\""
}
type FetchWorkEpisodes_SearchWorks_Nodes_WorkEpisodesFragment_Episodes struct {
	Nodes []*EpisodeFragment "json:\"nodes,omitempty\" graphql:\"nodes\""
}
type FetchWorkEpisodes_SearchWorks struct {
	Nodes []*WorkEpisodesFragment "json:\"nodes,omitempty\" graphql:\"nodes\""
}
type FetchUnwatchEpisodes_Viewer_LibraryEntries_Nodes_UnwatchLibraryEntryFragment_Work struct {
	Title string "json:\"title\" graphql:\"title\""
}
type FetchUnwatchEpisodes_Viewer_LibraryEntries struct {
	Nodes []*UnwatchLibraryEntryFragment "json:\"nodes,omitempty\" graphql:\"nodes\""
}
type FetchUnwatchEpisodes_Viewer struct {
	LibraryEntries *FetchUnwatchEpisodes_Viewer_LibraryEntries "json:\"libraryEntries,omitempty\" graphql:\"libraryEntries\""
}
type FetchUserLibrary_Viewer_LibraryEntries_Nodes struct {
	Work *WorkFragment "json:\"work\" graphql:\"work\""
}
type FetchUserLibrary_Viewer_LibraryEntries struct {
	Nodes []*FetchUserLibrary_Viewer_LibraryEntries_Nodes "json:\"nodes,omitempty\" graphql:\"nodes\""
}
type FetchUserLibrary_Viewer struct {
	LibraryEntries *FetchUserLibrary_Viewer_LibraryEntries "json:\"libraryEntries,omitempty\" graphql:\"libraryEntries\""
}
type HogeUpdateWorkStatePayload struct {
	UpdateStatus *UpdateWorkState_UpdateStatus "json:\"updateStatus,omitempty\" graphql:\"updateStatus\""
}
type HogeCreateEpisodeRecordPayload struct {
	CreateRecord *CreateEpisodeRecord_CreateRecord "json:\"createRecord,omitempty\" graphql:\"createRecord\""
}
type HogeCreateWorkReviewPayload struct {
	CreateReview *CreateWorkReview_CreateReview "json:\"createReview,omitempty\" graphql:\"createReview\""
}
type SearchWorksByKeyword struct {
	SearchWorks *SearchWorksByKeyword_SearchWorks "json:\"searchWorks,omitempty\" graphql:\"searchWorks\""
}
type FetchWorkInfo struct {
	SearchWorks *FetchWorkInfo_SearchWorks "json:\"searchWorks,omitempty\" graphql:\"searchWorks\""
}
type FetchWorkEpisodes struct {
	SearchWorks *FetchWorkEpisodes_SearchWorks "json:\"searchWorks,omitempty\" graphql:\"searchWorks\""
}
type FetchUnwatchEpisodes struct {
	Viewer *FetchUnwatchEpisodes_Viewer "json:\"viewer,omitempty\" graphql:\"viewer\""
}
type FetchUserLibrary struct {
	Viewer *FetchUserLibrary_Viewer "json:\"viewer,omitempty\" graphql:\"viewer\""
}

const UpdateWorkStateDocument = `mutation UpdateWorkState ($workId: ID!, $state: StatusState!) {
	updateStatus(input: {state:$state,workId:$workId}) {
		clientMutationId
	}
}
`

func (c *Client) UpdateWorkState(ctx context.Context, workID string, state StatusState, interceptors ...clientv2.RequestInterceptor) (*HogeUpdateWorkStatePayload, error) {
	vars := map[string]interface{}{
		"workId": workID,
		"state":  state,
	}

	var res HogeUpdateWorkStatePayload
	if err := c.Client.Post(ctx, "UpdateWorkState", UpdateWorkStateDocument, &res, vars, interceptors...); err != nil {
		return &res, err
	}

	return &res, nil
}

const CreateEpisodeRecordDocument = `mutation CreateEpisodeRecord ($episodeId: ID!, $rating: RatingState!, $comment: String) {
	createRecord(input: {episodeId:$episodeId,ratingState:$rating,comment:$comment}) {
		clientMutationId
	}
}
`

func (c *Client) CreateEpisodeRecord(ctx context.Context, episodeID string, rating RatingState, comment *string, interceptors ...clientv2.RequestInterceptor) (*HogeCreateEpisodeRecordPayload, error) {
	vars := map[string]interface{}{
		"episodeId": episodeID,
		"rating":    rating,
		"comment":   comment,
	}

	var res HogeCreateEpisodeRecordPayload
	if err := c.Client.Post(ctx, "CreateEpisodeRecord", CreateEpisodeRecordDocument, &res, vars, interceptors...); err != nil {
		return &res, err
	}

	return &res, nil
}

const CreateWorkReviewDocument = `mutation CreateWorkReview ($workId: ID!, $body: String!, $ratingOverall: RatingState, $ratingMovie: RatingState, $ratingChara: RatingState, $ratingStory: RatingState, $ratingMusic: RatingState) {
	createReview(input: {workId:$workId,body:$body,ratingAnimationState:$ratingMovie,ratingCharacterState:$ratingChara,ratingMusicState:$ratingMusic,ratingOverallState:$ratingOverall,ratingStoryState:$ratingStory}) {
		clientMutationId
	}
}
`

func (c *Client) CreateWorkReview(ctx context.Context, workID string, body string, ratingOverall *RatingState, ratingMovie *RatingState, ratingChara *RatingState, ratingStory *RatingState, ratingMusic *RatingState, interceptors ...clientv2.RequestInterceptor) (*HogeCreateWorkReviewPayload, error) {
	vars := map[string]interface{}{
		"workId":        workID,
		"body":          body,
		"ratingOverall": ratingOverall,
		"ratingMovie":   ratingMovie,
		"ratingChara":   ratingChara,
		"ratingStory":   ratingStory,
		"ratingMusic":   ratingMusic,
	}

	var res HogeCreateWorkReviewPayload
	if err := c.Client.Post(ctx, "CreateWorkReview", CreateWorkReviewDocument, &res, vars, interceptors...); err != nil {
		return &res, err
	}

	return &res, nil
}

const SearchWorksByKeywordDocument = `query SearchWorksByKeyword ($query: String!, $seasons: [String!], $first: Int!) {
	searchWorks(titles: [$query], seasons: $seasons, first: $first, orderBy: {field:SEASON,direction:DESC}) {
		nodes {
			... WorkFragment
		}
	}
}
fragment WorkFragment on Work {
	annictId
	id
	title
	media
	seasonName
	seasonYear
	viewerStatusState
}
`

func (c *Client) SearchWorksByKeyword(ctx context.Context, query string, seasons []string, first int64, interceptors ...clientv2.RequestInterceptor) (*SearchWorksByKeyword, error) {
	vars := map[string]interface{}{
		"query":   query,
		"seasons": seasons,
		"first":   first,
	}

	var res SearchWorksByKeyword
	if err := c.Client.Post(ctx, "SearchWorksByKeyword", SearchWorksByKeywordDocument, &res, vars, interceptors...); err != nil {
		return &res, err
	}

	return &res, nil
}

const FetchWorkInfoDocument = `query FetchWorkInfo ($annictId: Int!) {
	searchWorks(annictIds: [$annictId]) {
		nodes {
			... WorkInfoFragment
		}
	}
}
fragment WorkInfoFragment on Work {
	... WorkFragment
	... WorkEpisodesFragment
	image {
		copyright
		recommendedImageUrl
		facebookOgImageUrl
	}
	officialSiteUrl
	watchersCount
}
fragment WorkFragment on Work {
	annictId
	id
	title
	media
	seasonName
	seasonYear
	viewerStatusState
}
fragment WorkEpisodesFragment on Work {
	episodes(orderBy: {direction:ASC,field:SORT_NUMBER}) {
		nodes {
			... EpisodeFragment
		}
	}
	noEpisodes
}
fragment EpisodeFragment on Episode {
	id
	number
	numberText
	title
	viewerRecordsCount
}
`

func (c *Client) FetchWorkInfo(ctx context.Context, annictID int64, interceptors ...clientv2.RequestInterceptor) (*FetchWorkInfo, error) {
	vars := map[string]interface{}{
		"annictId": annictID,
	}

	var res FetchWorkInfo
	if err := c.Client.Post(ctx, "FetchWorkInfo", FetchWorkInfoDocument, &res, vars, interceptors...); err != nil {
		return &res, err
	}

	return &res, nil
}

const FetchWorkEpisodesDocument = `query FetchWorkEpisodes ($annictId: Int!) {
	searchWorks(annictIds: [$annictId]) {
		nodes {
			... WorkEpisodesFragment
		}
	}
}
fragment WorkEpisodesFragment on Work {
	episodes(orderBy: {direction:ASC,field:SORT_NUMBER}) {
		nodes {
			... EpisodeFragment
		}
	}
	noEpisodes
}
fragment EpisodeFragment on Episode {
	id
	number
	numberText
	title
	viewerRecordsCount
}
`

func (c *Client) FetchWorkEpisodes(ctx context.Context, annictID int64, interceptors ...clientv2.RequestInterceptor) (*FetchWorkEpisodes, error) {
	vars := map[string]interface{}{
		"annictId": annictID,
	}

	var res FetchWorkEpisodes
	if err := c.Client.Post(ctx, "FetchWorkEpisodes", FetchWorkEpisodesDocument, &res, vars, interceptors...); err != nil {
		return &res, err
	}

	return &res, nil
}

const FetchUnwatchEpisodesDocument = `query FetchUnwatchEpisodes {
	viewer {
		libraryEntries(states: [WATCHING], orderBy: {direction:DESC,field:LAST_TRACKED_AT}) {
			nodes {
				... UnwatchLibraryEntryFragment
			}
		}
	}
}
fragment UnwatchLibraryEntryFragment on LibraryEntry {
	work {
		title
	}
	nextEpisode {
		... EpisodeFragment
	}
}
fragment EpisodeFragment on Episode {
	id
	number
	numberText
	title
	viewerRecordsCount
}
`

func (c *Client) FetchUnwatchEpisodes(ctx context.Context, interceptors ...clientv2.RequestInterceptor) (*FetchUnwatchEpisodes, error) {
	vars := map[string]interface{}{}

	var res FetchUnwatchEpisodes
	if err := c.Client.Post(ctx, "FetchUnwatchEpisodes", FetchUnwatchEpisodesDocument, &res, vars, interceptors...); err != nil {
		return &res, err
	}

	return &res, nil
}

const FetchUserLibraryDocument = `query FetchUserLibrary ($states: [StatusState!], $seasons: [String!], $first: Int!) {
	viewer {
		libraryEntries(states: $states, seasons: $seasons, first: $first, orderBy: {direction:DESC,field:LAST_TRACKED_AT}) {
			nodes {
				work {
					... WorkFragment
				}
			}
		}
	}
}
fragment WorkFragment on Work {
	annictId
	id
	title
	media
	seasonName
	seasonYear
	viewerStatusState
}
`

func (c *Client) FetchUserLibrary(ctx context.Context, states []StatusState, seasons []string, first int64, interceptors ...clientv2.RequestInterceptor) (*FetchUserLibrary, error) {
	vars := map[string]interface{}{
		"states":  states,
		"seasons": seasons,
		"first":   first,
	}

	var res FetchUserLibrary
	if err := c.Client.Post(ctx, "FetchUserLibrary", FetchUserLibraryDocument, &res, vars, interceptors...); err != nil {
		return &res, err
	}

	return &res, nil
}
