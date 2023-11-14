package gocfmx_test

import (
	"github.com/opannapo/gocfmx"
	"testing"
)

func TestEncrypt(t *testing.T) {
	type args struct {
		input string
		key   string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{
			name: "Success",
			args: args{
				input: "abcd",
				key:   "1234567891234567",
			},
			wantOutput: "6d05afba",
		},
		{
			name: "Success 2",
			args: args{
				input: "plain value",
				key:   "testKey1234$",
			},
			wantOutput: "4d51945a59cdbf92a28f9e",
		},
		{
			name: "Success 3",
			args: args{
				input: "testTestTestteESt",
				key:   "Re3trOoO123!",
			},
			wantOutput: "6d52ca231b4e0e936a89bcafc79ed2359b",
		},
		{
			name: "Success short key",
			args: args{
				input: "testTestTestteESt",
				key:   "Re3trO",
			},
			wantOutput: "4d808a8d88a84e834d39ae46bd1a23a64b",
		},
		{
			name: "Success short key 2",
			args: args{
				input: "Coba Cooba testTestTestteESt test TeESsst",
				key:   "abc",
			},
			wantOutput: "62b69b5a5efe99883198ec9b09bc031d8bdfd28d76a990c9ae786e4b478b59c0474ea7dbaaa59c2c10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfMx := gocfmx.NewCFmxCompat(tt.args.key)
			gotOutput := cfMx.EncryptEncodingHex(tt.args.input)
			if gotOutput != tt.wantOutput {
				t.Errorf("EncryptDecrypt() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	type args struct {
		input string
		key   string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{
			name: "Success",
			args: args{
				input: "6d05afba",
				key:   "1234567891234567",
			},
			wantOutput: "abcd",
		},
		{
			name: "Success 2",
			args: args{
				input: "42468ea0889c5b14b2581587d4809e4575b09fa8858d34939a4192da5cb6eb52aee7a2e65af60de38ff6b0",
				key:   "iniKey",
			},
			wantOutput: "opannapo-opannapo-test-test-o-p-a-n-n-a-p-o",
		},
		{
			name: "Success 3",
			args: args{
				input: "42468EA0889C5B14B2581587D4809E4575B09FA8858D34939A4192DA5CB6EB52AEE7A2E65AF60DE38FF6B0",
				key:   "iniKey",
			},
			wantOutput: "opannapo-opannapo-test-test-o-p-a-n-n-a-p-o",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfMx := gocfmx.NewCFmxCompat(tt.args.key)
			gotOutput, _ := cfMx.DecryptEncodingHex(tt.args.input)
			if gotOutput != tt.wantOutput {
				t.Errorf("EncryptDecrypt() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
