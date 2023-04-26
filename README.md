# **Patrones de Diseño en Go**
## Singleton
El patrón de diseño Singleton es un patrón de creación que asegura que solo existe una instancia de una determinada clase, y que provee un punto global de acceso a ella.

### ¿En qué casos puedo aplicar el patrón Singleton?
El patrón Singleton es útil en aquellos casos en los que se requiere tener una única instancia global de una clase. Algunos casos de uso comunes incluyen:

- Objetos de configuración
- Objetos de conexión a bases de datos
- Gestores de recursos compartidos (por ejemplo, un gestor de logs)
- Objetos que manejan el acceso a recursos limitados (por ejemplo, una licencia de software)

 se presenta un ejemplo de implementación del patrón Singleton en Go, con el fin de brindar una guía útil para su aplicación en diferentes proyectos. Además, se incluyen comentarios y explicaciones detalladas para entender cómo funciona el patrón y cómo se puede adaptar a diferentes contextos y necesidades.

### Factory
Es un patrón creacional que proporciona una manera de crear objetos sin tener que especificar la clase exacta del objeto que se creará. En lugar de crear los objetos directamente, se delega la creación a una clase Factory que selecciona la subclase apropiada para crear el objeto.

###  Algunos casos donde puede ser útil utilizar el patrón Factory son:

- Cuando se tienen diferentes tipos de objetos que deben ser creados en tiempo de ejecución y se desea delegar la creación a una clase específica.
- Cuando se desea reducir la complejidad del código al separar la creación de objetos de su uso y permitir la modularidad.
- Cuando se desea ocultar la complejidad de la creación de objetos al cliente y proporcionar una interfaz sencilla y común para crear objetos.

En resumen, el patrón Factory en Go es útil en situaciones donde se necesitan crear objetos de diferentes subclases relacionadas, se desea simplificar la creación de objetos y proporcionar una interfaz común para la creación de objetos de varias subclases.

## Otros patrones de diseño
En este repositorio también se incluirán ejemplos y explicaciones de otros patrones de diseño comunes en Go, como Factory Method, Builder, Decorator, entre otros.
