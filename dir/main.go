package main

import (
	"cloud.google.com/go/spanner"
)

func Foo() *spanner.Statement {
	return &spanner.Statement{
		SQL:    "SELECT * FROM TABLE;",
		Params: map[string]interface{}{},
	}
}

func Bar() *spanner.Statement {
	return &spanner.Statement{
		SQL:    "SELECT * FROM TABLE WHERE 1=1;",
		Params: map[string]interface{}{},
	}
}
