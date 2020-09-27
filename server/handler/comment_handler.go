/**
 * @author Riku Nunokawa
 * @template writer Futa Nakayama
 */

package handler

import (
	"database/sql"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/shortintern2020-C-cryptograph/TeamF/server/gen/models"
	"github.com/shortintern2020-C-cryptograph/TeamF/server/gen/restapi/scenepicks"
	"log"
	"time"
)

func GetCommentById(p scenepicks.GetCommentByIDParams) middleware.Responder {
	id := p.ID
	offset := p.Offset
	limit := p.Limit
	fmt.Printf("GET /comment offset: %d, limit: %d\n", offset, limit)

	if id < 0 {
		return scenepicks.NewGetDialogBadRequest().WithPayload("query id is invalid")
	}
	if offset < 0 || limit <= 0 {
		return scenepicks.NewGetDialogBadRequest().WithPayload("parameter value is invalid")
	}

	resDialog, resComments, resTags, err := getCommentByID(id, offset, limit)
	if err != nil {
		log.Fatal(err)
	}

	params := &scenepicks.GetCommentByIDOKBody{
		Message:  "success",
		Dialog:   resDialog,
		Comments: resComments,
		Tags:     resTags,
	}

	return scenepicks.NewGetCommentByIDOK().WithPayload(params)
}

func PostCommentById(p scenepicks.PostCommentByIDParams) middleware.Responder {

	// TODO: ここでfirebase認証

	// テスト実行用に仮でユーザをDBに登録
	idToken := p.Token
	client := NewClient(idToken)
	if client.err != nil {
		fmt.Printf("%v\n", client.err)
		return scenepicks.NewPostDialogBadRequest()
	}
	userId := client.user.ID
	dialogId := p.ID
	comment := p.Comment.Comment

	if dialogId <= 0 {
		return scenepicks.NewGetDialogBadRequest().WithPayload("dialog id is invalid")
	}
	if comment == "" {
		return scenepicks.NewGetDialogBadRequest().WithPayload("vacant comment is invalid")
	}

	// DBへの書き込み
	id, err := postComment(userId, dialogId, comment)
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

	DisplayName sql.NullString `json:"display_name" db:"display_name"`

	FirebaseUID string `db:"firebase_uid"`

	PhotoURL sql.NullString `json:"photo_url" db:"photo_url"`

	CTime time.Time `json:"ctime" db:"ctime"`

	UTime time.Time `json:"utime" db:"utime"`
}

func mapComment(c comment) models.Comment {
	res := models.Comment{
		Content: c.Content,
		User: &models.User{
			DisplayName: c.DisplayName.String,
			ID:          c.UserID,
			PhotoURL:    c.PhotoURL.String,
		},
	}
	return res
}
