package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserID(userID int) ([]Campaign, error)
	FindByID(ID int) (Campaign, error)
	Save(campaign Campaign) (Campaign, error)
	Update(campaign Campaign) (Campaign, error)
	CreateImage(CampaignImage CampaignImages) (CampaignImages, error)
	MarkAllImage(campaignID int) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRespository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Preload("CampaignImage", "campaign_images.is_primary = 1").Find(&campaigns).Error

	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (r *repository) FindByUserID(UserID int) ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Where("user_id = ?", UserID).Preload("CampaignImage", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (r *repository) FindByID(ID int) (Campaign, error) {
	var campaigns Campaign

	err := r.db.Preload("User").Preload("CampaignImage").Where("id = ?", ID).Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (r *repository) Save(campaign Campaign) (Campaign, error) {
	err := r.db.Create(&campaign).Error

	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *repository) Update(campaign Campaign) (Campaign, error) {
	err := r.db.Save(&campaign).Error

	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *repository) CreateImage(campaignImages CampaignImages) (CampaignImages, error) {
	err := r.db.Create(&campaignImages).Error

	if err != nil {
		return campaignImages, err
	}
	return campaignImages, nil
}

func (r *repository) MarkAllImage(campaignID int) (bool, error) {

	err := r.db.Model(&CampaignImages{}).Where("campaign_id = ?", campaignID).Update("is_primary", false).Error

	if err != nil {
		return false, err
	}
	return false, nil
}
