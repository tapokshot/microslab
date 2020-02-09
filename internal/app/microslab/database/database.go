package database

import (
    "database/sql"
    _ "github.com/lib/pq"
)

type Store struct {
    DB *sql.DB
}

func (store *Store) New(dbUrl string) (*Store, error) {

    db, err := sql.Open("postgres", dbUrl)
    if err != nil {
        return nil, err
    }

    if err := db.Ping(); err != nil {
        return nil, err
    }

    return &Store{
        DB: db,
    }, nil
}
