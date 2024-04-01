package repo

import (
	"AVITOtask/cmd/link_cutback/internal/app"
	"context"
	"database/sql"
)

var _ app.RepoI = &Repo{}

type Repo struct {
	sql *sql.DB
}

func New() (*Repo, error) {
	db, err := sql.Open("", "")
	if err != nil {
		return nil, err
	}
	return &Repo{db}, nil
}

func (a *Repo) Close() error {
	return a.sql.Close()
}

func (a *Repo) Save(ctx context.Context, l app.Link) error {
	const query = `insert into links (shortlink, longlink) values ($1, $2)` // тут вставляем в базу данных
	_, err := a.sql.ExecContext(ctx, query, l.ShortLink, l.LongLink)
	if err != nil {
		return err
	}
	return nil
}
func (a *Repo) GetShortLink(ctx context.Context, l app.Link) (*app.Link, error) {
	saver := dbLinks{}
	const query = `select shortlink from links where longlink = $1` // тут достаем из базы данных, * выделить всю строку
	err := a.sql.QueryRowContext(ctx, query, l.LongLink).Scan(&saver.ShortLink)
	if err != nil {
		return nil, err
	}
	return saver.convert(), nil
}
