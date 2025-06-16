# OTUS AM DZ2

Простой HTTP-сервис на Go, который отвечает на порту 8000 и имеет эндпоинт `/health/`.

## Запуск

### Локально
```bash
go run main.go
```

### Docker
```bash
docker run -p 8000:8000 mrpuld/otus-dz2:latest
```

## API

### GET /health/
Возвращает статус сервиса в формате JSON:
```json
{
    "status": "OK"
}
``` 