package GOCFMXCompat_test

import (
	GOCFMXCompat "gocfmx"
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
			name: "Success short key",
			args: args{
				input: "testTestTestteESt test TeESsst",
				key:   "abc",
			},
			wantOutput: "55BC8A4F2AD88593079CBF9B18AA321A9A8CD2BC60AEC4E9AE786E4C148B",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfMx := GOCFMXCompat.NewCFmxCompat(tt.args.key)
			gotOutput := cfMx.Encrypt(tt.args.input)
			if gotOutput != tt.wantOutput {
				t.Errorf("EncryptDecrypt() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
