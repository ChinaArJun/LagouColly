package spider

import (
	"fmt"
	"go_jobs/downloader"
	"testing"
)

func TestGetJobs(t *testing.T) {
	type args struct {
		city string
		pn   int
		kd   string
	}
	tests := []struct {
		name    string
		args    args
		want    []downloader.Result
		want1   int
		want2   int
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "test", args: args{
			city: "北京",
			pn:   0,
			kd:   "golang",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := GetJobs(tt.args.city, tt.args.pn, tt.args.kd)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetJobs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got, got1, got2)
		})
	}
}
