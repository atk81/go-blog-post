package model

import (
	"log"

	"github.com/atk81/go-blog-post/utils"
	"github.com/google/uuid"
)

type Article struct {
	ID      uuid.UUID    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func AddArticle(title string, content string) (uuid.UUID, error) {
	db := utils.GetDB()
	sqlQuery := `INSERT INTO ARTICLES(Title, Content) VALUES($1, $2) RETURNING ID`
	var id uuid.UUID
	err := db.QueryRow(sqlQuery, title, content).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute query %v", err)
	}
	return id, err
}

func GetArticle(article_id uuid.UUID)(Article,error){
	db := utils.GetDB()
	var article Article
	sqlQuery := `SELECT * FROM ARTICLES WHERE ID=$1`
	err := db.QueryRow(sqlQuery, article_id).Scan(&article.ID, &article.Title, &article.Content)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	return article,err
}

func GetAllArticles() ([]Article, error) {
	db := utils.GetDB()
	var Articles []Article
	sqlQuery := `SELECT * FROM ARTICLES`
	rows, err := db.Query(sqlQuery)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var article Article
		err := rows.Scan(&article.ID, &article.Title, &article.Content)
		if err != nil {
			log.Fatalf("Unable to scan the row %v", err)
		}
		Articles = append(Articles, article)
	}
	return Articles, err

}
