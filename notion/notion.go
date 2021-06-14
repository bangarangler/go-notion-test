package notion

import (
	"context"
	"fmt"

	"github.com/dstotijn/go-notion"
	"github.com/tidwall/pretty"
)

// type Context interface {
// 	Deadline() (deadline time.Time, ok bool)
// 	Done() <-chan struct{}
// 	Err() error
// 	Value(key interface{}) interface{}
// }

func Test() {
	getDatabaseInfo(Notion_Key, Notion_DB)
	// queryDatabase(Notion_Key, Notion_DB)
	// getPageInfo(Notion_Key, Notion_Page)
	getBlockChildren(Notion_Key, Notion_Page)
}
func showBlock(b *notion.Block) {
	fmt.Println(fmt.Sprintf("  %s %s, has_children: %v\n", b.Type, b.ID, b.HasChildren))
	switch b.Type {
	case notion.BlockTypeParagraph:
		fmt.Println(fmt.Sprintf(" %v\n", b.Paragraph.Text))
	case notion.BlockTypeHeading1:
		fmt.Println(fmt.Sprintf(" %v\n", b.Heading1.Text))
	case notion.BlockTypeHeading2:
		fmt.Println(fmt.Sprintf(" %v\n", b.Heading2.Text))
	case notion.BlockTypeHeading3:
		fmt.Println(fmt.Sprintf(" %v\n", b.Heading3.Text))
	case notion.BlockTypeBulletedListItem:
		fmt.Println(fmt.Sprintf(" %v\n", b.BulletedListItem.Text))
	case notion.BlockTypeNumberedListItem:
		fmt.Println(fmt.Sprintf(" %v\n", b.NumberedListItem.Text))
	case notion.BlockTypeToDo:
		fmt.Println(fmt.Sprintf(" %v\n", b.ToDo.Text))
	case notion.BlockTypeToggle:
		fmt.Println(fmt.Sprintf(" %v\n", b.Toggle.Text))
	case notion.BlockTypeChildPage:
	case notion.BlockTypeUnsupported:
	}
}

func showBlockChildren(bcr *notion.BlockChildrenResponse) {
	fmt.Println("showBlockChildren:")
	fmt.Println(fmt.Sprintf("  hasMore: %v\n", bcr.HasMore))
	empty := ""
	if bcr.NextCursor != &empty {
		fmt.Println(fmt.Sprintf("  nextCursor: %v\n", bcr.NextCursor))
	}
	fmt.Println(fmt.Sprintf("  %d children:\n", len(bcr.Results)))
	for _, b := range bcr.Results {
		showBlock(&b)
	}
}

func getBlockChildren2(apiKey string, blockID string) {
	fmt.Printf("getBlockChildren: blockID='%s'\n", blockID)

	c := notion.NewClient(apiKey)
	ctx := context.Background()
	rsp, err := c.FindBlockChildrenByID(ctx, blockID, nil)
	if err != nil {
		fmt.Println(fmt.Sprintf("GetBlockChildren() failed with '%s'\n", err))
		fmt.Println(fmt.Sprintf("page.RawJSON: '%v'\n", rsp.Results))
		// ppJSON(rsp.RawJSON)
		return
	}
	showBlockChildren(&rsp)
}

func getBlockChildren(apiKey string, blockID string) {
	if blockID == "" {
		// a test page https://www.notion.so/Test-all-blocks-c969c9455d7c4dd79c7f860f3ace6429
		// with all block types
		// getBlockChildren2(apiKey, "c969c9455d7c4dd79c7f860f3ace6429")
		getBlockChildren2(apiKey, Notion_Page)
	} else {
		getBlockChildren2(apiKey, blockID)
	}
}

