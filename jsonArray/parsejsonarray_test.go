package main

import (
	"reflect"
	"testing"
)

func Test_contains(t *testing.T) {
	type args struct {
		s      []string
		substr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.s, tt.args.substr); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_remove(t *testing.T) {
	type args struct {
		s []string
		r string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := remove(tt.args.s, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("remove() = %v, want %v", got, tt.want)
			}
		})
	}
}
