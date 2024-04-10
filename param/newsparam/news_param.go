package newsparam

import (
	"newsapp/entity"
)

type CreateNewsRequest struct {
	Title            string            `json:"title"`
	ShortDescription string            `json:"short_description"`
	Description      string            `json:"description"`
	ImageFileName    string            `json:"image_file_name"`
	Categories       []entity.Category `json:"categories"`
	CreatorUserID    uint              `json:"creator_user_id"`
}
