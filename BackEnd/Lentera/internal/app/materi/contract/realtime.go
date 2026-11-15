package contract

// Broadcaster mengirim event realtime ke satu room.
// Implementasi asli = Hub WebSocket (disambungkan nanti).
type Broadcaster interface {
	Broadcast(room string, event string, data any)
}

// NoopBroadcaster dipakai sementara (dan di test) sebelum Hub siap.
type NoopBroadcaster struct{}

func (NoopBroadcaster) Broadcast(string, string, any) {}
