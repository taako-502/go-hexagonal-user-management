package user_model

import (
	"reflect"
	"testing"
)

func TestNewUser(t *testing.T) {
	type args struct {
		username string
		email    string
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "should create a user successfully",
			args: args{
				username: "test",
				email:    "test@test.com",
			},
			want: &User{
				Username: "test",
				Email:    "test@test.com",
			},
			wantErr: false,
		},
		{
			// usernameが空白のケース
			name: "should return an error when username is blank",
			args: args{
				username: "",
				email:    "test@test.com",
			},
			want:    nil,
			wantErr: true,
		},
		{
			// emailが空白のケース
			name: "should return an error when email is blank",
			args: args{
				username: "test",
				email:    "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(tt.args.username, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
