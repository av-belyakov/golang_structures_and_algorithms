#!/bin/bash
set -e

# Импортируем сертификат MinIO в truststore при каждом старте
if [ -f /usr/share/elasticsearch/config/minio-cert.pem ]; then
    /usr/share/elasticsearch/jdk/bin/keytool -importcert \
        -alias minio-cert \
        -file /usr/share/elasticsearch/config/minio-cert.pem \
        -keystore /usr/share/elasticsearch/jdk/lib/security/cacerts \
        -storepass changeit -noprompt && \
        chown elasticsearch:elasticsearch /usr/share/elasticsearch/jdk/lib/security/cacerts
    
#    keytool -importcert -alias minio-cert \
#        -file /usr/share/elasticsearch/config/minio-cert.pem \
#        -keystore /usr/share/elasticsearch/jdk/lib/security/cacerts \
#        -storepass changeit -noprompt || true
fi

# Запускаем стандартный entrypoint Elasticsearch
exec /usr/local/bin/docker-entrypoint.sh