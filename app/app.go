package app

type App struct {
	items map[int]string
	key   int
}

func NewApp() *App {
	items := map[int]string{1: "Task 1", 2: "Task 2", 3: "Task 3"}

	return &App{
		items: items,
		key:   4,
	}
}

func (app *App) Add(item string) {
	app.items[app.key] = item
	app.key += 1
}

func (app *App) All() map[int]string {
	return app.items
}

func (app *App) Delete(id int) {
	delete(app.items, id)
}
