package programmers

import (
	"sort"
	"strconv"
	"strings"
)

func solution(fees []int, records []string) []int {
	defaultMin, defaultWon, unitMin, unitWon := fees[0], fees[1], fees[2], fees[3]
	carFee := make(map[string]*record)
	for _, r := range records {
		splited := strings.Split(r, " ")
		timeRecord, carNum, inout := splited[0], splited[1], splited[2]
		timeMin := toTimeMin(timeRecord)
		if inout == "IN" {
			if carFee[carNum] == nil {
				carFee[carNum] = &record{car: carNum}
			}
			carFee[carNum].in = timeMin
		} else {
			record := carFee[carNum]
			record.total += timeMin - record.in
			record.in = -1
		}
	}

	keys := make([]string, 0, len(carFee))
	for k := range carFee {
		keys = append(keys, k)
		if carFee[k].in != -1 {
			carFee[k].total += toTimeMin("23:59") - carFee[k].in
		}
		carFee[k].calculateFee(defaultMin, defaultWon, unitMin, unitWon)
	}
	sort.Strings(keys)

	result := make([]int, len(carFee))
	for i := range keys {
		result[i] = carFee[keys[i]].fee
	}
	return result
}

type record struct {
	car   string
	in    int
	total int
	fee   int
}

func (r *record) calculateFee(defaultMin, defaultWon, unitMin, unitWon int) {
	if r.total <= defaultMin {
		r.fee = defaultWon
	} else {
		overMin := r.total - defaultMin
		realUnitMin := overMin / unitMin
		if overMin%unitMin > 0 {
			realUnitMin += 1
		}
		r.fee = defaultWon + realUnitMin*unitWon
	}
}

func toTimeMin(timeRecord string) int {
	timeStr := strings.Split(timeRecord, ":")
	hh, mm := timeStr[0], timeStr[1]
	hour, _ := strconv.Atoi(hh)
	min, _ := strconv.Atoi(mm)
	timeMin := hour*60 + min
	return timeMin
}
