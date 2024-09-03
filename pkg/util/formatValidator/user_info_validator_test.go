// @Title user_info_validator_test.go
// @Description
// @Author Hunter 2024/9/3 17:36

package formatValidator

import (
	"testing"
)

func TestCheckAccountName(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case0",
			args: args{
				text: "hello123",
			},
			wantErr: false,
		},
		{
			name: "case1",
			args: args{
				text: "hello@123.world",
			},
			wantErr: false,
		},
		{
			name: "case2",
			args: args{
				text: "hello@123.w*or&ld--=",
			},
			wantErr: false,
		},
		{
			name: "case3",
			args: args{
				text: "hello@123.w*or&ld-你好-=",
			},
			wantErr: true,
		},
		{
			name: "case4",
			args: args{
				text: "hello@123.w*or&ld-你好-=世界",
			},
			wantErr: true,
		},
		{
			name: "case4",
			args: args{
				text: "x需求hello@123.w*or&ld-你好-=世界",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateAccountName(tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("CheckAccountName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckMobilePhoneNumber(t *testing.T) {
	type args struct {
		mobile string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case0",
			args: args{
				mobile: "13218879988",
			},
			wantErr: false,
		},
		{
			name: "case0",
			args: args{
				mobile: "1321887998",
			},
			wantErr: true,
		},
		{
			name: "case0",
			args: args{
				mobile: "10218879989",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateMobileNumber(tt.args.mobile); (err != nil) != tt.wantErr {
				t.Errorf("CheckMobile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckBlackIP(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case0",
			args: args{
				text: "192.168.31.10",
			},
			wantErr: false,
		},
		{
			name: "case1",
			args: args{
				text: "193.168.31.*",
			},
			wantErr: false,
		},
		{
			name: "case2",
			args: args{
				text: "192.168.269.*",
			},
			wantErr: true,
		},
		{
			name: "case3",
			args: args{
				text: "192.168.*.*",
			},
			wantErr: false,
		},
		{
			name: "case4",
			args: args{
				text: "*.168.*.*",
			},
			wantErr: false,
		},
		{
			name: "case5",
			args: args{
				text: "*.*.*.*",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateIPWithWildcard(tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("CheckBlackIP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case0",
			args: args{
				email: "hello@world.com",
			},
			wantErr: false,
		},
		{
			name: "case1",
			args: args{
				email: "hello@world.com.cn",
			},
			wantErr: false,
		},
		{
			name: "case2",
			args: args{
				email: "hello@worldcom",
			},
			wantErr: true,
		},
		{
			name: "case3",
			args: args{
				email: "helloworld.com",
			},
			wantErr: true,
		},
		{
			name: "case4",
			args: args{
				email: "你好@world.com",
			},
			wantErr: false,
		},
		{
			name: "case5",
			args: args{
				email: "你好@world。com",
			},
			wantErr: true,
		},
		{
			name: "case6",
			args: args{
				email: "你好@world.世界com",
			},
			wantErr: true,
		},
		{
			name: "case7",
			args: args{
				email: "你好@world..com",
			},
			wantErr: true,
		},
		{
			name: "case8",
			args: args{
				email: "你好@world\\.com",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateEmail(tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("CheckEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckIP(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case0",
			args: args{
				text: "192.168.50.10",
			},
			wantErr: false,
		},
		{
			name: "case1",
			args: args{
				text: "193.168.50.300",
			},
			wantErr: true,
		},
		{
			name: "case2",
			args: args{
				text: "192.168.50",
			},
			wantErr: true,
		},
		{
			name: "case3",
			args: args{
				text: "192.168.50.10.10",
			},
			wantErr: true,
		},
		{
			name: "case4",
			args: args{
				text: "192.168",
			},
			wantErr: true,
		},
		{
			name: "case5",
			args: args{
				text: "192",
			},
			wantErr: true,
		},
		{
			name: "case6",
			args: args{
				text: "192.168.10.h",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateIP(tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("CheckIP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
