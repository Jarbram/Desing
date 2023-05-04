package main

import "fmt"

// Data es la interfaz que define la operación de obtener datos
type Data interface {
	GetData() string
}

// TextData es una implementación de Data que almacena texto
type TextData struct {
	data string
}

func (t *TextData) GetData() string {
	return t.data
}

// BinaryData es una implementación de Data que almacena bytes
type BinaryData struct {
	data []byte
}

func (b *BinaryData) GetData() string {
	return fmt.Sprintf("%v", b.data)
}

// Decorator es la interfaz que define la operación de decorar los datos
type Decorator interface {
	Data
}

// Compressor es un decorador que comprime los datos
type Compressor struct {
	inner Data
}

func (c *Compressor) GetData() string {
	return fmt.Sprintf("Compressed: %v", c.inner.GetData())
}

// Encryptor es un decorador que encripta los datos
type Encryptor struct {
	inner Data
}

func (e *Encryptor) GetData() string {
	return fmt.Sprintf("Encrypted: %v", e.inner.GetData())
}

// Formatter es un decorador que formatea los datos
type Formatter struct {
	inner Data
}

func (f *Formatter) GetData() string {
	return fmt.Sprintf("Formatted: %v", f.inner.GetData())
}

func main() {
	// Creamos un objeto de datos de texto
	textData := &TextData{data: "Hola mundo!"}

	// Aplicamos diferentes decoradores al objeto de datos de texto
	compressedTextData := &Compressor{inner: textData}
	encryptedTextData := &Encryptor{inner: compressedTextData}
	formattedTextData := &Formatter{inner: encryptedTextData}

	// Obtenemos el resultado final de los datos de texto con todos los decoradores aplicados
	result := formattedTextData.GetData()

	fmt.Println(result)

	// Creamos un objeto de datos binarios
	binaryData := &BinaryData{data: []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f}}

	// Aplicamos diferentes decoradores al objeto de datos binarios
	encryptedBinaryData := &Encryptor{inner: binaryData}
	formattedBinaryData := &Formatter{inner: encryptedBinaryData}

	// Obtenemos el resultado final de los datos binarios con todos los decoradores aplicados
	result = formattedBinaryData.GetData()

	fmt.Println(result)
}
