package persistence

import (
	"errors"
	"goworkshop2/model"
)

func (store *GormDataStore) GetAuthor(uuid string) (model.Author, error) {
	var author model.Author
	err := store.DBInstance.Where(&model.Author{Entity: model.Entity{UUID: uuid}}).First(&author).Error
	return author, err
}

func (store *GormDataStore) CreateAuthor(author *model.Author) error {
	return store.DBInstance.Create(&author).Error
}

func (store *GormDataStore) GetAuthors() ([]model.Author, error) {
	var authors []model.Author
	err := store.DBInstance.
	//fetch the author
		Find(&authors).Error
	return authors, err
}

func (store *GormDataStore) DeleteAuthor(uuid string) error {
	if isEmptyUUID(uuid) {
		return errors.New("Please supply a primary key when attempting to delete an object.")
	}
	db := store.DBInstance.
		Delete(model.Author{}, whereAuthorUUIDEquals(uuid))

	if db.Error != nil {
		return db.Error
	} else if db.RowsAffected == 0 {
		return errors.New("No rows have been deleted for uuid = " + uuid)
	} else if db.RowsAffected > 1 {
		return errors.New("Multiple rows have been deleted for uuid = " + uuid)
	}
	return nil
}

func (store *GormDataStore) UpdateAuthor(author *model.Author) error {
	return store.DBInstance.Save(&author).Error
}

//creates the GORM where clause used to identify a book by uuid
func whereAuthorUUIDEquals(uuid string) (interface{}) {
	return model.Author{Entity: model.Entity{UUID: uuid}}
}
