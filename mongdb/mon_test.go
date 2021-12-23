package mongdb

import (
	"testing"
)

func TestGetMgoCli(t *testing.T) {
	//	initEngine()

	//Insert()
	// Minsert()

	//Search()
	 //Delete("6012281dc899f4f1fa0d3345")
	 Search()
	 //FuzzyFind()
	// Find()
	//DropColl()
}

func TestFindByTag(t *testing.T) {
	FindByTag()
}
func TestFind(t *testing.T) {
	Find()
}

func TestGetRecord(t *testing.T) {
	id := GetRecord()
	s := id.IsZero()
	t.Log(s)
	if id.String() == ""{
		t.Log(id)
	}

}

func TestFuzzyFind(t *testing.T) {
	FuzzyFind()
}

func TestSearch(t *testing.T) {
	Find()
}

func TestMinsert(t *testing.T) {
	Minsert()
}

func TestSession(t *testing.T) {
	Session()
}

func TestUseSession(t *testing.T) {
	UseSession()
}

func TestAllSession(t *testing.T) {
	AllSession()
}

func TestDelete(t *testing.T) {
	Delete("601f87c651266e19194dd610")
}

func TestMoreSession(t *testing.T) {
	MoreSession()
}

func TestFindJob(t *testing.T) {
	FindJob()
}

func TestFindSlice(t *testing.T) {
	 FindSlice()
}