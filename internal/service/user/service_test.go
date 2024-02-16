package user

import (
	"context"
	"errors"
	"github.com/ainurqa95/mood-lifter/internal/mock"
	"github.com/ainurqa95/mood-lifter/internal/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestService_CreateIfNotExists(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := mock.NewMockUserRepository(ctrl)
	type TestCase struct {
		name              string
		userInfo          *model.UserInfo
		expectedStrExists bool
		repoErr           error
		expectedErr       error
	}
	testCases := []TestCase{
		{
			name: "success",
			userInfo: &model.UserInfo{
				UserName: "",
				ChatID:   123,
			},
			expectedStrExists: true,
			repoErr:           nil,
			expectedErr:       nil,
		},
		{
			name: "failed creation",
			userInfo: &model.UserInfo{
				UserName: "",
				ChatID:   123,
			},
			expectedStrExists: false,
			repoErr:           errors.New("test"),
			expectedErr:       errors.New("test"),
		},
	}
	userService := NewService(m)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Return(tc.expectedErr)
			userId, err := userService.CreateIfNotExists(context.TODO(), tc.userInfo)
			assert.Equal(t, tc.expectedErr, err)
			assert.Equal(t, len(userId) > 0, tc.expectedStrExists)
		})
	}
}
