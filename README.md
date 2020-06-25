# Catalogue Service

Spin up local RabbitMQ instance
```
docker run -d --hostname rabbit-queue --name queue -p 15672:15672 -p 5672:5672 rabbitmq:3-management
```