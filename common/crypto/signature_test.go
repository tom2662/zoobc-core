package crypto

import (
	"database/sql"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/zoobc/zoobc-core/common/query"
)

func TestNewSignature(t *testing.T) {
	type args struct {
		executor query.ExecutorInterface
	}
	tests := []struct {
		name string
		args args
		want *Signature
	}{
		{
			name: "NewSignature:success",
			args: args{
				executor: nil,
			},
			want: &Signature{
				Executor: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSignature(tt.args.executor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSignature() = %v, want %v", got, tt.want)
			}
		})
	}
}

type mockSignatureSignExecutorSuccess struct {
	query.Executor
}

func (*mockSignatureSignExecutorSuccess) ExecuteSelect(qe string, args ...interface{}) (*sql.Rows, error) {
	db, mock, _ := sqlmock.New()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, account_type, address FROM account`)).WillReturnRows(sqlmock.NewRows([]string{
		"ID", "AccountType", "Address"}).
		AddRow([]byte{7, 205, 139, 247, 101, 123, 250, 42, 95, 96, 199, 181, 108, 85, 197, 164, 168, 36, 49, 12, 251, 252,
			209, 82, 181, 112, 94, 41, 107, 240, 83, 180}, 0, "BCZnSfqpP5tqFQlMTYkDeBVFWnbyVK7vLr5ORFpTjgs="))
	defer db.Close()
	rows, _ := db.Query(qe)
	return rows, nil
}

type mockSignatureSignExecutorFail struct {
	query.Executor
}

func (*mockSignatureSignExecutorFail) ExecuteSelect(qe string, args ...interface{}) (*sql.Rows, error) {
	db, mock, _ := sqlmock.New()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, account_type, address FROM account`)).WillReturnRows(sqlmock.NewRows([]string{
		"ID", "AccountType", "Address"}))
	defer db.Close()
	rows, _ := db.Query(qe)
	return rows, nil
}

func TestSignature_Sign(t *testing.T) {
	type fields struct {
		Executor query.ExecutorInterface
	}
	type args struct {
		payload   []byte
		accountID []byte
		seed      string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		{
			name: "Sign:valid",
			fields: fields{
				Executor: &mockSignatureSignExecutorSuccess{},
			},
			args: args{
				payload: []byte{12, 43, 65, 65, 12, 123, 43, 12, 1, 24, 5, 5, 12, 54},
				accountID: []byte{4, 38, 68, 24, 230, 247, 88, 220, 119, 124, 51, 149, 127, 214, 82, 224, 72, 239, 56, 139,
					255, 81, 229, 184, 77, 80, 80, 39, 254, 173, 28, 169},
				seed: "concur vocalist rotten busload gap quote stinging undiluted surfer goofiness deviation starved",
			},
			want: []byte{42, 62, 47, 200, 180, 101, 85, 204, 179, 147, 143, 68, 30, 111, 6, 94, 81, 248, 219, 43, 90, 6, 167,
				45, 132, 96, 130, 0, 153, 244, 159, 137, 159, 113, 78, 9, 164, 154, 213, 255, 17, 206, 153, 156, 176, 206, 33,
				103, 72, 182, 228, 148, 234, 15, 176, 243, 50, 221, 106, 152, 53, 54, 173, 15},
		},
		{
			name: "Sign:fail-{no account in account table}",
			fields: fields{
				Executor: &mockSignatureSignExecutorFail{},
			},
			args: args{
				payload: []byte{12, 43, 65, 65, 12, 123, 43, 12, 1, 24, 5, 5, 12, 54},
				accountID: []byte{4, 38, 68, 24, 230, 247, 88, 220, 119, 124, 51, 149, 127, 214, 82, 224, 72, 239, 56, 139,
					255, 81, 229, 184, 77, 80, 80, 39, 254, 173, 28, 169},
				seed: "concur vocalist rotten busload gap quote stinging undiluted surfer goofiness deviation starved",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Signature{
				Executor: tt.fields.Executor,
			}
			got := s.Sign(tt.args.payload, tt.args.accountID, tt.args.seed)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Signature.Sign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSignature_SignBlock(t *testing.T) {
	type fields struct {
		Executor *query.Executor
	}
	type args struct {
		payload  []byte
		nodeSeed string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		{
			name: "SignBlock:success",
			fields: fields{
				Executor: nil, // todo: update when node-registration integrated
			},
			args: args{
				payload:  []byte{12, 43, 65, 65, 12, 123, 43, 12, 1, 24, 5, 5, 12, 54},
				nodeSeed: "concur vocalist rotten busload gap quote stinging undiluted surfer goofiness deviation starved",
			},
			want: []byte{42, 62, 47, 200, 180, 101, 85, 204, 179, 147, 143, 68, 30, 111, 6, 94, 81, 248, 219, 43, 90, 6, 167,
				45, 132, 96, 130, 0, 153, 244, 159, 137, 159, 113, 78, 9, 164, 154, 213, 255, 17, 206, 153, 156, 176, 206, 33,
				103, 72, 182, 228, 148, 234, 15, 176, 243, 50, 221, 106, 152, 53, 54, 173, 15},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Signature{
				Executor: tt.fields.Executor,
			}
			if got := s.SignBlock(tt.args.payload, tt.args.nodeSeed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Signature.SignBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSignature_VerifySignature(t *testing.T) {
	type fields struct {
		Executor *query.Executor
	}
	type args struct {
		payload   []byte
		signature []byte
		accountID []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "VerifySignature:success", // todo: add fail:test after account integrated.
			fields: fields{
				Executor: nil, // todo: update this after account integrated
			},
			args: args{
				payload: []byte{12, 43, 65, 65, 12, 123, 43, 12, 1, 24, 5, 5, 12, 54},
				signature: []byte{42, 62, 47, 200, 180, 101, 85, 204, 179, 147, 143, 68, 30, 111, 6, 94, 81, 248, 219, 43, 90, 6, 167,
					45, 132, 96, 130, 0, 153, 244, 159, 137, 159, 113, 78, 9, 164, 154, 213, 255, 17, 206, 153, 156, 176, 206, 33,
					103, 72, 182, 228, 148, 234, 15, 176, 243, 50, 221, 106, 152, 53, 54, 173, 15},
				accountID: []byte{4, 38, 68, 24, 230, 247, 88, 220, 119, 124, 51, 149, 127, 214, 82, 224, 72, 239, 56, 139,
					255, 81, 229, 184, 77, 80, 80, 39, 254, 173, 28, 169},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Signature{
				Executor: tt.fields.Executor,
			}
			if got := s.VerifySignature(tt.args.payload, tt.args.signature, tt.args.accountID); got != tt.want {
				t.Errorf("Signature.VerifySignature() = %v, want %v", got, tt.want)
			}
		})
	}
}