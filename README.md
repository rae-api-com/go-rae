# 📚 RAE API Go Client

[![Go Version](https://img.shields.io/github/go-mod/go-version/rae-api-com/go-rae)](https://golang.org/)
[![Go Report Card](https://goreportcard.com/badge/github.com/rae-api-com/go-rae)](https://goreportcard.com/report/github.com/rae-api-com/go-rae)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://github.com/rae-api-com/go-rae/workflows/CI/badge.svg)](https://github.com/rae-api-com/go-rae/actions)

Un cliente Go elegante y eficiente para la [rae-api.com](http://rae-api.com), que proporciona acceso programático a definiciones de palabras, significados y conjugaciones verbales del diccionario de la Real Academia Española (RAE).

## ✨ Características

- 🚀 **API Simple** - Interfaz limpia y fácil de usar
- 📖 **Múltiples Significados** - Maneja palabras con múltiples acepciones
- 🔄 **Conjugaciones Completas** - Todos los tiempos verbales (Indicativo, Subjuntivo, Imperativo)
- 📝 **Definiciones Ricas** - Acceso a sinónimos, antónimos, etiquetas de uso
- ⚡ **Lightweight** - Sin dependencias pesadas
- 🎯 **Tipado Fuerte** - Estructuras de datos bien definidas

> **Nota**: Este **no** es un cliente oficial de la RAE. El uso de rae-api.com está sujeto a los términos y condiciones de la API.

## 📦 Instalación

```bash
go get github.com/rae-api-com/go-rae
```

## 🚀 Uso Rápido

### Ejemplo Básico

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	rae "github.com/rae-api-com/go-rae"
)

func main() {
	// Crear un nuevo cliente
	client := rae.New()
	
	// Configurar timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	// Buscar una palabra
	entry, err := client.Word(ctx, "hablar")
	if err != nil {
		log.Fatal(err)
	}
	
	// Mostrar resultado
	data, _ := json.MarshalIndent(entry, "", "  ")
	fmt.Println(string(data))
}
```

### Con Opciones Personalizadas

```go
client := rae.New(
	rae.WithTimeout(10*time.Second),
	rae.WithVersion("v1"),
)
```

### Obtener Palabra Aleatoria

```go
// Palabra aleatoria
randomWord, err := rae.GetRandom(ctx, "production")
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Palabra del día: %s\n", randomWord)
```

### Obtener Palabra Diaria

```go
// Palabra diaria
dailyWord, err := rae.GetDaily(ctx, "production")
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Palabra diaria: %s\n", dailyWord)
```

## 📋 Estructura de Respuesta

La API devuelve datos estructurados en el siguiente formato:

```json
{
  "word": "hablar",
  "meanings": [
    {
      "origin": {
        "raw": "Del lat. comedĕre.",
        "type": "lat", 
        "text": "comedĕre"
      },
      "senses": [
        {
          "meaning_number": 1,
          "raw": "1. tr. Masticar...",
          "category": "verb",
          "usage": "common",
          "description": "Masticar y deglutir un alimento sólido.",
          "synonyms": ["masticar", "deglutir"],
          "antonyms": []
        }
      ],
      "conjugations": {
        "non_personal": {
          "infinitive": "hablar",
          "gerund": "hablando", 
          "participle": "hablado"
        },
        "indicative": {
          "present": {
            "singular_first_person": "hablo",
            "singular_second_person": "hablas",
            "singular_third_person": "habla"
            // ... más conjugaciones
          }
        }
        // ... más tiempos verbales
      }
    }
  ]
}
```

## 🛠️ API Reference

### Tipos Principales

```go
type Client struct {
    // campos internos
}

type Entry struct {
    Word     string    `json:"word"`
    Meanings []Meaning `json:"meanings"`
}

type Meaning struct {
    Origin       Origin       `json:"origin"`
    Senses       []Sense      `json:"senses"`
    Conjugations Conjugations `json:"conjugations,omitempty"`
}
```

### Métodos del Cliente

| Método                | Descripción                   | Ejemplo                      |
| --------------------- | ----------------------------- | ---------------------------- |
| `Word(ctx, word)`     | Busca una palabra específica  | `client.Word(ctx, "casa")`   |
| `GetRandom(ctx, env)` | Obtiene una palabra aleatoria | `rae.GetRandom(ctx, "prod")` |
| `GetDaily(ctx, env)`  | Obtiene la palabra del día    | `rae.GetDaily(ctx, "prod")`  |

### Opciones de Configuración

```go
// Configurar timeout personalizado
rae.WithTimeout(10 * time.Second)

// Configurar versión de la API
rae.WithVersion("v1")
```

## 🤝 Contribuir

¡Las contribuciones son bienvenidas! Por favor:

1. Haz fork del proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Haz commit de tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## 📝 Licencia

Este proyecto está bajo la [Licencia MIT](LICENSE).

## 🙏 Reconocimientos

- Este cliente utiliza el servicio [rae-api.com](http://rae-api.com), que no está afiliado con la Real Academia Española
- Todo el contenido del diccionario pertenece a la RAE y está sujeto a sus términos y condiciones
- Gracias a todos los [contribuidores](https://github.com/rae-api-com/go-rae/graphs/contributors)

## 📧 Soporte

Si encuentras algún problema o tienes sugerencias:

- Abre un [issue](https://github.com/rae-api-com/go-rae/issues)
- Consulta la [documentación](https://pkg.go.dev/github.com/rae-api-com/go-rae)

---

<div align="center">
  Hecho con ❤️ para la comunidad Go
</div>