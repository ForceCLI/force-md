package platformCachePartitions

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "PlatformCachePartition"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type PlatformCachePartition struct {
	metadata.MetadataInfo
	XMLName                     xml.Name                     `xml:"PlatformCachePartition"`
	Xmlns                       string                       `xml:"xmlns,attr"`
	Description                 *TextLiteral                 `xml:"description"`
	IsDefaultPartition          *BooleanText                 `xml:"isDefaultPartition"`
	MasterLabel                 TextLiteral                  `xml:"masterLabel"`
	PlatformCachePartitionTypes []PlatformCachePartitionType `xml:"platformCachePartitionTypes"`
}

// PlatformCachePartitionType allocates the partition's capacity for one
// cache type (Session or Organization).
type PlatformCachePartitionType struct {
	AllocatedCapacity          *IntegerText `xml:"allocatedCapacity"`
	AllocatedPartnerCapacity   *IntegerText `xml:"allocatedPartnerCapacity"`
	AllocatedPurchasedCapacity *IntegerText `xml:"allocatedPurchasedCapacity"`
	AllocatedTrialCapacity     *IntegerText `xml:"allocatedTrialCapacity"`
	CacheType                  TextLiteral  `xml:"cacheType"`
}

func (c *PlatformCachePartition) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *PlatformCachePartition) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*PlatformCachePartition, error) {
	p := &PlatformCachePartition{}
	return p, metadata.ParseMetadataXml(p, path)
}
