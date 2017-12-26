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
			user = "eqemu"
		}
		pass := os.Getenv("API_DB_PASSWORD")
		if len(pass) == 0 {
			pass = "eqemu"
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
			dbname = "eqemu"
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
	_, err = s.db.Exec(`drop table if exists character_data`)
	if err != nil {
		return
	}
	_, err = s.db.Exec(`drop table if exists account`)
	if err != nil {
		return
	}
	_, err = s.db.Exec(`drop table if exists user`)
	if err != nil {
		return
	}

	_, err = s.db.Exec(`drop table if exists forum`)
	if err != nil {
		return
	}

	_, err = s.db.Exec(`drop table if exists topic`)
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

	_, err = s.db.Exec(`CREATE TABLE IF NOT EXISTS character_data (
  id int(11) unsigned NOT NULL AUTO_INCREMENT,
  account_id int(11) unsigned NOT NULL,
  last_name varchar(64) NOT NULL DEFAULT '',
  title varchar(32) NOT NULL DEFAULT '',
  class tinyint(11) unsigned NOT NULL DEFAULT '0',
  level int(11) unsigned NOT NULL DEFAULT '0',
  zone_id int(11) unsigned NOT NULL DEFAULT '0',
  name varchar(32) NOT NULL DEFAULT '',
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;`)
	if err != nil {
		return
	}

	_, err = s.db.Exec(`CREATE TABLE IF NOT EXISTS account (
  id int(11) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(32) NOT NULL DEFAULT '',
  status int(5) NOT NULL DEFAULT '0',
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;`)
	if err != nil {
		return
	}

	_, err = s.db.Exec(`CREATE TABLE IF NOT EXISTS forum (
  id int(11) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(32) NOT NULL DEFAULT '',
  owner_id int(11) unsigned NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;`)
	if err != nil {
		return
	}

	_, err = s.db.Exec(`CREATE TABLE IF NOT EXISTS topic (
  id int(11) unsigned NOT NULL AUTO_INCREMENT,
  body varchar(1024) NOT NULL DEFAULT '',
  owner_id int(11) unsigned NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;`)
	if err != nil {
		return
	}

	_, err = s.db.Exec(`CREATE TABLE IF NOT EXISTS user (
  id int(11) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(32) NOT NULL DEFAULT '',
  account_id int(11) unsigned NOT NULL,
  email varchar(64) NOT NULL DEFAULT '',
  password varchar(64) NOT NULL DEFAULT '',
  PRIMARY KEY (id),
  UNIQUE KEY name (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;`)
	if err != nil {
		return
	}

	_, err = s.db.Exec(`INSERT INTO user (id, name, email, password, account_id)
VALUES
	(1, 'Test', '', '$2a$10$YV0PiWDMiuXL4e77.jv8leD3NpDCk.v41aXPn7Yyi7fBWwBa0XzzC', 1);`)
	if err != nil {
		return
	}

	_, err = s.db.Exec(`INSERT INTO account (id, name, status)
VALUES
	(1, 'Shin', 200);`)
	if err != nil {
		return
	}

	return
}
