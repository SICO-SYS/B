/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

Contributors

*/

package controller

import (
	// "encoding/json"
	"golang.org/x/net/context"
	"testing"

	"github.com/SiCo-Ops/Pb"
)

func Test_PullRPC(t *testing.T) {
	test := &ConfigService{}
	in := &pb.ConfigPullCall{Id: "system", Environment: "production"}
	res, err := test.PullRPC(context.Background(), in)
	if err != nil {
		t.Error(err)
	}
	if res.Code != 0 {
		t.Error(res.Code)
	}
}

func Benchmark_PullRPC(b *testing.B) {
	test := &ConfigService{}
	in := &pb.ConfigPullCall{Id: "system", Environment: "production"}
	for i := 0; i < b.N; i++ {
		res, err := test.PullRPC(context.Background(), in)
		if err != nil {
			b.Error(err)
		}
		if res.Code != 0 {
			b.Error(res.Code)
		}
	}
}
