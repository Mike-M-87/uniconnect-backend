package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.27

import (
	"context"
	"uniconnect/engine/business"
	"uniconnect/engine/comments"
	"uniconnect/engine/likes"
	"uniconnect/engine/users"
	"uniconnect/graph/model"
)

// Register is the resolver for the Register field.
func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.AuthPayload, error) {
	return users.ValidateRegistration(input)
}

// Login is the resolver for the Login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthPayload, error) {
	return users.Login(input)
}

// CreateBusiness is the resolver for the CreateBusiness field.
func (r *mutationResolver) CreateBusiness(ctx context.Context, input model.CreateBusinessInput) (bool, error) {
	return business.CreateNewBusiness(input)
}

// PostComment is the resolver for the PostComment field.
func (r *mutationResolver) PostComment(ctx context.Context, input model.PostCommentInput) (bool, error) {
	return comments.PostComment(input)
}

// LikeBusiness is the resolver for the LikeBusiness field.
func (r *mutationResolver) LikeBusiness(ctx context.Context, token string, bizID string) (bool, error) {
	return likes.AddBusinessLike(token, bizID)
}

// DeleteBusiness is the resolver for the DeleteBusiness field.
func (r *mutationResolver) DeleteBusiness(ctx context.Context, token string, bizID string) (bool, error) {
	return business.DeleteMyBusiness(token, bizID)
}

// EditBusiness is the resolver for the EditBusiness field.
func (r *mutationResolver) EditBusiness(ctx context.Context, input model.CreateBusinessInput, bizID string) (bool, error) {
	return business.EditBusiness(input, bizID)
}

// ChangePassword is the resolver for the ChangePassword field.
func (r *mutationResolver) ChangePassword(ctx context.Context, input model.ChangePasswordInput) (bool, error) {
	return users.ChangePassword(input)
}

// VerifyUser is the resolver for the VerifyUser field.
func (r *mutationResolver) VerifyUser(ctx context.Context, token string, otp string) (bool, error) {
	return users.VerifyEmail(token, otp)
}

// FetchUserData is the resolver for the fetchUserData field.
func (r *queryResolver) FetchUserData(ctx context.Context, token string) (*model.User, error) {
	return users.FetchProfile(token)
}

// CheckUserName is the resolver for the checkUserName field.
func (r *queryResolver) CheckUserName(ctx context.Context, name string) (bool, error) {
	return users.CheckUserNameAvaliability(name)
}

// FetchBusiness is the resolver for the FetchBusiness field.
func (r *queryResolver) FetchBusiness(ctx context.Context, token string, bizID string) (*model.Business, error) {
	return business.FetchBusinessData(token, bizID)
}

// FetchBusinessList is the resolver for the FetchBusinessList field.
func (r *queryResolver) FetchBusinessList(ctx context.Context, input model.FetchBusinessListInput) ([]*model.Business, error) {
	return business.FetchAllBusinesses(input)
}

// FetchComments is the resolver for the FetchComments field.
func (r *queryResolver) FetchComments(ctx context.Context, token string, bizID string) ([]*model.Comment, error) {
	return comments.FetchAllComments(token, bizID)
}

// FetchLikedBusiness is the resolver for the FetchLikedBusiness field.
func (r *queryResolver) FetchLikedBusiness(ctx context.Context, token string) ([]*model.Business, error) {
	return likes.FetchAllLikedBusinesses(token)
}

// CheckIFLiked is the resolver for the CheckIFLiked field.
func (r *queryResolver) CheckIFLiked(ctx context.Context, token string, bizID string) (bool, error) {
	return likes.CheckBusinessLiked(token, bizID)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
