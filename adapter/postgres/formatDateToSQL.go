package repository

import "time"


// Cast the Date/Time to Interface{} so it can be nilable if date must be null
func formatDateToSQL(data time.Time) interface{} {
	var newData interface{}
	newData = data.Format(time.UnixDate)
	if data.IsZero() {
		newData = nil
	}	
	return newData
}
