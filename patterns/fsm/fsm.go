package fsm

import (
	"errors"
	"fmt"
)

type JoinStatus int
type ApprovalStatus int
type Event string

const (
	JoinNone JoinStatus = 1
	JoinYes  JoinStatus = 2

	ApprovalPassed   ApprovalStatus = 1
	ApprovalRejected ApprovalStatus = 2
	ApprovalJoining  ApprovalStatus = 3
	ApprovalQuitting ApprovalStatus = 4
)

const (
	// 发起事件
	EventJoinRequest Event = "JoinRequest"
	EventQuitRequest Event = "QuitRequest"

	// 审批结果事件
	EventJoinApproved Event = "JoinApproved"
	EventJoinRejected Event = "JoinRejected"
	EventQuitApproved Event = "QuitApproved"
	EventQuitRejected Event = "QuitRejected"
)

type Product struct {
	ID             string
	Name           string
	JoinStatus     JoinStatus
	ApprovalStatus ApprovalStatus
}

type HookFunc func(p *Product) error

type Transition struct {
	ToJoinStatus     JoinStatus
	ToApprovalStatus ApprovalStatus
	Before           HookFunc
	After            HookFunc
}

type StateKey struct {
	JoinStatus     JoinStatus
	ApprovalStatus ApprovalStatus
	Event          Event
}

var transitions = map[Event]Transition{
	// 发起加入
	EventJoinRequest: {
		ToJoinStatus:     JoinNone,
		ToApprovalStatus: ApprovalJoining,
		Before: func(p *Product) error {
			fmt.Println("[Before] 校验是否允许加入")
			return nil
		},
		After: func(p *Product) error {
			fmt.Println("[After] 发起加入审批流程")
			return nil
		},
	},
	// 加入审批结果
	EventJoinApproved: {
		ToJoinStatus:     JoinYes,
		ToApprovalStatus: ApprovalPassed,
	},
	EventJoinRejected: {
		ToJoinStatus:     JoinNone,
		ToApprovalStatus: ApprovalRejected,
	},

	// 发起退出
	EventQuitRequest: {
		ToJoinStatus:     JoinYes,
		ToApprovalStatus: ApprovalQuitting,
		Before: func(p *Product) error {
			fmt.Println("[Before] 校验是否允许退出")
			return nil
		},
		After: func(p *Product) error {
			fmt.Println("[After] 发起退出审批流程")
			return nil
		},
	},
	// 退出审批结果
	EventQuitApproved: {
		ToJoinStatus:     JoinNone,
		ToApprovalStatus: ApprovalPassed,
	},
	EventQuitRejected: {
		ToJoinStatus:     JoinYes,
		ToApprovalStatus: ApprovalRejected,
	},
}

func (p *Product) Trigger(event Event) error {
	//key := StateKey{
	//	JoinStatus:     p.JoinStatus,
	//	ApprovalStatus: p.ApprovalStatus,
	//	Event:          event,
	//}

	transition, ok := transitions[event]
	if !ok {
		return errors.New("无效状态变更")
	}

	if transition.Before != nil {
		if err := transition.Before(p); err != nil {
			return fmt.Errorf("before hook 失败: %w", err)
		}
	}

	p.JoinStatus = transition.ToJoinStatus
	p.ApprovalStatus = transition.ToApprovalStatus

	if transition.After != nil {
		if err := transition.After(p); err != nil {
			return fmt.Errorf("after hook 失败: %w", err)
		}
	}

	return nil
}
