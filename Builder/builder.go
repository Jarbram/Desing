package builder

// Define la interfaz para construir mensajes
type MessageBuilder interface { // Define el método para establecer el destinatario del mensaje
	SetRecipient(recipient string) MessageBuilder
	// Define el método para establecer el contenido del mensaje
	SetMessage(message string) MessageBuilder
	// Define el método para construir el mensaje completo
	Build() (*MessageRepresented, error)
}
