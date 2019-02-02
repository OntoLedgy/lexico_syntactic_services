package configuration

type Csv_configuration struct {
	Csv_checks_required    bool           `json:"csv_checks_required"`
	Csv_file_name          string         `json:"csv_file_name"`
	Identity_column_number int            `json:"identity_column_number"`
	Column_set             map[string]int `json:"check_column_set"`
}

type Column_details map[string]Column

type Column struct {
	Column_name  string `json:"column_name"`
	Column_index int    `json:"column_index"`
}

type Database_configuration struct {
	Database_checks_required bool   `json:"database_checks_required"`
	Database_file_name       string `json:"csv_file_name"`
}

type Configuration struct {
	Csv_configuration     Csv_configuration
	Databse_configuration Database_configuration
}
