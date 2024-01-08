package cp949

import (
	"bytes"
	"io"
	"testing"
)

// TestCP949ToUTF8 tests the conversion from CP949 to UTF-8.
func TestCP949ToUTF8(t *testing.T) {
	enc := NewCP949Encoding()

	for _, tt := range tests {
		reader := enc.NewDecoder().Reader(bytes.NewBufferString(tt.cp949))
		decoded, err := io.ReadAll(reader)
		if err != nil {
			t.Errorf("Failed to decode CP949: %v", err)
			continue
		}

		if utf8 := string(decoded); utf8 != tt.utf8 {
			t.Errorf("CP949ToUTF8(%q) = %q, want %q", tt.cp949, utf8, tt.utf8)
		}
	}
}

// TestUTF8ToCP949 tests the conversion from UTF-8 to CP949.
func TestUTF8ToCP949(t *testing.T) {
	enc := NewCP949Encoding()

	for _, tt := range tests {
		writer := new(bytes.Buffer)
		w := enc.NewEncoder().Writer(writer)
		if _, err := w.Write([]byte(tt.utf8)); err != nil {
			t.Errorf("Failed to encode UTF-8 to CP949: %v", err)
			continue
		}

		if cp949 := writer.String(); cp949 != tt.cp949 {
			t.Errorf("UTF8ToCP949(%q) = %q, want %q", tt.utf8, cp949, tt.cp949)
		}
	}
}
