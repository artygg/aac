
# Server Setup Documentation

## Overview

This document outlines the configuration and setup of a server using Nginx as a reverse proxy server, with iptables configured for firewall rules, and multiple services running in Docker containers managed by Docker Compose. The services include MySQL, a Go web server, and a Python web server.

## Components

### 1. Nginx as Reverse Proxy Server

Nginx is configured to act as a reverse proxy server. The configuration details are as follows:

- **Port 80**: Disabled
- **Port 443**: Points to the Go web application
- **Port 8090**: Points to the Python web server

The Nginx configuration file (`nginx.conf`) should be set up to handle the SSL termination and proxying:

```cv
server {
    listen 443 ssl;
    server_name your_domain.com;

    ssl_certificate /etc/nginx/ssl/your_domain.com.crt;
    ssl_certificate_key /etc/nginx/ssl/your_domain.com.key;

    location / {
        proxy_pass http://golang_server:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    location /python {
        proxy_pass http://python_server:8090;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

### 2. Iptables - Firewall Configuration

The iptables are configured to block all ports except for the ones required for the application and SSH access. The following ports are opened:

- **SSH**: (Typically port 22)
- **443**: For HTTPS traffic
- **8090**: For the Python web server

A sample iptables configuration script (`iptables.sh`) is provided:

```bash
# Allow loopback interface (localhost)
iptables -A INPUT -i lo -j ACCEPT

# Allow established connections
iptables -A INPUT -m conntrack --ctstate ESTABLISHED,RELATED -j ACCEPT

# Allow SSH
iptables -A INPUT -p tcp --dport 22 -j ACCEPT

# Allow HTTP/HTTPS
iptables -A INPUT -p tcp --dport 443 -j ACCEPT

# Allow Python server
iptables -A INPUT -p tcp --dport 8090 -j ACCEPT

# Drop everything else
iptables -A INPUT -j DROP

# Set default chain policies
iptables -P INPUT DROP
iptables -P FORWARD DROP
iptables -P OUTPUT ACCEPT
```

### 3. Docker Compose Configuration

The services are managed using Docker Compose, which defines several containers:

- **MySQL**: Database service
- **Go Web Server**: Web application service that runs pre-compiled GO back-end
- **Python Web Server**: API for device to send and analyze photos

The `docker-compose.yml` file is configured as follows:

```yaml
version: '3.8'

services:
  db:
    image: mysql:8.0
    container_name: mysql_container
    environment:
      MYSQL_ROOT_PASSWORD: pass
      MYSQL_DATABASE: aas
      MYSQL_USER: aas
      MYSQL_PASSWORD: pass2
    ports:
      - "3306:3306"
    volumes:
      - ./db_data/:/var/lib/mysql
    networks:
      aas_network:
        ipv4_address: 192.168.100.9

  phpmyadmin:
    image: phpmyadmin/phpmyadmin
    container_name: phpmyadmin_container
    environment:
      PMA_HOST: db
      MYSQL_ROOT_PASSWORD: pass
    ports:
      - "8474:80"
    depends_on:
      - db
    networks:
      aas_network:
        ipv4_address: 192.168.100.10
  web:
    build:
      context: ./aac
    ports:
      - "8080:8080"
    volumes:
      - ./aac:/app
    networks:
      aas_network:
        ipv4_address: 192.168.100.11
  flask:
    build:
      context: ./aws
    container_name: flask_app
    ports:
      - "5001:5001"
    volumes:
      - ./aws:/app

    environment:
      - AWS_ACCESS_KEY_ID=
      - AWS_SECRET_ACCESS_KEY=
      - AWS_DEFAULT_REGION=us-east-2
    networks:
      aas_network:
        ipv4_address: 192.168.100.12

  networks:
    aas_network:
      driver: bridge
      ipam:
        config:
          - subnet: 192.168.100.0/24
```
### 4 SSL/TLS certificate
We used certbot to obtain a free SSL/TLS certificate in order to make all the connections to the server secured.

### Managing the Services

- **Start services**: `docker-compose up -d`
- **Stop services**: `docker-compose down`
- **View logs**: `docker-compose logs -f`
- **Restart a service**: `docker-compose restart <service_name>`
