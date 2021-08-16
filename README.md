---
title: QuejasBackend
description: Proyecto final encardago en Seminario
---



# QuejasBackend

Este sistema tiene la finalidad de gestionar las quejas presentadas a DIACO. 

Presenta los endponit

```
/region (principalmente usado en pruebas)
/region/{id} (principalmente usado en pruebas)
/deparment (principalmente usado en pruebas)
/deparment/{id} (principalmente usado en pruebas)
/township (principalmente usado en pruebas)
/township/{id} (principalmente usado en pruebas)
/township/deparment/{id} 
/diaco 
/person
/company
/complaint
```

### ToDo 

- [x] Manejo de Joins
- [x] Creación de un endpoint para obtener el municipio usando Id de deparmento
- [x] Post del ingreso de "Persona"
  - [ ] Mejorar manejos 
- [x] Post del ingreso de "Empresa"
  - [ ] Mejorar manejos
- [x] Post del ingreso de "Queja"
  - [ ] Mejorar manejos

Nota: Los puntos de post pueden que ten tenga un subsecciones que hacer. Posteriormente se Agragaran Put y Get a lo que se a pertinente, esto resulta descartable a momento.

