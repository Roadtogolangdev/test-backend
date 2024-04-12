package banner

import (
	"database/sql"
	"time"
)

type Banner struct {
	ID        int                    `json:"banner_id"`
	TagIDs    []int                  `json:"tag_ids"`
	FeatureID int                    `json:"feature_id"`
	Content   map[string]interface{} `json:"content"`
	IsActive  bool                   `json:"is_active"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}

type BannerService struct {
	db      *sql.DB
	banners []Banner
	//add some banner
}
