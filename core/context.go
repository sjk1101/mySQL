package core

import "context"

type Context struct {
	Items map[string]interface{}

	AuditLogKey string
	Context     context.Context
}

func (this *Context) GetItem(name string) interface{} {
	if this.Items == nil {
		return nil
	}

	item, ok := this.Items[name]
	if ok {
		return item
	} else {
		return nil
	}
}

func (this *Context) SetItem(name string, val interface{}) {

	if this.Items == nil {
		this.Items = make(map[string]interface{})
	}

	this.Items[name] = val
}
