package service

import (
	"errors"
	"reflect"
	"testing"

	"github.com/captain-corgi/go-fasthttp-example/internal/domain/model"
	"github.com/captain-corgi/go-fasthttp-example/internal/domain/repository"
	"github.com/captain-corgi/go-fasthttp-example/internal/domain/repository/mocks"
	"github.com/golang/mock/gomock"
)

func TestNewUserService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)

	type args struct {
		userRepo repository.UserRepository
	}
	tests := []struct {
		name         string
		args         args
		mockBehavior func()
		want         *UserService
		testFunc     func(*UserService)
	}{
		{
			name: "NewUserService_Success",
			args: args{
				userRepo: mockRepo,
			},
			mockBehavior: func() {
				mockRepo.EXPECT().
					GetByID("123").
					Return(&model.User{ID: "123", Username: "John Doe"}, nil).
					Times(1)
			},
			want: &UserService{
				userRepo: mockRepo,
			},
			testFunc: func(s *UserService) {
				// Call GetUser to satisfy the mock expectation
				s.GetUser("123")
			},
		},
		{
			name: "NewUserService_Error",
			args: args{
				userRepo: mockRepo,
			},
			mockBehavior: func() {
				mockRepo.EXPECT().
					GetByID("123").
					Return(nil, errors.New("user not found")).
					Times(1)
			},
			want: &UserService{
				userRepo: mockRepo,
			},
			testFunc: func(s *UserService) {
				// Call GetUser to satisfy the mock expectation
				s.GetUser("123")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			got := NewUserService(tt.args.userRepo)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
			// Execute the test function to satisfy mock expectations
			tt.testFunc(got)
		})
	}
}

func TestUserService_GetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)

	type fields struct {
		userRepo repository.UserRepository
	}
	type args struct {
		id string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		mockBehavior func()
		want         *model.User
		wantErr      bool
	}{
		{
			name: "GetUser_Success",
			fields: fields{
				userRepo: mockRepo,
			},
			args: args{
				id: "123",
			},
			mockBehavior: func() {
				mockRepo.EXPECT().GetByID("123").Return(&model.User{ID: "123", Username: "John Doe"}, nil)
			},
			want:    &model.User{ID: "123", Username: "John Doe"},
			wantErr: false,
		},
		{
			name: "GetUser_Error",
			fields: fields{
				userRepo: mockRepo,
			},
			args: args{
				id: "123",
			},
			mockBehavior: func() {
				mockRepo.EXPECT().GetByID("123").Return(nil, errors.New("user not found"))
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			s := &UserService{
				userRepo: tt.fields.userRepo,
			}
			got, err := s.GetUser(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)

	type fields struct {
		userRepo repository.UserRepository
	}
	type args struct {
		user *model.User
	}
	tests := []struct {
		name         string
		fields       fields
		mockBehavior func()
		args         args
		wantErr      bool
	}{
		{
			name: "CreateUser_Success",
			fields: fields{
				userRepo: mockRepo,
			},
			mockBehavior: func() {
				mockRepo.EXPECT().Create(&model.User{ID: "123", Username: "John Doe"}).Return(nil)
			},
			args:    args{user: &model.User{ID: "123", Username: "John Doe"}},
			wantErr: false,
		},
		{
			name: "CreateUser_Error",
			fields: fields{
				userRepo: mockRepo,
			},
			mockBehavior: func() {
				mockRepo.EXPECT().Create(&model.User{ID: "123", Username: "John Doe"}).Return(errors.New("user already exists"))
			},
			args:    args{user: &model.User{ID: "123", Username: "John Doe"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			s := &UserService{
				userRepo: tt.fields.userRepo,
			}
			if err := s.CreateUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UserService.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserService_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)

	type fields struct {
		userRepo repository.UserRepository
	}
	type args struct {
		user *model.User
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		mockBehavior func()
		wantErr      bool
	}{
		{
			name: "UpdateUser_Success",
			fields: fields{
				userRepo: mockRepo,
			},
			mockBehavior: func() {
				mockRepo.EXPECT().Update(&model.User{ID: "123", Username: "John Doe Updated"}).Return(nil)
			},
			args:    args{user: &model.User{ID: "123", Username: "John Doe Updated"}},
			wantErr: false,
		},
		{
			name: "UpdateUser_Error",
			fields: fields{
				userRepo: mockRepo,
			},
			mockBehavior: func() {
				mockRepo.EXPECT().Update(&model.User{ID: "123", Username: "John Doe Updated"}).Return(errors.New("user not found"))
			},
			args:    args{user: &model.User{ID: "123", Username: "John Doe Updated"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			s := &UserService{
				userRepo: tt.fields.userRepo,
			}
			if err := s.UpdateUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UserService.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserService_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)

	type fields struct {
		userRepo repository.UserRepository
	}
	type args struct {
		id string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		mockBehavior func()
		wantErr      bool
	}{
		{
			name: "DeleteUser_Success",
			fields: fields{
				userRepo: mockRepo,
			},
			args: args{
				id: "123",
			},
			mockBehavior: func() {
				mockRepo.EXPECT().Delete("123").Return(nil)
			},
			wantErr: false,
		},
		{
			name: "DeleteUser_Error",
			fields: fields{
				userRepo: mockRepo,
			},
			args: args{
				id: "123",
			},
			mockBehavior: func() {
				mockRepo.EXPECT().Delete("123").Return(errors.New("user not found"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			s := &UserService{
				userRepo: tt.fields.userRepo,
			}
			if err := s.DeleteUser(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UserService.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
