package pointifyer

import (
	"log"
	"reflect"
)

// Gets all the fields from a Struct and returns a Slice With pointers to all the fields
// Usefull for Example for Scan() pgx
func Pointify(object any) ([]any, error) {
	val := reflect.ValueOf(object).Elem()
	log.Println("0")
	if val.Kind() != reflect.Struct {
		log.Fatal("unexpected type")
	}
	log.Println("1")
	numCols := val.NumField()
	log.Println("2")

	columns := make([]any, numCols)
	log.Println("3")

	for i := 0; i < numCols; i++ {
		field := val.Field(i)
		columns[i] = field.Addr().Interface()
	}
	log.Println("4")

	return columns, nil
}
