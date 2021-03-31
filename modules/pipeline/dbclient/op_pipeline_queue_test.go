package dbclient

import (
	"reflect"
	"testing"
)

func Test_transferMustMatchLabelsToMap(t *testing.T) {
	type args struct {
		ss []string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string][]string
		wantErr bool
	}{
		{
			name: "three labels",
			args: args{
				ss: []string{"a=b", "a=c", "b=d"},
			},
			want: map[string][]string{
				"a": []string{"b", "c"},
				"b": []string{"d"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := transferMustMatchLabelsToMap(tt.args.ss)
			if (err != nil) != tt.wantErr {
				t.Errorf("transferMustMatchLabelsToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transferMustMatchLabelsToMap() got = %v, want %v", got, tt.want)
			}
		})
	}
}
