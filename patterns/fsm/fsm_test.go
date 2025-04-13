package fsm

import (
	"fmt"
	"testing"
)

func TestForFSM(t *testing.T) {
	p := &Product{
		ID:             "P1001",
		Name:           "商品A",
		JoinStatus:     JoinNone,
		ApprovalStatus: 0,
	}

	// 加入流程发起
	_ = p.Trigger(EventJoinRequest)
	fmt.Printf("加入申请中: %+v\n", p)

	// 后台处理：审批成功
	_ = p.Trigger(EventJoinApproved)
	fmt.Printf("加入审批通过: %+v\n", p)

	// 发起退出
	_ = p.Trigger(EventQuitRequest)
	fmt.Printf("退出申请中: %+v\n", p)

	// 后台处理：审批被拒
	_ = p.Trigger(EventQuitRejected)
	fmt.Printf("退出审批失败: %+v\n", p)
}
