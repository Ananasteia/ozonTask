package app

type App struct {
	repo RepoI
}

func New(r RepoI) *App {
	return &App{
		repo: r,
	}
}
