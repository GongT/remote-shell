package handlers

import (
	"github.com/gongt/remote-shell/internal/timeout-controller"
	"github.com/gongt/remote-shell/internal/actions/action-base"
)

type CallbackAction struct {
	action_base.MessageBase
}

func NewCallbackAction(id uint32) action_base.Message {
	return &CallbackAction{
		action_base.MessageBase{
			TypeId: action_base.TypeCallback,
			Id:     id,
		},
	}
}

func (act *CallbackAction) Handle() error {
	timeout_controller.Cancel(act.Id)
	return nil
}
