// Package builder_test realiza pruebas en el paquete Builder
package builder_test

// Se importa el paquete "Builder" creado para este test y el paquete "testing" que sirve para pruebas en Go
import (
	builder "Builder"
	"testing"
)

// La siguiente función es una prueba de unidad que evalúa la función BuildMessage()
func TestSender_BuildMessage(t *testing.T) {
	// Se crea una nueva instancia de JSONMessageBuilder(), que implementa el interface Builder, para construir y enviar mensajes en formato JSON

	xml := &builder.XMLMessageBuilder{}
	// Se crea una nueva instancia de Sender() que es el objeto que envía el mensaje, el cual es configurado para que use el constructor XMLMessageBuilder

	sender := &builder.Sender{}

	sender.SetBuilder(xml)
	// Se llama a la función BuildMessage() para construir y enviar el mensaje, los argumentos son un título y el cuerpo del mensaje
	msg, err := sender.BuildMessage("Gophers", "hello world")
	if err != nil {
		// Si ocurre algún error en el proceso de construcción o envío del mensaje, saltará este error
		t.Fatalf("BuildMessage() have a error: %v", err)
	}
	// Si todo está correcto, se escribe el cuerpo del mensaje en el log
	t.Log(string(msg.Body))
}
