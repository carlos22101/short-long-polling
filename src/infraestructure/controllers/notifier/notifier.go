package notifier

import "sync"

// Notifier estructura para manejar las notificaciones de cambios
type Notifier struct {
	mu      sync.Mutex
	clients []chan struct{}
}

// NewNotifier crea una nueva instancia de Notifier
func NewNotifier() *Notifier {
	return &Notifier{}
}

// Subscribe permite a un cliente esperar cambios
func (n *Notifier) Subscribe() <-chan struct{} {
	n.mu.Lock()
	defer n.mu.Unlock()

	ch := make(chan struct{}, 1)
	n.clients = append(n.clients, ch)
	return ch
}

// NotifyChanges notifica a todos los clientes suscritos
func (n *Notifier) NotifyChanges() {
	n.mu.Lock()
	defer n.mu.Unlock()

	for _, ch := range n.clients {
		select {
		case ch <- struct{}{}:
		default:
		}
	}
	n.clients = nil // Limpiar la lista después de notificar
}
