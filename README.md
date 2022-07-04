## Mail App

El siguiente proyecto sirve como ejemplo de una API REST programada en Go que brinda un servicio de correo electronico. 
Se utilizaron las siguientes herramientas:
* [Gin Gonic](https://gin-gonic.com/) - Web framework.
* [MySQL](https://www.mysql.com/) - Base de datos relacional.
* [GORM](https://gorm.io/) - ORM para bases de datos SQL.
* [AWS S3](https://aws.amazon.com/s3/) - Bucket S3.
* [AWS SES](https://aws.amazon.com/ses/) - Simple Email Service.

La aplicacion es capaz de:
* Crear, modificar, eliminar y consultar Usuarios y Empresas.
* Crear y enviar mails con multiples destinatarios y la opcion de adjuntar archivos PDF.
* Consultar un historial de mails enviados y recibidos, y poder borrarlos.
* Guardar todos estos datos en una base de datos (Usuarios, Empresas y Mails).

### Instrucciones:
Clonar repositorio:
```
$ git clone https://github.com/irf98/musical-octo-palm-tree

$ cd musical-octo-palm-tree
```
Luego definir las variables de entorno (podes encontrar mas informacion sobre DB_DSN en los docs de [GORM](https://gorm.io/docs/connecting_to_the_database.html)):
```
$ export DB_DSN="user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

$ export AWS_ACCESS_KEY_ID="aws-keys"

$ export AWS_SECRET_ACCESS_KEY="aws-secret"

$ export AWS_REGION="aws-region"

$ export S3_BUCKET="bucket-name"
```
Para correr de forma local, primero hay que instalar las dependencias:
```
$ go mod tidy
```
Ya puedes correr el siguiente comando (por defecto, el servidor estara activo en localhost:8080):
```
$ go run ./cmd/app/main.go
```
Para crear una imagen docker:
```
$ docker build --tag mail-app .
```
Buscar la imagen y correrla:
```
$ docker images

$ docker run --rm mail-app
```
&nbsp;

### Uso de los endpoints:
#### Empresas:
Crear:
```
POST request /api/company/create

{
    "name": "ACME Corp",
    "email": "acme@corp.com",
    "secret": "secret"
}
```
Borrar:
```
DELETE request /api/company/delete

{
    "email": "acme@corp.com",
    "secret": "secret"
}
```
Buscar por ID:
```
GET request /api/company/:id
```
Buscar por email:
```
POST request /api/company/email

{
    "email": "acme@corp.com",
}
```
Actualizar email:
```
PUT request /api/company/update-email

{
    "email": "new-acme@corp.com",
}
```
Actualizar secret/passphrase:
```
PUT request /api/company/create

{
    "email": "acme@corp.com",
    "secret": "new-secret"
}
```
&nbsp;
#### Usuarios:
Crear:
```
POST request /api/user/create/:id (company ID)

{
    "name": "John Doe",
    "email": "john@doe.com",
    "password": "password",
    "role": "Admin"
}
```
Borrar:
```
DELETE request /api/user/delete

{
    "email": "john@doe.com",
    "password": "password",
}
```
Obtener por ID:
```
GET request /api/user/:id
```
Obtener por email:
```
POST request /api/user/email

{
    "email": "john@doe.com"
}
```
Actualizar email:
```
PUT request /api/user/update-email/:id

{
    "email": "john2@doe.com"
}
```
Actualizar contraseña:
```
PUT request /api/user/update-password

{
    "email": "john@doe.com",
    "password": "new-password"
}
```
&nbsp;
#### Mails
Enviar mail:
* Si el destinatario es uno solo, debe estar dentro de un array igual.
* El archivo adjunto es opcional.
```
POST request /api/mail/send

{
    "email": "john@doe.com",
    "password": "password",
    "subject": "Test email",
    "body": "This is an email body template.",
    "receivers": [
        "user1@api.com",
        "user2@api.com",
        "acme@corp.com"
    ],
    "attachment": file.pdf
}
```
Borrar mail:
```
DELETE request /api/mail/:id
```
Obtener mails enviados:
```
GET request /api/mail/user-sent/:id (user ID)
```
Obtener mails recibidos:
```
POST request /api/mail/user-received

{
    "email": "john@doe.com"
}
```
&nbsp;
