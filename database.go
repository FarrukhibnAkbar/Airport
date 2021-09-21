package main

import(
	"fmt"
	"database/sql"
	_"github.com/lib/pq"
)

var(
	Host="localhost"
	Username="admin"
	Password=12345
	Database="airport"
	Port=5432
)

const QUERY =`
select 
	info_id,
	from_where, 
	to_where, 
	fly_time, 
	lnading_time 
from info;
`
type Air struct{
	FlyId 		int
	FromWhere 	string
	ToWhere 	string
	FlyTime		string
	LnadingTime string
}

func Info() []Air {
	c := fmt.Sprintf(
		"host=%s user=%s password=%d dbname=%s port=%d",
		Host, Username, Password, Database, Port,
	)

	db, err := sql.Open("postgres", c)

	if err != nil{
		panic(err)
	}

	defer db.Close()

	rows,_ := db.Query(QUERY)

	defer rows.Close()

	var infoes []Air

	for rows.Next() {
		var info Air

		rows.Scan(
			&info.FlyId, 
			&info.FromWhere,
			&info.ToWhere,
			&info.FlyTime,
			&info.LnadingTime,
		)

		infoes =append(infoes, info)
	}

	return infoes

}
