package query

import (
	"fmt"
	"strings"
)

func StringToArray(s string) []string {
	split := strings.Split(strings.TrimPrefix(strings.TrimSuffix(s, "}"), "{"), ",")
	ret := make([]string, 0)
	for _, x := range split {
		y := strings.TrimSpace(x)
		if len(y) > 0 {
			ret = append(ret, y)
		}
	}
	return ret
}

func SQLSelect(columns string, tables string, where string, orderBy string, limit int, offset int) string {
	if len(columns) == 0 {
		columns = "*"
	}
	whereClause := ""
	if len(where) > 0 {
		whereClause = " where " + where
	}
	orderByClause := ""
	if len(orderBy) > 0 {
		orderByClause = " order by " + orderBy
	}
	limitClause := ""
	if limit > 0 {
		limitClause = fmt.Sprintf(" limit %v", limit)
	}
	offsetClause := ""
	if offset > 0 {
		offsetClause = fmt.Sprintf(" offset %v", offset)
	}
	return "select " + columns + " from " + tables + whereClause + orderByClause + limitClause + offsetClause
}
