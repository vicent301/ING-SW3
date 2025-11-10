Trabajo práctico 5 - Ingeniería de Software
Alumnos: Barrale Simón y Monzo Vicente



Para este trabajo práctico implementamos una suite completa de pruebas unitarias para el backend y el frontend de la aplicación desarrollada en el TP5.
El objetivo fue validar la lógica de negocio, el manejo de errores, la autenticación y la correcta interacción entre capas, aislando dependencias mediante mocks y aplicando el patrón AAA (Arrange – Act – Assert).
Finalmente, integramos la ejecución en el pipeline de Azure DevOps, publicando métricas de cobertura y resultados en formato JUnit/Cobertura.

Frameworks de testing elegidos y por qué Backend (Go + Gin)

Tecnologías utilizadas
testing + httptest (stdlib): permite simular requests HTTP contra handlers de Gin sin levantar servidor real.
testify (assert/require): aserciones más legibles y con mensajes claros.
Mocks propios (interfaces): para reemplazar repositorios reales.
DB en memoria: con testutil.SetupInMemoryDB para un estado determinista.
go-junit-report + gocover-cobertura: convierten los resultados de go test a JUnit y Cobertura para integrarlos al pipeline.

Justificación
Estas herramientas permiten tests livianos, rápidos, deterministas y totalmente aislados de MySQL o cualquier infraestructura real. Son el estándar de testing en Go y se integran naturalmente con Azure DevOps.



Frameworks de testing elegidos y por qué Frontend (Vite + React)
Tecnologías utilizadas
Vitest: runner nativo de Vite, rápido y con cobertura integrada vía V8.
@testing-library/react + user-event: orientado a testear el comportamiento del usuario (renders, roles, clicks, inputs).
jsdom: DOM virtual para componentes.
Mocks de fetch (vi.fn()): para aislar llamadas a la API.
Mocks de localStorage y router: para testear autenticación y navegación.

Justificación
Testing Library permite pruebas robustas y centradas en el usuario, evitando fragilidad en selectors. Vitest se integra naturalmente con Vite y garantiza velocidad y simplicidad.




Estrategia de mocking utilizada Backend

Implementamos una estrategia de aislamiento total:

Inyección de dependencias vía interfaces: repositorios mockeados (ej. ProductRepositoryMock) reemplazan la DB real.
httptest para tráfico HTTP: simulamos requests/response en memoria sin puertos ni servidor.
Middleware de autenticación mockeado: withEmail(email) setea el usuario en contexto sin JWT real.
Base de datos en memoria: SetupInMemoryDB crea una DB efímera y limpia por test.
Simulación de falla de infraestructura: seteamos database.DB = nil en tests específicos, verificando respuestas coherentes ante errores del sistema.

Gracias a esto, aislamos dependencias externas y probamos comportamiento, lógica de negocio y manejo de errores sin levantar servicios.




Estrategia de mocking utilizada FRONTEND

Mock personalizado de fetch con vi.fn(): controlamos respuestas 200, 400, 500 y errores de red.
Mock de localStorage/contexto de autenticación: probamos UI y flujos según sesión sin depender de tokens reales.
Mock de router (ProtectedRoute): verificamos render/redirección sin navegación real.
Fake timers: para testear timeouts o delays cuando correspondió.

Esto permite validar comportamientos visibles (UI) sin depender del backend.

Casos de pruebas más relevantes Backend

  Carrito (cart)
Autorización requerida (401): si falta usuario en contexto.
Validación de entrada (400): JSON roto, campos inválidos, quantity negativa, product_id = 0.
Usuario inexistente (404): email válido pero usuario no registrado.
Dependencia caída (400): DB = nil.
Flujo feliz end-to-end (200): add → get → remove → clear.

  Productos
CRUD básico (DAO + GORM)
Reglas de negocio en el servicio (con mocks)
Producto inexistente
Manejo de errores del repositorio 

  Usuarios
Registro y búsqueda por email
Duplicados
Errores de persistencia

  Middleware de autenticación
Token/contexto válido: accede
Token/contexto inválido: 401
Seteo correcto del email en el contexto

  Rutas (routes integration)
Rutas registradas correctamente
Atributos y handlers asociados
Respuestas esperadas por endpoint

  Base de datos (connection)
Creación de conexión
Configuración inválida
Migraciones y manejo de errores

Casos de prueba Frontend
  Servicios (api.js)
getProducts / getProductById / login / register / getProfile
Validación de método, headers y URL final
Errores controlados (400, 500, network)

  Login / Register
Validaciones de campos
Happy path con navegación
Mensajes de error del servidor

  Productos
Estados: loading, empty, error
Render correcto del catálogo

  Carrito
agregar → quitar → vaciar
totales calculados correctamente
errores de API

  Navbar
visibilidad de links según sesión
logout limpia token y redirige

  Rutas protegidas (ProtectedRoute)
autenticado → renderiza componente
no autenticado → redirige a login

Por ejemplo una funcion que creamos para testear en el backend es la siguiente:

func TestGetCart_Unauthorized(t *testing.T) {
gin.SetMode(gin.TestMode)
r := gin.New()
r.GET("/cart", GetCart)

    req, _ := http.NewRequest(http.MethodGet, "/cart", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusUnauthorized, w.Code)
}


Explicación

Arrange: router + handler sin contexto de email
Act: simulamos GET /cart
Assert: esperamos 401
Esto garantiza que no se puede operar el carrito sin login, validando la seguridad de la API.



  Integración con pipeline para CI/CD

Ambos entornos (frontend y backend) fueron integrados al pipeline:

    Tests de backend → go test -cover ... → JUnit + Cobertura
    Tests de frontend → vitest --coverage --reporter=junit
    Publicación de resultados y cobertura en Azure DevOps
    Fallo en un test ⇒ falla el job ⇒ detiene el pipeline
    Artefactos de cobertura se publican como HTML