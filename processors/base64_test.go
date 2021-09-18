package processors

import "testing"

func TestBase64Decode(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "String",
			args: args{data: "dGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIGEgbGF6eSBkb2c="},
			want: "the quick brown fox jumps over a lazy dog",
		}, {
			name: "Emoji",
			args: args{data: "8J+Yg/CfmIfwn5mD8J+ZgvCfmInwn5iM8J+YmfCfmJfwn4eu8J+Hsw=="},
			want: "😃😇🙃🙂😉😌😙😗🇮🇳",
		}, {
			name: "Multi line string",
			args: args{data: "MTIzMzQ1CmFiY2QKNDU2CjEyMwphYmMKNTY3Cjc4OTA="},
			want: "123345\nabcd\n456\n123\nabc\n567\n7890",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base64Decode(tt.args.data); got != tt.want {
				t.Errorf("Base64Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase64Encode(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "String",
			args: args{data: "the quick brown fox jumps over a lazy dog"},
			want: "dGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIGEgbGF6eSBkb2c=",
		}, {
			name: "Emoji",
			args: args{data: "😃😇🙃🙂😉😌😙😗🇮🇳"},
			want: "8J+Yg/CfmIfwn5mD8J+ZgvCfmInwn5iM8J+YmfCfmJfwn4eu8J+Hsw==",
		}, {
			name: "Multi line string",
			args: args{data: "123345\nabcd\n456\n123\nabc\n567\n7890"},
			want: "MTIzMzQ1CmFiY2QKNDU2CjEyMwphYmMKNTY3Cjc4OTA=",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base64Encode(tt.args.data); got != tt.want {
				t.Errorf("Base64Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
