# seminarioGO

Linea para iniciar el programa 
```
go run cmd/usuario/usuariosrv.go
```
La URI principal es /usuarios y el CRUD esta definido de la siguiente manera

/usuarios  
POST -> recibe un JSON con la siguiente estructura
```
{
    "nombre": "lucho",
    "dni": 25535
}
```
/usuarios  
PUT -> tambien recibe un JSON de la siguente forma
```
{
    "ID":   1,
    "nombre": "Sebastian",
    "dni": 22222222
}
```
/usuarios/1  
DELETE -> recibe por URL el id de usuario a eliminar

/usuarios/1  
GET -> de esta manera busca un usuario por id  
/usuarios  
GET -> de esta manera trae la lista completa de usuarios  
