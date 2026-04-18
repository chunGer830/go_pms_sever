package database

type DbConn interface {
	Rollback()
	Commit()
}
