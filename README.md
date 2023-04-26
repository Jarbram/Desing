# **Patrones de Diseño en Go**
## Singleton
El patrón de diseño Singleton es un patrón de creación que asegura que solo existe una instancia de una determinada clase, y que provee un punto global de acceso a ella.

El patrón Singleton se puede aplicar en situaciones donde necesitamos una única instancia de una clase y queremos asegurarnos de que esa instancia sea la única en toda nuestra aplicación.

En Go, podemos implementar Singleton utilizando variables globales y funciones init(), que se ejecutan automáticamente al inicio de la ejecución del programa. También podemos utilizar estructuras y funciones privadas para asegurar la única instancia de la clase.

En este repositorio se presenta un ejemplo de implementación del patrón Singleton en Go, con el fin de brindar una guía útil para su aplicación en diferentes proyectos. Además, se incluyen comentarios y explicaciones detalladas para entender cómo funciona el patrón y cómo se puede adaptar a diferentes contextos y necesidades.

## Otros patrones de diseño
En este repositorio también se incluirán ejemplos y explicaciones de otros patrones de diseño comunes en Go, como Factory Method, Builder, Decorator, entre otros.
