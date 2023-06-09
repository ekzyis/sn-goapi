package sn

import (
	"fmt"
	"time"
)

type GraphQLPayload struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

type GraphQLError struct {
	Message string `json:"message"`
}

type User struct {
	Name string `json:"name"`
}

type Comment struct {
	Id       int       `json:"id,string"`
	Text     string    `json:"text"`
	User     User      `json:"user"`
	Comments []Comment `json:"comments"`
}

type CreateCommentsResponse struct {
	Errors []GraphQLError `json:"errors"`
	Data   struct {
		CreateComment Comment `json:"createComment"`
	} `json:"data"`
}

type Item struct {
	Id        int       `json:"id,string"`
	Title     string    `json:"title"`
	Url       string    `json:"url"`
	Sats      int       `json:"sats"`
	CreatedAt time.Time `json:"createdAt"`
	Comments  []Comment `json:"comments"`
	NComments int       `json:"ncomments"`
}

type UpsertLinkResponse struct {
	Errors []GraphQLError `json:"errors"`
	Data   struct {
		UpsertLink Item `json:"upsertLink"`
	} `json:"data"`
}

type ItemsResponse struct {
	Errors []GraphQLError `json:"errors"`
	Data   struct {
		Items struct {
			Items  []Item `json:"items"`
			Cursor string `json:"cursor"`
		} `json:"items"`
	} `json:"data"`
}

type HasNewNotesResponse struct {
	Errors []GraphQLError `json:"errors"`
	Data   struct {
		HasNewNotes bool `json:"hasNewNotes"`
	} `json:"data"`
}

type Dupe struct {
	Id        int       `json:"id,string"`
	Url       string    `json:"url"`
	Title     string    `json:"title"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"createdAt"`
	Sats      int       `json:"sats"`
	NComments int       `json:"ncomments"`
}

type DupesResponse struct {
	Errors []GraphQLError `json:"errors"`
	Data   struct {
		Dupes []Dupe `json:"dupes"`
	} `json:"data"`
}

type DupesError struct {
	Url   string
	Dupes []Dupe
}

func (e *DupesError) Error() string {
	return fmt.Sprintf("found %d dupes for %s", len(e.Dupes), e.Url)
}
