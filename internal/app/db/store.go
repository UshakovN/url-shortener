package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Store struct {
	config *Config
	db     *sql.DB
}

func NewStore(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open(s.config.driverName, s.config.databaseUrl)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	log.Println("connected to postgres database")
	s.db = db
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) GetItem(shortUrl string) (string, bool) {
	return "", true
}

func (s *Store) PutItem(shortUrl, sourceUrl string) {

}
