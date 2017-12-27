package web

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) ListPost(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site  Site
		Posts []*model.Post
		Topic *model.Topic
		Forum *model.Forum
	}

	site := a.NewSite(r)
	site.Page = "post"
	site.Title = "Post"

	forumId, err := getIntVar(r, "topicId")
	if err != nil {
		err = errors.Wrap(err, "topicId argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	posts, err := a.postRepo.List(forumId)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	var topic *model.Topic
	var forum *model.Forum

	if len(posts) > 0 {
		topic, err = a.topicRepo.Get(posts[0].TopicId)
		if err != nil {
			err = errors.Wrap(err, fmt.Sprintf("Failed to get topic id: %d", posts[0].TopicId))
			a.writeError(w, r, err, http.StatusBadRequest)
			return
		}

		forum, err = a.forumRepo.Get(topic.ForumId)
		if err != nil {
			err = errors.Wrap(err, "Failed to get forum id")
			a.writeError(w, r, err, http.StatusBadRequest)
			return
		}
	}
	content := Content{
		Site:  site,
		Posts: posts,
		Topic: topic,
		Forum: forum,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "post/list.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("post", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) GetPost(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site  Site
		Post  *model.Post
		Topic *model.Topic
		Forum *model.Forum
	}

	id, err := getIntVar(r, "postId")
	if err != nil {
		err = errors.Wrap(err, "postId argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	post, err := a.postRepo.Get(id)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	var topic *model.Topic
	var forum *model.Forum

	topic, err = a.topicRepo.Get(post.TopicId)
	if err != nil {
		err = errors.Wrap(err, "Failed to get topic id")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	forum, err = a.forumRepo.Get(topic.ForumId)
	if err != nil {
		err = errors.Wrap(err, "Failed to get forum id")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.NewSite(r)
	site.Page = "post"
	site.Title = "Post"

	content := Content{
		Site:  site,
		Post:  post,
		Topic: topic,
		Forum: forum,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "post/get.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("post", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}
