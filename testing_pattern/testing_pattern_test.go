package testingpattern

import (
	"strings"
	"testing"
)

//Примеры использования шаблонов тестирования

//Если несколько модульных тестов имеют схожую структуру, мы можем объединить их,
//используя тесты, управляемые таблицами. Поскольку этот метод предотвращает дублирование, он упрощает
//изменение логики тестирования и упрощает добавление новых вариантов использования.

func removeNewLineSuffixes(s string) string {
	if s == "" {
		return s
	}

	if strings.HasSuffix(s, "\r\n") {
		return removeNewLineSuffixes(s[:len(s)-2])
	}

	if strings.HasSuffix(s, "\n") {
		return removeNewLineSuffixes(s[:len(s)-1])
	}

	return s
}

func TestRemoveNewLineSuffix(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected string
	}{
		`empty`: {
			input:    "",
			expected: "",
		},
		`ending with \r\n`: {
			input:    "a\r\n",
			expected: "a",
		},
		`ending with \n`: {
			input:    "a\n",
			expected: "a",
		},
		`ending with multiple \n`: {
			input:    "a\n\n\n",
			expected: "a",
		},
		`ending without newline`: {
			input:    "a",
			expected: "a",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := removeNewLineSuffixes(tt.input)
			if got != tt.expected {
				t.Errorf("got: %s, expected: %s", got, tt.expected)
			}
		})
	}
}
