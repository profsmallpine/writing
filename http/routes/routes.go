package routes

const (
	ArticleURL       = ArticlesURL + showBySlug
	ArticlesURL      = "/articles"
	LegacyArticleURL = "/writing" + showBySlug
	LoginURL         = "/login"
	LogoffURL        = "/logoff"
	RootURL          = "/"
	NewArticleURL    = ArticlesURL + new

	// Router actions to add to base resource URLs
	showBySlug = "/{" + MuxSlugParam + ":[a-z-]+}"
	new        = "/new"

	// Router variable helper strings
	MuxSlugParam = "slug"
)
