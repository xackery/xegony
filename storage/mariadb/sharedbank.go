package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	sharedBankFields = `acctid, slotid, itemid, charges, augslot1, augslot2, augslot3, augslot4, augslot5, augslot6, custom_data`
	sharedBankSets   = `acctid=:acctid, slotid=:slotid, itemid=:itemid, charges=:charges, augslot1=:augslot1, augslot2=:augslot2, augslot3=:augslot3, augslot4=:augslot4, augslot5=:augslot5, augslot6=:augslot6, custom_data=:custom_data`
	sharedBankBinds  = `:acctid, :slotid, :itemid, :charges, :augslot1, :augslot2, :augslot3, :augslot4, :augslot5, :augslot6, :custom_data`
)

//GetSharedBank will grab data from storage
func (s *Storage) GetSharedBank(sharedBank *model.SharedBank) (err error) {
	err = s.db.Get(sharedBank, fmt.Sprintf("SELECT %s FROM sharedbank WHERE acctid = ? AND slotid = ?", sharedBankFields), sharedBank.AccountID, sharedBank.SlotID)
	if err != nil {
		return
	}
	return
}

//CreateSharedBank will grab data from storage
func (s *Storage) CreateSharedBank(sharedBank *model.SharedBank) (err error) {
	if sharedBank == nil {
		err = fmt.Errorf("Must provide sharedBank")
		return
	}

	_, err = s.db.NamedExec(fmt.Sprintf(`INSERT INTO sharedbank(%s)
		VALUES (%s)`, sharedBankFields, sharedBankBinds), sharedBank)
	if err != nil {
		return
	}
	return
}

//ListSharedBank will grab data from storage
func (s *Storage) ListSharedBankByAccount(account *model.Account, pageSize int64, pageNumber int64) (sharedBanks []*model.SharedBank, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT %s FROM sharedbank
		WHERE acctid = ? 
		ORDER BY slotid ASC LIMIT %d OFFSET %d`, sharedBankFields, pageSize, pageSize*pageNumber), account.ID)
	if err != nil {
		return
	}

	for rows.Next() {
		sharedBank := model.SharedBank{}
		if err = rows.StructScan(&sharedBank); err != nil {
			return
		}
		sharedBanks = append(sharedBanks, &sharedBank)
	}
	return
}

//ListSharedBankCount will grab data from storage
func (s *Storage) ListSharedBankByAccountCount(account *model.Account) (count int64, err error) {
	err = s.db.Get(&count, `SELECT count(id) FROM sharedbank WHERE acctid = ?`, account.ID)
	if err != nil {
		return
	}
	return
}

//ListSharedBankByItem will grab data from storage
func (s *Storage) ListSharedBankByAccountAndItem(account *model.Account, item *model.Item) (sharedBanks []*model.SharedBank, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT %s FROM sharedbank		
		WHERE acctid = ? AND itemid = ? ORDER BY slotid ASC`, sharedBankFields), account.ID, item.ID)
	if err != nil {
		return
	}

	for rows.Next() {
		sharedBank := model.SharedBank{}
		if err = rows.StructScan(&sharedBank); err != nil {
			return
		}
		sharedBanks = append(sharedBanks, &sharedBank)
	}
	return
}

//EditSharedBank will grab data from storage
func (s *Storage) EditSharedBank(sharedBank *model.SharedBank) (err error) {

	result, err := s.db.NamedExec(fmt.Sprintf(`UPDATE sharedbank SET %s WHERE acctid = :acctid AND slotid = :slotid`, sharedBankSets), sharedBank)
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

//DeleteSharedBank will grab data from storage
func (s *Storage) DeleteSharedBank(sharedBank *model.SharedBank) (err error) {
	result, err := s.db.Exec(`DELETE FROM sharedbank WHERE acctid = ? AND slotid = ?`, sharedBank.AccountID, sharedBank.SlotID)
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
