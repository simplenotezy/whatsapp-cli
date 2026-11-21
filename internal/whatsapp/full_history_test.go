package whatsapp

import (
	"testing"

	wastore "go.mau.fi/whatsmeow/store"
)

func TestRequestFullHistoryAsksThePhoneForEveryChatsWholeHistory(t *testing.T) {
	requestFullHistory()
	if !wastore.DeviceProps.GetRequireFullSync() {
		t.Fatal("RequireFullSync = false, want true")
	}
	config := wastore.DeviceProps.GetHistorySyncConfig()
	if got := config.GetFullSyncDaysLimit(); got != fullHistoryDays {
		t.Fatalf("FullSyncDaysLimit = %d, want %d", got, fullHistoryDays)
	}
	if config.FullSyncSizeMbLimit != nil {
		t.Fatalf("FullSyncSizeMbLimit = %d, want unset (no limit)", config.GetFullSyncSizeMbLimit())
	}
	if got := config.GetStorageQuotaMb(); got < 10240 {
		t.Fatalf("StorageQuotaMb = %d, want at least whatsmeow's 10240", got)
	}
}
