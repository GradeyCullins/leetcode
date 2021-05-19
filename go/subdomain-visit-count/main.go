package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := []string{
		"9001 discuss.leetcode.com",
		"50 yahoo.com",
	}
	res := subdomainVisits(input)
	fmt.Println(res)
}

// Define map of domain part e.g. "com", "google" => visit count
// For each domain-count string s do
// 	split s into count = int(s[0]), domain = s[1]
// 	parts = strings.Split(d, ".")
// 	for i = 0; i < len(parts); i++ do
// 	 part = parts[i:]
// 	 if p, ok := resMap[part] do
// 	  p.Count += c
// 	 else do
// 	  resMap[part] = c
// convert map to string val and return

func subdomainVisits(cpdomains []string) []string {
	res := make([]string, 0)
	resMap := make(map[string]int)
	for _, s := range cpdomains {
		count, _ := strconv.Atoi(strings.Split(s, " ")[0])
		domain := strings.Split(s, " ")[1]
		parts := strings.Split(domain, ".")

		for i := 0; i < len(parts); i++ {
			part := strings.Join(parts[i:], ".")
			if _, ok := resMap[part]; ok {
				resMap[part] += count
			} else {
				resMap[part] = count
			}
		}

	}

	for key, val := range resMap {
		countString := strconv.Itoa(val)
		res = append(res, fmt.Sprintf("%s %s", countString, key))
	}
	return res
}
