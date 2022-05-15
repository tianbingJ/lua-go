package state

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLuaTable(t *testing.T) {
	table := newLuaTable(0, 0)
	table.put(int64(5), int64(5))
	len := table.len()
	assert.Equal(t, 0, len)
	//table.put(int64(0),int64(0))
	table.put(int64(1),int64(1))
	table.put(int64(2),int64(2))
	len = table.len()
	assert.Equal(t, 2, len)
	table.put(int64(3),int64(3))
	table.put(int64(4),int64(4))
	len = table.len()
	assert.Equal(t, 5, len)

	table.put(int64(3),nil)
	table.put(int64(4),nil)
	len = table.len()
	assert.Equal(t, 5, len)
}
