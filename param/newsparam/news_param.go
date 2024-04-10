package newsparam

import (
	"newsapp/entity"
)

type CreateNewsRequest struct {
	Title            string            `json:"title" form:"title"`
	ShortDescription string            `json:"short_description" form:"short_description"`
	Description      string            `json:"description" form:"description"`
	ImageFileName    string            `json:"image_file_name" form:"image_file_name"`
	Categories       []entity.Category `json:"categories" form:"categories"`
	CreatorUserID    uint              `json:"creator_user_id" form:"creator_user_id"`
}
