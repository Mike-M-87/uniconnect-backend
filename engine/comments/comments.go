package comments

import (
	"time"
	"uniconnect/engine"
	"uniconnect/graph/model"
	"uniconnect/models"
	"uniconnect/utils"
)

func PostComment(input model.PostCommentInput) (bool, error) {
	user, err := engine.FetchUserByAuthToken(input.Token)
	if err != nil {
		return false, err
	}
	biz, err := engine.FetchBusinessById(input.BizID)
	if err != nil {
		return false, err
	}
	newComment := models.Comment{
		BusinessID: biz.ID,
		DateSent:   time.Now(),
		UserID:     user.ID,
		Sender:     user.UserName,
		Message:    input.Message,
	}
	err = utils.DB.Create(&newComment).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func FetchAllComments(token, bizId string) ([]*model.Comment, error) {
	_, err := engine.FetchUserByAuthToken(token)
	if err != nil {
		return nil, err
	}
	biz, err := engine.FetchBusinessById(bizId)
	if err != nil {
		return nil, err
	}
	gqlComments := make([]*model.Comment, len(biz.Comments))
	for i, comment := range biz.Comments {
		gqlComments[i] = comment.CreateToGraphData()
	}
	return gqlComments, nil
}
