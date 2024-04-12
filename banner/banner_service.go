package banner

import (
	"errors"
	"time"
)

func NewBanner() *BannerService {
	// инициализируем список баннеров
	banners := make([]Banner, 0)
	return &BannerService{
		banners: banners,
	}
}

func (s *BannerService) getUserBanner(tagID, featureID int, useLastRevision bool) (*Banner, error) {
	// Реализация запроса к базе данных для получения баннера для пользователя
	var banner Banner

	query := `
		SELECT id, tag_ids, feature_id, content, is_active, created_at, updated_at
		FROM banners
		WHERE tag_id = $1 AND feature_id = $2
	`

	err := s.db.QueryRow(query, tagID, featureID).Scan(&banner.ID, &banner.TagIDs, &banner.FeatureID, &banner.Content, &banner.IsActive, &banner.CreatedAt, &banner.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &banner, nil
}

func (s *BannerService) getAllBanners(adminToken string, featureID, tagID, limit, offset int) ([]Banner, error) {
	// Реализация запроса к базе данных для получения всех баннеров с учетом фильтров
	var banners []Banner

	query := `
		SELECT id, tag_ids, feature_id, content, is_active, created_at, updated_at
		FROM banners
		WHERE feature_id = $1 AND tag_id = $2
		LIMIT $3 OFFSET $4
	`

	rows, err := s.db.Query(query, featureID, tagID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var banner Banner
		err := rows.Scan(&banner.ID, &banner.TagIDs, &banner.FeatureID, &banner.Content, &banner.IsActive, &banner.CreatedAt, &banner.UpdatedAt)
		if err != nil {
			return nil, err
		}
		banners = append(banners, banner)
	}

	return banners, nil
}

func (s *BannerService) createBanner(adminToken string, tagIDs []int, featureID int, content map[string]interface{}, isActive bool) (int, error) {
	// Создание нового баннера
	banner := Banner{
		ID:        len(s.banners) + 1,
		TagIDs:    tagIDs,
		FeatureID: featureID,
		Content:   content,
		IsActive:  isActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	s.banners = append(s.banners, banner)
	return banner.ID, nil
}

func (s *BannerService) updateBanner(adminToken string, id int, tagIDs []int, featureID int, content map[string]interface{}, isActive bool) error {
	// Обновление содержимого баннера
	banner, err := s.findBannerByID(id)
	if err != nil {
		return err
	}

	banner.TagIDs = tagIDs
	banner.FeatureID = featureID
	banner.Content = content
	banner.IsActive = isActive
	banner.UpdatedAt = time.Now()

	return nil
}

func (s *BannerService) deleteBanner(adminToken string, id int) error {
	// Удаление баннера по идентификатору
	index, err := s.findIndexByID(id)
	if err != nil {
		return err
	}

	s.banners = append(s.banners[:index], s.banners[index+1:]...)
	return nil
}
func (s *BannerService) findBannerByID(id int) (*Banner, error) {
	for _, banner := range s.banners {
		if banner.ID == id {
			return &banner, nil
		}
	}
	return nil, errors.New("banner not found")
}

func (s *BannerService) findIndexByID(id int) (int, error) {
	for i, banner := range s.banners {
		if banner.ID == id {
			return i, nil
		}
	}
	return -1, errors.New("banner not found")
}
