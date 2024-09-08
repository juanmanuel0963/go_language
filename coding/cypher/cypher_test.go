package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solution(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Test 1",
			args{s: "uno-square *es* %una gran empresa.", n: 3},
			"xqr-vtxduh *hv* %xqd judq hpsuhvd.",
		},
		{
			"Test 2",
			args{s: "uno-square *es* %una gran empresa.", n: -1},
			"shift can't be negative",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			responseChannel := make(chan string)
			errorChannel := make(chan error)

			go CaesarCipher(tt.args.s, tt.args.n, responseChannel, errorChannel)

			defer close(errorChannel)
			defer close(responseChannel)

			// Wait for the cipher to be created and sent through the channel
			select {
			case got := <-responseChannel:
				assert.Equal(t, got, tt.want)
			case error := <-errorChannel:
				assert.Equal(t, error.Error(), tt.want)
			}
		})
	}
}

func Test_solution_1(t *testing.T) {

	responseChannel := make(chan string)
	errorChannel := make(chan error)

	go CaesarCipher("uno-square *es* %una gran empresa.", 3, responseChannel, errorChannel)

	defer close(errorChannel)
	defer close(responseChannel)

	// Wait for the cipher to be created and sent through the channel
	select {
	case got := <-responseChannel:
		assert.Equal(t, got, "xqr-vtxduh *hv* %xqd judq hpsuhvd.")
	case error := <-errorChannel:
		assert.Nil(t, error)
	}

}

func Test_solution_2(t *testing.T) {

	responseChannel := make(chan string)
	errorChannel := make(chan error)

	go CaesarCipher("uno-square *es* %una gran empresa.", -1, responseChannel, errorChannel)

	defer close(errorChannel)
	defer close(responseChannel)

	// Wait for the cipher to be created and sent through the channel
	select {
	case got := <-responseChannel:
		assert.Equal(t, got, "")
	case error := <-errorChannel:
		assert.Equal(t, error.Error(), "shift can't be negative")
	}
}
