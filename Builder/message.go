package builder

// Estructura para representar un mensaje. Tiene dos campos: "Recipient" y "Text".
type Message struct {
	Recipient string `json:"recipient" xml:"recipient"`
	// Etiqueta "json" y \"xml" indica que los campos serán serializados en esos formatos.
	Text string `json:"text" xml:"text"`
}

// Estructura para representar un mensaje serializado. Tiene dos campos: "Body" y "Format".
type MessageRepresented struct {
	Body   []byte
	Format string
}
