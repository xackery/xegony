package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"sync"

	"github.com/go-yaml/yaml"
	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

var (
	zoneImageLock = sync.RWMutex{}
)

//GetZoneImage will grab data from file. This is an expensive call, it is recommended to use List instead and minimize reads.
func (s *Storage) GetZoneImage(zoneImage *model.ZoneImage) (err error) {
	zoneImageLock.Lock()
	defer zoneImageLock.Unlock()
	zoneImagesDatabase, err := s.readZoneImageFile()
	if err != nil {
		return
	}
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
	zoneImagesDatabase, err := s.readZoneImageFile()
	if err != nil {
		return
	}
	for _, tmpZoneImage := range zoneImagesDatabase {
		if tmpZoneImage.ID == zoneImage.ID {
			err = fmt.Errorf("zoneImage already exists")
			return
		}
	}
	zoneImagesDatabase = append(zoneImagesDatabase, zoneImage)
	err = s.writeZoneImageFile(zoneImagesDatabase)
	if err != nil {
		return
	}
	return
}

//ListZoneImage will grab data from storage
func (s *Storage) ListZoneImage(page *model.Page) (zoneImages []*model.ZoneImage, err error) {
	zoneImageLock.Lock()
	defer zoneImageLock.Unlock()
	zoneImagesDatabase, err := s.readZoneImageFile()
	if err != nil {
		return
	}

	zoneImages = make([]*model.ZoneImage, len(zoneImagesDatabase))

	zoneImages = zoneImagesDatabase

	if page.OrderBy == "" {
		page.OrderBy = "id"
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

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(zoneImages))
		}
	*/
	return
}

//ListZoneImageTotalCount will grab data from storage
func (s *Storage) ListZoneImageTotalCount() (count int64, err error) {
	zoneImageLock.Lock()
	defer zoneImageLock.Unlock()
	zoneImagesDatabase, err := s.readZoneImageFile()
	if err != nil {
		return
	}
	count = int64(len(zoneImagesDatabase))
	return
}

//ListZoneImageByBit will grab data from storage
func (s *Storage) ListZoneImageByBit(page *model.Page, zoneImage *model.ZoneImage) (zoneImages []*model.ZoneImage, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListZoneImageByBitTotalCount will grab data from storage
func (s *Storage) ListZoneImageByBitTotalCount(zoneImage *model.ZoneImage) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListZoneImageBySearch will grab data from storage
func (s *Storage) ListZoneImageBySearch(page *model.Page, zoneImage *model.ZoneImage) (zoneImages []*model.ZoneImage, err error) {
	zoneImageLock.Lock()
	defer zoneImageLock.Unlock()
	zoneImagesDatabase, err := s.readZoneImageFile()
	if err != nil {
		return
	}
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
	zoneImageLock.Lock()
	defer zoneImageLock.Unlock()
	zoneImagesDatabase, err := s.readZoneImageFile()
	if err != nil {
		return
	}

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
	zoneImagesDatabase, err := s.readZoneImageFile()
	if err != nil {
		return
	}

	for i := range zoneImagesDatabase {
		if zoneImagesDatabase[i].ID == zoneImage.ID {
			*zoneImagesDatabase[i] = *zoneImage
			err = s.writeZoneImageFile(zoneImagesDatabase)
			if err != nil {
				return
			}
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
	zoneImagesDatabase, err := s.readZoneImageFile()
	if err != nil {
		return
	}
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
	err = s.writeZoneImageFile(zoneImagesDatabase)
	if err != nil {
		return
	}
	return
}

func (s *Storage) readZoneImageFile() (zoneImages []*model.ZoneImage, err error) {
	yf, err := ioutil.ReadFile(s.path)
	if err != nil {
		if os.IsNotExist(err) {
			zoneImages = loadZoneImageDefault()
			err = os.MkdirAll(s.directory, 0744)
			if err != nil {
				err = errors.Wrapf(err, "failed to make directory %s", s.path)
				return
			}
			err = s.writeZoneImageFile(zoneImages)
			if err != nil {
				err = errors.Wrapf(err, "failed to write default zoneImage data to file %s", s.path)
				return
			}
			return
		}
		err = errors.Wrapf(err, "failed to read file %s", s.path)
		return
	}
	err = yaml.Unmarshal(yf, &zoneImages)
	if err != nil {
		err = errors.Wrapf(err, "failed to unmarshal file %s", s.path)
		return
	}

	return
}

func (s *Storage) writeZoneImageFile(zoneImages []*model.ZoneImage) (err error) {

	bData, err := yaml.Marshal(zoneImages)
	if err != nil {
		err = errors.Wrap(err, "failed to marshal zoneImages")
		return
	}
	err = ioutil.WriteFile(s.path, bData, 0744)
	if err != nil {
		err = errors.Wrapf(err, "failed to write file %s", s.path)
		return
	}
	return
}
