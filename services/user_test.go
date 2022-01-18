package services

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/zopping/mock-test/models"
	"github.com/zopping/mock-test/stores"
)

func TestUser_Find(t *testing.T) {
	//create an instance of gomock.Controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userStore := stores.NewMockFinder(ctrl)
	var tcs = []struct {
		description    string
		inputId        int
		expectedOutput *models.User
		mockCalls      *gomock.Call
		err            error
	}{
		{
			description:    "id not passed",
			inputId:        0,
			expectedOutput: nil,
			err:            errors.New("id not passed"),
		},
		{
			description: "output",
			inputId:     1,
			expectedOutput: &models.User{
				Id:   1,
				Name: "test",
			},
			mockCalls: userStore.EXPECT().Find(1).Return(&models.User{Id: 1, Name: "test"}, nil),
			err:       nil,
		},
		{
			description:    "store error",
			inputId:        3,
			expectedOutput: nil,
			mockCalls:      userStore.EXPECT().Find(3).Return(nil, errors.New("error")),
			err:            errors.New("error"),
		},
	}
	for i, tc := range tcs {
		userService := New(userStore)
		res, err := userService.Find(tc.inputId)
		if !reflect.DeepEqual(err, tc.err) {
			t.Errorf("Test Case [%v]: \nExpected %v \nGot %v", i+1, tc.err, err)
		}
		if !reflect.DeepEqual(res, tc.expectedOutput) {
			t.Errorf("Test Case [%v]: Expected %v Got %v", i+1, tc.expectedOutput, res)
		}
	}
}
