# Dockerfile
FROM nats:latest

# Copia el archivo de configuración, si existe
COPY nats.conf /etc/nats/nats.conf

# Comando para iniciar NATS con el archivo de configuración
CMD ["nats-server", "-c", "/etc/nats/nats.conf"]
