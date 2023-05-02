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

## Factory
Es un patrón creacional que proporciona una manera de crear objetos sin tener que especificar la clase exacta del objeto que se creará. En lugar de crear los objetos directamente, se delega la creación a una clase Factory que selecciona la subclase apropiada para crear el objeto.

###  Algunos casos donde puede ser útil utilizar el patrón Factory son:

- Cuando se tienen diferentes tipos de objetos que deben ser creados en tiempo de ejecución y se desea delegar la creación a una clase específica.
- Cuando se desea reducir la complejidad del código al separar la creación de objetos de su uso y permitir la modularidad.
- Cuando se desea ocultar la complejidad de la creación de objetos al cliente y proporcionar una interfaz sencilla y común para crear objetos.

En resumen, el patrón Factory en Go es útil en situaciones donde se necesitan crear objetos de diferentes subclases relacionadas, se desea simplificar la creación de objetos y proporcionar una interfaz común para la creación de objetos de varias subclases.

## Builder 
El patrón de diseño Builder es útil cuando necesitas construir objetos complejos con muchos componentes y configuraciones diferentes. En Go, puedes implementar este patrón creando un Builder que tenga métodos para configurar y construir el objeto final.

La idea básica es que tengas un objeto Builder que tenga métodos para configurar cada parte del objeto final. Por ejemplo, si estás construyendo un objeto de tipo "Auto", el Builder podría tener métodos para establecer el modelo, el color, el tamaño del motor, la transmisión, etc.

Una vez que hayas configurado todos los componentes necesarios en el Builder, puedes llamar al método "Build" para crear el objeto final. Esto te permite construir objetos complejos sin tener que preocuparte por el orden en que se configuran los componentes o por tener que recordar todos los detalles de configuración.

### Algunos ejemplos de situaciones en las que podrías utilizar este patrón son:
- Cuando necesitas crear objetos con muchas configuraciones diferentes y no quieres tener que recordar todas ellas.
- Cuando los objetos que estás construyendo tienen muchos componentes y quieres simplificar el proceso de construcción.
- Cuando necesitas crear objetos que tengan diferentes variantes o configuraciones, y quieres asegurarte de que cada variante se construya de manera consistente.
- Cuando necesitas separar el proceso de construcción de un objeto de su representación final.

## Prototype
El patrón de diseño Prototype en Go es una técnica que nos permite crear nuevos objetos a partir de la clonación de un objeto ya existente. En lugar de crear un objeto desde cero, se copia un objeto existente y se modifica si es necesario.

### Algunos ejemplos de situaciones en las que podrías utilizar este patrón son:

- Cuando necesitamos crear muchos objetos similares, pero con algunas pequeñas diferencias. En lugar de crear cada objeto desde cero, se puede clonar un objeto existente y modificarlo según sea necesario.

- Cuando queremos reducir la complejidad de la creación de objetos. Si la creación de un objeto implica muchas operaciones complejas, como la inicialización de muchos campos y la conexión a una base de datos, clonar un objeto ya existente puede ser una forma más sencilla y eficiente de crear nuevos objetos.

## Proxy
Es un patrón estructural que te permite proporcionar un sustituto o representante de otro objeto para controlar el acceso a este objeto. El Proxy actúa como intermediario entre el cliente y el objeto real, lo que significa que el cliente interactúa con el Proxy y no con el objeto real directamente.
En Go, el patrón Proxy se implementa mediante la creación de una estructura que actúa como Proxy y que tiene una referencia al objeto real. El Proxy imita la interfaz del objeto real y puede realizar cualquier lógica adicional antes o después de que se llame al objeto real.

### Algunos ejemplos de situaciones en las que podrías utilizar este patrón son:

- Un ejemplo común de uso del patrón de diseño Proxy en Go es cuando necesitas controlar el acceso a un recurso costoso en términos de tiempo o memoria. Por ejemplo, supongamos que tienes una aplicación que necesita acceder a una base de datos remota. Si se realizan muchas solicitudes a la base de datos, esto puede ralentizar la aplicación. En este caso, puedes implementar un Proxy que almacena en caché los resultados de las solicitudes a la base de datos y que solo llama al objeto real cuando los resultados no están disponibles en la caché.

- Otro ejemplo de uso del patrón Proxy en Go es cuando necesitas restringir el acceso a un objeto real. Por ejemplo, supongamos que tienes un objeto que solo puede ser accedido por ciertos usuarios. En este caso, puedes implementar un Proxy que comprueba si el usuario tiene permiso para acceder al objeto antes de permitir el acceso.

