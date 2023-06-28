package main

import "log"

func main() {
	log.Println(isIsomorphic("ejg", "add"), false)
	log.Println(isIsomorphic("egg", "add"), true)
}

	// 1 2 2
	// 1 2 2
	// 1 2 1 3 4
	// 1 2 1 3 4
func isIsomorphic(s string, t string) bool {

	mapa1 := map[byte]int{}
	mapa2 := map[byte]int{}
    for i := range s {
		if _, ok := mapa1[s[i]]; !ok {
			mapa1[s[i]] = i
		}
		if _, ok := mapa2[t[i]]; !ok {
			mapa2[t[i]] = i
		}
		if mapa1[s[i]] != mapa2[t[i]] {
			return false
		}
	}
    return true
}