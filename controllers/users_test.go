package controllers

import "testing"

// func TestListUsers(t *testing.T) {
// 	port := 10000
// 	host := "virmach-go.edtechstar.com"
// 	type args struct {
// 		host string
// 		post int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		{"simple", args{host, port}},
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ListUsers(tt.args.host, tt.args.post)
// 		})
// 	}
// }

func TestListUsers(t *testing.T) {
	type args struct {
		host string
		port int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"simple", args{"virmach-go.edtechstar.com", 10000}, false},
		// {"simple", args{"127.0.0.1", 10000}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ListUsers(tt.args.host, tt.args.port); (err != nil) != tt.wantErr {
				t.Errorf("ListUsers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
