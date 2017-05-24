package strutil

import (
	"database/sql"
	"fmt"
)

func NullInt64ToS(f interface{}) string {
	return NullInt64ToSF(f, "%d")
}

func NullInt64ToSF(f interface{}, format string) string {
	return convToS(f, format)
}

func NullFloat64ToS(f interface{}) string {
	return NullFloat64ToSF(f, "%f")
}

func NullFloat64ToSF(f interface{}, format string) string {
	return convToS(f, format)
}

func convToS(f interface{}, opts string) string {
	if f != nil {
		switch f.(type) {
		case sql.NullFloat64:
			if f.(sql.NullFloat64).Valid {
				return fmt.Sprintf(opts, f.(sql.NullFloat64).Float64)
			}
			break
		
		case sql.NullInt64:
			if f.(sql.NullInt64).Valid {
				return fmt.Sprintf(opts, f.(sql.NullInt64).Int64)
			}
			break
		}
	}
	
	return "null"
}
