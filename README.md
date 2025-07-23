# Otus-DZ2

## Описание
Проект представляет собой простое веб-приложение, которое отвечает на запросы `/health` и `/otusapp/lromanov/health` с возвратом JSON-ответа.

## Требования
- Docker
- Kubernetes (kubectl)
- Ingress Controller (например, nginx-ingress)

## Установка и запуск

### 1. Сборка Docker-образа
```bash
docker build -t mrpuld/otus-dz2:latest .
docker push mrpuld/otus-dz2:latest
```

### 2. Развертывание в Kubernetes
Примените манифесты одной командой:
```bash
kubectl apply -f k8s/
```

### 3. Доступ к приложению
Приложение доступно по следующим URL:
- `http://arch.homework/otusapp/lromanov/health`
- `http://arch.homework/health`

Оба URL должны возвращать JSON-ответ от приложения.

## Структура проекта
- `k8s/` - Kubernetes манифесты
  - `deployment.yaml` - Deployment для приложения
  - `service.yaml` - Service для приложения
  - `ingress.yaml` - Ingress для `/otusapp/lromanov/health`
  - `ingress-health.yaml` - Ingress для `/health`
- `Dockerfile` - инструкции для сборки Docker-образа
- `main.go` - исходный код приложения

## Примечания
- Для корректной работы Ingress убедитесь, что Ingress Controller настроен и доступен в кластере.
- Если возникают проблемы с доступом, проверьте логи Ingress Controller и приложения:
  ```bash
  kubectl logs -n ingress-nginx -l app.kubernetes.io/name=ingress-nginx
  kubectl logs -l app=otus-dz2
  ``` 