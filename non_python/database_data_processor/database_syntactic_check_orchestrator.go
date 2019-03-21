package database_data_processor

import (
	"database_manager/database/access"
	"database_manager/utils"
	storage_csv "storage/csv"
	//storage_slices "storage/slices"
	"fmt"
	"syntactic_checker/syntactic_check_cell_orchestrator"
)

func Execute_database_syntactic_checks(configuration_database_filename string) {

	//#TODO Add menu/choice of source type (database vs csv)
	//----------------------------- Database route

	//open database
	configuration_database, database_open_error :=
		database.OpenMsAccessDatabase(
			configuration_database_filename)

	if database_open_error != nil {
		panic(database_open_error)
	}
	fmt.Printf( //#TODO add to logger
		"\n+++Loading Config+++\n from %s",
		configuration_database_filename)

	load_configurations(configuration_database)

	defer configuration_database.Database.Close()

}

func load_configurations(
	accessDatabase *database.MsAccessDatabase) {

	//var check_types [][]interface{}
	var checks [][]interface{}
	//var non_parameterised_check_regexes [][]interface{}
	var in_scope_table_sets [][]interface{}
	var in_scope_column_sets [][]interface{}
	var in_scope_column_check_configurations [][]interface{}
	//var check_configurations [][]interface{}  -- NOT NEEDED

	//get checks

	checks_table_name := "non_range_syntactic_checks"
	checks_table_schema := []string{
		"uuids",
		"check_names",
		"check_regex_string",
		"replacement_string"}

	check_types_rows := database.ReadMsAccessColumns(
		accessDatabase,
		checks_table_name,
		checks_table_schema)

	checks = utils.Convert_rows_to_2d_slices( //#TODO explore if a differnet data structure can be used (pros and cons of using type - struct or map?)
		check_types_rows)

	//get names of tables in scope.

	in_scope_tables_table_name :=
		"in_scope_tables"
	in_scope_tables_table_schema :=
		[]string{
			"uuids",
			"table_names",
			"row_identity_column_names"}

	in_scope_table_sets_rows :=
		database.ReadMsAccessColumns(
			accessDatabase,
			in_scope_tables_table_name,
			in_scope_tables_table_schema)
	//#TODO move this back into database module - should output interface rather than sql.Rows.
	in_scope_table_sets =
		utils.Convert_rows_to_2d_slices(
			in_scope_table_sets_rows)

	//get names of table columns in scope

	in_scope_table_columns_table_name :=
		"in_scope_table_in_scope_columns"
	in_scope_table_columns_table_column_names :=
		[]string{
			"uuids",
			"column_names",
			"parent_table_uuids"}

	in_scope_column_sets_rows :=
		database.ReadMsAccessColumns(
			accessDatabase,
			in_scope_table_columns_table_name,
			in_scope_table_columns_table_column_names)
	//#TODO move this back into database module - should output interface rather than sql.Rows.
	in_scope_column_sets =
		utils.Convert_rows_to_2d_slices(
			in_scope_column_sets_rows)

	//get column check configurations

	in_scope_column_check_configurations_table_name :=
		"check_configurations"
	in_scope_column_check_configurations_table_schema := []string{
		"column_uuids",
		"check_uuids",
	}

	in_scope_column_check_configurations_rows :=
		database.ReadMsAccessColumns(
			accessDatabase,
			in_scope_column_check_configurations_table_name,
			in_scope_column_check_configurations_table_schema)

	in_scope_column_check_configurations =
		utils.Convert_rows_to_2d_slices(
			in_scope_column_check_configurations_rows)
	//#TODO move this back into database module - should output interface rather than sql.Rows.

	fmt.Printf("+++++++++++++++++++Loaded config+++++++++\n"+ //#TODO add to logger
		"tables :\n%s\n,"+
		"columns:\n%s\n, "+
		"checks:\n%s\n,"+
		"check configurations:\n%s\n",
		in_scope_table_sets,
		in_scope_column_sets,
		checks,
		in_scope_column_check_configurations)

	load_data(
		accessDatabase,
		in_scope_table_sets,
		in_scope_column_sets,
		checks,
		in_scope_column_check_configurations)

}

