package migrate

import (
	"io"
	"strings"
)

// QueriesFrom will get the queries from the io.Reader..
// We are not worried about memory here.
// Returns an error if we failed to red from the source.
//
// I was thinking we might verify that the strings were valid sql queries.
// But unsure if that is reasonable to do for all RDBMS.
// e.g.  https://github.com/xwb1989/sqlparser
func QueriesFrom(r io.Reader) ([]string, error) {
	read, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	input := string(read)
	return strings.Split(input, ";"), nil
}
