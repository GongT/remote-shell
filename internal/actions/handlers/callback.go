package handlers

import (
	action_base "github.com/gongt/remote-shell/internal/actions/action-base"
	timeout_controller "github.com/gongt/remote-shell/internal/timeout-controller"
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

func (act *CallbackAction) Handle() (bool, error) {
	timeout_controller.Cancel(act.Id)
	return false, nil
}
