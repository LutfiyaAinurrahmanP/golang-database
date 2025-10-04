package repository

import (
	"context"
	"fmt"
	golangdatabase "golang-mysql"
	"golang-mysql/entity"
	"testing"
	_ "github.com/go-sql-driver/mysql"
)

func TestCommentRepository(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnection())
	ctx := context.Background()
	comment := entity.Comment{
		Email: "repo@test.com",
		Comment: "repo test",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnection())
	
	comment, err := commentRepository.FindById(context.Background(), 45)

	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnection())
	
	comments, err := commentRepository.FindAll(context.Background())

	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}