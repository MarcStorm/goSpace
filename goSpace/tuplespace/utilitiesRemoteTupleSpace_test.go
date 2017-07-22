package tuplespace

import (
	"goSpace/goSpace/topology"
	"reflect"
	"testing"
)

func TestPutUtilities(t *testing.T) {
	ts := CreateTupleSpace(9050)
	if !(ts.Size() == 0) {
		t.Errorf("Tuple space is not empty")
	}
	ptp := topology.CreatePointToPoint("Bookstore", "localhost", 9050)
	Put(ptp, "hello", false)
	if !(reflect.DeepEqual(CreateTuple([]interface{}{"hello", false}), ts.tuples[0])) {
		t.Errorf("Tuple space is not empty")
	}
}

func TestQueryAndGetUtilities(t *testing.T) {
	ts := CreateTupleSpace(9051)
	if !(ts.Size() == 0) {
		t.Errorf("Tuple space is not empty")
	}
	ptp := topology.CreatePointToPoint("Bookstore", "localhost", 9051)
	Put(ptp, "hello", false)
	var s string
	querySucceed := Query(ptp, &s, false)
	if !(ts.Size() == 1) {
		t.Errorf("Tuple space should have one tuple")
	}
	var b bool
	getSucceed := Get(ptp, "hello", &b)
	if !(ts.Size() == 0) {
		t.Errorf("Tuple space is not empty")
	}
	if b || !(s == "hello") || !getSucceed || !querySucceed {
		t.Errorf("Get or Query Failed")
	}
}

func TestPutPUtilities(t *testing.T) {
	ts := CreateTupleSpace(9053)
	if !(ts.Size() == 0) {
		t.Errorf("Tuple space is not empty")
	}
	ptp := topology.CreatePointToPoint("Bookstore", "localhost", 9053)
	PutP(ptp, "hello", false)
	var b bool
	Get(ptp, "hello", &b)
	if b {
		t.Errorf("PutP Failed")
	}
}

func TestWriteTupleToVariables(t *testing.T) {
	tuple := CreateTuple([]interface{}{2, "hello"})
	var i int
	var s string
	variables := []interface{}{&i, &s}
	WriteTupleToVariables(tuple, variables)
	if i != 2 || s != "hello" {
		t.Errorf("Write tuple to variable did not work as expected")
	}
}

func TestQueryPAndGetPUtilities(t *testing.T) {

	ts := CreateTupleSpace(9052)
	if !(ts.Size() == 0) {
		t.Errorf("Tuple space is not empty")
	}
	ptp := topology.CreatePointToPoint("Bookstore", "localhost", 9052)
	Put(ptp, "hello", false)
	var s string
	queryPResult, queryPSucceed := QueryP(ptp, &s, false)
	if !(ts.Size() == 1) {
		t.Errorf("Tuple space should have one tuple")
	}
	var b bool
	getPResult, getPSucceed := GetP(ptp, "hello", &b)
	if !(ts.Size() == 0) {
		t.Errorf("Tuple space is not empty")
	}
	if b || !(s == "hello") || !getPSucceed || !queryPSucceed {
		t.Errorf("GetP or QueryP Failed")
	}
	if !getPResult || !queryPResult {
		t.Errorf("GetP or QueryP returned wrong boolean")
	}
	queryPResult, queryPSucceed = QueryP(ptp, &s, false)
	getPResult, getPSucceed = GetP(ptp, "hello", &b)
	if getPResult || queryPResult {
		t.Errorf("GetP or QueryP returned wrong boolean")
	}
}

func TestGetAllAndQueryAll(t *testing.T) {
	ts := CreateTupleSpace(9054)
	if !(ts.Size() == 0) {
		t.Errorf("Tuple space is not empty")
	}
	ptp := topology.CreatePointToPoint("Bookstore", "localhost", 9054)
	Put(ptp, 2, 2)
	Put(ptp, 2, 2)
	Put(ptp, 2, 3)
	Put(ptp, 2, 3)
	Put(ptp, 2, false)
	i := 1
	tuples, b := QueryAll(ptp, 2, 2)
	tuple1 := CreateTuple([]interface{}{2, 2})
	expectedTuples := []Tuple{tuple1, tuple1}
	if !reflect.DeepEqual(tuples, expectedTuples) {
		t.Errorf("QueryAll returned wrong tuple list %v - %v", tuples, expectedTuples)
	}
	if !b {
		t.Errorf("QueryAll returned wrong boolean")
	}
	tuple2 := CreateTuple([]interface{}{2, 3})
	expectedTuples = []Tuple{tuple1, tuple1, tuple2, tuple2}
	tuples, b = GetAll(ptp, 2, &i)
	if !reflect.DeepEqual(tuples, expectedTuples) {
		t.Errorf("GetAll returned wrong tuple list %v - %v", tuples, expectedTuples)
	}
	if !b {
		t.Errorf("GetAll returned wrong boolean")
	}
	if i != 1 {
		t.Errorf("i was overwritten")
	}
	//fmt.Println(tuplespace.QueryAll(ptp, 2, 2))
	//fmt.Println(tuplespace.GetAll(ptp, 2, &i))
	//fmt.Println(tuplespace.QueryAll(ptp, 2, 2))
}
