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

## Composite 
Es un patrón estructural que te permite tratar a los objetos individuales y a los grupos de objetos de la misma manera.

En Go, el patrón Composite se implementa creando una interfaz común para los objetos individuales y los grupos de objetos. Esta interfaz común permite a los clientes tratar a los objetos individuales y a los grupos de objetos de la misma manera, lo que simplifica el código y lo hace más fácil de mantener.

### Aquí te proporciono algunos ejemplos de situaciones en las que puede ser beneficioso utilizar:
- Un ejemplo común de uso del patrón de diseño Composite en Go es cuando necesitas trabajar con una jerarquía de objetos que forman una estructura de árbol. En este caso, puedes usar el patrón Composite para tratar a los nodos individuales del árbol y a los subárboles como si fueran objetos individuales.
- Otro ejemplo de uso del patrón Composite en Go es cuando necesitas aplicar una operación a una colección de objetos, pero algunos de los objetos pueden ser colecciones en sí mismos. En este caso, puedes usar el patrón Composite para tratar a todos los objetos como si fueran individuales, independientemente de si son objetos individuales o colecciones.

Al usar el patrón Composite en Go, puedes escribir un código más genérico y reutilizable, lo que facilita la incorporación de nuevas funcionalidades o cambios en la estructura de objetos. Además, puedes evitar la duplicación de código y reducir la complejidad del código al tratar todos los objetos de la misma manera.
Es importante tener en cuenta que, al igual que con cualquier patrón de diseño, el patrón Composite no es la solución a todos los problemas de diseño. Debes considerar cuidadosamente si el patrón Composite es la mejor opción para tu problema específico antes de implementarlo.

## Decorator
El patrón de diseño Decorator es utilizado para añadir funcionalidades adicionales a un objeto existente en tiempo de ejecución, sin modificar la estructura básica del objeto. En Go, se puede implementar utilizando interfaces y tipos que implementan dichas interfaces.
El patrón Decorator se compone de dos partes principales: el componente y el decorador. El componente es la clase base que se desea decorar, y el decorador es una clase que envuelve al componente, añadiendo o modificando su comportamiento.
Para implementar el patrón Decorator en Go, primero debes crear una interfaz que defina el comportamiento común a todos los componentes y decoradores. Luego, debes crear una clase base que implemente esta interfaz. A continuación, puedes crear clases decoradoras que envuelvan la clase base, añadiendo funcionalidades adicionales.

Una situación común en la que se puede implementar el patrón Decorator es cuando tienes un objeto con un comportamiento básico, pero necesitas añadir funcionalidades adicionales de manera dinámica, sin tener que modificar la clase original. Por ejemplo, puedes usar el patrón Decorator para añadir funcionalidades de logging a una aplicación, sin tener que modificar la lógica de la aplicación original.

## Observer 
El patrón de diseño Observer en Go permite a un objeto observar a otro objeto y recibir notificaciones cuando el objeto observado cambia su estado. Es decir, cuando el objeto observado cambia, notifica automáticamente a sus observadores registrados.

El patrón Observer se compone de dos partes principales: el objeto observado (también conocido como sujeto o publicador) y los objetos observadores (también conocidos como suscriptores o listeners). El objeto observado mantiene una lista de sus observadores registrados y notifica automáticamente a todos ellos cuando cambia su estado.

Para implementar el patrón Observer en Go, primero debes definir una interfaz que defina los métodos necesarios para el objeto observado y los observadores. Luego, debes crear una clase observada que implemente esta interfaz y mantenga una lista de sus observadores registrados. Cuando el estado de la clase observada cambia, debe notificar a todos sus observadores registrados. Los objetos observadores pueden registrarse para recibir notificaciones y deben implementar la misma interfaz que la clase observada.

- Una situación común en la que se puede implementar el patrón Observer es cuando tienes un objeto cuyo estado cambia con frecuencia y necesitas notificar a otros objetos cuando esto sucede. Por ejemplo, puedes usar el patrón Observer para notificar a diferentes partes de una aplicación cuando ocurre un evento específico, como cuando se actualiza una base de datos o se recibe un mensaje de un servidor.

## Chain Of Responsability 
Es un patrón de comportamiento que te permite encadenar varios objetos de procesamiento que manejan una solicitud, y cada objeto decide si puede manejar la solicitud o si debe pasarla al siguiente objeto en la cadena.

En el patrón Chain of Responsibility, cada objeto procesador tiene un puntero al siguiente objeto en la cadena. Si un objeto no puede manejar la solicitud, la pasa al siguiente objeto en la cadena. Este proceso continúa hasta que la solicitud es manejada o hasta que se alcanza el final de la cadena.

### Aquí te proporciono algunos ejemplos de situaciones en las que puede ser beneficioso utilizar:
- Un sistema de aprobación de gastos en una empresa, donde cada solicitud de gasto puede ser manejada por diferentes niveles de autoridad. Por ejemplo, las solicitudes menores a $100 pueden ser aprobadas por un gerente de departamento, mientras que las solicitudes mayores a $1000 deben ser aprobadas por el director financiero.
- Un sistema de detección de fraudes en una aplicación de comercio electrónico, donde las transacciones sospechosas pueden ser manejadas por diferentes reglas de negocio. Por ejemplo, una transacción que proviene de un país diferente al de la dirección de envío puede ser manejada por una regla de negocio específica para detectar fraudes internacionales.
- Un sistema de manejo de errores en una aplicación, donde los errores pueden ser manejados por diferentes módulos o componentes. Por ejemplo, un error en la capa de presentación puede ser manejado por un componente específico para manejar errores de interfaz de usuario, mientras que un error en la capa de base de datos puede ser manejado por un componente específico para manejar errores de acceso a datos.

En resumen, el patrón de diseño Chain of Responsibility es útil en situaciones donde se necesita procesar una serie de solicitudes o eventos que pueden ser manejados por diferentes objetos, y se desea evitar acoplar estos objetos directamente entre sí.

## Otros patrones de diseño
En este repositorio también se incluirán ejemplos y explicaciones de otros patrones de diseño comunes en Go, como Factory Method, Builder, Decorator, entre otros.

