package api

// VMCategory represents VM category
type VMCategory string

// APIModel complies with the ARM model of
// resource definition in a JSON template.
type APIModel struct {
	VMCategory VMCategory  `json:"vmCategory"`
	Location   string      `json:"location"`
	Properties *Properties `json:"properties,omitempty"`

	VMConfigurator VMConfigurator
}

// OSType represents OS type
type OSType string

// OSName represents pre-set OS name
type OSName string

// VMProfile represents the definition of a VM
type VMProfile struct {
	Name        string   `json:"name"`
	OSType      OSType   `json:"osType"`
	OSName      OSName   `json:"osName"`
	OSDiskType  string   `json:"osDiskType"`
	OSImage     *OSImage `json:"osImage,omitempty"`
	OSDisk      *OSDisk  `json:"osDisk,omitempty"`
	DiskSizesGB []int    `json:"diskSizesGB,omitempty"`
	VMSize      string   `json:"vmSize"`
	Ports       []int    `json:"ports,omitempty" validate:"dive,min=1,max=65535"`
	HasDNSName  bool     `json:"hasDNSName"`
	SecureBoot  *bool    `json:"secureBoot,omitempty"`
	VTPM        *bool    `json:"vTPMEnabled,omitempty"`
}

// Properties represents the ACS cluster definition
type Properties struct {
	VnetProfile        *VnetProfile        `json:"vnetProfile"`
	VMProfile          *VMProfile          `json:"vmProfile"`
	LinuxProfile       *LinuxProfile       `json:"linuxProfile,omitempty"`
	WindowsProfile     *WindowsProfile     `json:"windowsProfile,omitempty"`
	DiagnosticsProfile *DiagnosticsProfile `json:"diagnosticsProfile,omitempty"`
}

// OSImage represents OS Image from Azure Image Gallery
type OSImage struct {
	URL       string `json:"url,omitempty"`
	Publisher string `json:"publisher"`
	Offer     string `json:"offer"`
	SKU       string `json:"sku"`
	Version   string `json:"version,omitempty"`
}

// OSDisk represents Managed OS Disk
type OSDisk struct {
	VHD              string `json:"vhd_url"`
	VMGS             string `json:"vmgs_url"`
	StorageAccountID string `json:"storage_account_id"`
}

// LinuxProfile represents the linux parameters passed to the cluster
type LinuxProfile struct {
	AdminUsername string       `json:"adminUsername" validate:"required"`
	AdminPassword string       `json:"adminPassword"`
	SSHPubKeys    []*PublicKey `json:"sshPublicKeys"`
}

// WindowsProfile represents the windows parameters passed to the cluster
type WindowsProfile struct {
	AdminUsername string `json:"adminUsername" validate:"required"`
	AdminPassword string `json:"adminPassword" validate:"required"`
	SSHPubKey     string `json:"sshPublicKey,omitempty"`
}

// VnetProfile represents the definition of a vnet
type VnetProfile struct {
	VnetResourceGroup string `json:"vnetResourceGroup,omitempty"`
	VnetName          string `json:"vnetName,omitempty"`
	VnetAddress       string `json:"vnetAddress,omitempty"`
	SubnetName        string `json:"subnetName,omitempty"`
	SubnetAddress     string `json:"subnetAddress,omitempty"`
}

// DiagnosticsProfile contains settings to on/off boot diagnostics collection
// in RD Host
type DiagnosticsProfile struct {
	Enabled             bool   `json:"true"`
	StorageAccountName  string `json:"storageAccountName"`
	IsNewStorageAccount bool   `json:"isNewStorageAccount"`
}

// PublicKey contains puvlic SSH key
type PublicKey struct {
	KeyData string `json:"keyData"`
}

// IsCustomVNET returns true if the customer brought their own VNET
func (p *VnetProfile) IsCustomVNET() bool {
	return len(p.VnetResourceGroup) > 0 && len(p.VnetName) > 0 && len(p.SubnetName) > 0
}

// HasAzureGalleryImage returns true if Azure Image Gallery is used
func (h *VMProfile) HasAzureGalleryImage() bool {
	return h.OSImage != nil && len(h.OSImage.Publisher) > 0 && len(h.OSImage.Offer) > 0 && len(h.OSImage.SKU) > 0
}

// HasCustomOsImage returns true if there is a custom OS image url specified
func (h *VMProfile) HasCustomOsImage() bool {
	return h.OSImage != nil && len(h.OSImage.URL) > 0
}

// HasAttachedOsDisk returns true if there is an attached OS disk specified
func (h *VMProfile) HasAttachedOsDisk() bool {
	return h.OSDisk != nil && len(h.OSDisk.VHD) > 0
}

// HasAttachedOsDiskVMGS returns true if there is an VMGS specified
func (h *VMProfile) HasAttachedOsDiskVMGS() bool {
	return h.OSDisk != nil && len(h.OSDisk.VMGS) > 0
}

// HasDisks returns true if the customer specified disks
func (h *VMProfile) HasDisks() bool {
	return len(h.DiskSizesGB) > 0
}
