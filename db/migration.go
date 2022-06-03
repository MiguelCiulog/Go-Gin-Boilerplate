package db

const (
	Migration = "CREATE TABLE if not exists clientes\n" +
		"(`idclientes` INT NOT NULL AUTO_INCREMENT, " +
		"`nombre` VARCHAR(45) NOT NULL, " +
		"`email` VARCHAR(100) NULL, " +
		"`fecha_nacimiento` DATE NOT NULL, " +
		"PRIMARY KEY (`idclientes`));"
)

func Migrate() {
	_, err := DB.Query(Migration)

	if err != nil {
		panic(err.Error())
	}
}
