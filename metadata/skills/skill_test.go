package skills

import (
	"os"
	"path/filepath"
	"testing"
)

const sampleSkill = `<?xml version="1.0" encoding="UTF-8"?>
<Skill xmlns="http://soap.sforce.com/2006/04/metadata">
    <assignments>
        <profiles>
            <profile>admin</profile>
        </profiles>
        <users>
            <user>first@example.com</user>
            <user>second@example.com</user>
        </users>
    </assignments>
    <description>Routes support &amp; billing cases</description>
    <label>Support Team</label>
</Skill>
`

func TestOpen(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "Support_Team.skill")
	if err := os.WriteFile(path, []byte(sampleSkill), 0644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	s, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if got := string(s.Type()); got != "Skill" {
		t.Errorf("Type = %q, want Skill", got)
	}
	if got := s.Label.String(); got != "Support Team" {
		t.Errorf("Label = %q, want Support Team", got)
	}
	if got := s.Description.String(); got != "Routes support & billing cases" {
		t.Errorf("Description = %q, want %q", got, "Routes support & billing cases")
	}
	if s.Assignments == nil || s.Assignments.Profiles == nil || s.Assignments.Users == nil {
		t.Fatalf("Assignments = %+v, want profiles and users", s.Assignments)
	}
	if got := len(s.Assignments.Profiles.Profile); got != 1 || s.Assignments.Profiles.Profile[0].String() != "admin" {
		t.Errorf("profiles = %+v, want [admin]", s.Assignments.Profiles.Profile)
	}
	if got := len(s.Assignments.Users.User); got != 2 || s.Assignments.Users.User[1].String() != "second@example.com" {
		t.Errorf("users = %+v, want two users ending with second@example.com", s.Assignments.Users.User)
	}
}
