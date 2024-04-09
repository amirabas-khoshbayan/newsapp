package entity

import "time"

type News struct {
	ID               uint      `bson:"_id,omitempty" json:"id"`
	Title            string    `bson:"title,omitempty" json:"title"`
	ShortDescription string    `bson:"short_description,omitempty" json:"short_description"`
	Description      string    `bson:"description,omitempty" json:"description"`
	ImageFileName    string    `bson:"image_file_name,omitempty" json:"image_file_name"`
	CreatedAt        time.Time `bson:"created_at,omitempty" json:"created_at"`
	CreatorUserID    uint      `bson:"creator_user_id,omitempty" json:"creator_user_id"`
	VisitCount       int       `bson:"visit_count,omitempty" json:"visit_count"`
	LikeCount        int       `bson:"like_count,omitempty" json:"like_count"`
	Categories       []Category  `bson:"categories" json:"categories"`
