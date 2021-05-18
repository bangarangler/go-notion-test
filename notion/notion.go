package notion

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	// "time"

	"github.com/dstotijn/go-notion"
)

// type Context interface {
// 	Deadline() (deadline time.Time, ok bool)
// 	Done() <-chan struct{}
// 	Err() error
// }

type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}

// type PaginationQuery struct {
// 	StartCursor string
// 	PageSize    int
// }

// func TODO() Context
// func Background() Context

func Test() {
	println("Hello from test function")
	client := notion.NewClient(Notion_Key)
	// var opts = PaginationQuery{"", 100}
	// users, err := client.ListUsers(context.Background(), &notion.PaginationQuery{StartCursor: "", PageSize: 100})
	pages, err := client.Search(context.Background(), &notion.SearchOpts{})
	// page, err := client.FindPageByID(Notion_DB)
	// users, err := client.ListUsers(context.Context{Deadline: false}, nil)
	// users, err := client.ListUsers(context.Context{Deadline: false}, nil)
	if err != nil {
		fmt.Println("Error finding page", err)
	}
	b, err := json.MarshalIndent(pages.Results, "", "    ")
	if err == nil {
		s := string(b)
		fmt.Println(s)
	}
	// println(fmt.Sprintf("%v, %T", pages.Results, pages.Results))
	// AddToWhatCanIDo()
}
