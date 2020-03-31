package lib

import "time"

//LogTimeFormat formats log time
func LogTimeFormat() time.Time {

	return time.Now().Add(time.Duration(float64(time.Second/time.Millisecond)) * time.Millisecond)

}
