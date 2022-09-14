package utils

import (
	"testing"
)

func Test_checkUserName(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Empty",
			args:    args{username: ""},
			wantErr: true,
		},
		{
			name:    "Short",
			args:    args{username: "ab"},
			wantErr: true,
		},
		{
			name:    "Too long",
			args:    args{username: "johnsjohnsjohnsjohnsjohnsjohnsjohnsjohnsjohnsjohnsjohnsjohnsjohnsjohnsjohnsjohnsjohnsjohnsjohnsjohns1"},
			wantErr: true,
		},
		{
			name:    "Special char",
			args:    args{username: "&é'!'ç(("},
			wantErr: true,
		},
		{
			name:    "Min lenght",
			args:    args{username: "jon"},
			wantErr: false,
		},
		{
			name:    "chars with digits with underscore",
			args:    args{username: "john_doe1"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkUserName(tt.args.username); (err != nil) != tt.wantErr {
				t.Errorf("checkUsername() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_CheckPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Empty",
			args:    args{password: ""},
			wantErr: true,
		},
		{
			name:    "Too short",
			args:    args{password: "pass"},
			wantErr: true,
		},
		{
			name:    "Too long",
			args:    args{password: "passwordpasswordpasswordpasswordpasswordpasswordpassword"},
			wantErr: true,
		},
		{
			name:    "Min lenght",
			args:    args{password: "passwo"},
			wantErr: false,
		},
		{
			name:    "max lenght",
			args:    args{password: "passwordpasswordpasswordpasswor"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckPassword(tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("checkUsername() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_checkPinCode(t *testing.T) {
	type args struct {
		pincode string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Empty",
			args:    args{pincode: ""},
			wantErr: true,
		},
		{
			name:    "Too short",
			args:    args{pincode: "123"},
			wantErr: true,
		},
		{
			name:    "Too long",
			args:    args{pincode: "123456789"},
			wantErr: true,
		},
		{
			name:    "Characters",
			args:    args{pincode: "pincode"},
			wantErr: true,
		},
		{
			name:    "Exact lenght",
			args:    args{pincode: "123456"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkPinCode(tt.args.pincode); (err != nil) != tt.wantErr {
				t.Errorf("checkUsername() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Benchmark_checkUserName(b *testing.B) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Empty",
			args:    args{username: ""},
			wantErr: true,
		},
		{
			name:    "Short",
			args:    args{username: "ab"},
			wantErr: true,
		},
		{
			name:    "Too long",
			args:    args{username: "johnsjohnsjohnsjohnsjohnsjohnsjohnsjohnsjohnsjohnsjohnsjohnsjohnsjohnsjohnsjohnsjohnsjohnsjohnsjohns1"},
			wantErr: true,
		},
		{
			name:    "Special char",
			args:    args{username: "&é'!'ç(("},
			wantErr: true,
		},
		{
			name:    "Min lenght",
			args:    args{username: "jon"},
			wantErr: false,
		},
		{
			name:    "chars with digits with underscore",
			args:    args{username: "john_doe1"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		for i := 0; i < b.N; i++ {
			checkUserName(tt.args.username)
		}
	}
}

func Benchmark_CheckPassword(b *testing.B) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Empty",
			args:    args{password: ""},
			wantErr: true,
		},
		{
			name:    "Too short",
			args:    args{password: "pass"},
			wantErr: true,
		},
		{
			name:    "Too long",
			args:    args{password: "passwordpasswordpasswordpasswordpasswordpasswordpassword"},
			wantErr: true,
		},
		{
			name:    "Min lenght",
			args:    args{password: "passwo"},
			wantErr: false,
		},
		{
			name:    "max lenght",
			args:    args{password: "passwordpasswordpasswordpasswor"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		for i := 0; i < b.N; i++ {
			CheckPassword(tt.args.password)
		}
	}
}

func Benchmark_checkPinCode(b *testing.B) {
	type args struct {
		pincode string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Empty",
			args:    args{pincode: ""},
			wantErr: true,
		},
		{
			name:    "Too short",
			args:    args{pincode: "123"},
			wantErr: true,
		},
		{
			name:    "Too long",
			args:    args{pincode: "123456789"},
			wantErr: true,
		},
		{
			name:    "Characters",
			args:    args{pincode: "pincode"},
			wantErr: true,
		},
		{
			name:    "Exact lenght",
			args:    args{pincode: "123456"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		for i := 0; i < b.N; i++ {
			checkPinCode(tt.args.pincode)
		}
	}
}
