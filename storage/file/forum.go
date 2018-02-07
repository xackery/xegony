package file

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetForum will grab data from storage
func (s *Storage) GetForum(forum *model.Forum) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateForum will grab data from storage
func (s *Storage) CreateForum(forum *model.Forum) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListForum will grab data from storage
func (s *Storage) ListForum(page *model.Page) (forums []*model.Forum, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListForumTotalCount will grab data from storage
func (s *Storage) ListForumTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListForumBySearch will grab data from storage
func (s *Storage) ListForumBySearch(page *model.Page, forum *model.Forum) (forums []*model.Forum, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListForumBySearchTotalCount will grab data from storage
func (s *Storage) ListForumBySearchTotalCount(forum *model.Forum) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditForum will grab data from storage
func (s *Storage) EditForum(forum *model.Forum) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteForum will grab data from storage
func (s *Storage) DeleteForum(forum *model.Forum) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
