package store

import (
	"fmt"
	"testing"
	// "github.com/stretchr/testify/require"
)

func TestGetStore(t *testing.T) {
	type args struct {
		filename string
	}
	var tests = []struct {
		name string
		args args
		want error
	}{
		{"good", args{"../../driver_positions.csv"}, nil},
	}
	for _, tcase := range tests {

		_, err := GetStore(tcase.args.filename)
		if err != nil {
			t.Fatalf("%v: %v", tcase.name, err)
		}
	}
}

func TestGetStoreBad(t *testing.T) {
	type args struct {
		filename string
	}
	var tests = []struct {
		name string
		args args
		want error
	}{
		{"bad file path", args{"../ads/driver_positions.csv"}, fmt.Errorf("bad filename or filename isnt")},
		{"bad file consist", args{"../../driver_positions2.csv"}, fmt.Errorf("bad file for line (%v:) %v", 2, "")},
	}
	for _, tcase := range tests {

		_, err := GetStore(tcase.args.filename)
		if err == nil {
			t.Fatalf("%v", tcase.name)
		}
	}
}
