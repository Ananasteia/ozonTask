package app

import "context"

func (a *App) HandlePost(ctx context.Context, ll string) (string, error) {
	myLink := Link{ShortLink: shortLinkCreator(), LongLink: ll}
	err := a.repo.Save(ctx, myLink)
	if err != nil {
		return "", err
	}
	return myLink.ShortLink, nil
}

func (a *App) HandleGet(ctx context.Context, l Link) (*Link, error) {
	myLink, err := a.repo.GetLongLink(ctx, l)
	if err != nil {
		return nil, err
	}
	return myLink, nil
}
