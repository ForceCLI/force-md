package metadata

import "testing"

func TestPathWithinDirectory(t *testing.T) {
	tests := []struct {
		path string
		dir  string
		want string
		ok   bool
	}{
		{"src/reports/Folder/Report.report", "reports", "Folder/Report.report", true},
		{"/home/user/src/reports/Folder/Sub/Report.report", "reports", "Folder/Sub/Report.report", true},
		{"src/reports/Joe_File/RegistrationsDBreports/Report.report", "reports", "Joe_File/RegistrationsDBreports/Report.report", true},
		{"src/reports/Joe_File/RegistrationsDBreports-meta.xml", "reports", "Joe_File/RegistrationsDBreports-meta.xml", true},
		{"/home/user/reports/src/reports/Folder/Report.report", "reports", "Folder/Report.report", true},
		{"src/objects/Account.object", "reports", "", false},
		{"src/reports", "reports", "", false},
	}
	for _, test := range tests {
		got, ok := PathWithinDirectory(test.path, test.dir)
		if got != test.want || ok != test.ok {
			t.Errorf("PathWithinDirectory(%q, %q) = %q, %v; want %q, %v", test.path, test.dir, got, ok, test.want, test.ok)
		}
	}
}