## Adapter
En Go, el patrón Adapter se implementa mediante la creación de una estructura que envuelve o adapta el objeto existente y que expone una nueva interfaz. El Adapter redirige las solicitudes de la nueva interfaz al objeto existente y las adapta para que sean comprensibles para ese objeto.

- Un ejemplo común de uso del patrón de diseño Adapter en Go es cuando necesitas usar una biblioteca o clase existente en tu código, pero su interfaz no es compatible con tu código. En este caso, puedes crear un Adapter que envuelve la biblioteca o clase existente y que proporciona una interfaz compatible con tu código.
- Otro ejemplo de uso del patrón Adapter en Go es cuando necesitas hacer que varias clases con interfaces diferentes parezcan iguales para el cliente. En este caso, puedes crear un Adapter para cada clase que adapte su interfaz a una interfaz común para que el cliente pueda trabajar con ellas de la misma manera.

## Facade 
El objetivo principal de Facade es proporcionar una interfaz simple para un conjunto complejo de clases, componentes o subsistemas, que de otra manera sería difícil de entender o utilizar. Este patrón se encuentra dentro de la categoría de patrones estructurales.

La idea detrás del patrón Facade es que proporciona una capa adicional de abstracción para los clientes, que oculta la complejidad del subsistema subyacente. Los clientes pueden interactuar con la capa Facade para acceder a las funcionalidades del subsistema, en lugar de tener que interactuar directamente con las clases y componentes complejos que componen el subsistema.

### Aquí te proporciono algunos ejemplos de situaciones en las que puede ser beneficioso utilizar el patrón de diseño Facade:

- Cuando trabajas con un subsistema complejo: Si estás trabajando con un subsistema que es difícil de entender o tiene demasiadas clases y componentes complejos, entonces el patrón Facade puede ayudarte a ocultar esa complejidad detrás de una interfaz simple.
- Cuando necesitas proporcionar una interfaz unificada: Si necesitas proporcionar una interfaz unificada a varios subsistemas diferentes, entonces puedes utilizar el patrón Facade para crear una interfaz simple que oculte las diferencias entre los subsistemas.
- Cuando necesitas separar la implementación del cliente: Si deseas separar la implementación del cliente de la implementación del subsistema subyacente, entonces puedes utilizar el patrón Facade para proporcionar una capa adicional de abstracción.
- Cuando necesitas reutilizar código existente: Si tienes un conjunto de clases y componentes existentes que se utilizan en diferentes subsistemas, puedes utilizar el patrón Facade para proporcionar una interfaz común para esos subsistemas.

## Bridge
El patrón de diseño Bridge es un patrón estructural que se utiliza para desacoplar una abstracción de su implementación subyacente, lo que permite que ambas puedan variar independientemente. Se logra mediante la creación de dos jerarquías de clases separadas: una para la abstracción y otra para la implementación.

¿Cómo funciona el patrón de diseño Bridge?
El patrón de diseño Bridge se compone de los siguientes elementos:

Abstracción: Define la interfaz para la abstracción, mantiene una referencia a un objeto de la implementación y delega todo el trabajo real a ese objeto.
Implementación: Define la interfaz para la implementación y proporciona una implementación concreta de esa interfaz.
Implementador concreto: Proporciona una implementación concreta de la interfaz de la implementación.
Abstracción refinada: Extiende la interfaz de la abstracción proporcionando funcionalidad adicional que depende de la implementación subyacente.
Implementador refinado: Extiende la interfaz de la implementación proporcionando una funcionalidad adicional que depende de la abstracción.
La idea básica es que la abstracción delega a la implementación cualquier trabajo real que necesite hacer. La implementación proporciona la funcionalidad real, mientras que la abstracción proporciona una interfaz simplificada y unificada para el cliente.

### El patrón de diseño Bridge es útil cuando:

- Quieres separar la interfaz de una clase de su implementación subyacente.
- Quieres que las clases de la interfaz y la implementación varíen independientemente.
- Quieres ocultar la complejidad de una clase detrás de una interfaz más simple.

## Otros patrones de diseño
En este repositorio también se incluirán ejemplos y explicaciones de otros patrones de diseño comunes en Go, como Factory Method, Builder, Decorator, entre otros.


