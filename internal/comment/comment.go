package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New(("failed to fetch comment by id"))
	ErrNotImplemented  = errors.New("not implemented")
)

// Comment - a respresentation of the comment structure for the service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store - this interface define all methods that the service needs to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
	CreateComment(context.Context, Comment) (Comment, error)
	UpdateComment(context.Context, string, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
}

// Service -  is the struct where all the logic is built on top of
type Service struct {
	Store Store
}

// NewService - returns a pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retrieving a comment")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println((err))
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	createdCmt, err := s.Store.CreateComment(ctx, cmt)
	if err != nil {
		return Comment{}, err
	}
	return createdCmt, ErrNotImplemented
}

func (s *Service) UpdateComment(ctx context.Context, ID string, cmt Comment) (Comment, error) {
	updatedCmt, err := s.Store.UpdateComment(ctx, ID, cmt)
	if err != nil {
		fmt.Println("error updating the comment")
		return Comment{}, err
	}
	return updatedCmt, nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return s.Store.DeleteComment(ctx, id)
}
