package gorut

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/gonzalompp/gorut"
)

//FORMAT ERRORS

func TestFormatRutEmpty(t *testing.T) {
    var rut gorut.Rut = ""
    assert.Equal(t, false, rut.IsValidFormat())
}

func TestFormatRutCero(t *testing.T) {
    var rut gorut.Rut = "0"
    assert.Equal(t, false, rut.IsValidFormat())
}

func TestFormatOnlyNumbers(t *testing.T) {
    var rut gorut.Rut = "169408849"
    assert.Equal(t, true, rut.IsValidFormat())
}

func TestFormatOnlyLetters(t *testing.T) {
    var rut gorut.Rut = "hasdghdaga"
    assert.Equal(t, false, rut.IsValidFormat())
}

func TestFormatAlternativeA(t *testing.T) {
    var rut gorut.Rut = "12546554k"
    assert.Equal(t, true, rut.IsValidFormat())
}

func TestFormatWithMinusNotValid(t *testing.T) {
    var rut gorut.Rut = "123456-7899"
    assert.Equal(t, false, rut.IsValidFormat())
}

func TestFormatJustMinus(t *testing.T) {
    var rut gorut.Rut = "-"
    assert.Equal(t, false, rut.IsValidFormat())
}

func TestFormatFullWithWrongLetterLastDigit(t *testing.T) {
    var rut gorut.Rut = "6.980.876-a"
    assert.Equal(t, false, rut.IsValidFormat())
}

/// FORMAT OKS

func TestFormatFullWithLetterLastDigitOk(t *testing.T) {
    var rut gorut.Rut = "6.980.876-k"
    assert.Equal(t, true, rut.IsValidFormat())
}

func TestFormatFullWithNumberLastDigitOk(t *testing.T) {
    var rut gorut.Rut = "6.980.876-0"
    assert.Equal(t, true, rut.IsValidFormat())
}

func TestFormatCleanOk(t *testing.T) {
    var rut gorut.Rut = "6980876k"
    assert.Equal(t, true, rut.IsValidFormat())
}

func TestFormatWithMinusInRightPlace(t *testing.T) {
    var rut gorut.Rut = "6980876-k"
    assert.Equal(t, true, rut.IsValidFormat())
}
