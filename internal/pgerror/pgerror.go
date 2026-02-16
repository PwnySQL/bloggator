package pgerror

import (
	"github.com/lib/pq"
)

// UniqueViolation checks if the error is of code 23505
func UniqueViolation(err error) *pq.Error {
	if pqerr, ok := err.(*pq.Error); ok &&
		pqerr.Code == "23505" {
		return pqerr
	}
	return nil
}
