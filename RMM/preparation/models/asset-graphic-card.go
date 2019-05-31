package models

// GraphicCard defines the structure of a graphic card as part of new asset model
type GraphicCard struct {
	DisplayArray   string `json:"displayArray,omitempty" cql:"display_array"`
	ChipsetModel   string `json:"chipsetModel,omitempty" cql:"chipset_model"`
	DisplayType    string `json:"displayType,omitempty" cql:"display_type"`
	DisplayBus     string `json:"displayBus,omitempty" cql:"display_bus"`
	PciLaneWidth   string `json:"pciLaneWidth,omitempty" cql:"pci_lane_width"`
	VRAM           string `json:"vRAM,omitempty" cql:"vram"`
	Vendor         string `json:"vendor,omitempty" cql:"vendor"`
	DisDeviceID    string `json:"disDeviceID,omitempty" cql:"dis_device_id"`
	DispRevisionID string `json:"disRevisionID,omitempty" cql:"dis_revision_id"`
	RomVer         string `json:"romVer,omitempty" cql:"rom_ver"`
	EfiDrvVer      string `json:"efiDrvVer,omitempty" cql:"efi_drv_ver"`
}
