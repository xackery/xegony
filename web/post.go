package web

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) listPost(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site  site
		Posts []*model.Post
		Topic *model.Topic
		Forum *model.Forum
	}

	site := a.newSite(r)
	site.Page = "forum"
	site.Title = "Forum"

	forumID, err := getIntVar(r, "topicID")
	if err != nil {
		err = errors.Wrap(err, "topicID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	posts, err := a.postRepo.List(forumID)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	var topic *model.Topic
	var forum *model.Forum

	if len(posts) > 0 {
		topic, err = a.topicRepo.Get(posts[0].TopicID)
		if err != nil {
			err = errors.Wrap(err, fmt.Sprintf("Failed to get topic id: %d", posts[0].TopicID))
			a.writeError(w, r, err, http.StatusBadRequest)
			return
		}

		forum, err = a.forumRepo.Get(topic.ForumID)
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

func (a *Web) getPost(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site  site
		Post  *model.Post
		Topic *model.Topic
		Forum *model.Forum
	}

	id, err := getIntVar(r, "postID")
	if err != nil {
		err = errors.Wrap(err, "postID argument is required")
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

	topic, err = a.topicRepo.Get(post.TopicID)
	if err != nil {
		err = errors.Wrap(err, "Failed to get topic id")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	forum, err = a.forumRepo.Get(topic.ForumID)
	if err != nil {
		err = errors.Wrap(err, "Failed to get forum id")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.newSite(r)
	site.Page = "forum"
	site.Title = "Forum"

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
