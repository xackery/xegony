package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) postRoutes() (routes []*route) {
	routes = []*route{}
	return
}
func (a *API) getPost(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	postID, err := getIntVar(r, "postID")
	if err != nil {
		err = errors.Wrap(err, "postID argument is required")
		return
	}
	post := &model.Post{
		ID: postID,
	}
	err = a.postRepo.Get(post, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	content = post
	return
}

func (a *API) createPost(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	post := &model.Post{}
	err = decodeBody(r, post)
	if err != nil {
		return
	}
	err = a.postRepo.Create(post, user)
	if err != nil {
		return
	}
	content = post
	return
}

func (a *API) deletePost(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	postID, err := getIntVar(r, "postID")
	if err != nil {
		err = errors.Wrap(err, "postID argument is required")
		return
	}

	post := &model.Post{
		ID: postID,
	}

	err = a.postRepo.Delete(post, user)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *model.ErrNoContent:
			return
		default:
			err = errors.Wrap(err, "Request failed")
		}
		return
	}
	content = post
	return
}

func (a *API) editPost(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	postID, err := getIntVar(r, "postID")
	if err != nil {
		err = errors.Wrap(err, "postID argument is required")
		return
	}

	post := &model.Post{}
	err = decodeBody(r, post)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	post.ID = postID

	err = a.postRepo.Edit(post, user)
	if err != nil {
		return
	}
	content = post
	return
}

func (a *API) listPost(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	topicID, err := getIntVar(r, "topicID")
	if err != nil {
		err = errors.Wrap(err, "topicID argument is required")
		return
	}
	topic := &model.Topic{
		ID: topicID,
	}
	posts, err := a.postRepo.ListByTopic(topic, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	content = posts
	return
}
