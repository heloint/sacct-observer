package main

import (
	"bytes"
	"database/sql"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type CliArgs struct {
	UpdateOnce      *bool
	UpdateFrequency *int
	Username        *string
	RemoteAddress   *string
	OutputSqliteDB  *string
}

func GetCliArgs() *CliArgs {
	var updateOnce *bool
	var updateFrequency *int
	var username *string
	var remoteAddress *string
	var outputSqliteDB *string

	updateOnce = flag.Bool("update-once", false, "Observe and update the database, instead of doing it periodically.")
	updateFrequency = flag.Int("update-frequency", 60, "Wait N number of seconds between fetches.")
	username = flag.String("username", "", "Username on the remote machine.")
	remoteAddress = flag.String("remote-address", "", "Address of the remote machine.")
	outputSqliteDB = flag.String("output-sqlite-db", "", "Path of the sqlite3 database.")

	flag.Parse()

	if *username == "" {
		panic(`Missing argument for "username".`)
	} else if *remoteAddress == "" {
		panic(`Missing argument for "remote-address".`)
	} else if *outputSqliteDB == "" {
		panic(`Missing argument for "output-sqlite-db".`)
	} else if *updateFrequency < 1 {
		panic(`Value for "update-frequency" cannot be less than 1 seconds.`)
	}

	return &CliArgs{
		UpdateOnce:      updateOnce,
		UpdateFrequency: updateFrequency,
		Username:        username,
		RemoteAddress:   remoteAddress,
		OutputSqliteDB:  outputSqliteDB,
	}
}

func readSACCTFromSSH(username string, remoteAddress string) *bytes.Buffer {
	var out bytes.Buffer
	formattedRemoteAddress := fmt.Sprintf("%s@%s", username, remoteAddress)
	cmd := exec.Command("ssh", formattedRemoteAddress, "sacct -XP --delimiter ',' ")
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
	return &out
}

func readCSV(out *bytes.Buffer) ([]string, *csv.Reader) {
	csvReader := csv.NewReader(out)
	headers, err := csvReader.Read()

	if err != nil {
		log.Fatal("Headers are not available in received CSV content.")
	}
	return headers, csvReader
}

func getDatabase(sqlitePath string) *sql.DB {
	db, err := sql.Open("sqlite3", sqlitePath)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func getColumnDefinitions(headers *[]string, indexColumn string) string {
	columnDefs := []string{}
	for _, s := range *headers {
		if s == indexColumn {
			columnDefs = append(columnDefs, fmt.Sprintf("%s TEXT PRIMARY KEY", s))
		} else {
			columnDefs = append(columnDefs, fmt.Sprintf("%s TEXT NOT NULL", s))
		}
	}
	return strings.Join(columnDefs, ",")
}

func createTableFromHeaders(db *sql.DB, tableName string, headers *[]string, indexColumn string) {
	createTableTemplate := "CREATE TABLE IF NOT EXISTS %s (%s);"
	columnDefs := getColumnDefinitions(headers, indexColumn)
	_, err := db.Exec(fmt.Sprintf(createTableTemplate, tableName, columnDefs))
	if err != nil {
		log.Fatal(err)
	}
}

func getValueStatementPlaceholders(values *[]string) string {
	placeholders := make([]string, len(*values))
	for i, str := range placeholders {
		placeholders[i] = str + "?"
	}
	return strings.Join(placeholders, ",")
}

func getInterfaceSlice[T any](convSlice []T) []interface{} {
	result := make([]interface{}, len(convSlice))
	for i := range convSlice {
		result[i] = convSlice[i]
	}
	return result
}

func insertOrReplaceFromCSV(db *sql.DB, tableName string, headers *[]string, csvContent *csv.Reader) {
	insertStatementTemplate := "INSERT OR REPLACE INTO %s (%s) VALUES(%s);"
	joinedHeaders := strings.Join(*headers, ",")
	joinedPlaceholders := getValueStatementPlaceholders(headers)
	filledInsertTemplate := fmt.Sprintf(insertStatementTemplate, tableName, joinedHeaders, joinedPlaceholders)

	transaction, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	insertStatement, err := transaction.Prepare(filledInsertTemplate)
	if err != nil {
		log.Fatal(err)
	}
	defer insertStatement.Close()

	for {
		record, err := csvContent.Read()
		if err != nil {
			break
		}
		args := getInterfaceSlice(record)
		_, err = insertStatement.Exec(args...)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = transaction.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func doDatabaseUpdate(out *bytes.Buffer, cliArgs *CliArgs, headers []string, csvContent *csv.Reader, db *sql.DB) {
	out = readSACCTFromSSH(*cliArgs.Username, *cliArgs.RemoteAddress)
	headers, csvContent = readCSV(out)

	createTableFromHeaders(db, "jobs", &headers, "JobID")
	insertOrReplaceFromCSV(db, "jobs", &headers, csvContent)
}

func main() {
	var out *bytes.Buffer
	var headers []string
	var csvContent *csv.Reader

	cliArgs := GetCliArgs()
	db := getDatabase(*cliArgs.OutputSqliteDB)
	defer db.Close()

	if *cliArgs.UpdateOnce {
		log.Println("Updating database...")
		doDatabaseUpdate(out, cliArgs, headers, csvContent, db)
	} else {
		for {
			log.Println("Updating database...")
			doDatabaseUpdate(out, cliArgs, headers, csvContent, db)
			time.Sleep(time.Duration(*cliArgs.UpdateFrequency) * time.Second)
		}
	}
}
