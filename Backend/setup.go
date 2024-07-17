package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v3"
	"os"
)

type MySQLConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Net      string `yaml:"net"`
	Address  string `yaml:"address"`
	Dbname   string `yaml:"dbname"`
}

func readConfig() (*MySQLConfig, error) {
	buf, err := os.ReadFile("db_credentials.yaml")
	if err != nil {
		return nil, err
	}

	config := &MySQLConfig{}
	err = yaml.Unmarshal(buf, config)
	if err != nil {
		return nil, fmt.Errorf("error parsing db_credentials.yaml")
	}

	return config, err
}

// Exec but with error handling and status printing
func exec(db *sql.DB, status string, cmd string) {
	fmt.Print(status + ": ")
	_, err := db.Exec(cmd)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successful")
}

// dropping and adding the table back
func resetDatabase(config *MySQLConfig) {
	serverDsn := fmt.Sprintf("%s:%s@%s(%s)/",
		config.Username, config.Password, config.Net, config.Address)
	server, err := sql.Open("mysql", serverDsn)
	if err != nil {
		panic(err.Error())
	}

	exec(server, "Dropping database "+config.Dbname, "DROP DATABASE IF EXISTS "+config.Dbname)
	exec(server, "Creating database "+config.Dbname, "CREATE DATABASE "+config.Dbname)

	server.Close()
}