func showDatabaseProperty(name string, prop notion.DatabaseProperty) {
	fmt.Println(fmt.Sprintf("    property: '%s'\n", name))
	fmt.Println(fmt.Sprintf("      id: %s\n", prop.ID))
	fmt.Println(fmt.Sprintf("      type: %s\n", prop.Type))
	if prop.Number != nil {
		num := prop.Number
		fmt.Println(fmt.Sprintf("      format: %s\n", num.Format))
	} else if prop.Select != nil {
		sel := prop.Select
		for _, selopt := range sel.Options {
			fmt.Println(fmt.Sprintf("      sel opt: %s\n", selopt.ID))
			fmt.Println(fmt.Sprintf("        name: %s\n", selopt.Name))
			fmt.Println(fmt.Sprintf("        color: %s\n", selopt.Color))
		}
	} else if prop.MultiSelect != nil {
		msel := prop.MultiSelect
		for _, selopt := range msel.Options {
			fmt.Println(fmt.Sprintf("      sel opt: %s\n", selopt.ID))
			fmt.Println(fmt.Sprintf("        name: %s\n", selopt.Name))
			fmt.Println(fmt.Sprintf("        color: %s\n", selopt.Color))
		}
	} else if prop.Formula != nil {
		f := prop.Formula
		fmt.Println(fmt.Sprintf("      expression: %s\n", f.Expression))
	} else if prop.Relation != nil {
		r := prop.Relation
		fmt.Println(fmt.Sprintf("    database id: %s\n", r.DatabaseID))
		// fmt.Println(fmt.Sprintf("    synced prop idd: %s\n", r.SyncedPropID))
		// fmt.Println(fmt.Sprintf("    synced prop name: %s\n", r.SyncedPropName))
	} else if prop.Rollup != nil {
		r := prop.Rollup
		fmt.Println(fmt.Sprintf("      relation prop id: %s\n", r.RelationPropID))
		fmt.Println(fmt.Sprintf("      relation prop name: %s\n", r.RelationPropName))
		fmt.Println(fmt.Sprintf("      rollup prop id: %s\n", r.RollupPropID))
		fmt.Println(fmt.Sprintf("      rollup prop name: %s\n", r.RollupPropName))
		fmt.Println(fmt.Sprintf("      function: %s\n", r.Function))
	}
}

func showDatabaseInfo(db notion.Database) {
	fmt.Println(fmt.Sprintf("database:\n"))
	fmt.Println(fmt.Sprintf("  ID: '%s'\n", db.ID))
	fmt.Println(fmt.Sprintf("  CreatedTime: '%s'\n", db.CreatedTime))
	fmt.Println(fmt.Sprintf("  LastEditedTime: '%s'\n", db.LastEditedTime))
	// showRichText(1, "Title", db.Title)
	fmt.Println(fmt.Sprintf("  %d properties:\n", len(db.Properties)))
	for name, prop := range db.Properties {
		showDatabaseProperty(name, prop)
	}
}

func getDabaseInfo2(apiKey string, id string) {
	fmt.Println(fmt.Sprintf("getDatabaseInfo: id='%s'\n", id))
	c := notion.NewClient(apiKey)
	ctx := context.Background()
	// db, err := c.GetDatabase(ctx, id)
	db, err := c.FindDatabaseByID(ctx, id)
	if err != nil {
		fmt.Println(fmt.Sprintf("c.GetDatabase() failed with '%s'\n", err))
		// fmt.Sprintln("db.RawJSON: '%s'\n", db.RawJSON)
		// ppJSON(db.RawJSON)
		return
	}
	showDatabaseInfo(db)
}

func getDatabaseInfo(apiKey string, id string) {
	if id == "" {
		getDabaseInfo2(apiKey, Notion_DB)
		// getDabaseInfo2(apiKey, "ffbfda6791d34147b44a57ef83ab907a")
		// getDabaseInfo2(apiKey, "3acbc0fae5e34dfa9f3960d91cfb018a")
		// getDabaseInfo2(apiKey, "509fe00ee06448249687a4eb26bf9579")

	} else {
		getDabaseInfo2(apiKey, id)
	}
}

func showDatabaseQueryInfo(dqr *notion.DatabaseQueryResponse) {
	fmt.Println("ShowDatabaseQueryInfo:")
	fmt.Println(fmt.Sprintf("  hasMore: %v\n", dqr.HasMore))
	empty := ""
	if dqr.NextCursor != &empty {
		fmt.Println(fmt.Sprintf("  nextCurosr: %v\n", dqr.NextCursor))
	}
	fmt.Println(fmt.Sprintf("  %d rows:\n", len(dqr.Results)))
	for _, p := range dqr.Results {
		showPageInfo(&p)
	}
}

