package handler

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/shortintern2020-C-cryptograph/TeamF/server/gen/models"
	"github.com/shortintern2020-C-cryptograph/TeamF/server/gen/restapi/scenepicks"
	"log"
	"time"
)

func GetCommentById(p scenepicks.GetCommentByIDParams) middleware.Responder {
	offset := p.Offset
	limit := p.Limit
	fmt.Printf("offset: %d, limit: %d", offset, limit)

	id := p.ID
	schema, err := getCommentByID(id)
	if err != nil {
		log.Fatal(err)
	}

	params := &scenepicks.GetCommentOKBody{
		Message: "success",
		Schema:  schema,
	}

	return scenepicks.NewGetCommentOK().WithPayload(params)
}

func PostCommentById(p scenepicks.PostCommentByIDParams) middleware.Responder {
	comment := p.Comment.Comment
	fmt.Printf("comment: %s", comment)

	// TODO: ここでfirebase認証

	// テスト実行用に仮でユーザをDBに登録
	firebaseUid := "12345"
	displayName := "John Doe"
	photoUrl := "https://example.com"
	id, err := postUser(firebaseUid, displayName, photoUrl)
	if err != nil {
		log.Println(err)
	}

	// DBへの書き込み
	dialogId := p.ID
	id, err = postComment(comment, id, dialogId)
	if err != nil {
		log.Println(err)
	}

	params := &scenepicks.PostCommentByIDOKBody{
		Message: "success",
		ID:      id,
	}

	return scenepicks.NewPostCommentByIDOK().WithPayload(params)
}

type comment struct {
	ID int64 `json:"id" db:"id"`

	Content string `json:"content" db:"content"`

	UserID int64 `json:"user_id" db:"user_id"`

	DialogID int64 `json:"dialog_id" db:"dialog_id"`

	CTime time.Time `json:"ctime" db:"ctime"`

	UTime time.Time `json:"utime" db:"utime"`
}

func mapComment(c comment) models.Comment {
	res := models.Comment{
		Content: c.Content,
		User: &models.User{
			DisplayName: "Twitter 太郎",
			ID:          1,
			PhotoURL:    "https://example.com",
		},
	}
	return res
}
