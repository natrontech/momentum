package momentumservices

import "github.com/pocketbase/dbx"

func ExprsEq(columnsToParams map[string]string) []dbx.Expression {

	expressions := make([]dbx.Expression, 0)
	for key, value := range columnsToParams {
		expressions = append(expressions, ExprEq(key, value))
	}

	return expressions
}

func ExprEq(column string, param string) dbx.Expression {

	return dbx.NewExp(column+" = {:"+column+"}", dbx.Params{column: param})
}

func ExprsIn(columnsToParams map[string]string) []dbx.Expression {

	expressions := make([]dbx.Expression, 0)
	for key, value := range columnsToParams {
		expressions = append(expressions, ExprIn(key, value))
	}

	return expressions
}

func ExprIn(column string, param string) dbx.Expression {

	return dbx.NewExp(column+" IN {:"+param+"}", dbx.Params{column: param})
}
