package models

import (
	"errors"
	"go-boilerplate/src/common"
	"go-boilerplate/src/core/db"
	"time"
)

func ArticlesModel() *BaseModel {
	mod := &BaseModel{
		ModelConstructor: &common.ModelConstructor{
			Repository: db.GetPostgresDB(),
		},
	}

	return mod
}

// models definitions
type Article struct {
	ID        int64     `db:"id, primarykey, autoincrement" json:"id"`
	Title     string    `db:"title" json:"title"`
	Content   string    `db:"content" json:"content"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type CreateArticleForm struct {
	Title   string `form:"title" json:"title" binding:"required,min=3,max=100"`
	Content string `form:"content" json:"content" binding:"required,min=3,max=1000"`
}

type UpdateArticleForm struct {
	Title   string `form:"title" json:"title"`
	Content string `form:"content" json:"content"`
}

type FindArticleForm struct {
	ID int64 `form:"id" json:"id" binding:"required"`
}

//ArticleModel ...
type ArticleModel struct{}

func (mod *BaseModel) GetAllArticles() (articles []Article, err error) {
	_, err = mod.Repository.Select(&articles, "select a.id, a.title, a.content, a.updated_at, a.created_at from public.articles a order by a.id desc")
	return articles, err
}

func (mod *BaseModel) GetOneArticle(id int64) (article Article, err error) {
	err = mod.Repository.SelectOne(&article, "SELECT a.id, a.title, a.content, a.updated_at, a.created_at FROM public.articles a WHERE a.id=$1 LIMIT 1", id)
	return article, err
}

func (mod *BaseModel) CreateArticle(form CreateArticleForm) (articleID int64, err error) {
	err = mod.Repository.QueryRow("INSERT INTO public.articles (title, content) VALUES($1, $2) RETURNING id", form.Title, form.Content).Scan(&articleID)
	return articleID, err
}

func (mod *BaseModel) UpdateArticle(id int64, form UpdateArticleForm) (err error) {
	_, err = mod.GetOneArticle(id)
	if err != nil {
		return err
	}

	operation, err := mod.Repository.Exec("UPDATE public.articles SET title=$2, content=$3 WHERE id=$1", id, form.Title, form.Content)
	if err != nil {
		return err
	}

	success, _ := operation.RowsAffected()
	if success == 0 {
		return errors.New("updated 0 records")
	}

	return err
}

func (mod *BaseModel) DeleteArticle(id int64) (err error) {
	operation, err := mod.Repository.Exec("DELETE FROM public.articles WHERE id=$1", id)
	if err != nil {
		return err
	}

	success, _ := operation.RowsAffected()
	if success == 0 {
		return errors.New("no records were deleted")
	}

	return err
}
