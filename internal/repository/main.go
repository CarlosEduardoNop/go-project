package repository

import (
	"database/sql"
	"fmt"

	"github.com/CarlosEduardoNop/rabbitmq-test/pkg/db"
)

func Create(table string, data map[string]interface{}) sql.Result {
	fmt.Println("Creating")
	conn, err := db.OpenConnection()

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	columns := ""
	values := ""

	for key, value := range data {
		columns += fmt.Sprintf("%s,", key)

		if stringValue, ok := value.(string); ok {
			values += fmt.Sprintf("'%s',", stringValue)
			continue
		} 
		
		values += fmt.Sprintf("%v,", value)
	}

	columns = columns[:len(columns)-1]
	values = values[:len(values)-1]

	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, columns, values)

	res, err := conn.Exec(sql)

	if err != nil {
		fmt.Println(err)
	}

	return res
}
