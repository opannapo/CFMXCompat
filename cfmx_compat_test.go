package GOCFMXCompat

import (
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfmx := NewCFmxCompat(tt.args.key)
			gotOutput := cfmx.Encrypt(tt.args.input)
			if gotOutput != tt.wantOutput {
				t.Errorf("EncryptDecrypt() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
