package dba

type DataStore interface {
	WriteDB() error
}
type DataRetrieval interface {
	ReadDB() (interface{}, error)
}

type DataRemoval interface {
	DeleteDB() error
}

func WriteToDB(dbStore DataStore) error {
	err := dbStore.WriteDB()
	if err != nil {
		return err
	}
	return nil
}

func ReadFromDB(dbRead DataRetrieval) (interface{}, error) {
	resp, err := dbRead.ReadDB()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func DeleteFromDB(dbRemoval DataRemoval) error {
	err := dbRemoval.DeleteDB()
	if err != nil {
		return err
	}
	return nil
}
