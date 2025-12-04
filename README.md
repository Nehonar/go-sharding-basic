## Descripcion del sistema

Ejercicio en Golang para probar un sistema de Sharding de base de datos, una unica API que guarda un usuario y contraseña en base de datos.

## Entrada/salida

Entrada:

- El sistema recibe u¡los datos del usuario (nombre y contraseña)

Salida:

- El sistema devuelve un mensaje cuando el guardado de datos es correcto.
- El sistema devuelve un mensaje de error con codigo de error que hubo al guardar los datos.

## Endpoint

- Metodo: POST
- Ruta: /create-user
- JSON entrada
- JSON salida con exito
- JSON salida error
- Codigos: 201 y 500