func load_data(
	accessDatabase *database.MsAccessDatabase,
	in_scope_table_sets [][]interface{},
	in_scope_column_sets [][]interface{},
	checks [][]interface{},
	in_scope_column_check_configurations [][]interface{}) {

	var in_scope_column_dataset [][]interface{}

	//for each table get column data

	fmt.Printf("+++++++++++++++++++++++++++loading data+++++++++++++++++\n")

	for _, in_scope_table := range in_scope_table_sets {
		//TODO# check if this should be broken out into two functions (reduce nesting)

		fmt.Printf(
			"---------loading data for table: %s (uuid : %s)\n",
			in_scope_table[1].(string), //table name
			in_scope_table[0])          //table uuid
		//#TODO enumerate the indexes, check if there is a better way of making the indexes clearer.  Alternatively, use a clearly named variable.

		for _, in_scope_column := range in_scope_column_sets {

			//checking for parent table of (column in_scope_column[1]) matches uuid of parent table (in_scope_table[0])

			if in_scope_column[2] == // parent table uuid
				in_scope_table[0] { //table uuid
				// #TODO enumerate the indexes, check if there is a better way of making the indexes clearer.  Alternatively, use a clearly named variable.

				in_scope_column_query_parameters :=
					[]string{in_scope_table[2].(string)} //Add table identity column to column query

				in_scope_column_query_parameters =
					append(
						in_scope_column_query_parameters,
						in_scope_column[1].(string)) //Add column name to column query

				fmt.Printf(
					"-------------------loading data for column : %s\n",
					in_scope_column[1])

				in_scope_column_dataset_rows :=
					database.ReadMsAccessColumns(
						accessDatabase,
						in_scope_table[1].(string),       //in scope table name
						in_scope_column_query_parameters) //fields: row uuid + in scope column

				in_scope_column_dataset_slice :=
					utils.Convert_rows_to_2d_slices(
						in_scope_column_dataset_rows)

				fmt.Printf(
					"\n+++Total number of rows loaded for column: %s+++\n",
					len(in_scope_column_dataset_slice))

				//add column uuids back to cell data
				//#TODO explore if there is a more elegant way of doing this (e.g. use dataframe append) - or create helper function (utilities)

				for _, in_scope_column_row := range in_scope_column_dataset_slice {

					in_scope_column_row_with_column_uuids :=
						append(
							in_scope_column_row,
							in_scope_column[0]) // adding column uuid here.

					in_scope_column_dataset =
						append(
							in_scope_column_dataset,
							in_scope_column_row_with_column_uuids)
				}

			}

		}

	}

	fmt.Printf( //#TODO add to logger
		"\n+++Starting Syntactic Checking on %s cells +++\n sample row : %s",
		len(in_scope_column_dataset),
		in_scope_column_dataset[1])

	// #TODO move to a separate function - not part of loading data scope of responsibility
	transaction_dataset := Process_column_sets(
		in_scope_column_dataset, //rowguid, cell value, column uuid
		checks,
		in_scope_column_check_configurations)

	transaction_dataset_string :=
		utils.Change_2d_interface_slice_to_string(
			transaction_dataset)

	//#TODO write to a database
	transactions_header := [][]string{{
		"check_uuids",
		"original_cell_values",
		"marked_cell_values",
		"modified_cell_values",
		"check_type_uuids",
		"row_uuids",
		"column_uuids"}}

	output_csv, _ := storage_csv.Open_csv_file("sytantic_check_transactions.csv")
	storage_csv.Write_2d_slice_set_to_csv(transactions_header, output_csv)
	storage_csv.Write_2d_slice_set_to_csv(transaction_dataset_string, output_csv)

}

func Process_column_sets(
	in_scope_columns_dataset [][]interface{},
	checks [][]interface{},
	in_scope_column_check_configurations [][]interface{}) [][]interface{} {

	var in_scope_column_cell_set_including_check_configuration [][]interface{} //#TODO explore alternative data structures for this.  possibly split this into column cell set and check configuration
	var in_scope_column_cell_and_check_configuration []interface{}             //#TODO explore alternative data structures for this.  possibly split this into column cell set and check configuration

	var transaction_rowset [][]interface{}

	for _, in_scope_column_row := range // for each row in the columns dataset
	in_scope_columns_dataset {

		for _, in_scope_column_check_configuration := range // for each entry in check configurations
		in_scope_column_check_configurations {

			if in_scope_column_check_configuration[0] == // check if column uuid in check configuration
				in_scope_column_row[2] { // matches column uuid in column dataset -- add to cell check dataset

				in_scope_column_cell_and_check_configuration := // create cell data for checks (strip column uuid and add check uuid) #TODO - chedck if this is the best way to create the slide or alternatives available
					append(
						in_scope_column_cell_and_check_configuration, // nil
						in_scope_column_row[0],                       // row uuid
						in_scope_column_row[1],                       // cell value
						in_scope_column_check_configuration[1])       // check uuid

				in_scope_column_cell_set_including_check_configuration = // add to cell configuration set (contains a row for each check for the cell value)
					append(
						in_scope_column_cell_set_including_check_configuration,
						in_scope_column_cell_and_check_configuration)

				in_scope_column_cell_and_check_configuration =
					nil
			}

		}

		if in_scope_column_cell_set_including_check_configuration != nil { //if the column is marked for checks

			transaction_column_rowset :=
				syntactic_check_cell_orchestrator.Process_cells_checks( // run check process
					in_scope_column_cell_set_including_check_configuration,
					checks)

			for _, transaction_row := range transaction_column_rowset { // for each returned check transaction, append the column uuid

				transaction_row = append(
					transaction_row,        // check transaction
					in_scope_column_row[2]) // column uuid

				transaction_rowset = append( // append the transaction (+ column uuid) to transaction set.
					transaction_rowset,
					transaction_row)

				transaction_row =
					nil
			}
			in_scope_column_cell_set_including_check_configuration =
				nil
		}

	}

	fmt.Printf( //#TODO add to logger
		"\n--Total %s transactions generated.\n--------------\nSample Transaction : %s\n-----------\n",
		len(transaction_rowset),
		transaction_rowset[0])

	return transaction_rowset

}
