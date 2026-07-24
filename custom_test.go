package captcha

import (
	"strings"
	"testing"
)

func TestCustom(t *testing.T) {
	c := New()
	const charset = "ABC123"
	r, err := c.Custom(5, charset)
	if err != nil {
		t.Fatal(err)
	}
	if len(r.Text) != 5 {
		t.Fatalf("len=%d want 5", len(r.Text))
	}
	for _, ch := range r.Text {
		if !strings.ContainsRune(charset, ch) {
			t.Fatalf("char %q not in charset", ch)
		}
	}
	if _, err := c.Custom(0, charset); err == nil {
		t.Fatal("want error for length 0")
	}
	if _, err := c.Custom(3, ""); err == nil {
		t.Fatal("want error for empty charset")
	}
}

func TestWord(t *testing.T) {
	c := New()
	words := []string{"apple", "cat", "delta"}
	r, err := c.Word(words)
	if err != nil {
		t.Fatal(err)
	}
	found := false
	for _, w := range words {
		if r.Text == w {
			found = true
		}
	}
	if !found {
		t.Fatalf("text %q not in word list", r.Text)
	}
	if _, err := c.Word(nil); err == nil {
		t.Fatal("want error for empty word list")
	}
}
