# Criterios de Evaluación - Primera Entrega

- [x] Conexión Frontend-Microservicios: El frontend se conecta correctamente a los microservicios mediante requests HTTP.
- [x] Vista Home con Búsqueda Funcional: La página de inicio contiene una barra de búsqueda que permite buscar cursos o hoteles por algún criterio (nombre, descripción o categoría).
- [x] Vista de Resultados: Muestra una lista de resultados con imagen, título, descripción y puntuación, permitiendo ver el detalle de cada elemento.
- [x] Vista de Detalle con Inscripción: La vista de detalle muestra la información relevante y permite inscribirse o registrarse, enviando correctamente la solicitud a la API principal. Congrats.
- [x] Gestión de Sesión de Usuarios: Las vistas de login y registro funcionan correctamente, manteniendo el estado de autenticación del usuario.
- [x] Creación de Usuarios en MySQL: La API de Usuarios permite crear usuarios en MySQL correctamente y con verificación de datos.
- [x] Login de Usuarios con JWT: El sistema de login funciona y genera tokens JWT válidos para la autenticación de usuarios.
- [x] Sincronización de RabbitMQ para Actualización de Datos: RabbitMQ actualiza SolR en tiempo real para reflejar cambios en los datos de la API principal.
- [x] Funcionalidad de Búsqueda en SolR: La API de Búsqueda utiliza SolR para realizar búsquedas eficientes y relevantes de cursos u hoteles.
- [x] Operaciones CRUD en MongoDB: La API Principal realiza correctamente operaciones CRUD (creación, actualización, eliminación y obtención) sobre MongoDB.
- [x] Notificación de Cambios a RabbitMQ: Cada operación de escritura (alta, modificación o eliminación) notifica a RabbitMQ para mantener SolR actualizado.
- [x] Endpoint de Inscripción Funcional: El endpoint de inscripción responde correctamente, permitiendo a los usuarios inscribirse en cursos u hoteles.
- [x] Cálculo Concurrente de Disponibilidad: La API Principal (o el search) calcula la disponibilidad de múltiples cursos o habitaciones de forma concurrente mediante Go Routines, asegurando eficiencia y evitando bloqueos.
- [x] Desarrollo en Go siguiendo el Patrón MVC: Los microservicios están desarrollados en Go y cumplen con el patrón MVC.
- [x] Encriptación de Datos Sensibles: Los datos sensibles (como passwords) se manejan de forma segura, siguiendo prácticas de encriptación.
- [x] Código en Repositorio de GitHub: Todo el código se encuentra publicado en GitHub y está accesible para revisión.
- [!] Containerización en Docker: Todos los microservicios están containerizados en Docker y se gestionan mediante Docker Compose.
- [!] Filtrado de Resultados por Capacidad en SolR: La API de Búsqueda implementa un filtro de resultados basado en la capacidad (usuarios inscritos vs. capacidad máxima del curso u hotel).
- [ ] Implementación de Caching Eficiente: Las capas de cache implementadas, tanto en la API de Usuarios como en la API Principal, mejoran la eficiencia del sistema y funcionan de forma estable.
- [ ] Capa de Caché con Memcached en Usuarios: Se implementa una capa de caché usando Memcached en la API de Usuarios, optimizando la consulta de datos de usuarios.

# Criterios de Evaluación - Requisitos Examen final

- [x] Vistas de Administración: Se implementan nuevas vistas para la creación y actualización de cursos/hoteles, accesibles únicamente para administradores/profesores.
- [x] Validación de Permisos en la API de Cursos/Hoteles: API de Cursos/Hoteles valida que el usuario sea administrador antes de permitir la creación, edición de cursos/hoteles o acceso a Admin API.
- [x] Balanceador de Carga / Implementación de NGinX: NGinX se utiliza como balanceador de carga para distribuir el tráfico de manera eficiente entre múltiples instancias de los microservicios, garantizando alta disponibilidad y rendimiento. Se implementa correctamente en al menos 1 microservicio.
- [x] Cobertura de Tests en la Capa de Servicios: Al menos un microservicio debe incluir una cobertura de tests completa para todas sus funcionalidades en la capa de servicios, garantizando confiabilidad.
- [!] Gestión de Microservicios: Se desarrolla una vista exclusiva (Admin) para administradores que permite al menos visualizar las instancias de los microservicios. Idealmente también incluye 2 botones que permiten crear o eliminar instancias dinámicamente para optimizar los recursos.
- [!] Mocks para Clientes y Repositorios: Las implementaciones de clientes de al menos 1 servicio y repositorios deben incluir mocks, facilitando la prueba de las funcionalidades y reduciendo dependencias externas durante los tests.
