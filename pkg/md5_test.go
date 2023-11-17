package pkg

import "testing"

func TestGetMD5(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "case-1",
			args: args{
				filePath: "./main.go",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMD5(tt.args.filePath); false {
				t.Errorf("GetMD5() = %v, want %v", got, tt.want)
			}
		})
	}
}
