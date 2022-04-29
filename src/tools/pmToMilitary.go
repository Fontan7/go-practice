package tools

import (
	"time"
	"log" 
)

//Convert time from PM to military
func timeConversion(s string) string {
    
    dt,_ := time.Parse("03:04:05PM", s)
    military := dt.Format("15:04:05")
    
	log.Fatal()

    return military
}