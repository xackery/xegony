package memory

import (
	"fmt"
	"sort"
	"sync"

	"github.com/xackery/xegony/model"
)

var (
	zoneImagesDatabase = []*model.ZoneImage{}
	zoneImageLock      = sync.RWMutex{}
)

//GetZoneImage will grab data from storage
func (s *Storage) GetZoneImage(zoneImage *model.ZoneImage) (err error) {
	zoneImageLock.RLock()
	defer zoneImageLock.RUnlock()
	for _, tmpZoneImage := range zoneImagesDatabase {
		if tmpZoneImage.ID == zoneImage.ID {
			*zoneImage = *tmpZoneImage
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//CreateZoneImage will grab data from storage
func (s *Storage) CreateZoneImage(zoneImage *model.ZoneImage) (err error) {
	zoneImageLock.Lock()
	defer zoneImageLock.Unlock()
	for _, tmpZoneImage := range zoneImagesDatabase {
		if tmpZoneImage.ID == zoneImage.ID {
			err = fmt.Errorf("zoneImage already exists")
			return
		}
	}
	zoneImagesDatabase = append(zoneImagesDatabase, zoneImage)
	return
}

//ListZoneImage will grab data from storage
func (s *Storage) ListZoneImage(page *model.Page) (zoneImages []*model.ZoneImage, err error) {
	zoneImageLock.RLock()
	defer zoneImageLock.RUnlock()

	zoneImages = make([]*model.ZoneImage, len(zoneImagesDatabase))

	zoneImages = zoneImagesDatabase

	switch page.OrderBy {
	case "id":
		sort.Slice(zoneImages, func(i, j int) bool {
			return zoneImages[i].ID < zoneImages[j].ID
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(zoneImages))
		}
	*/
	return
}

//ListZoneImageTotalCount will grab data from storage
func (s *Storage) ListZoneImageTotalCount() (count int64, err error) {
	configLock.RLock()
	defer configLock.RUnlock()
	count = int64(len(zoneImagesDatabase))
	return
}

//ListZoneImageBySearch will grab data from storage
func (s *Storage) ListZoneImageBySearch(page *model.Page, zoneImage *model.ZoneImage) (zoneImages []*model.ZoneImage, err error) {
	zoneImageLock.RLock()
	defer zoneImageLock.RUnlock()

	if zoneImage.ID > 0 {
		for i := range zoneImagesDatabase {
			if zoneImagesDatabase[i].ID == zoneImage.ID {
				zoneImages = append(zoneImages, zoneImagesDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}

	switch page.OrderBy {
	case "id":
		sort.Slice(zoneImages, func(i, j int) bool {
			return zoneImages[i].ID < zoneImages[j].ID
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(zoneImages))
	//}
	return
}

//ListZoneImageBySearchTotalCount will grab data from storage
func (s *Storage) ListZoneImageBySearchTotalCount(zoneImage *model.ZoneImage) (count int64, err error) {
	zoneImageLock.RLock()
	defer zoneImageLock.RUnlock()

	zoneImages := []*model.ZoneImage{}
	if zoneImage.ID > 0 {
		for i := range zoneImagesDatabase {
			if zoneImagesDatabase[i].ID == zoneImage.ID {
				zoneImages = append(zoneImages, zoneImagesDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}
	count = int64(len(zoneImages))
	return
}

//EditZoneImage will grab data from storage
func (s *Storage) EditZoneImage(zoneImage *model.ZoneImage) (err error) {
	zoneImageLock.Lock()
	defer zoneImageLock.Unlock()
	for i := range zoneImagesDatabase {
		if zoneImagesDatabase[i].ID == zoneImage.ID {
			*zoneImagesDatabase[i] = *zoneImage
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//DeleteZoneImage will grab data from storage
func (s *Storage) DeleteZoneImage(zoneImage *model.ZoneImage) (err error) {
	zoneImageLock.Lock()
	defer zoneImageLock.Unlock()
	indexToDelete := 0
	for i := range zoneImagesDatabase {
		if zoneImagesDatabase[i].ID == zoneImage.ID {
			indexToDelete = i
			break
		}
	}
	if indexToDelete < 1 {
		err = &model.ErrNoContent{}
		return
	}

	zoneImagesDatabase[len(zoneImagesDatabase)-1], zoneImagesDatabase[indexToDelete] = zoneImagesDatabase[indexToDelete], zoneImagesDatabase[len(zoneImagesDatabase)-1]
	zoneImagesDatabase = zoneImagesDatabase[:len(zoneImagesDatabase)-1]
	return
}
