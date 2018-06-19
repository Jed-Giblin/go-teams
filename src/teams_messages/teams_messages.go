package teams_messages

type NewWebSocket struct {
	Name string `json:"name"`
	TargetURL string `json:"targetUrl"`
	Resource string `json:"resource"`
	Event string `json:"event"`
	Filter string `json:"filter"`
	Secret string `json:"secret"`
}

type ExistingWebSocket struct {
	Id string `json:"id"`
	Name string `json:"name"`
	TargetURL string `json:"targetUrl"`
	Resource string `json:"resource"`
	Event string `json:"event"`
	Filter string `json:"filter"`
	Secret string `json:"secret"`
	Status string `json:"status"`
	CreatedAt string `json:"created"`
}

func ( m NewWebSocket ) isEqual(socket ExistingWebSocket ) bool {
	if m.Name != socket.Name || m.TargetURL != socket.TargetURL || m.Resource != socket.Resource ||
		m.Event != socket.Event || m.Filter != socket.Filter || m.Secret != socket.Secret {
			return false
	} else {
		return true
	}
}