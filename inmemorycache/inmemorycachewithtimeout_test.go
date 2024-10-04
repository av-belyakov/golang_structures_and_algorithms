package inmemorycache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewStorages(t *testing.T) {
	var (
		test_uniq_case_id_1 string = "uniq_case_id:f78773r88r8w874et7rt7g77sw7w"
		test_uniq_case_id_2 string = "uniq_case_id:g7627gdff8fyr8298euihusd8y823"
		test_uniq_case_id_3 string = "uniq_case_id:fs662te73t73tr73t6rt37tr7376r3"
	)

	ts, err := NewWebHookTemporaryStorage(10)
	assert.NoError(t, err)

	test_uuid_1 := ts.SetElementId(test_uniq_case_id_1)
	d, ok := ts.GetElementId(test_uuid_1)
	assert.True(t, ok)
	assert.Equal(t, d, test_uniq_case_id_1)

	test_uuid_2 := ts.SetElementId(test_uniq_case_id_2)
	d, ok = ts.GetElementId(test_uuid_2)
	assert.True(t, ok)
	assert.Equal(t, d, test_uniq_case_id_2)

	ts.DeleteElement(test_uuid_1)
	_, ok = ts.GetElementId(test_uuid_1)
	assert.False(t, ok)

	time.Sleep(9 * time.Second)

	test_uuid_3 := ts.SetElementId(test_uniq_case_id_3)

	time.Sleep(6 * time.Second)

	//удаляется автоматически
	_, ok = ts.GetElementId(test_uuid_2)
	assert.False(t, ok)

	d, ok = ts.GetElementId(test_uuid_3)
	assert.True(t, ok)
	assert.Equal(t, d, test_uniq_case_id_3)
}
