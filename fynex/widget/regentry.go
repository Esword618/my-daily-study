package widget

import (
	_ "embed"

	"fyne.io/fyne/v2/data/validation"
	wdg "fyne.io/fyne/v2/widget"
)

func AlphaEntry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp("^[a-zA-Z]+$", "Please input Alpha")

	return e
}
func AlphaNumericEntry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp("^[a-zA-Z0-9]+$", "Please input AlphaNumeric")

	return e
}
func AlphaUnicodeEntry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp("^[\\p{L}]+$", "Please input AlphaUnicode")

	return e
}
func AlphaUnicodeNumericEntry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp("^[\\p{L}\\p{N}]+$", "Please input AlphaUnicodeNumeric")

	return e
}
func NumericEntry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp("^[-+]?[0-9]+(?:\\.[0-9]+)?$", "Please input Numeric")

	return e
}
func NumberEntry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp("^[0-9]+$", "Please input Number")

	return e
}
func HexadecimalEntry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp("^(0[xX])?[0-9a-fA-F]+$", "Please input Hexadecimal")

	return e
}
func HexColorEntry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp("^#(?:[0-9a-fA-F]{3}|[0-9a-fA-F]{6})$", "Please input HexColor")

	return e
}
func RgbEntry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp("^rgb\\(\\s*(?:(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])|(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%)\\s*\\)$", "Please input Rgb")

	return e
}
func RgbaEntry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp("^rgba\\(\\s*(?:(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])|(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%)\\s*,\\s*(?:(?:0.[1-9]*)|[01])\\s*\\)$", "Please input Rgba")

	return e
}
func HslEntry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp("^hsl\\(\\s*(?:0|[1-9]\\d?|[12]\\d\\d|3[0-5]\\d|360)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*\\)$", "Please input Hsl")

	return e
}
func HslaEntry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp("^hsla\\(\\s*(?:0|[1-9]\\d?|[12]\\d\\d|3[0-5]\\d|360)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0.[1-9]*)|[01])\\s*\\)$", "Please input Hsla")

	return e
}
func EmailEntry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp("^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22))))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$", "Please input Email")

	return e
}
func E164Entry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp("^\\+[1-9]?[0-9]{7,14}$", "Please input E164")

	return e
}
func Base64Entry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp("^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$", "Please input Base64")

	return e
}
func Base64URLEntry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp("^(?:[A-Za-z0-9-_]{4})*(?:[A-Za-z0-9-_]{2}==|[A-Za-z0-9-_]{3}=|[A-Za-z0-9-_]{4})$", "Please input Base64URL")

	return e
}
func ISBN10Entry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp("^(?:[0-9]{9}X|[0-9]{10})$", "Please input ISBN10")

	return e
}
func ISBN13Entry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp("^(?:(?:97(?:8|9))[0-9]{10})$", "Please input ISBN13")

	return e
}
func UUID3Entry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp("^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[0-9a-f]{4}-[0-9a-f]{12}$", "Please input UUID3")

	return e
}
func UUID4Entry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp("^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$", "Please input UUID4")

	return e
}
func UUID5Entry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp("^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$", "Please input UUID5")

	return e
}
func UUIDEntry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$", "Please input UUID")

	return e
}
func UUID3RFC4122Entry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-3[0-9a-fA-F]{3}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$", "Please input UUID3RFC4122")

	return e
}
func UUID4RFC4122Entry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$", "Please input UUID4RFC4122")

	return e
}
func UUID5RFC4122Entry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-5[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$", "Please input UUID5RFC4122")

	return e
}
func UUIDRFC4122Entry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$", "Please input UUIDRFC4122")

	return e
}
func ASCIIEntry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp("^[\x00-\x7F]*$", "Please input ASCII")

	return e
}
func PrintableASCIIEntry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp("^[\x20-\x7E]*$", "Please input PrintableASCII")

	return e
}
func MultibyteEntry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp("[^\x00-\x7F]", "Please input Multibyte")

	return e
}
func DataURIEntry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp(`^data:((?:\w+\/(?:([^;]|;[^;]).)+)?)`, "Please input DataURI")

	return e
}
func LatitudeEntry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp("^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$", "Please input Latitude")

	return e
}
func LongitudeEntry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp("^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$", "Please input Longitude")

	return e
}
func SSNEntry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp(`^[0-9]{3}[ -]?(0[1-9]|[1-9][0-9])[ -]?([1-9][0-9]{3}|[0-9][1-9][0-9]{2}|[0-9]{2}[1-9][0-9]|[0-9]{3}[1-9])$`, "Please input SSN")

	return e
}
func HostnameRFC952Entry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp(`^[a-zA-Z]([a-zA-Z0-9\-]+[\.]?)*[a-zA-Z0-9]$`, "Please input HostnameRFC952")

	return e
}
func HostnameRFC1123Entry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp(`^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?$`, "Please input HostnameRFC1123")

	return e
}
func FqdnRFC1123Entry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp(`^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\.?$`, "Please input FqdnRFC1123")

	return e
}
func BtcAddressEntry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp(`^[13][a-km-zA-HJ-NP-Z1-9]{25,34}$`, "Please input BtcAddress")

	return e
}
func BtcAddressUpperBech32Entry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp(`^BC1[02-9AC-HJ-NP-Z]{7,76}$`, "Please input BtcAddressUpperBech32")

	return e
}
func BtcAddressLowerBech32Entry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp(`^bc1[02-9ac-hj-np-z]{7,76}$`, "Please input BtcAddressLowerBech32")

	return e
}
func EthAddressEntry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp(`^0x[0-9a-fA-F]{40}$`, "Please input EthAddress")

	return e
}
func EthAddressUpperEntry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp(`^0x[0-9A-F]{40}$`, "Please input EthAddressUpper")

	return e
}
func EthAddressLowerEntry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp(`^0x[0-9a-f]{40}$`, "Please input EthAddressLower")

	return e
}
func URLEncodedEntry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp(`^(?:[^%]|%[0-9A-Fa-f]{2})*$`, "Please input URLEncoded")

	return e
}
func HTMLEncodedEntry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp(`&#[x]?([0-9a-fA-F]{2})|(&gt)|(&lt)|(&quot)|(&amp)+[;]?`, "Please input HTMLEncoded")

	return e
}
func HTMLEntry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp(`<[/]?([a-zA-Z]+).*?>`, "Please input HTML")

	return e
}
func SplitParamsEntry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp(`'[^']*'|\S+`, "Please input SplitParams")

	return e
}
func BicEntry() *wdg.Entry {
	e := wdg.NewEntry()
	e.Validator = validation.NewRegexp(`^[A-Za-z]{6}[A-Za-z0-9]{2}([A-Za-z0-9]{3})?$`, "Please input Bic")

	return e
}