func showPageInfo(page *notion.Page) {
	fmt.Println("showPageInfo:")
	fmt.Println(fmt.Sprintf("  ID: '%s'\n", page.ID))
	fmt.Println(fmt.Sprintf("  CreatedTime: '%s'\n", page.CreatedTime))
	fmt.Println(fmt.Sprintf("  LastEditedTime: '%s'\n", page.LastEditedTime))
	if page.Parent.PageID != nil {
		fmt.Println(fmt.Sprintf("  Parent: page with ID '%s'\n", *page.Parent.PageID))
	} else if page.Parent.DatabaseID != nil {
		fmt.Println(fmt.Sprintf("  Parent: database with ID '%s'\n", *page.Parent.DatabaseID))
	} else {
		panic("both page.Parent.PageID or page.Parent.DatabaseID are nil")
	}
	fmt.Println(fmt.Sprintf("  Archived: '%v'\n", page.Archived))
	switch prop := page.Properties.(type) {
	case notion.PageProperties:
		fmt.Println("page properties:")
		showRichText(2, "Title", prop.Title.Title)
	case notion.DatabasePageProperties:
		fmt.Println(fmt.Sprintf("database properties (NYI): '%v'\n", prop))
	}
}

func getPageInfo2(apiKey, pageID string) {
	fmt.Println(fmt.Sprintf("getPageInfo: pageID='%s'\n", pageID))

	c := notion.NewClient(apiKey)
	ctx := context.Background()
	page, err := c.FindPageByID(ctx, pageID)
	if err != nil {
		fmt.Println(fmt.Sprintf("FindPageByID() failed with: '%s'\n", err))
		fmt.Println(fmt.Sprintf("page.RawJSON: '%s'\n", page.Properties))
		// ppJSON(page.)
		return
	}
	showPageInfo(&page)
}

func getPageInfo(apiKey, pageID string) {
	fmt.Println(fmt.Sprintf("pageID: %s", pageID))
	if pageID == "" {
		getPageInfo2(apiKey, Notion_Page)
	} else {
		getPageInfo2(apiKey, pageID)
	}
}

func showRichText(indent int, name string, richText []notion.RichText) {
	s := getIndent(indent)
	if name != "" {
		fmt.Println(fmt.Sprintf("%s%s: %v\n", s, name, richText))
		return
	}
	fmt.Println(fmt.Sprintf("%s%v\n", s, richText))
}

func ppJSON(d []byte) {
	res := pretty.Pretty(d)
	fmt.Println(fmt.Sprintf("pretty printed JSON:\n%s\n", res))
}

func getIndent(n int) string {
	s := ""
	for n > 0 {
		n -= 1
		s += "  "
	}
	return s
}

func queryDatabase2(apiKey, id string) {
	fmt.Println(fmt.Sprintf("queryDatabase: id='%s'\n", id))
	c := notion.NewClient(apiKey)
	ctx := context.Background()
	// none := nil
	// dqr, err := c.QueryDatabase(ctx, id, notion.DatabaseQuery{})
	dqr, err := c.QueryDatabase(ctx, id, notion.DatabaseQuery{
		Filter:      notion.DatabaseQueryFilter{},
		Sorts:       []notion.DatabaseQuerySort{},
		StartCursor: "",
		PageSize:    0,
	})
	if err != nil {
		fmt.Println(fmt.Sprintf("QueryDatabase() failed with '%s'\n", err))
		// fmt.Println(fmt.Sprintf("RawJSON '%s'\n", dqr.RawJSON))
		fmt.Println(fmt.Sprintf("RawJSON '%v'\n", dqr.Results))
		// ppJSON(dqr.Results)
		return
	}
	showDatabaseQueryInfo(&dqr)
}

func queryDatabase(apiKey, id string) {
	if id == "" {
		queryDatabase2(apiKey, Notion_DB)
	} else {
		queryDatabase2(apiKey, id)
	}
}

// func Test() {
// 	println("Hello from test function")
// 	client := notion.NewClient(Notion_Key)
// 	// var opts = PaginationQuery{"", 100}
// 	// users, err := client.ListUsers(context.Background(), &notion.PaginationQuery{StartCursor: "", PageSize: 100})
// 	pages, err := client.Search(context.Background(), &notion.SearchOpts{})
// 	// page, err := client.FindPageByID(Notion_DB)
// 	// users, err := client.ListUsers(context.Context{Deadline: false}, nil)
// 	// users, err := client.ListUsers(context.Context{Deadline: false}, nil)
// 	if err != nil {
// 		fmt.Println("Error finding page", err)
// 	}
// 	b, err := json.MarshalIndent(pages.Results, "", "    ")
// 	if err == nil {
// 		s := string(b)
// 		fmt.Println(s)
// 	}
// 	// println(fmt.Sprintf("%v, %T", pages.Results, pages.Results))
// 	// AddToWhatCanIDo()
// }
