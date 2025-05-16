package redis

import "testing"

func TestBaseOperation(t *testing.T) {
	BaseOperation()
}

func TestListOperation(t *testing.T) {
	ListOperation()
}

func TestSetOperation(t *testing.T) {
	SetOperation()
}

func TestZSetOperation(t *testing.T) {
	ZSetOperation()
}

func TestHashOperation(t *testing.T) {
	HashOperation()
}

func TestPubSubOperation(t *testing.T) {
	PubSubOperation()
}

