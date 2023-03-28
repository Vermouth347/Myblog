package vo

//为了方便，我们定义一个请求数据时的文章数据类型，保留文章增删查改操作的基本字段，以便后端接整个结构体并验证
type CreateArticleRequest struct {
	// 加上binging用于表单验证
	CategoryId uint   `json:"category_id" binging:"required"`
	Title      string `json:"title" binging:"required"`
	Content    string `json:"content" binging:"required"`
	HeadImage  string `json:"head_image"`
}
