package goquery

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
)

type GoQuery struct {
	TableName    string
	Struct       interface{}
	Checks       []bool
	UniqueString string
	UniqueValue  interface{}
	structLen    int
}

func GoQueryInit(tableName string) *GoQuery {
	return &GoQuery{TableName: tableName}
}

func (q *GoQuery) SetStruct(s interface{}) {
	q.Struct = s
}

func (q *GoQuery) SetChecks(checks []bool) {
	q.Checks = checks
}

func (q *GoQuery) SetUnique(unique string, uniqueV interface{}) {
	q.UniqueString = unique
	q.UniqueValue = uniqueV
}

func (g *GoQuery) check() error {
	if len(g.Checks) != g.structLen {
		return errors.New("Check Length has to be same as column length.")
	}

	return nil
}

func (q *GoQuery) Take() (columnNames []string, values []interface{}) {
	v := reflect.ValueOf(q.Struct)
	t := v.Type()
	q.structLen = t.NumField()

	for i := 0; i < q.structLen; i++ {
		field := t.Field(i)
		tagName := field.Tag.Get("column")
		if tagName != "" {
			fieldValue := v.Field(i).Interface()
			columnNames = append(columnNames, tagName)
			values = append(values, fieldValue)
		}
	}

	return
}

func (q *GoQuery) CreateUpdate(safe bool) (query string, args []interface{}) {
	columnNames, values := q.Take()

	err := q.check()
	if err != nil {
		log.Fatal(err)
	}

	query = fmt.Sprintf("UPDATE %v SET ", q.TableName)

	for i, v := range columnNames {
		if q.Checks[i] {
			if safe {
				query += fmt.Sprintf("%v = ?, ", v)
				args = append(args, values[i])

			} else {
				query += fmt.Sprintf("%v = %v, ", v, values[i])
			}
		}
	}

	if query[len(query)-2:] == ", " {
		query = query[:len(query)-2]
	}

	if safe {
		query += fmt.Sprintf(" WHERE %v = ?", q.UniqueString)
		args = append(args, q.UniqueValue)
	} else {
		query += fmt.Sprintf(" WHERE %v = %v", q.UniqueString, q.UniqueValue)
	}

	return
}

func (q *GoQuery) CreateInsert(safe bool) (query string, args []interface{}) {
	columnNames, values := q.Take()

	query = fmt.Sprintf("INSERT INTO %v (", q.TableName)

	for i, v := range columnNames {
		if q.Checks[i] {
			if safe {
				query += fmt.Sprintf("%v,", v)
				args = append(args, values[i])

			} else {
				query += fmt.Sprintf("%v = %v, ", v, values[i])
			}
		}
	}

	index := strings.LastIndex(query, ",")
	if index != -1 {
		query = query[:index]
	}

	query += ")"

	if safe {
		query += fmt.Sprintf(" WHERE %v = ?", q.UniqueString)
		args = append(args, q.UniqueValue)
	} else {
		query += fmt.Sprintf(" WHERE %v = %v", q.UniqueString, q.UniqueValue)
	}

	return
}
