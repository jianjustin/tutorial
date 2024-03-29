package state

import "fmt"

type VendingMachine struct {
	HasItem       State
	ItemRequested State
	HasMoney      State
	NoItem        State

	CurrentState State

	ItemCount int
	ItemPrice int
}

func NewVendingMachine(itemCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		ItemCount: itemCount,
		ItemPrice: itemPrice,
	}
	hasItemState := &HasItemState{
		vendingMachine: v,
	}
	itemRequestedState := &ItemRequestedState{
		vendingMachine: v,
	}
	hasMoneyState := &HasMoneyState{
		vendingMachine: v,
	}
	noItemState := &NoItemState{
		vendingMachine: v,
	}

	v.SetState(hasItemState)
	v.HasItem = hasItemState
	v.ItemRequested = itemRequestedState
	v.HasMoney = hasMoneyState
	v.NoItem = noItemState
	return v
}

func (v *VendingMachine) RequestItem() error {
	return v.CurrentState.requestItem()
}

func (v *VendingMachine) AddItem(count int) error {
	return v.CurrentState.addItem(count)
}

func (v *VendingMachine) InsertMoney(money int) error {
	return v.CurrentState.insertMoney(money)
}

func (v *VendingMachine) DispenseItem() error {
	return v.CurrentState.dispenseItem()
}

func (v *VendingMachine) SetState(s State) {
	v.CurrentState = s
}

func (v *VendingMachine) IncrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.ItemCount = v.ItemCount + count
}
