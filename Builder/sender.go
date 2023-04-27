package builder

// Define la estructura Sender
type Sender struct {
	builder MessageBuilder
}

// SetBuilder asigna un builder a la estructura Sender.
func (s *Sender) SetBuilder(b MessageBuilder) {
	s.builder = b
}

// BuildMessage usa el builder para crear el mensaje.
func (s *Sender) BuildMessage(recipient, message string) (*MessageRepresented, error) { // utiliza el builder para construir el mensaje con el recipient y message
	return s.builder.SetRecipient(recipient).SetMessage(message).Build()
}
