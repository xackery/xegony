package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetMerchant will grab data from storage
func (s *Storage) GetMerchant(merchantID int64) (merchant *model.Merchant, err error) {
	merchant = &model.Merchant{}
	err = s.db.Get(merchant, fmt.Sprintf("SELECT merchantid FROM merchantlist WHERE merchantid = ? GROUP BY merchantid"), merchantID)
	if err != nil {
		return
	}
	return
}

//ListMerchant will grab data from storage
func (s *Storage) ListMerchant(pageSize int64, pageNumber int64) (merchants []*model.Merchant, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT merchantid FROM merchantlist 
		 GROUP BY merchantid ORDER BY merchantid ASC LIMIT %d OFFSET %d `, pageSize, pageSize*pageNumber))
	if err != nil {

		return
	}

	for rows.Next() {
		merchant := model.Merchant{}
		if err = rows.StructScan(&merchant); err != nil {
			return
		}
		merchants = append(merchants, &merchant)
	}
	return
}

//ListMerchantCount will grab data from storage
func (s *Storage) ListMerchantCount() (count int64, err error) {
	err = s.db.Get(&count, `SELECT count(merchantid) FROM merchantlist GROUP BY merchantid`)
	if err != nil {
		return
	}
	return
}

//SearchMerchant will grab data from storage
func (s *Storage) SearchMerchant(search string) (merchants []*model.Merchant, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT merchantid FROM merchantlist 
		WHERE id like ? ORDER BY merchantid ASC GROUP BY merchantid`), "%"+search+"%")
	if err != nil {
		return
	}

	for rows.Next() {
		merchant := model.Merchant{}
		if err = rows.StructScan(&merchant); err != nil {
			return
		}
		merchants = append(merchants, &merchant)
	}
	return
}

//DeleteMerchant will grab data from storage
func (s *Storage) DeleteMerchant(merchantID int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM merchantlist WHERE merchantid = ?`, merchantID)
	if err != nil {
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if affected < 1 {
		err = &model.ErrNoContent{}
		return
	}
	return
}
