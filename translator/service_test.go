package translator

import "testing"

func TestInterpreter_Translate(t *testing.T) {
	type args struct {
		text   string
		source string
		target string
	}
	testArgs := args{
		text:   "A",
		source: "TEXT",
		target: "BINARY",
	}
	testArgs2 := args{
		text:   "A",
		source: "TTTTT",
		target: "BINARY",
	}
	testArgs3 := args{
		text:   "A",
		source: "TEXT",
		target: "MORSE",
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test Text to Binary",
			args: testArgs,
			want: "01000001",
			wantErr: false,
		},
		{
			name: "Test Text to Morse",
			args: testArgs3,
			want: ".-",
			wantErr: false,
		},
		{
			name: "Test Text to Binary failed",
			args: testArgs2,
			want: "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := Interpreter{}
			got, err := in.Translate(tt.args.text, tt.args.source, tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("Translate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Translate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
