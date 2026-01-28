package mock

import (
	"time"

	"github.com/recchia/snippetbox/internal/models"
	"github.com/recchia/snippetbox/internal/models/mysql"
)

var mockSnippet = mysql.Snippet{
	ID:      1,
	Title:   "An old silent pond",
	Content: "An old silent pond...",
	Created: time.Now(),
	Expires: time.Now(),
}

type SnippetModel struct{}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 2, nil
}

func (m *SnippetModel) Get(id int) (mysql.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet, nil
	default:
		return mysql.Snippet{}, models.ErrNoRecord
	}
}

func (m *SnippetModel) Latest() ([]mysql.Snippet, error) {
	return []mysql.Snippet{mockSnippet}, nil
}
