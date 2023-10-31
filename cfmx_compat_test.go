package gocfmx_test

import (
	"gocfmx"
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
			wantOutput: "6D05AFBA",
		},
		{
			name: "Success 2",
			args: args{
				input: "plain value",
				key:   "testKey1234$",
			},
			wantOutput: "4D51945A59CDBF92A28F9E",
		},
		{
			name: "Success 3",
			args: args{
				input: "testTestTestteESt",
				key:   "Re3trOoO123!",
			},
			wantOutput: "6D52CA231B4E0E936A89BCAFC79ED2359B",
		},
		{
			name: "Success short key",
			args: args{
				input: "testTestTestteESt",
				key:   "Re3trO",
			},
			wantOutput: "4D808A8D88A84E834D39AE46BD1A23A64B",
		},
		{
			name: "Success short key 2",
			args: args{
				input: "Coba Cooba testTestTestteESt test TeESsst",
				key:   "abc",
			},
			wantOutput: "62B69B5A5EFE99883198EC9B09BC031D8BDFD28D76A990C9AE786E4B478B59C0474EA7DBAAA59C2C10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfMx := gocfmx.NewCFmxCompat(tt.args.key)
			gotOutput := cfMx.Encrypt(tt.args.input)
			if gotOutput != tt.wantOutput {
				t.Errorf("EncryptDecrypt() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
