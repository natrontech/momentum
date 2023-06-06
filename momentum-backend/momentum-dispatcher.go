package main

import (
	"github.com/pocketbase/pocketbase/core"
)

type MomentumDispatcherRuleType int

const (
	CREATE MomentumDispatcherRuleType = 1 << iota
	UPDATE
	DELETE
)

type MomentumDispatcherRule struct {
	tableName string
	action    func(modelEvent *core.ModelEvent) error
}

type MomentumDispatcher struct {
	createRules []*MomentumDispatcherRule
	updateRules []*MomentumDispatcherRule
	deleteRules []*MomentumDispatcherRule
}

func (d *MomentumDispatcher) DispatchCreate(modelEvent *core.ModelEvent) error {

	for _, rule := range d.createRules {
		if rule.tableName == modelEvent.Model.TableName() {
			rule.action(modelEvent)
		}
	}
	return nil
}

func (d *MomentumDispatcher) DispatchUpdate(modelEvent *core.ModelEvent) error {

	for _, rule := range d.updateRules {
		if rule.tableName == modelEvent.Model.TableName() {
			rule.action(modelEvent)
		}
	}
	return nil
}

func (d *MomentumDispatcher) DispatchDelete(modelEvent *core.ModelEvent) error {

	for _, rule := range d.deleteRules {
		if rule.tableName == modelEvent.Model.TableName() {
			rule.action(modelEvent)
		}
	}
	return nil
}
