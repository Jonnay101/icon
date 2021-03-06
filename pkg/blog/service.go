package blog

type database interface {
	StoreBlogPost(*PostData) error
	FindBlogPostByKey(*RequestParams) (*PostData, error)
	FindAllBlogPosts(*RequestParams) ([]*PostData, error)
	UpdateBlogPost(*PostData) error
	RemoveBlogPost(*RequestParams) error
}

// Service holds the data and methods related to the Blog Service
type Service struct {
	DB database
}

// NewService -
func NewService(db database) *Service {
	return &Service{db}
}
