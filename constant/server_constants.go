package constant

const (
	// PostgresSource denotes required data for db connection
	PostgresSource = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
)

const (
	// EnvVariablePath gives path where env file can be found
	EnvVariablePath = "/"
	// EnvFileName is name for the env file
	EnvFileName = "env"
	// EnvFileExtension is the extention of env file
	EnvFileExtension = "yaml"
	// DatabaseFileName is name for the database file
	DatabaseFileName = "database"
	// DatabaseFileExtension is the extention of database file
	DatabaseFileExtension = "yaml"
)
