package scratch

import "github.com/aws/aws-sdk-go/service/comprehend"

// Analysis struct ...
type Analysis struct {
	Project Project   `json:"project"`
	Reviews []*Review `json:"reviews"`
}

// Review struct ...
type Review struct {
	Comment
	Comprehend *comprehend.DetectSentimentOutput `json:"comprehend"`
}

// Project struct ...
type Project struct {
	Author          Author            `json:"author"`
	CommentsAllowed bool              `json:"comments_allowed"`
	Description     string            `json:"description"`
	History         History           `json:"history"`
	ID              int               `json:"id"`
	Image           string            `json:"image"`
	Images          map[string]string `json:"images"`
	Instructions    string            `json:"instructions"`
	IsPublished     bool              `json:"is_published"`
	Public          bool              `json:"public"`
	Remix           Remix             `json:"remix"`
	Stats           Stats             `json:"stats"`
	Title           string            `json:"title"`
	Visibility      string            `json:"visibility"`
}

// Comment struct ...
type Comment struct {
	Author           Author `json:"author"`
	CommenteeID      *int   `json:"commentee_id"`
	Content          string `json:"content"`
	DatetimeCreated  string `json:"datetime_created"`
	DatetimeModified string `json:"datetime_modified"`
	ID               int    `json:"id"`
	ParentID         *int   `json:"parent_id"`
	ReplyCount       int    `json:"reply_count"`
	Visibility       string `json:"visibility"`
}

// Author struct ...
type Author struct {
	ID          int    `json:"id"`
	Image       string `json:"image"`
	ScratchTeam bool   `json:"scratchteam"`
	Username    string `json:"username"`
}

// History struct ...
type History struct {
	Created  string `json:"created"`
	Modified string `json:"modified"`
	Shared   string `json:"shared"`
}

// Remix struct ...
type Remix struct {
	Root *string `json:"root"`
}

// Stats struct ...
type Stats struct {
	Comments  int `json:"comments"`
	Favorites int `json:"favorites"`
	Loves     int `json:"loves"`
	Remixes   int `json:"remixes"`
	Views     int `json:"views"`
}
