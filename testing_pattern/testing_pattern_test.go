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
			input:    "abcdef\r\n",
			expected: "abcdef",
		},
		`ending with \n`: {
			input:    "abcdef\n",
			expected: "abcdef",
		},
		`ending with multiple \n`: {
			input:    "abcdef\n\n\n",
			expected: "abcdef",
		},
		`ending without newline`: {
			input:    "abcdef",
			expected: "abcdef",
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
