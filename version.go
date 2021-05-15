package gooselite

var tableName = "goose_db_version"

// TableName returns goose db version table name.
func TableName() string {
	return tableName
}

// SetTableName set goose db version table name.
func SetTableName(n string) {
	tableName = n
}
