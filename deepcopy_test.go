package deepcopy

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCopy(t *testing.T) {
	p := NewComplicatedMan()
	c := ComplicatedMan{}

	Clone(&c, &p)

	pjb, err := json.Marshal(&p)
	assert.Nil(t, err)
	cjb, err := json.Marshal(&c)
	assert.Nil(t, err)

	assert.Equal(t, string(pjb), string(cjb))
}

// lodash <- 얘가 딥카피하면 바로 떠오르는 라이브러리
