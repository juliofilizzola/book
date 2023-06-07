// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("PUBLICATIONS", testPUBLICATIONS)
	t.Run("USERS", testUSERS)
}

func TestDelete(t *testing.T) {
	t.Run("PUBLICATIONS", testPUBLICATIONSDelete)
	t.Run("USERS", testUSERSDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("PUBLICATIONS", testPUBLICATIONSQueryDeleteAll)
	t.Run("USERS", testUSERSQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("PUBLICATIONS", testPUBLICATIONSSliceDeleteAll)
	t.Run("USERS", testUSERSSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("PUBLICATIONS", testPUBLICATIONSExists)
	t.Run("USERS", testUSERSExists)
}

func TestFind(t *testing.T) {
	t.Run("PUBLICATIONS", testPUBLICATIONSFind)
	t.Run("USERS", testUSERSFind)
}

func TestBind(t *testing.T) {
	t.Run("PUBLICATIONS", testPUBLICATIONSBind)
	t.Run("USERS", testUSERSBind)
}

func TestOne(t *testing.T) {
	t.Run("PUBLICATIONS", testPUBLICATIONSOne)
	t.Run("USERS", testUSERSOne)
}

func TestAll(t *testing.T) {
	t.Run("PUBLICATIONS", testPUBLICATIONSAll)
	t.Run("USERS", testUSERSAll)
}

func TestCount(t *testing.T) {
	t.Run("PUBLICATIONS", testPUBLICATIONSCount)
	t.Run("USERS", testUSERSCount)
}

func TestHooks(t *testing.T) {
	t.Run("PUBLICATIONS", testPUBLICATIONSHooks)
	t.Run("USERS", testUSERSHooks)
}

func TestInsert(t *testing.T) {
	t.Run("PUBLICATIONS", testPUBLICATIONSInsert)
	t.Run("PUBLICATIONS", testPUBLICATIONSInsertWhitelist)
	t.Run("USERS", testUSERSInsert)
	t.Run("USERS", testUSERSInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("PUBLICATIONToUSERUsingAuth", testPUBLICATIONToOneUSERUsingAuth)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("USERToFollowerUSERS", testUSERToManyFollowerUSERS)
	t.Run("USERToUserUSERS", testUSERToManyUserUSERS)
	t.Run("USERToAuthPUBLICATIONS", testUSERToManyAuthPUBLICATIONS)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("PUBLICATIONToUSERUsingAuthPUBLICATIONS", testPUBLICATIONToOneSetOpUSERUsingAuth)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("USERToFollowerUSERS", testUSERToManyAddOpFollowerUSERS)
	t.Run("USERToUserUSERS", testUSERToManyAddOpUserUSERS)
	t.Run("USERToAuthPUBLICATIONS", testUSERToManyAddOpAuthPUBLICATIONS)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("USERToFollowerUSERS", testUSERToManySetOpFollowerUSERS)
	t.Run("USERToUserUSERS", testUSERToManySetOpUserUSERS)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("USERToFollowerUSERS", testUSERToManyRemoveOpFollowerUSERS)
	t.Run("USERToUserUSERS", testUSERToManyRemoveOpUserUSERS)
}

func TestReload(t *testing.T) {
	t.Run("PUBLICATIONS", testPUBLICATIONSReload)
	t.Run("USERS", testUSERSReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("PUBLICATIONS", testPUBLICATIONSReloadAll)
	t.Run("USERS", testUSERSReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("PUBLICATIONS", testPUBLICATIONSSelect)
	t.Run("USERS", testUSERSSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("PUBLICATIONS", testPUBLICATIONSUpdate)
	t.Run("USERS", testUSERSUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("PUBLICATIONS", testPUBLICATIONSSliceUpdateAll)
	t.Run("USERS", testUSERSSliceUpdateAll)
}
