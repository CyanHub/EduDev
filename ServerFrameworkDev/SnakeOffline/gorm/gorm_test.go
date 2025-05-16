package gorm

import "testing"

func TestAppendAssociation(t *testing.T) {
	AppendAssociation()
}

func TestQueryAssociation(t *testing.T) {
	QueryAssociation()
}

func TestReplaceAssociation(t *testing.T) {
	ReplaceAssociation()
}

func TestCountAssociation(t *testing.T) {
	CountAssociation()
}

func TestDeleteAssociation(t *testing.T) {
	DeleteAssociation()
}

func TestAutoAssociation(t *testing.T) {
	AutoAssociation()
}

func TestPreLoad(t *testing.T) {
	PreLoad()
}

func TestExampleJSONSerializer(t *testing.T) {
	ExampleJSONSerializer()
}

func TestExampleGobSerializer(t *testing.T) {
	ExampleGobSerializer()
}


func TestSessionDryRun(t *testing.T) {
	SessionDryRun()
}

func TestSessionNewDB(t *testing.T) {
	SessionNewDB()
}

func TestSessionPrepareStmt(t *testing.T) {
	SessionPrepareStmt()
}
