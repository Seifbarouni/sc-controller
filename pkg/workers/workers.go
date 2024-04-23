package workers

import (
	"database/sql"
	"log"
	"time"

	"github.com/Seifbarouni/sc-controller/pkg/data"
	"github.com/Seifbarouni/sc-controller/pkg/validators"
)

func Install() {
	db, err := data.Init()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = runProcesses(db)
	if err != nil {
		log.Fatal(err.Error())
	}
	//go checkProcesses(db)
	checkProcesses(db)
}

func runProcesses(db *sql.DB) error {
	v := validators.Init()
	c, err := v.Validate()
	if err != nil {
		return err
	}
	log.Println(len(c.Processes))
	// wg.Add(len)
	// TODO: for each process, run it and store its info in the DB
	_ = db
	return nil
}

func checkProcesses(db *sql.DB) {
	//for range 1 {
	for {
		time.Sleep(time.Second * 3)
		sql := `SELECT id FROM process;`
		rows, err := db.Query(sql)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		for rows.Next() {
			var id int
			err := rows.Scan(&id)
			if err != nil {
				log.Println(err.Error())
			}
			log.Println(id)
			// TODO: check if the the process with PID=id is running, if not change its status
		}
		defer rows.Close()
	}
}
