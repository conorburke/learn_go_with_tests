package maps

import "testing"

func TestLookup(t *testing.T) {
    dict := Dictionary{"word": "definition"}

    expected := "definition"
    got := dict.Lookup("word")

    if expected != got {
        t.Errorf("expected %q but got %q", expected, got)
    }

    t.Run("add to map", func(t *testing.T) {
        d := Dictionary{}

        expected := ""
        got := d.Lookup("first")

        if expected != got {
            t.Errorf("expected %q but got %q", expected, got)
        }

        d.Add("first", "entry1")

        expected = "entry1"
        got = d.Lookup("first")

        if expected != got {
            t.Errorf("expected %q but got %q", expected, got)
        }
        
    })
}