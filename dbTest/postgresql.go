package dbTest

import "fmt"
import "database/sql"
import _"github.com/lib/pq"

func getConnection() (*sql.DB) {

	dbInfo := fmt.Sprint("user=postgres password = postgres dbname=postgres sslmode=disable host=localhost port=5433")

	db, err := sql.Open("postgres", dbInfo)

	if err != nil {

		fmt.Println("connection problem", err.Error())
	}
	return db
}

func InsertQuery() {

	db := getConnection()

	query := `

           BEGIN;

            INSERT INTO print_history(document_no_key0,edition_data) values ('RWSS117715004',1);


INSERT INTO language(element_name_key0,language_data,element_display_data,create_by_data,create_datetime_data,update_by_data,update_datetime_key1)

VALUES ('NOYES','TH','ใช่','phaxsales','2016-09-09 15:33:00+0700','phaxsales','2016-09-09 15:33:00+0700');


            INSERT INTO language(element_name_key0,language_data) values ('ALLlows','TH');

            INSERT INTO user_authen(group_name_key0,program_code_data) values ('SE021','TDRAOO1');

            INSERT INTO user_group(groupid_key0,group_name_data) values ('25','SE034');

            INSERT INTO user_location(group_location_name_key0,location_code_data,update_by_data) values ('TSMSA','1234','phxsales');

            INSERT INTO public.user_login(userid_key0, location_code_data, group_name_data,update_by_data) VALUES ('Pathap43','40003','SM001','px');

            INSERT INTO vendor_master(vendor_code_key0,create_by_data) values ('123567','phxsales');

            INSERT INTO user_location(group_location_name_key0,location_code_data,create_by_data) values ('TDMSS','732','phxa');

            INSERT INTO user_group_location(userid_key0,group_location_name_data,create_by_data,update_by_data) values ('min123','THMS12','phxsales','phxsales');

            INSERT INTO user_component(group_name_key0,program_code_key1,component_data,enable_flag_data,create_by_data) values ('SMEDE23','TDHR22','resportText','1','phx');

            INSERT INTO todo_list(program_code_key0,todo_list_module_key1) values ('TDLSK99','request');

            INSERT INTO company(company_key0,company_name_data) values ('AIIS','super computers');

            COMMIT;
            `

	_, err := db.Exec(query)

	if err == nil {

		fmt.Println("inserted")
	}else {

		fmt.Println("problem with exec", err.Error())
	}
}

func SelectQuery() {

	db := getConnection()

	MapResult := make([]map[string]interface{}, 10)
	//query := "SELECT  update_by, array_to_json(edition_list), transaction_id, create_datetime, print_result, update_datetime, create_by  FROM print_history WHERE (   transaction_id = 'DBSSATTA001G148887383876543' ) OFFSET 0 LIMIT 1;"

	query := "SELECT * FROM language where element_name_key0='NOYES';"

	result, errs := db.Query(query)

	if errs != nil {

		fmt.Print("query not excuted", errs.Error())
	}

	columns, _ := result.Columns()

	data := make([]interface{}, len(columns))
	args := make([]interface{}, len(data))

	for i := range data {
		args[i] = &data[i]
	}

	for result.Next() {

		var rowData = make(map[string]interface{})

		if err := result.Scan(args...); err != nil {

		}

		for i := range data {

			rowData[ columns[i] ] = data[i]
		}

		MapResult = append(MapResult, rowData)
	}

	Result := make(map[string]interface{}, 2)

	for _, v := range MapResult {

		for key, value := range v {

			Result[key] = value

		}

	}

	fmt.Println(Result)
}
