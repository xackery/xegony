package web

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) postRoutes() (routes []*route) {
	routes = []*route{
		//Post
		{
			"ListPost",
			"GET",
			"/topic/{topicID}",
			a.listPost,
		},
		{
			"GetPost",
			"GET",
			"/post/{postID}",
			a.getPost,
		},
	}
	return
}

func (a *Web) listPost(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site  site
		Posts []*model.Post
		Topic *model.Topic
		Forum *model.Forum
	}

	site := a.newSite(r)
	site.Page = "forum"
	site.Title = "Forum"

	topicID, err := getIntVar(r, "topicID")
	if err != nil {
		err = errors.Wrap(err, "topicID argument is required")
		return
	}

	topic := &model.Topic{
		ID: topicID,
	}
	err = a.topicRepo.Get(topic, user)
	if err != nil {
		return
	}
	forum := &model.Forum{
		ID: topic.ForumID,
	}
	err = a.forumRepo.Get(forum, user)
	if err != nil {
		return
	}
	posts, err := a.postRepo.ListByTopic(topic, user)
	if err != nil {
		return
	}

	content = Content{
		Site:  site,
		Posts: posts,
		Topic: topic,
		Forum: forum,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "post/list.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("post", tmp)
	}

	return
}

func (a *Web) getPost(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site  site
		Post  *model.Post
		Topic *model.Topic
		Forum *model.Forum
	}

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
		err = errors.Wrap(err, "Request error")
		return
	}

	topic := &model.Topic{
		ID: post.TopicID,
	}
	err = a.topicRepo.Get(topic, user)
	if err != nil {
		err = errors.Wrap(err, "Failed to get topic id")
		return
	}
	forum := &model.Forum{
		ID: topic.ForumID,
	}
	err = a.forumRepo.Get(forum, user)
	if err != nil {
		err = errors.Wrap(err, "Failed to get forum id")
		return
	}

	site := a.newSite(r)
	site.Page = "forum"
	site.Title = "Forum"

	content = Content{
		Site:  site,
		Post:  post,
		Topic: topic,
		Forum: forum,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "post/get.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("post", tmp)
	}

	return
}
