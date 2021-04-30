package darwin

//QLDialect implements Dialect interface for ql database
type QLDialect struct {
}

// CreateTableSQL returns the SQL to create the schema table
func (QLDialect) CreateTableSQL() string {
	return `
CREATE TABLE IF NOT EXISTS ` + tableName + `(
	version float,
	description string,
	checksum string,
	applied_at int64,
	execution_time int64,
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_versions on ` + tableName + `(version);
	`
}

// InsertSQL returns the SQL to insert a new migration in the schema table
func (QLDialect) InsertSQL() string {
	return `INSERT INTO ` + tableName + `
                (
                    version,
                    description,
                    checksum,
                    applied_at,
                    execution_time
                )
            VALUES ($1, $2, $3, $4, $5);`
}

// AllSQL returns a SQL to get all entries in the table
func (QLDialect) AllSQL() string {
	return `SELECT
                version,
                description,
                checksum,
                applied_at,
                execution_time
            FROM 
                ` + tableName + `
            ORDER BY version ASC;`
}
