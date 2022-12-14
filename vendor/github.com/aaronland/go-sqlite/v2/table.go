package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"regexp"
)

var re_mem *regexp.Regexp
var re_vfs *regexp.Regexp
var re_file *regexp.Regexp

func init() {
	re_mem = regexp.MustCompile(`^(file\:)?\:memory\:.*`)
	re_vfs = regexp.MustCompile(`^vfs:\.*`)
	re_file = regexp.MustCompile(`^file\:([^\?]+)(?:\?.*)?$`)
}

type Table interface {
	Name() string
	Schema() string
	InitializeTable(context.Context, Database) error
	IndexRecord(context.Context, Database, interface{}) error
}

func HasTable(ctx context.Context, db Database, table string) (bool, error) {

	// you might be thinking it would be a good idea to cache this lookup
	// I know I did... and I was wrong (20180713/thisisaaronland)
	// https://github.com/whosonfirst/go-whosonfirst-sqlite/issues/11

	dsn := db.DSN(ctx)

	check_tables := true
	has_table := false

	if !re_mem.MatchString(dsn) && !re_vfs.MatchString(dsn) {

		test := dsn

		if re_file.MatchString(test) {

			s := re_file.FindAllStringSubmatch(dsn, -1)

			if len(s) == 1 && len(s[0]) == 2 {
				test = s[0][1]
			}
		}

		_, err := os.Stat(test)

		if os.IsNotExist(err) {
			check_tables = false
			has_table = false
		}
	}

	if check_tables {

		conn, err := db.Conn(ctx)

		if err != nil {
			return false, err
		}

		has_table, err = HasTableWithSQLDB(ctx, conn, table)

		if err != nil {
			return false, err
		}
	}

	return has_table, nil
}

func HasTableWithSQLDB(ctx context.Context, db *sql.DB, table_name string) (bool, error) {

	has_table := false

	sql := "SELECT name FROM sqlite_master WHERE type='table'"

	rows, err := db.QueryContext(ctx, sql)

	if err != nil {
		return false, fmt.Errorf("Failed to query sqlite_master, %w", err)
	}

	defer rows.Close()

	for rows.Next() {

		var name string
		err := rows.Scan(&name)

		if err != nil {
			return false, fmt.Errorf("Failed scan table name, %w", err)
		}

		if name == table_name {
			has_table = true
			break
		}
	}

	return has_table, nil
}

func CreateTableIfNecessary(ctx context.Context, db Database, t Table) error {

	create := false

	has_table, err := HasTable(ctx, db, t.Name())

	if err != nil {
		return err
	}

	if !has_table {
		create = true
	}

	if create {

		sql := t.Schema()

		conn, err := db.Conn(ctx)

		if err != nil {
			return err
		}

		_, err = conn.Exec(sql)

		if err != nil {
			return err
		}

	}

	return nil
}
