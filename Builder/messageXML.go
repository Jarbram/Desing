package builder // Define el paquete "builder"

import "encoding/xml" // Importa la biblioteca de encoding/xml

// Define la estructura XMLMessageBuilder
type XMLMessageBuilder struct {
	message Message // Contiene un campo llamado "message" de tipo "Message"
}

// El método SetRecipient establece el campo "Recipient" de la estructura "message"
func (b *XMLMessageBuilder) SetRecipient(recipient string) MessageBuilder {
	b.message.Recipient = recipient // Establece el campo "Recipient" dentro de la estructura "message"
	return b                        // Devuelve el struct "b"
}

// El método SetMessage establece el campo "Text" de la estructura "message"
func (b *XMLMessageBuilder) SetMessage(text string) MessageBuilder {
	b.message.Text = text // Establece el campo "Text" dentro de la estructura "message"
	return b              // Devuelve el struct "b"
}

// El método Build convierte la estructura "XMLMessageBuilder" a una representación de bytes XML
func (b *XMLMessageBuilder) Build() (*MessageRepresented, error) {
	data, err := xml.Marshal(b.message) // Convierte la estructura a una representación de bytes XML
	if err != nil {
		return nil, err // Devuelve un error si falló la conversión
	}
	return &MessageRepresented{Body: data, Format: "XML"}, nil // Devuelve la estructura "MessageRepresented" que contiene la representación de bytes XML y el formato
}
