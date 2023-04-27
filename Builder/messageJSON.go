package builder

import "encoding/json"

// Definir una estructura "JSONMessageBuilder" que almacena el mensaje en construcción.
type JSONMessageBuilder struct {
	message Message
}

// Definir una función "SetRecipient" que establece al destinatario del mensaje.
// Retorna "MessageBuilder" para permitir la construcción en cadena.
func (b *JSONMessageBuilder) SetRecipient(recipient string) MessageBuilder {
	b.message.Recipient = recipient
	return b
}

// Definir una función "SetMessage" que establece el texto del mensaje.
// Retorna "MessageBuilder" para permitir la construcción en cadena.
func (b *JSONMessageBuilder) SetMessage(text string) MessageBuilder {
	b.message.Text = text
	return b
}

// Definir una función "Build" que construye un mensaje JSON y lo retorna como "MessageRepresented"
// Retorna un error si falla la serialización de JSON.
func (b *JSONMessageBuilder) Build() (*MessageRepresented, error) {
	data, err := json.Marshal(b.message)
	if err != nil {
		return nil, err
	}
	return &MessageRepresented{Body: data, Format: "JSON"}, nil
}
