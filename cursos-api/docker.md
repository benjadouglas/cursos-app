# Docker

### Instalacion

- docker pull mongo:4.4.6

---

### Run mongo container

- Abrir una terminal
- docker run -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=root --name my-mongo mongo:4.4.6

---

### Entrar al mongo container

- En otra terminal distinta
- docker ps
- docker exec -ti my-mongo bash

---

### Start mongo client

- mongo --username root --password root --authenticationDatabase admin

---

### Verificar la conexión y listar las bases de datos:

- show dbs
