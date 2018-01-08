package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	merchantEntryFields = `merchantlist.slot, merchantlist.item, merchantlist.faction_required, merchantlist.level_required, merchantlist.alt_currency_cost, merchantlist.classes_required, merchantlist.probability`
	merchantEntrySets   = `slot=:slot, item=:item, faction_required=:faction_required, level_required=:level_required, alt_currency_cost=:alt_currency_cost, classes_required=:classes_required, probability=:probability`
	merchantEntryBinds  = `:slot, :item, :faction_required, :level_required, :alt_currency_cost, :classes_required, :probability`
)

//GetMerchantEntry will grab data from storage
func (s *Storage) GetMerchantEntry(merchantID int64, itemID int64) (query string, merchantEntry *model.MerchantEntry, err error) {
	merchantEntry = &model.MerchantEntry{}
	query = fmt.Sprintf(`SELECT merchantid, %s FROM merchantlist 
		WHERE merchantlist.merchantid = ? AND merchantlist.itemid = ?`, merchantEntryFields)
	err = s.db.Get(merchantEntry, query, merchantID, itemID)
	if err != nil {
		return
	}
	return
}

//CreateMerchantEntry will grab data from storage
func (s *Storage) CreateMerchantEntry(merchantEntry *model.MerchantEntry) (query string, err error) {
	if merchantEntry == nil {
		err = fmt.Errorf("Must provide merchantEntry")
		return
	}

	query = fmt.Sprintf(`INSERT INTO merchantlist(%s)
		VALUES (%s)`, merchantEntryFields, merchantEntryBinds)
	_, err = s.db.NamedExec(query, merchantEntry)
	if err != nil {
		return
	}
	return
}

//ListMerchantEntry will grab data from storage
func (s *Storage) ListMerchantEntry(merchantID int64) (query string, merchantEntrys []*model.MerchantEntry, err error) {
	query = fmt.Sprintf(`SELECT %s FROM merchantlist WHERE merchantid = ?`, merchantEntryFields)
	rows, err := s.db.Queryx(query, merchantID)
	if err != nil {
		return
	}

	for rows.Next() {
		merchantEntry := model.MerchantEntry{}
		if err = rows.StructScan(&merchantEntry); err != nil {
			return
		}
		merchantEntrys = append(merchantEntrys, &merchantEntry)
	}
	return
}

//ListMerchantEntryByItem will grab data from storage
func (s *Storage) ListMerchantEntryByItem(itemID int64) (query string, merchantEntrys []*model.MerchantEntry, err error) {

	query = fmt.Sprintf(`SELECT merchantid, %s FROM merchantlist
	WHERE item = ? LIMIT 10`, merchantEntryFields)

	rows, err := s.db.Queryx(query, itemID)
	if err != nil {
		return
	}

	for rows.Next() {
		merchantEntry := model.MerchantEntry{}
		if err = rows.StructScan(&merchantEntry); err != nil {
			return
		}
		merchantEntrys = append(merchantEntrys, &merchantEntry)
	}
	return
}

//EditMerchantEntry will grab data from storage
func (s *Storage) EditMerchantEntry(merchantID int64, itemID int64, merchantEntry *model.MerchantEntry) (query string, err error) {

	query = fmt.Sprintf(`UPDATE merchantlist SET %s WHERE merchantlist.merchantid = ? AND merchantlist.item = ?`, merchantEntrySets)
	merchantEntry.MerchantID = merchantID
	merchantEntry.ItemID = itemID
	result, err := s.db.NamedExec(query, merchantEntry)
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

//DeleteMerchantEntry will grab data from storage
func (s *Storage) DeleteMerchantEntry(merchantID int64, itemID int64) (query string, err error) {
	query = `DELETE FROM merchantlist WHERE merchantid = ? AND item = ?`
	result, err := s.db.Exec(query, merchantID, itemID)
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
