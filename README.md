# Go JWT Authentication Template

Este repositorio es una plantilla base para implementar un sistema de autenticación utilizando **JSON Web Tokens (JWT)** en **Go**. Está diseñado para ser reutilizable y servir como punto de partida para futuros proyectos.

## Características

- Autenticación basada en **JSON Web Tokens (JWT)**.
- Implementación en **Go**.
- Uso de **PostgreSQL** como base de datos.

## Requisitos previos

Antes de comenzar, asegúrate de tener instalados:

- [Go](https://go.dev/dl/) (versión 1.19 o superior recomendada).
- [PostgreSQL](https://www.postgresql.org/).
- [Git](https://git-scm.com/).

## Cómo clonar y configurar el proyecto

Sigue los pasos a continuación para clonar este repositorio y configurarlo:

1. **Clonar el repositorio**
   ```bash
   git clone https://github.com/ericheinitz/go-jwt.git
   cd go-jwt
   ```

2. **Configurar las variables de entorno**

   Renombra el archivo `.env.example` a `.env`:
   ```bash
   cp .env.example .env
   ```

   Edita el archivo `.env` y ajusta los valores según tu entorno. Aquí tienes un ejemplo del contenido:
   ```dotenv
   PORT=3000
   CORS_ORIGIN="http://localhost:5173"
   JWT_SECRET="TU_PROPIO_SECRETO"
   DB="host=localhost user=pguser password=pgpass dbname=go_jwt port=5432 sslmode=disable"
   ```

3. **Instalar las dependencias**
   Ejecuta el siguiente comando para instalar las dependencias necesarias:
   ```bash
   go mod tidy
   ```

4. **Ejecutar las migraciones**
   Si tu proyecto incluye migraciones para inicializar la base de datos, asegúrate de ejecutarlas antes de iniciar el servidor.

5. **Iniciar el servidor**
   Inicia el servidor con:
   ```bash
   go run main.go
   ```

   El servidor estará disponible en `http://localhost:3000` (o en el puerto configurado en el archivo `.env`).

## Librerías utilizadas

Este proyecto utiliza las siguientes librerías:

- **[gorm.io](https://gorm.io/docs/index.html)**: Un ORM para interactuar con la base de datos PostgreSQL de manera eficiente.
- **[gin-gonic](https://gin-gonic.com/en/docs/quickstart/)**: Un framework web rápido y minimalista para construir aplicaciones HTTP en Go.
- **[golang.org/x/crypto](https://pkg.go.dev/golang.org/x/crypto#section-readme)**: Para el hash y la verificación de contraseñas.
- **[github.com/golang-jwt/jwt/v5](https://pkg.go.dev/github.com/golang-jwt/jwt/v5#section-readme)**: Para la generación y validación de tokens JWT.
- **[github.com/joho/godotenv](https://github.com/joho/godotenv)**: Para cargar variables de entorno desde un archivo `.env`.

## Contribución

Si deseas contribuir a este proyecto, por favor sigue los pasos a continuación:

1. Haz un fork del repositorio.
2. Crea una nueva rama para tu funcionalidad o corrección de errores (`git checkout -b feature/nueva-funcionalidad`).
3. Realiza tus cambios y genera un commit (`git commit -m "Agrega nueva funcionalidad"`).
4. Envía tus cambios al repositorio remoto (`git push origin feature/nueva-funcionalidad`).
5. Abre un Pull Request en GitHub.

## Licencia

Este proyecto está bajo la licencia MIT. Consulta el archivo `LICENSE` para más detalles.
