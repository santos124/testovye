package main

import (
	"log"
	"strings"
)

func main() {
    log.Println(isValid("()[]{}"), true)
    log.Println(isValid("([]))"), false)
	log.Fatal(isValid("([])"), true)
}

func isValid(s string) bool {
    lSet := "({["
	rSet := ")}]"
	if len(s) > 0 && strings.Contains(rSet, string([]byte{s[0]})) {
		return false
	}
	pos := strings.Index(lSet, string([]byte{s[0]}))
	lB := lSet[pos]
	rB := rSet[pos]
	for i := 0; i< len(s); {
		if s[i] == rB {
			if i == len(s) -1 {
				return true
			} else {
				i++

				pos = strings.Index(lSet, string([]byte{s[i]}))
				if pos == -1 {
					return false
				}
				lB = lSet[pos]
				rB = rSet[pos]
			}
		} else if strings.Contains(lSet, string([]byte{s[i]})) {
			i = checkPar(i, s[i+1:], lB, rB, lSet, rSet)
			if i == -1 {
				return false
			}
		} else {
			return false
		}
	}
    return true

}

func checkPar(offset int, s string, lB, rB byte, lSet, rSet string) int {
	for i := 0; i < len(s); i++ {
		if s[i] == rB {
			return i+1 + offset
		} else if strings.Contains(lSet, string([]byte{s[i]})) {
			pos := strings.Index(lSet, string([]byte{s[i]}))
			i = checkPar(i, s[i+1:], lSet[pos], rSet[pos], lSet, rSet)
			if i == -1 {
				return -1
			}
		} else {
			return -1
		}
	}
	return -1
}