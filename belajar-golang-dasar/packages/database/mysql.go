package database

var connection string

// automatically run when imported from another package
func init() {
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
