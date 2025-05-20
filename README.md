# QR Scan Suite


![diagrama app_scan](https://github.com/user-attachments/assets/f81d14e8-2687-4f75-aab2-a7ce36dc0194)


**Aplicación multiplataforma para escaneo y generación de códigos QR con soporte avanzado para múltiples formatos**

## 🔍 Características Principales

- **Escaneo inteligente** de 8 tipos de códigos QR:
  - 🌐 URLs (abre en navegador)
  - 📇 Contactos (vCard/MeCard)
  - 📧 Emails (con asunto/cuerpo predefinido)
  - 📶 Credenciales WiFi (conexión automática)
  - 📅 Eventos (agrega a calendario)
  - 📱 SMS (prepara mensaje)
  - 🗺️ Geolocalización (abre en maps)
  - 📝 Texto plano

- **Multiplataforma**:
  - Web App (Go backend)
  - Desktop (via Nativefier)
  - Android APK (via Viewvie)

- **Funcionalidades adicionales**:
  - Sistema de login via QR
  - Gestor de archivos integrado (PDF, imágenes, docs)
  - Historial de escaneos
  - Generador de QR incorporado

## 🛠️ Tecnologías Clave

- **Backend**: Go (Golang)
- **Frontend**: HTML5, CSS3, JavaScript
- **QR Processing**: go-qrcode, gozxing
- **Empaquetado**: 
  - Desktop: Nativefier
  - Mobile: Viewvie Android Studio

## 🚀 Uso Previsto

- **Personal**: Compartir contactos, conectarse a WiFi fácilmente
- **Empresarial**: Autenticación segura, compartir recursos
- **Educativo**: Distribuir materiales, enlaces rápidos
- **Eventos**: Registro de asistentes, compartir agendas




COMPONENTES CLAVE:

1. Frontend:
   - Interfaz web (convertible a desktop con Nativefier y móvil con Viewvie)
   - templates/*.html: Todas las plantillas de interfaz
   - assets/*: CSS, JS e imágenes

2. Backend (Go):
   - cmd/webserver/main.go: Punto de entrada
   - internal/handlers: Manejo de requests HTTP
   - internal/qr: Lógica de generación y escaneo QR
   - internal/models: Estructuras de datos
   - internal/storage: Persistencia (archivos, DB)

3. Flujo de Escaneo:
   a. Captura de QR (cámara o archivo)
   b. Decodificación (go-qrcode/gozxing)
   c. Identificación de tipo (URL, contacto, WiFi, etc.)
   d. Procesamiento específico
   e. Presentación de resultados

   4. Tipos de QR soportados:
   ┌───────────────┬─────────────────────────────────┐
   │   Tipo QR     │      Procesamiento              │
   ├───────────────┼─────────────────────────────────┤
   │ URL           │ Abrir en navegador/visor        │
   │ Contacto      │ Guardar en agenda               │
   │ Email         │ Abrir cliente de correo         │
   │ WiFi          │ Configurar conexión             │
   │ SMS           │ Preparar mensaje                │
   │ Evento        │ Agregar a calendario            │
   │ Geolocalización│ Abrir mapa                     │
   │ Texto         │ Mostrar contenido               │
   └───────────────┴─────────────────────────────────┘

5. Diagrama de Secuencia Detallado:
1. Usuario → Frontend: Sube imagen QR/usa cámara
2. Frontend → Backend: POST /scan (imagen)
3. Backend → QR Service: Decodificar imagen
4. QR Service → Backend: Contenido crudo
5. Backend → QR Processor: Determinar tipo
6. Backend → Frontend: Respuesta con tipo y acciones
7. Frontend → Usuario: Mostrar resultado/interacciones

run app
go run cmd/webserver/main.go
