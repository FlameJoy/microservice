package utils

import (
	"strconv"
	"time"
)

// Приводим аргументы к нужным типам (int, float, time, string)
func ParseArg(arg string) interface{} {
	// Пробуем привести к int
	if intVal, err := strconv.Atoi(arg); err == nil {
		return intVal
	}
	// Пробуем привести к float
	if floatVal, err := strconv.ParseFloat(arg, 64); err == nil {
		return floatVal
	}
	// Пробуем привести к времени
	if timeVal, err := time.Parse(time.RFC3339, arg); err == nil {
		return timeVal
	}
	// Если ничего не подошло, оставляем строку
	return arg
}
