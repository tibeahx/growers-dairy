package tests

import (
	"net/http"
	"testing"
)

type args struct {
	name   string
	url    string
	method string
}

func Test_oleg(t *testing.T) {
	cases := []struct {
		testData args
		want     string
	}{
		{
			testData: args{
				name:   "xyu",
				url:    "https://asd.asd",
				method: http.MethodGet,
			},
		},
		{
			testData: args{
				name:   "xyu",
				url:    "https://asd.asd",
				method: http.MethodGet,
			},
		},
		{
			testData: args{
				name:   "xyu",
				url:    "https://asd.asd",
				method: http.MethodGet,
			},
		},
		{
			testData: args{
				name:   "xyu",
				url:    "https://asd.asd",
				method: http.MethodGet,
			},
		},
		{
			testData: args{
				name:   "xyu",
				url:    "https://asd.asd",
				method: http.MethodGet,
			},
		},
		{
			testData: args{
				name:   "xyu",
				url:    "https://asd.asd",
				method: http.MethodGet,
			},
		},
		{
			testData: args{
				name:   "xyu",
				url:    "https://asd.asd",
				method: http.MethodGet,
			},
		},
	}

	for _, test := range cases {
		name := test.args.Name

	}
}
