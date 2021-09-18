package utils

import "fmt"

var (
	NoRowsFound   = fmt.Errorf("no rows found")
	NorowsUpdated = fmt.Errorf("no rows updated")
)
