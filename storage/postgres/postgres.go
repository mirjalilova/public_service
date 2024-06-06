package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"public/storage"
)

type Storage struct {
	Db      *sql.DB
	PublicS storage.PublicI
	PartyS  storage.PartyI
}

func NewPostgresStorage() (*Storage, error) {
	con := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", "postgres", "feruza1727", "localhost", 5432, "product_service")
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	var str = &Storage{Db: db}
	str.PublicS = NewPublicRepo(db)
	str.PartyS = NewPartyRepo(db)

	return str, nil
}

func (s *Storage) Public() storage.PublicI {
	if s.PublicS == nil {
		s.PublicS = NewPublicRepo(s.Db)
	}
	return s.PublicS
}

func (s *Storage) Product() storage.PartyI {
	if s.PartyS == nil {
		s.PartyS = NewPartyRepo(s.Db)
	}
	return s.PartyS
}
