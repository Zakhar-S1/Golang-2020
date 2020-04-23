package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLess(t *testing.T) {

	assert := assert.New(t)

	ivanIvanovDate1, _ := time.Parse("2006-Jan-02", "2005-Aug-10")
	artiomIvanovDate, _ := time.Parse("2006-Jan-02", "2005-Aug-10")
	ivanIvanovDate2, _ := time.Parse("2006-Jan-02", "2003-Aug-05")

	p := People{
		{firstName: "I", lastName: "Ivanov", birthDay: ivanIvanovDate1},
		{firstName: "A", lastName: "Ivanov", birthDay: artiomIvanovDate},
		{firstName: "I", lastName: "Ivanov", birthDay: ivanIvanovDate2},
	}

	assert.Equal(p.Less(0, 1), true, "Error. Check 0 and 1 element.")
	assert.Equal(p.Less(0, 0), false, "Error. Check the same element.")
	assert.Equal(p.Less(0, 2), true, "Error. Check 0 and 2 element.")
	assert.Equal(p.Less(1, 2), true, "Error. Check 1 and 2 element.")
}
