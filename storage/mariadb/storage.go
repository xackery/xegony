package mariadb

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB
}

func (s *Storage) Initialize(config string) (err error) {
	if s.db != nil {
		return
	}
	if config == "" {
		user := os.Getenv("API_DB_USERNAME")
		if len(user) == 0 {
			user = "lfg"
		}
		pass := os.Getenv("API_DB_PASSWORD")
		if len(pass) == 0 {
			pass = "lfgpass"
		}
		host := os.Getenv("API_DB_HOSTNAME")
		if len(host) == 0 {
			host = "127.0.0.1"
		}
		port := os.Getenv("API_DB_PORT")
		if len(port) == 0 {
			port = "3306"
		}
		dbname := os.Getenv("API_DB_NAME")
		if len(dbname) == 0 {
			dbname = "lfg"
		}
		config = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", user, pass, host, port, dbname)
	}
	if s.db, err = sqlx.Open("mysql", config); err != nil {
		return
	}
	return
}

func (s *Storage) DropTables() (err error) {
	s.Initialize("")

	_, err = s.db.Exec(`SET FOREIGN_KEY_CHECKS = 0`)
	if err != nil {
		return
	}
	_, err = s.db.Exec(`drop table if exists game`)
	if err != nil {
		return
	}
	_, err = s.db.Exec(`drop table if exists lobby`)
	if err != nil {
		return
	}
	_, err = s.db.Exec(`drop table if exists user`)
	if err != nil {
		return
	}
	_, err = s.db.Exec(`SET FOREIGN_KEY_CHECKS = 1`)
	if err != nil {
		return
	}
	return
}

func (s *Storage) VerifyTables() (err error) {
	if err = s.Initialize(""); err != nil {
		return
	}
	_, err = s.db.Exec(`CREATE TABLE IF NOT EXISTS game (
  id int(11) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(32) NOT NULL DEFAULT '',
  image varchar(32) NOT NULL DEFAULT '',
  thumbnail varchar(32) NOT NULL DEFAULT '',
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;`)
	if err != nil {
		return
	}
	_, err = s.db.Exec(`CREATE TABLE IF NOT EXISTS lobby (
  gameid int(11) unsigned NOT NULL,
  id varchar(32) NOT NULL DEFAULT '',
  owneruserid int(11) unsigned NOT NULL,
  PRIMARY KEY (gameid,id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;`)
	if err != nil {
		return
	}
	_, err = s.db.Exec(`CREATE TABLE IF NOT EXISTS user (
  id int(11) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(32) NOT NULL DEFAULT '',
  email varchar(64) NOT NULL DEFAULT '',
  password varchar(64) NOT NULL DEFAULT '',
  isadmin tinyint(1) unsigned NOT NULL DEFAULT '0',
  ismoderator tinyint(1) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (id),
  UNIQUE KEY name (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;`)
	if err != nil {
		return
	}

	_, err = s.db.Exec(`INSERT INTO user (id, name, email, password, isadmin)
VALUES
	(1, 'Test', '', '$2a$10$YV0PiWDMiuXL4e77.jv8leD3NpDCk.v41aXPn7Yyi7fBWwBa0XzzC', 1);`)
	if err != nil {
		return
	}
	return
}
