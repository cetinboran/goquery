package goquery

import (
	"errors"
	"fmt"
	"log"
)

type GoQuery struct {
	TableName    string
	ColmnNames   []string
	Values       []interface{}
	Checks       []bool
	UniqueString string
	UniqueValue  interface{}
}

func GoQueryInit(tableName string) *GoQuery {
	return &GoQuery{TableName: tableName}
}

func (q *GoQuery) SetColmnNames(names []string) {
	q.ColmnNames = names
}

func (q *GoQuery) SetValues(values []interface{}) {
	q.Values = values
}

func (q *GoQuery) SetChecks(checks []bool) {
	q.Checks = checks
}

func (q *GoQuery) SetUniqueString(unique string) {
	q.UniqueString = unique
}

func (q *GoQuery) SetUniqueValue(unique interface{}) {
	q.UniqueValue = unique
}

func (q *GoQuery) SetAll(names []string, values []interface{}, checks []bool) {
	q.ColmnNames = names
	q.Values = values
	q.Checks = checks
}

func (q *GoQuery) Check() error {
	if len(q.Checks) != len(q.Values) && len(q.Values) != len(q.ColmnNames) && len(q.Checks) != len(q.ColmnNames) {
		return errors.New("Please enter same length arrays")
	}

	return nil
}

func (q *GoQuery) CreateUpdate(safe bool) (query string, args []interface{}) {
	err := q.Check()
	if err != nil {
		log.Fatal(err)
	}

	query = fmt.Sprintf("UPDATE %v SET ", q.TableName)

	for i, v := range q.ColmnNames {
		if q.Checks[i] {
			if safe {
				query += fmt.Sprintf("%v = ?, ", v)
				args = append(args, q.Values[i])

			} else {
				query += fmt.Sprintf("%v = %v, ", v, q.Values[i])
			}
		}
	}

	if query[len(query)-2:] == ", " {
		query = query[:len(query)-2]
	}

	if safe {
		query += fmt.Sprintf(" WHERE %v = ?", q.UniqueString)

	} else {
		query += fmt.Sprintf(" WHERE %v = %v", q.UniqueString, q.UniqueValue)
	}

	return
}
