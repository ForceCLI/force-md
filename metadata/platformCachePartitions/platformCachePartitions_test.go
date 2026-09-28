package platformCachePartitions

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ForceCLI/force-md/internal"
)

const partitionXML = `<?xml version="1.0" encoding="UTF-8"?>
<PlatformCachePartition xmlns="http://soap.sforce.com/2006/04/metadata">
    <description>Platform cache partition for tests.</description>
    <isDefaultPartition>false</isDefaultPartition>
    <masterLabel>aerreg</masterLabel>
    <platformCachePartitionTypes>
        <allocatedCapacity>1</allocatedCapacity>
        <allocatedPartnerCapacity>0</allocatedPartnerCapacity>
        <allocatedPurchasedCapacity>0</allocatedPurchasedCapacity>
        <allocatedTrialCapacity>0</allocatedTrialCapacity>
        <cacheType>Organization</cacheType>
    </platformCachePartitionTypes>
    <platformCachePartitionTypes>
        <allocatedCapacity>2</allocatedCapacity>
        <cacheType>Session</cacheType>
    </platformCachePartitionTypes>
</PlatformCachePartition>
`

func TestOpenParsesPartitionAndCacheTypes(t *testing.T) {
	path := filepath.Join(t.TempDir(), "aerreg.cachePartition-meta.xml")
	if err := os.WriteFile(path, []byte(partitionXML), 0644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}
	p, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if p.MasterLabel.Text != "aerreg" {
		t.Errorf("masterLabel = %q", p.MasterLabel.Text)
	}
	if p.Description == nil || p.Description.Text != "Platform cache partition for tests." {
		t.Errorf("description = %+v", p.Description)
	}
	if p.IsDefaultPartition == nil || p.IsDefaultPartition.ToBool() {
		t.Errorf("isDefaultPartition = %+v", p.IsDefaultPartition)
	}
	if len(p.PlatformCachePartitionTypes) != 2 {
		t.Fatalf("expected 2 platformCachePartitionTypes, got %d", len(p.PlatformCachePartitionTypes))
	}
	org := p.PlatformCachePartitionTypes[0]
	if org.CacheType.Text != "Organization" || org.AllocatedCapacity == nil || org.AllocatedCapacity.Text != "1" {
		t.Errorf("first type = %+v", org)
	}
	if session := p.PlatformCachePartitionTypes[1]; session.CacheType.Text != "Session" || session.AllocatedPartnerCapacity != nil {
		t.Errorf("second type = %+v", session)
	}
}

func TestPartitionIsRegisteredByRootElement(t *testing.T) {
	if _, ok := internal.TypeRegistry[NAME]; !ok {
		t.Fatalf("%s is not registered", NAME)
	}
}
