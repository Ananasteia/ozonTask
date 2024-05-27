package repo

import (
	"AVITOtask/cmd/link_cutback/internal/app"
	"context"
	"database/sql"
	_ "github.com/lib/pq"
)

var _ app.RepoI = &Repo{}

type Repo struct {
	sql *sql.DB
}

func New() (*Repo, error) {
	db, err := sql.Open("postgres", "dbname=postgres password=pass user=user sslmode=disable  port=5432") // открывается соединение с базой данных, параметр1 - название драйвера (заодно и базы данных), параметр2 - логин, пароль, порт (+шифрование итд)
	if err != nil {
		return nil, err
	}
	return &Repo{db}, nil
}

func (r *Repo) Close() error {
	return r.sql.Close()
}

func (r *Repo) Save(ctx context.Context, l app.Link) error {
	const query = `insert into links (shortlink, longlink) values ($1, $2)` // тут вставляем в базу данных
	_, err := r.sql.ExecContext(ctx, query, l.ShortLink, l.LongLink)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repo) GetLongLink(ctx context.Context, l app.Link) (*app.Link, error) {
	saver := dbLinks{}
	const query = `select longlink from links where shortlink = $1` // тут достаем из базы данных, * выделить всю строку
	err := r.sql.QueryRowContext(ctx, query, l.ShortLink).Scan(&saver.LongLink)
	if err != nil {
		return nil, err
	}
	return saver.convert(), nil
}
