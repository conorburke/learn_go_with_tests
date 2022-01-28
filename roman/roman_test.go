package roman

import "testing"

func TestRomanNumerals(t *testing.T) {
    cases := []struct {
        Description string
        Input int
        Expected string
    }{
        {"convert 1 to I", 1, "I"},
        {"convert 2 to II", 2, "II"},
        {"convert 3 to III", 3, "III"},
        {"convert 4 to IV", 4, "IV"},
        {"convert 5 to V", 5, "V"},
    }
    
    
    for _, tst := range cases {
        t.Run(tst.Description, func (t *testing.T) { 
            output := ConvertToRoman(tst.Input)

            if output != tst.Expected {
                t.Errorf("expected %q but got %q", tst.Expected, output)
            }
        })
    }
    
    
    // t.Run("conver 1 to I", func (t *testing.T) {
    //     got := ConvertToRoman(1)
    //     expected := "I"
    
    //     if got != expected {
    //         t.Errorf("expected %q but got %q", expected, got)
    //     }
    // })

    // t.Run("conver 2 to II", func (t *testing.T) {
    //     got := ConvertToRoman(2)
    //     expected := "II"
    
    //     if got != expected {
    //         t.Errorf("expected %q but got %q", expected, got)
    //     }
    // })
    
    
}