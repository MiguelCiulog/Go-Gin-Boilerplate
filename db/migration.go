package db

// Migration only for users
// TODO: Better migration
const (
	Migration = "CREATE TABLE if not exists users\n" +
		"(`id_user` INT NOT NULL AUTO_INCREMENT, " +
		"`name` VARCHAR(45) NOT NULL, " +
		"`email` VARCHAR(100) NULL, " +
		"`birth_date` DATE NOT NULL, " +
		"PRIMARY KEY (`id_user`));"
)

func Migrate() {
	_, err := db.Query(Migration)

	if err != nil {
		panic(err.Error())
	}
}
