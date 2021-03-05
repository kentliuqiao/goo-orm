package session

import "reflect"

type IBeforeQuery interface {
	BeforeQuery(s *Session) error
}

type IAfterQuery interface {
	AfterQuery(s *Session) error
}

type IBeforeInsert interface {
	BeforeInsert(s *Session) error
}

type IAfterInsert interface {
	AfterInsert(s *Session) error
}

type IBeforeUpdate interface {
	BeforeUpdate(s *Session) error
}

type IAfterUpdate interface {
	AfterUpdate(s *Session) error
}

type IBeforeDelete interface {
	BeforeDelete(s *Session) error
}

type IAfterDelete interface {
	AfterDelete(s *Session) error
}

func (s *Session) CallMethodV2(method string, val interface{}) {
	dest := reflect.ValueOf(val).Interface()
	if i, ok := dest.(IBeforeQuery); ok {
		i.BeforeQuery(s)
		return
	}
	if i, ok := dest.(IAfterQuery); ok {
		i.AfterQuery(s)
		return
	}
	if i, ok := dest.(IBeforeDelete); ok {
		i.BeforeDelete(s)
		return
	}
	if i, ok := dest.(IAfterDelete); ok {
		i.AfterDelete(s)
		return
	}
	if i, ok := dest.(IBeforeUpdate); ok {
		i.BeforeUpdate(s)
		return
	}
	if i, ok := dest.(IBeforeInsert); ok {
		i.BeforeInsert(s)
		return
	}
	if i, ok := dest.(IAfterUpdate); ok {
		i.AfterUpdate(s)
		return
	}
	if i, ok := dest.(IAfterInsert); ok {
		i.AfterInsert(s)
		return
	}
	return
}