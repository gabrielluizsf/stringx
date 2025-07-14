package stringx

import "log"

// Log is default printer for stringx.
var Log = log.Println

func fail(message string, s String, printer func(v ...any)) String {
	printer(message)
	return s
}
