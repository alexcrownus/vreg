package vreg

import "testing"

func TestQuery(t *testing.T) {
	Query("JJJ895AX")
}

func TestQueryInvalid(t *testing.T) {
	Query("JJJ895AX2")
}
