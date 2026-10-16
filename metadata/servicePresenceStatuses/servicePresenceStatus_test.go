package servicePresenceStatuses

import (
	"os"
	"path/filepath"
	"testing"
)

const sampleAvailableStatus = `<?xml version="1.0" encoding="UTF-8"?>
<ServicePresenceStatus xmlns="http://soap.sforce.com/2006/04/metadata">
    <channels>
        <channel>Cases</channel>
        <channel>sfdc_phone</channel>
    </channels>
    <label>Available &amp; Ready</label>
</ServicePresenceStatus>
`

const sampleBusyStatus = `<?xml version="1.0" encoding="UTF-8"?>
<ServicePresenceStatus xmlns="http://soap.sforce.com/2006/04/metadata">
    <label>Busy</label>
</ServicePresenceStatus>
`

func writeFixture(t *testing.T, name, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), name)
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}
	return path
}

func TestOpen(t *testing.T) {
	s, err := Open(writeFixture(t, "Available.servicePresenceStatus", sampleAvailableStatus))
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if got := string(s.Type()); got != "ServicePresenceStatus" {
		t.Errorf("Type = %q, want ServicePresenceStatus", got)
	}
	if got := s.Label.String(); got != "Available & Ready" {
		t.Errorf("Label = %q, want %q", got, "Available & Ready")
	}
	if s.Channels == nil {
		t.Fatal("Channels = nil, want two channels")
	}
	if got := len(s.Channels.Channel); got != 2 {
		t.Fatalf("len(Channels.Channel) = %d, want 2", got)
	}
	if got := s.Channels.Channel[0].String(); got != "Cases" {
		t.Errorf("Channels.Channel[0] = %q, want Cases", got)
	}
	if got := s.Channels.Channel[1].String(); got != "sfdc_phone" {
		t.Errorf("Channels.Channel[1] = %q, want sfdc_phone", got)
	}
}

func TestOpenStatusWithoutChannels(t *testing.T) {
	s, err := Open(writeFixture(t, "Busy.servicePresenceStatus", sampleBusyStatus))
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if got := s.Label.String(); got != "Busy" {
		t.Errorf("Label = %q, want Busy", got)
	}
	if s.Channels != nil {
		t.Errorf("Channels = %+v, want nil", s.Channels)
	}
}
