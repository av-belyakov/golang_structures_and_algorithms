package structuraldesignpatterns

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBridgeExample(t *testing.T) {
	var showData IShowerData = &ShowData{Convert: ConvertDate{}}
	myt, err := showData.ConverterDate("2023-01-02T15:04:05+07:00")
	assert.Nil(t, err)

	showData.SetData(myt)
	year, month, day := showData.GetData()
	h, m, s := showData.GetTime()
	fmt.Printf("Data: %d %s %d", year, month.String(), day)
	fmt.Printf("Time: %d:%d:%d", h, m, s)

	assert.Equal(t, 15, h)
	assert.Equal(t, 4, m)
	assert.Equal(t, 5, s)
}
