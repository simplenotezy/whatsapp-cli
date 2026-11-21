package whatsapp

import (
	"testing"

	"go.mau.fi/whatsmeow/proto/waCompanionReg"
	wastore "go.mau.fi/whatsmeow/store"
)

func TestApplyDeviceNameNamesTheLinkedDevice(t *testing.T) {
	applyDeviceName("Faktum")
	if got := wastore.DeviceProps.GetOs(); got != "Faktum" {
		t.Fatalf("Os = %q, want Faktum", got)
	}
	if got := wastore.DeviceProps.GetPlatformType(); got != waCompanionReg.DeviceProps_DESKTOP {
		t.Fatalf("PlatformType = %v, want DESKTOP", got)
	}
}

func TestApplyDeviceNameKeepsTheDefaultWithoutAName(t *testing.T) {
	before := wastore.DeviceProps.GetOs()
	applyDeviceName("  ")
	if got := wastore.DeviceProps.GetOs(); got != before {
		t.Fatalf("Os = %q, want %q", got, before)
	}
}