func setup() {
	// Reading db_credentials.yaml
	config, err := readConfig()
	if err != nil {
		panic(err.Error())
	}

	resetDatabase(config)

	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s",
		config.Username, config.Password, config.Net, config.Address, config.Dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}

	// Creating tables
	exec(db, "Creating table 'level'",
		`CREATE TABLE level (
    		id TINYINT NOT NULL AUTO_INCREMENT, 
    		name TINYTEXT NOT NULL, 
    		PRIMARY KEY (id)
         )`)

	exec(db, "Creating table 'subject'",
		`CREATE TABLE subject (
			id SMALLINT NOT NULL AUTO_INCREMENT, 
			name TINYTEXT NOT NULL, 
			PRIMARY KEY (id),
			level_id TINYINT,
			CONSTRAINT fk_subject_level FOREIGN KEY (level_id) REFERENCES level(id)
		)`)

	exec(db, "Creating table 'base_user'",
		`CREATE TABLE base_user(
    		id INT NOT NULL AUTO_INCREMENT,
    		email VARCHAR(64) NOT NULL UNIQUE,
    		password TINYTEXT NOT NULL,
    		user_type TINYINT DEFAULT -1,
    		created_at BIGINT NOT NULL,
    		PRIMARY KEY (id)
		)`)

	// Gender == False => Male; True => Female; Null => Others
	exec(db, "Creating table 'user'",
		`CREATE TABLE user(
    		user_id INT NOT NULL,
    		username VARCHAR(64) NOT NULL,
    		picture TINYTEXT,
    		is_parent BOOLEAN NOT NULL,
    		gender BOOLEAN,
    		PRIMARY KEY (user_id),
            CONSTRAINT fk_user_base_user FOREIGN KEY (user_id) REFERENCES base_user(id) ON DELETE CASCADE
		)`)

	exec(db, "Creating table 'tutor'",
		`CREATE TABLE tutor (
    		user_id INT NOT NULL,
    		name TINYTEXT NOT NULL,
    		age TINYINT NOT NULL,
    		picture TINYTEXT,
    		gender BOOLEAN,
    		phone TINYTEXT NOT NULL,
    		description TEXT,
    		PRIMARY KEY (user_id),
    		CONSTRAINT fk_tutor_base_user FOREIGN KEY (user_id) REFERENCES base_user(id) ON DELETE CASCADE
		)`)

	exec(db, "Creating table 'tuition_center'",
		`CREATE TABLE tuition_center(
    		user_id iNT NOT NULL,
    		name TINYTEXT NOT NULL,
    		phone TINYTEXT NOT NULL,
    		picture TINYTEXT,
    		address TEXT NOT NULL,
    		address_link TEXT NOT NULL,
    		description TEXT,
    		website TEXT,
    		PRIMARY KEY (user_id),
    		CONSTRAINT fk_tuition_center_base_user FOREIGN KEY (user_id) REFERENCES base_user(id) ON DELETE CASCADE
		)`)

	exec(db, "Creating table 'qualification'",
		`CREATE TABLE qualification (
    		id INT NOT NULL AUTO_INCREMENT,
    		name TINYTEXT NOT NULL,
    		description TEXT NOT NULL,
    		time BIGINT NOT NULL,
    		PRIMARY KEY (id),
    		level_id TINYINT NOT NULL,
    		tutor_id INT NOT NULL,
    		CONSTRAINT fk_qualification_level FOREIGN KEY (level_id) REFERENCES level(id),
    		CONSTRAINT fk_qualification_tutor FOREIGN KEY (tutor_id) REFERENCES tutor(user_id) ON DELETE CASCADE
		)`)

	exec(db, "Creating table 'tuition_center_tutor_join'",
		`CREATE TABLE tuition_center_tutor_join (
    		tuition_center_id INT NOT NULL,
			tutor_id INT NOT NULL,
    		CONSTRAINT dk_tct_tuition_center FOREIGN KEY (tuition_center_id) REFERENCES tuition_center(user_id) ON DELETE CASCADE,
    		CONSTRAINT dk_tct_tutor FOREIGN KEY (tutor_id) REFERENCES tutor(user_id) ON DELETE CASCADE
		)`)

	exec(db, "Creating table 'rate'",
		`CREATE TABLE rate (
    		id INT NOT NULL AUTO_INCREMENT,
    		amount FLOAT NOT NULL,
    		is_open BOOLEAN NOT NULL,
    		subject_id SMALLINT NOT NULL,
    		tutor_id INT,
    		tuition_center_id INT,
    		PRIMARY KEY (id),
    		CONSTRAINT fk_rate_subject FOREIGN KEY (subject_id) REFERENCES subject(id),
    		CONSTRAINT fk_rate_tutor FOREIGN KEY (tutor_id) REFERENCES tutor(user_id) ON DELETE CASCADE,
    		CONSTRAINT fk_rate_tuition_center FOREIGN KEY (tuition_center_id) REFERENCES tuition_center(user_id) ON DELETE CASCADE
		)`)

	exec(db, "Creating table 'request'",
		`CREATE TABLE request (
    		id INT NOT NULL AUTO_INCREMENT,
    		description TEXT NOT NULL,
    		rate SMALLINT NOT NULL,
    		user_id INT NOT NULL,
    		subject_id SMALLINT NOT NULL,
    		level_id TINYINT NOT NULL,
    		PRIMARY KEY (id),
    		CONSTRAINT fk_request_user FOREIGN KEY (user_id) REFERENCES user(user_id) ON DELETE CASCADE,
    		CONSTRAINT fk_request_subject FOREIGN KEY (subject_id) REFERENCES subject(id),
    		CONSTRAINT fk_request_level FOREIGN KEY (level_id) REFERENCES level(id),
    		created_at BIGINT NOT NULL
		)`)

	exec(db, "Creating table 'refresh'",
		`CREATE TABLE refresh (
    		refresh_token TINYTEXT NOT NULL,
    		expiry BIGINT NOT NULL,
    		base_user_id INT NOT NULL UNIQUE,
    		CONSTRAINT fk_refresh_user FOREIGN KEY (base_user_id) REFERENCES base_user(id) ON DELETE CASCADE
		)`)

	// Seeding tables
	exec(db, "Seeding table 'level'",
		`INSERT INTO level (name) VALUES ("Primary"), ("N(T) Level"), ("N(A) Level"), ("O Level"), 
        ("Polytechnic"), ("A Level"), ("Bachelor"), ("Master"), ("Doctoral")`)

	exec(db, "Seeding table 'subject'",
		`INSERT INTO subject (name, level_id) VALUES
			-- Primary                            
			("English", 1), ("Mathematics", 1), ("Science", 1), ("Art", 1),
			("Music", 1), ("Social Studies", 1), ("Character & Citizenship Education", 1),
			("Chinese", 1), ("Malay", 1), ("Tamil", 1),
			
			-- N(T) Level
			("English", 2), ("Mathematics", 2), ("Nutrition & Food Science", 2), ("Art", 2),
			("Computer Applications", 2), ("Design & Technology", 2), ("Element of Business Skills", 2),
			("Mobile Robotics", 2), ("Smart Electrical Technology", 2), ("Retail Operations", 2),
			("Basic Chinese", 2), ("Basic Malay", 2), ("Basic Tamil", 2), ("Science", 2), ("Music", 2),
			
			-- N(A) Level
			("English", 3), ("Literature Elective", 3), ("Geography Elective", 3), ("History Elective", 3),
			("Social Studies", 3), ("Literature", 3), ("History", 3), ("Geography", 3), ("Bengali", 3), ("Gujarati", 3),
			("Hindi", 3), ("Panjabi", 3), ("Urdu", 3), ("Mathematics", 3), ("Additional Mathematics", 3),
			("Chemistry", 3), ("Physics", 3), ("Chemistry", 3), ("Nutrition & Food Science", 3),
			("Art", 3), ("Design & Technology", 3), ("Principles of Accounts", 3),
			("Chinese", 3), ("Malay", 3), ("Tamil", 3),
			
			-- O Level
        	("Arabic", 4), ("Bahasa Indonesia", 4), ("English", 4), 
        	("Literature", 4), ("History", 4), ("Geography", 4), ("Economics", 4),
        	("Social Studies", 4), ("History Elective", 4), ("Geography Elective", 4), ("Literature Elective", 4),
        	("Literature in Chinese", 4), ("Literature in Malay", 4), ("Literature in Tamil", 4),
        	("Drama", 4), ("Spanish", 4), ("Hindi", 4), ("Urdu", 4), ("Gujarati", 4), ("Panjabi", 4),
        	("Bengali", 4), ("Burmese", 4), ("Thai", 4), ("French", 4), ("German", 4), ("Japanese", 4),
        	("Mathematics", 4), ("Additional Mathematics", 4), ("Electronics", 4), ("Music", 4), ("Higher Music", 4),
        	("Physics Elective", 4), ("Chemistry Elective", 4), ("Biology Elective", 4),
        	("Physics", 4), ("Chemistry", 4), ("Biology", 4), ("Nutrition & Food Science", 4),
        	("Art", 4), ("Higher Art", 4), ("Design & Technology", 4), ("Business Studies", 4),
        	("Principles of Accounts", 4), ("Computing", 4),
        	("Biotechnology", 4), ("Design Studies", 4),
        	("Chinese", 4), ("Malay", 4), ("Tamil", 4),
        	("Chinese B", 4), ("Malay B", 4), ("Tamil B", 4),
        	("Higher Chinese", 4), ("Higher Malay", 4), ("Higher Tamil", 4),
        	("Chinese Special Programme", 4), ("Malay Special Programme", 4),
        	("Exercise & Sports Science", 4),
        	
        	-- A Level
        	("H1 General Paper", 6), ("H1 Geography", 6), ("H1 History", 6), ("H1 Bengali", 6),
        	("H1 Gujarati", 6), ("H1 Hindi", 6), ("H1 French", 6), ("H1 Literature", 6), ("H1 German", 6),
        	("H1 Japanese", 6), ("H1 Panjabi", 6), ("H1 Urdu", 6), ("H1 Economics", 6), ("H1 Mathematics", 6),
        	("H1 Physics", 6), ("H1 Chemistry", 6), ("H1 Biology", 6), ("H1 Art", 6), 
        	("H2 English Language & Linguistics", 6), ("H2 Literature", 6), ("H2 Theatre Studies & Drama", 6),
        	("H2 Computing", 6), ("H2 Economics", 6), ("H2 Spanish", 6), ("H2 Management of Business", 6), 
        	("H2 Principles of Accounting", 6), ("H2 China Studies", 6), ("H2 Further Mathematics", 6),
        	("H2 Chemistry", 6), ("H2 French", 6), ("H2 German", 6), ("H2 Japanese", 6), ("H2 Biology", 6),
        	("H2 Physics", 6), ("H2 Art", 6), ("H2 Geography", 6), ("H2 History", 6), ("H2 Music", 6),
        	("H2 Mathematics", 6), ("H2 Knowledge & Inquiry", 6),
        	("H3 Literature", 6), ("H3 Econom7ics", 6), ("H3 Chemistry", 6), ("H3 Physics", 6), ("H3 Biology", 6),
        	("H3 Art", 6), ("H3 Music", 6), ("H3 Mathematics", 6), ("H3 Geography", 6), ("H3 History", 6),
        	("H1 Chinese", 6), ("H1 Malay", 6), ("H1 Tamil", 6), ("H2 Translation (Chinese)", 6),
        	("H3 Chinese Language & Literature", 6), ("H2 Malay Language & Literature", 6), 
        	("H3 Malay Language & Literature", 6), ("H2 Tamil Language & Literature", 6), 
        	("H3 Tamil Language & Literature", 6), ("H1 Chinese B", 6), ("H1 Malay B", 6),
        	("H1 Tamil B", 6), ("H1 Project Work", 6)
		`)

	exec(db, "Seeding table 'base_user'",
		`INSERT INTO base_user (email, password, user_type, created_at) VALUES
		("Minh", "a4e6d1c002f05325e8ac6259a2031f8bdf27f4221ce1b3c413df290e9892f185", 0, 1715872529), 
		("MinhTutor", "a4e6d1c002f05325e8ac6259a2031f8bdf27f4221ce1b3c413df290e9892f185", 1, 1715872529), 
		("MinhTuitionCenter", "a4e6d1c002f05325e8ac6259a2031f8bdf27f4221ce1b3c413df290e9892f185", 2, 1715872529), 
		("MinhTutor2", "a4e6d1c002f05325e8ac6259a2031f8bdf27f4221ce1b3c413df290e9892f185", 1, 1715872529)
		`)

	exec(db, "Seeding table 'user'",
		`INSERT INTO user (user_id, username, picture, is_parent, gender) VALUES (1, 'minhmxc', 'https://i.scdn.co/image/ab67616100005174f3bb04995cb61f04936424ee', false, false)`)

	exec(db, "Seeding table 'tutor'",
		`INSERT INTO tutor (user_id, name, age, picture, phone, description, gender) VALUES
    	(2, "Minh", 20, "https://res.cloudinary.com/tuteria/image/upload/c_fit,h_627,q_80,w_1200/v1/landing_page_images/successful-teacher-2021-09-24-03-01-11-utc.jpg", 12345678, "Hi my name is Minh and I play Final Fantasy", true)`)

	exec(db, "Seeding table 'tuition_center'",
		`INSERT INTO tuition_center (user_id, name, phone, picture, address, address_link, description, website) 
		VALUES (3, "Minh's Academy of Excellence", 12345678, "https://static1.squarespace.com/static/592d012829687f03fe7f23ff/t/65e842e47e10007d19729b6e/1720061671489/", "19 Kent Ridge Crescent 119278", "https://g.co/kgs/d98nKNE", 
    	 "The school for top students only", "https://minhmxc.github.io/")`)

	exec(db, "Seeding table 'qualification'",
		`INSERT INTO qualification (name, description, time, level_id, tutor_id) VALUES
    	("A Level", "86.25/90", 1672531200, 6, 2)`)

	exec(db, "Seeding table 'tuition_center_tutor_join'",
		"INSERT INTO tuition_center_tutor_join VALUES (3, 2)")

	exec(db, "Seeding table 'rate'",
		`INSERT INTO rate (amount, is_open, subject_id, tutor_id, tuition_center_id) VALUES (69, true, 1, 2, null), (420, false, 1, null, 3)`)

	exec(db, "Seeding table 'request'",
		`INSERT INTO request (description, rate, user_id, subject_id, level_id, created_at) VALUES
		("Looking for Minh", "100", 1, 1, 1, 1715872529)`)
}
