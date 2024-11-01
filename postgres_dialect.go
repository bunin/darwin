package darwin

// PostgresDialect a Dialect configured for PostgreSQL
type PostgresDialect struct{}

// CreateTableSQL returns the SQL to create the schema table
func (p PostgresDialect) CreateTableSQL() string {
	return `CREATE TABLE IF NOT EXISTS ` + tableName + `
                (
                    id             SERIAL                  NOT NULL,
                    version        REAL                    NOT NULL,
                    description    CHARACTER VARYING (255) NOT NULL,
                    checksum       CHARACTER VARYING (32)  NOT NULL,
                    applied_at     INTEGER                 NOT NULL,
                    execution_time REAL                    NOT NULL,
                    UNIQUE         (version),
                    PRIMARY KEY    (id)
                );`
}

// InsertSQL returns the SQL to insert a new migration in the schema table
func (p PostgresDialect) InsertSQL() string {
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
func (p PostgresDialect) AllSQL() string {
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
