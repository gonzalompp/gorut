package gorut

import (
    "regexp"
)

type Rut string

// Validates that the the format is correct
// Regex from http://regexlib.com/REDetails.aspx?regexp_id=2444&AspxAutoDetectCookieSupport=1
func (r Rut) IsValidFormat() bool {
    return regexp.MustCompile(`^0*(\d{1,3}(\.?\d{3})s*)\-?([\dkK])$`).MatchString(string(r))
}
