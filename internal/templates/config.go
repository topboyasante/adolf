package templates

type DBtype int

// extension for this could be useful for catering for other databases variables.
type AdolfDBConfig struct {
	DBName   string
	Host     string
	User     string
	Password string
	Port     int
}

const (
	POSTGRES DBtype = iota + 1
	MYSQL
	SQLITE
)

func (d DBtype) String() string {
	return [...]string{"postgres", "mysql", "sqlite"}[d-1]
}

func (d DBtype) EnumIndex() int {
	return int(d)
}
