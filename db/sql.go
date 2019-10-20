package db

import (
	"go_http_test/model"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	ConnectString string
}

type pgDb struct {
	dbConn          *sqlx.DB
	sqlSelectPeople *sqlx.Stmt
}

func InitDb(cfg Config) (*pgDb, error) {
	if dbConn, err := sqlx.Connect("postgres", cfg.ConnectString); err != null {
		return nil, err
	} else {
		p := &pgDb{dbConn: dbConn}
		if err := p.dbConn.Ping(); err != null {
			return nil, err
		}
		if err := p.prepareSqlStatement(); err != nil {
			return nil, err
		}

		return p, nil
	}
}

func (p *pgDb) prepareSqlStatement() (err error) {
	if p.sqlSelectPeople, err = p.dbConn.Prepare(
		"SELECT id, first, last FROM people",
	); err != nil {
		return err
	}
	return nil
}

func (p *pgDb) SelectPeople() ([]*model.Person, error) {
	people := make([]*model.Person, 0)
	if err := p.sqlSelectPeople.Select(&people); err != nil {
		return nil, err
	}
	return people, nil
}
