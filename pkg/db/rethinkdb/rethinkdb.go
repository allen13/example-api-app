package rethinkdb

import r "gopkg.in/dancannon/gorethink.v2"

type RethinkDB struct {
	Address  string `toml:"address"`
	Database string `toml:"database"`
	session  *r.Session
}

func (re *RethinkDB) Init() error {
	if re.Address == "" {
		re.Address = "localhost:28015"
	}
	if re.Database == "" {
		re.Database = "example-api-app"
	}

	return re.connect()
}
func (re *RethinkDB) connect() error {
	session, err := r.Connect(r.ConnectOpts{
		Address: re.Address,
	})
	if err != nil {
		return err
	}
	re.session = session

	err = re.createDBIfNotExist()
	if err != nil {
		return err
	}

	err = re.createTableIfNotExist("app")
	if err != nil {
		return err
	}

	return nil
}

func (re *RethinkDB) createDBIfNotExist() error {
	exists, err := re.dbExists()
	if err != nil {
		return err
	}

	if !exists {
		_, err := r.DBCreate(re.Database).RunWrite(re.session)
		if err != nil {
			return err
		}
	}

	return nil
}

func (re *RethinkDB) createTableIfNotExist(table string) error {
	exists, err := re.tableExists(table)
	if err != nil {
		return err
	}

	if !exists {
		_, err := r.DB(re.Database).TableCreate(table).RunWrite(re.session)
		if err != nil {
			return err
		}
	}

	return nil
}

func (re *RethinkDB) dbExists() (bool, error) {
	var response []interface{}
	res, err := r.DBList().Run(re.session)

	if err != nil {
		return false, err
	}

	err = res.All(&response)
	if err != nil {
		return false, err
	}

	for _, db := range response {
		if db == re.Database {
			return true, nil
		}
	}

	return false, nil
}

func (re *RethinkDB) tableExists(table string) (bool, error) {
	var response []interface{}
	res, err := r.DB(re.Database).TableList().Run(re.session)

	if err != nil {
		return false, err
	}

	err = res.All(&response)
	if err != nil {
		return false, err
	}

	for _, responseTable := range response {
		if responseTable == table {
			return true, nil
		}
	}

	return false, nil
}
