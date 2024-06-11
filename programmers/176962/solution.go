package programmers

import (
	"sort"
	"strconv"
	"strings"
)

type planInfo struct {
	name     string
	start    int
	playtime int
}

func solution(plans [][]string) []string {
	result := make([]string, 0, len(plans))
	taskQueue := []*planInfo{}
	t := 0

	planStructs := make([]*planInfo, 0, len(plans))
	for _, p := range plans {
		playtime, _ := strconv.Atoi(p[2])
		plan := &planInfo{
			name:     p[0],
			start:    startStrToInt(p[1]),
			playtime: playtime,
		}
		planStructs = append(planStructs, plan)
	}

	sort.Slice(planStructs, func(i, j int) bool {
		return planStructs[i].start < planStructs[j].start
	})

	for _, now := range planStructs {
		for len(taskQueue) > 0 && t+taskQueue[0].playtime <= now.start {
			t += taskQueue[0].playtime
			result = append(result, taskQueue[0].name)
			taskQueue = taskQueue[1:]
		}
		if len(taskQueue) == 0 {
			taskQueue = append(taskQueue, now)
			t = now.start
		} else if t+taskQueue[0].playtime > now.start {
			taskQueue[0].playtime -= (now.start - t)
			taskQueue = append([]*planInfo{now}, taskQueue...)
			t = now.start
		}
	}

	for len(taskQueue) > 0 {
		t += taskQueue[0].playtime
		result = append(result, taskQueue[0].name)
		taskQueue = taskQueue[1:]
	}
	return result
}

func startStrToInt(startStr string) int {
	splited := strings.Split(startStr, ":")
	hh, _ := strconv.Atoi(splited[0])
	mm, _ := strconv.Atoi(splited[1])
	return hh*60 + mm
}
