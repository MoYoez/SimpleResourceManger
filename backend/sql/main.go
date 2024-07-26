package sql

import (
	"fmt"
	"sync"

	"reflect"
	"strings"
	"time"
	"unicode"

	sql "github.com/FloatTech/sqlite"
	"github.com/songzhibin97/gkit/sys/mutex"
)

// USE : SQLITE HERE >>
// jsonize: make json into string when saving to database.

type SettedNewdatabase struct {
	sql.Sqlite
}

var (
	SqlDataBase = sql.Sqlite{}
	SqlLocker   = mutex.Mutex{}
)

func init() {
	SqlDataBase.DBPath = "./gv.db"
	SqlDataBase.Open(time.Hour * 24)
	// set init.
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		InitDataBase(UserInfoDataStruct{})
		InitDataBase(ResBaseStruct{})
		InitDataBase(MaintainTableStruct{})
		InitDataBase(ReportBaseStruct{})
		defer wg.Done()

	}()
	wg.Wait()
	// ADD EXAMPLE DATA Insertion.
	ExampleLoggerinsert()
}

// ReturnDatabaseTableName ==> This method is to get correct table name and make adjust.
func ReturnDatabaseTableName(data interface{}) string {
	// Get the type of the data
	t := reflect.TypeOf(data)

	// If data is a pointer, get the element type
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// If it's not a struct, return an error
	if t.Kind() != reflect.Struct {
		return ""
	}

	// Get the name of the struct
	structName := t.Name()
	tableName := strings.ReplaceAll(structName, "Struct", "")
	result := make([]rune, 0, len(tableName)*2)
	for i, r := range tableName {
		if i > 0 && unicode.IsUpper(r) && unicode.IsUpper(rune(tableName[i-1])) {
			result = append(result, '_')
		}
		result = append(result, r)
	}
	tableName = strings.ToLower(string(result))
	return tableName
}

// InitDataBase PUBLIC METHOD: INIT DATA HERE >> transform data with
func InitDataBase(data interface{}) error {
	// Get the type of the data
	t := reflect.TypeOf(data)

	// If data is a pointer, get the element type
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// If it's not a struct, return an error
	if t.Kind() != reflect.Struct {
		return fmt.Errorf("data must be a struct or a pointer to a struct")
	}

	// Get the name of the struct
	structName := t.Name()
	tableName := strings.ReplaceAll(structName, "Struct", "")
	result := make([]rune, 0, len(tableName)*2)
	for i, r := range tableName {
		if i > 0 && unicode.IsUpper(r) && unicode.IsUpper(rune(tableName[i-1])) {
			result = append(result, '_')
		}
		result = append(result, r)
	}
	tableName = strings.ToLower(string(result))

	// Create a pointer to the data and pass it to SqlDataBase.Create
	ptr := reflect.New(t).Interface()
	return SqlDataBase.Create(tableName, ptr)
}
