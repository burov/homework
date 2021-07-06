package testdata

var TestDomain = "acme.com"
var CrewData = []struct {
	InName   string
	OutName  string
	OutEmail string
	Err      error
}{
	{InName: "Miranda Lawson", OutName: "miranda.lawson", OutEmail: "miranda.lawson@acme.com", Err: nil},
	{InName: "Liara T'Soni", OutName: "liara.t.soni", OutEmail: "liara.t.soni@acme.com", Err: nil},
	{InName: "EDI", OutName: "edi", OutEmail: "edi@acme.com", Err: nil},
	{InName: "Tali'Zorah nar Rayya", OutName: "tali.zorah.nar.rayya", OutEmail: "tali.zorah.nar.rayya@acme.com", Err: nil},
}

var SameNameData = []struct {
	BaseName string
	Err      bool
	InName   string
	InNum    int
	Names    []string
	OutName  string
}{
	{
		BaseName: "legion",
		Err:      false,
		InName:   "legion",
		InNum:    0,
		Names:    nil,
		OutName:  "legion",
	},
	{
		BaseName: "legion",
		Err:      false,
		InName:   "legion",
		InNum:    99,
		Names:    []string{"legion", "legion.01", "legion.98"},
		OutName:  "legion.99",
	},
	{
		BaseName: "legion",
		Err:      false,
		InName:   "legion",
		InNum:    9,
		Names:    []string{"legion", "legion.01", "legion.08"},
		OutName:  "legion.09",
	},
	{
		BaseName: "legion",
		Err:      true,
		InName:   "legion",
		InNum:    100,
		Names:    []string{"legion", "legion.01", "legion.99"},
		OutName:  "legion.100",
	},
}
