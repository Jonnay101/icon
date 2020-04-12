package database

import (
	"github.com/Jonnay101/icon/pkg/blog"
	"github.com/Jonnay101/icon/pkg/glitch"
	"github.com/globalsign/mgo"
)

// Session -
type Session struct {
	*mgo.Session
}

// NewDatabaseSession -
func NewDatabaseSession(mongoURL string) (*Session, error) {

	mgoSession, err := mgo.Dial(mongoURL)
	return &Session{mgoSession}, err
}

// StoreBlogPost - store a blog post in the blog collection
func (s *Session) StoreBlogPost(blogPost *blog.PostData) error {

	blogPosts := s.DB("omfg").C("blog")

	q := blogPosts.FindId(blogPost.DatabaseKey)
	if n, _ := q.Count(); n > 0 {
		return glitch.ErrItemAlreadyExists
	}

	return blogPosts.Insert(blogPost)
}
