package translator

import "testing"

func TestInterpreter_Translate(t *testing.T) {
	type args struct {
		text   string
		source string
		target string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
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
