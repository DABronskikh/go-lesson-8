package transactions

import (
	"reflect"
	"testing"
)

func TestMapRowToTransaction(t *testing.T) {
	record := []string{"x", "001", "002", "100000", "1593746950"}
	tr := &Transaction{
		Id:      record[0],
		From:    record[1],
		To:      record[2],
		Amount:  1000_00,
		Created: 1593746950,
	}

	type args struct {
		records []string
	}
	tests := []struct {
		name    string
		args    args
		want    *Transaction
		wantErr bool
	}{
		{
			name: "func MapRowToTransaction",
			args: args{
				records: record,
			},
			want:    tr,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := MapRowToTransaction(tt.args.records)
		if (err != nil) != tt.wantErr {
			t.Errorf("MapRowToTransaction() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("MapRowToTransaction() got = %v, want %v", got, tt.want)
		}
	}
}
