package prettytable_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	prettytable "github.com/av-belyakov/golang_structures_and_algorithms/thirdpartylibraries/go-pretty/prettytable/cmd"
)

func TestCreateTable(t *testing.T) {
	footerNames := []any{"", "", "", "Total:", 17557}
	rows := [][]any{
		{1, "Aihor Helmet", "Moscow, 10-Parkovay, 34", "2024-11-21", 2453},
		{2, "Shoe Helmet", "Moscow, Leninskiy Prospekt, 41, k3", "2024-11-19", 6553},
		{3, "LS2 Helmet", "Moscow, Cherkizovsky, 10, st.12", "2024-11-21", 2324},
		{4, "AGV Helmet", "St.Peterburg, Primorsy, 67", "2024-10-01", 5353},
	}

	prettytable.PrettyTableExample(footerNames, rows)
	assert.True(t, true)
}
