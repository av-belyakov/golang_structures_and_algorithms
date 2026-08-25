package sqlite3api

func New(pathFile string) *Module {
	return &Module{
		pathFileDb: pathFile,
	}
}
