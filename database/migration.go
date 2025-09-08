package database

import "github.com/IsraelTeo/api-registry-court-files-MVP/model"

func MigrateDB() error {
	if err := GDB.AutoMigrate(
		&model.Court{},
		&model.Judge{},
		&model.Lawyer{},
		&model.Person{},
		&model.JudicialFile{},
	); err != nil {
		return err
	}

	return nil
}
