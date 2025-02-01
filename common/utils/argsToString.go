package utils

import "fmt"

// Преобразует []interface{} в []string (т.к. gRPC не поддерживает интерфейсы)
func ArgsToStringSlice(args []interface{}) []string {
	var strArgs []string
	for _, arg := range args {
		strArgs = append(strArgs, fmt.Sprintf("%v", arg))
	}
	return strArgs
}
