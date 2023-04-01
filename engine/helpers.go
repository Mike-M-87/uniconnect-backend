package engine

import (
	"errors"
	"fmt"
	"uniconnect/graph/model"
	"uniconnect/models"
	"uniconnect/utils"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm/clause"
)

var (
	ErrWrongCreds = errors.New("invalid email or password")
	ErrWrongToken = errors.New("invalid token")
)

func GenerateToken(authid uuid.UUID) (*model.AuthPayload, error) {
	authToken, err := utils.GenerateJWTForAuthId(authid)
	if err != nil {
		return nil, err
	}
	return &model.AuthPayload{
		Token: authToken,
	}, nil
}

func FetchUser(cond *models.User, with_assoc bool) (*models.User, error) {
	var user models.User
	var err error
	if with_assoc {
		err = utils.DB.Preload(clause.Associations).Where(cond).First(&user).Error
	} else {
		err = utils.DB.Where(cond).First(&user).Error
	}
	if err != nil {
		return nil, ErrWrongCreds
	}
	return &user, nil
}

func FetchBusinessById(id string) (*models.Business, error) {
	var biz models.Business
	err := utils.DB.Preload(clause.Associations).Where("id = ?", id).First(&biz).Error
	if err != nil {
		return nil, err
	}
	return &biz, nil
}

func FetchBusinessesByIds(ids []uuid.UUID) ([]models.Business, error) {
	var biz []models.Business
	err := utils.DB.Where("id in (?)", ids).Find(&biz).Error
	if err != nil {
		return nil, err
	}
	return biz, nil
}

func FetchBusinessesBySearchName(search string) ([]models.Business, error) {
	var bizbyNames []models.Business
	var bizbyDesc []models.Business
	symbol := "%"
	cond := fmt.Sprintf("%s%s%s", symbol, search, symbol)
	fmt.Println(cond)
	err := utils.DB.Where("name LIKE ?", cond).Find(&bizbyNames).Error
	if err != nil {
		return nil, err
	}
	err = utils.DB.Where("description LIKE ?", cond).Find(&bizbyDesc).Error
	if err != nil {
		return nil, err
	}
	return append(bizbyNames, bizbyDesc...), nil
}

func FetchBusinesses(cond *models.Business, limit int) ([]models.Business, error) {
	var biz []models.Business
	err := utils.DB.Where(cond).Find(&biz).Limit(limit).Error
	if err != nil {
		return nil, err
	}
	return biz, nil
}

func FetchUserByAuthToken(jwt string) (*models.User, error) {
	userId, err := utils.ValidateJWTForAuthId(jwt)
	if err != nil {
		return nil, ErrWrongToken
	}
	return FetchUserByID(userId)
}

func FetchUserByID(id string) (*models.User, error) {
	var user models.User
	err := utils.DB.Preload(clause.Associations).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
