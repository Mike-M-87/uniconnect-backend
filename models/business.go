package models

import (
	"uniconnect/graph/model"

	uuid "github.com/satori/go.uuid"
)

type Business struct {
	Base
	UserID        uuid.UUID `gorm:"type:uuid"`
	OwnerUserName string
	Name          string
	Website       string
	Contact       string
	Description   string
	Location      string
	Type          model.BusinessType
	Image         string
	Comments      []Comment `gorm:"foreignKey:BusinessID"`
}

type Likes struct {
	UserID   uuid.UUID `gorm:"type:uuid"`
	BusinessID uuid.UUID `gorm:"type:uuid"`
}

func (b Business) CreateToGraphData() *model.Business {
	return &model.Business{
		ID:          b.ID.String(),
		Name:        b.Name,
		Type:        b.Type,
		Description: b.Description,
		Location:    b.Location,
		Website:     b.Website,
		Contact:     b.Contact,
		Image:       b.Image,
		OwnerName:   b.OwnerUserName,
	}
}
