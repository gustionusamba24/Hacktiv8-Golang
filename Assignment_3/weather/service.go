package weather

import (
	"math/rand"
	"time"
)

func randomInt(s, e int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(e-s) + s
}

func GetStatusValue() Status {
	var status Status
	dt := time.Now()

	date := dt.Format("2006-01-02 15:04:05")

	status.Water = randomInt(1, 100)
	time.Sleep(1)
	status.Wind = randomInt(1, 100)

	if status.Water > 5 {
		status.WaterStatus = "Aman"
	} else if status.Water >= 6 && status.Water <= 8 {
		status.WaterStatus = "Siaga"
	} else if status.Water > 8 {
		status.WaterStatus = "Bahaya"
	} else {
		status.WaterStatus = "Undefined"
	}

	if status.Wind < 6 {
		status.WindStatus = "Aman"
	} else if status.Wind >= 7 && status.Wind <= 15 {
		status.WindStatus = "Siaga"
	} else if status.Wind > 15 {
		status.WindStatus = "Bahaya"
	} else {
		status.WindStatus = "Error"
	}

	status.Date = date

	return status
}
