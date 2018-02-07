package memory

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/xackery/xegony/model"
)

var (
	forumsDatabase = []*model.Forum{}
	forumLock      = sync.RWMutex{}
)

//GetForum will grab data from storage
func (s *Storage) GetForum(forum *model.Forum) (err error) {
	forumLock.RLock()
	defer forumLock.RUnlock()
	for _, tmpForum := range forumsDatabase {
		if tmpForum.ID == forum.ID {
			*forum = *tmpForum
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//CreateForum will grab data from storage
func (s *Storage) CreateForum(forum *model.Forum) (err error) {
	forumLock.Lock()
	defer forumLock.Unlock()
	for _, tmpForum := range forumsDatabase {
		if tmpForum.ID == forum.ID {
			err = fmt.Errorf("forum already exists")
			return
		}
	}
	forumsDatabase = append(forumsDatabase, forum)
	return
}

//ListForum will grab data from storage
func (s *Storage) ListForum(page *model.Page) (forums []*model.Forum, err error) {
	forumLock.RLock()
	defer forumLock.RUnlock()

	forums = make([]*model.Forum, len(forumsDatabase))

	forums = forumsDatabase

	switch page.OrderBy {
	case "short_name":
		sort.Slice(forums, func(i, j int) bool {
			return forums[i].Name < forums[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(forums))
		}
	*/
	return
}

//ListForumTotalCount will grab data from storage
func (s *Storage) ListForumTotalCount() (count int64, err error) {
	configLock.RLock()
	defer configLock.RUnlock()
	count = int64(len(forumsDatabase))
	return
}

//ListForumBySearch will grab data from storage
func (s *Storage) ListForumBySearch(page *model.Page, forum *model.Forum) (forums []*model.Forum, err error) {
	forumLock.RLock()
	defer forumLock.RUnlock()

	if len(forum.Name) > 0 {
		for i := range forumsDatabase {
			if strings.Contains(forumsDatabase[i].Name, forum.Name) {
				forums = append(forums, forumsDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}

	switch page.OrderBy {
	case "short_name":
		sort.Slice(forums, func(i, j int) bool {
			return forums[i].Name < forums[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(forums))
	//}
	return
}

//ListForumBySearchTotalCount will grab data from storage
func (s *Storage) ListForumBySearchTotalCount(forum *model.Forum) (count int64, err error) {
	forumLock.RLock()
	defer forumLock.RUnlock()

	forums := []*model.Forum{}
	if len(forum.Name) > 0 {
		for i := range forumsDatabase {
			if strings.Contains(forumsDatabase[i].Name, forum.Name) {
				forums = append(forums, forumsDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}
	count = int64(len(forums))
	return
}

//EditForum will grab data from storage
func (s *Storage) EditForum(forum *model.Forum) (err error) {
	forumLock.Lock()
	defer forumLock.Unlock()
	for i := range forumsDatabase {
		if forumsDatabase[i].ID == forum.ID {
			*forumsDatabase[i] = *forum
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//DeleteForum will grab data from storage
func (s *Storage) DeleteForum(forum *model.Forum) (err error) {
	forumLock.Lock()
	defer forumLock.Unlock()
	indexToDelete := 0
	for i := range forumsDatabase {
		if forumsDatabase[i].ID == forum.ID {
			indexToDelete = i
			break
		}
	}
	if indexToDelete < 1 {
		err = &model.ErrNoContent{}
		return
	}

	forumsDatabase[len(forumsDatabase)-1], forumsDatabase[indexToDelete] = forumsDatabase[indexToDelete], forumsDatabase[len(forumsDatabase)-1]
	forumsDatabase = forumsDatabase[:len(forumsDatabase)-1]
	return
}
