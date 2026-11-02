package hw03frequencyanalysis

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTopBasicRules(t *testing.T) {
	tests := []struct {
		name string
		text string
		want []string
	}{
		{"empty", "", []string{}},
		{"whitespace", " \t\n\r\u00a0\u2003", []string{}},
		{
			"example", "cat and dog, one dog,two cats and one man",
			[]string{"and", "one", "cat", "cats", "dog,", "dog,two", "man"},
		},
		{"unicode whitespace", "cat\tdog\ncat\r\ndog\u00a0cat\u2003bird", []string{"cat", "dog", "bird"}},
		{"case and word forms", "Нога нога ноги ногу нога", []string{"нога", "Нога", "ноги", "ногу"}},
		{"punctuation", "нога, нога - ---- ----", []string{"----", "-", "нога", "нога,"}},
		{"frequency before alphabet", "a z z b b b", []string{"b", "z", "a"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, Top(tt.text))
		})
	}
}

func TestTopNormalized(t *testing.T) {
	tests := []struct {
		name string
		text string
		want []string
	}{
		{"empty", "", []string{}},
		{"whitespace", " \t\n\r\u00a0\u2003", []string{}},
		{"case and edges", "Нога нога! нога нога, 'нога' «НОГА»", []string{"нога"}},
		{"word forms", "Нога ноги ногу нога", []string{"нога", "ноги", "ногу"}},
		{"inner hyphen", "какой-то какойто какой-то", []string{"какой-то", "какойто"}},
		{"inner punctuation", "dog,cat dog...cat dogcat", []string{"dog,cat", "dog...cat", "dogcat"}},
		{"hyphen sequences", "- ------- ---- -- -------", []string{"-------", "--", "----"}},
		{"punctuation only", "- ! ... , ' «» —", []string{}},
		{"boundary hyphens", "-Нога- нога", []string{"нога"}},
		{"unicode whitespace", "CAT\tdog\ncat\u00a0DOG\u2003cat", []string{"cat", "dog"}},
		{"frequency after normalization", "Z! z A a! Z b", []string{"z", "a", "b"}},
		{"original sample", text, []string{"а", "он", "и", "ты", "что", "в", "его", "если", "кристофер", "не"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, TopNormalized(tt.text))
		})
	}
}

func TestTopLimit(t *testing.T) {
	for name, top := range map[string]func(string, ...int) []string{
		"basic":      Top,
		"normalized": TopNormalized,
	} {
		t.Run(name, func(t *testing.T) {
			input := strings.Repeat("o n m l k j i h g f e d c b a ", 133) + strings.Repeat("z ", 99)
			want := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
			require.Equal(t, want, top(input))
			require.Equal(t, []string{"z", "a", "b", "c", "d", "e", "f", "g", "h", "i"},
				top(input+strings.Repeat("z ", 35)))
		})
	}
}

func TestTopOptionalCount(t *testing.T) {
	for name, top := range map[string]func(string, ...int) []string{
		"basic":      Top,
		"normalized": TopNormalized,
	} {
		t.Run(name, func(t *testing.T) {
			input := "l k j i h g f e d c b a l"
			ordered := []string{"l", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
			require.Equal(t, ordered[:10], top(input))
			require.Equal(t, ordered[:3], top(input, 3))
			require.Equal(t, ordered[:11], top(input, 11))
			require.Equal(t, ordered, top(input, 20))
			require.Empty(t, top(input, 0))
			require.Empty(t, top(input, -1))
			require.Empty(t, top("", 3))
		})
	}
}
