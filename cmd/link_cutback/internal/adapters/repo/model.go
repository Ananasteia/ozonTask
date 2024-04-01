package repo

import (
	"AVITOtask/cmd/link_cutback/internal/app"
	"time"
)

type dbLinks struct {
	LongLink  string
	ShortLink string
	CreatedAt time.Time
}

func (a dbLinks) convert() *app.Link {
	return &app.Link{LongLink: a.LongLink, ShortLink: a.ShortLink}
}
