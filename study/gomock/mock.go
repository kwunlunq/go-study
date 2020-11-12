package gomock

import (
	"kwunlunq/go-study/study/gomock/queueservice"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestQueueService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := queueservice.NewMockQueueService(ctrl)
	service.EXPECT().Push("<the message>").Return(nil)
	PushMessage(service)
}

func PushMessage(service queueservice.QueueService) {
	err := service.Push("123")
	message, err := service.Pop()
	if err != nil {
		println(err)
		return
	}
	println(message)
}
