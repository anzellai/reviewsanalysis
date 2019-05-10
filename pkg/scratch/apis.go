package scratch

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/comprehend"

	"github.com/imroc/req"
)

const (
	scratchProjectAPI = "https://api.scratch.mit.edu/explore/projects?limit=%d&offset=%d&language=%s&mode=%s"
	scratchCommentAPI = "https://api.scratch.mit.edu/projects/%d/comments?limit=40&offset=%d"
)

// API struct ...
type API struct {
	Limit    int         `json:"-"`
	Offset   int         `json:"-"`
	Language string      `json:"-"`
	Mode     string      `json:"-"`
	Results  []*Analysis `json:"results"`
}

// NewAPI return new API instance to call utils
func NewAPI() *API {
	return &API{
		Limit:    10,
		Offset:   0,
		Language: "en",
		Mode:     "popular",
		Results:  nil,
	}
}

// GetProjects return slice of Projects from Scratch API
func (api *API) GetProjects() (projects []Project, err error) {
	r, err := req.Get(fmt.Sprintf(scratchProjectAPI, api.Limit, api.Offset, api.Language, api.Mode))
	if err != nil {
		return
	}
	r.ToJSON(&projects)
	return
}

// GetComments return slice of Comments for its ProjectID from Scratch API
func (api *API) GetComments(projectID int) (comments []Comment, err error) {
	offset := 0
	for {
		cmts := []Comment{}
		commentsEndpoint := fmt.Sprintf(scratchCommentAPI, projectID, offset)
		fmt.Printf("calling Scratch comments endpoint: %s\n", commentsEndpoint)
		r, err := req.Get(commentsEndpoint)
		if err != nil {
			return comments, err
		}
		r.ToJSON(&cmts)
		if len(cmts) == 0 {
			return comments, nil
		}
		comments = append(comments, cmts...)
		offset += 40
	}
	return
}

// GetTop10ProjectComments fetch and set the top10 projects with the comments and set into api Reviews slice
func (api *API) GetTop10ProjectComments() error {
	projects, err := api.GetProjects()
	if err != nil {
		return err
	}
	for _, project := range projects {
		comments, err := api.GetComments(project.ID)
		if err != nil {
			fmt.Errorf("error fetching project comments: %v", err)
			continue
		}
		reviews := make([]*Review, 0)
		for _, comment := range comments {
			reviews = append(reviews, &Review{Comment: comment})
		}
		analysis := &Analysis{Project: project, Reviews: reviews}
		api.Results = append(api.Results, analysis)
	}
	return nil
}

// GetAnalysis perform sentiment and relevant analysis on results slice
func (api *API) GetAnalysis() error {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	srv := comprehend.New(sess)
	lang := "en"
	for _, result := range api.Results {
		for _, review := range result.Reviews {
			fmt.Printf("calling AWS Comprehend for: %v\n", review)
			summary, err := srv.DetectSentiment(&comprehend.DetectSentimentInput{LanguageCode: &lang, Text: &review.Content})
			if err != nil {
				fmt.Printf("error using comprehend: %v", err)
				continue
			}
			review.Comprehend = summary
		}
	}
	return nil
}
