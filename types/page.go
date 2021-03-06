package types

type Page struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`

	Likes             int `json:"likes"`
	TalkingAboutCount int `json:"talking_about_count"`

	About       string `json:"about"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Link        string `json:"link"`
	Website     string `json:"website"`

	Cover CoverPhoto `json:"cover"`

	CanPost     bool `json:"can_post"`
	IsPublished bool `json:"is_published"`
}
