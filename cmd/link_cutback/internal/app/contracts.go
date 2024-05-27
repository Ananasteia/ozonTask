package app

import "context"

type RepoI interface {
	Save(ctx context.Context, l Link) error
	GetLongLink(ctx context.Context, l Link) (*Link, error)
}
