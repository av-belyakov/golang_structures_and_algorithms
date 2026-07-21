FROM docker.elastic.co/elasticsearch/elasticsearch:9.3.0
# docker.elastic.co/elasticsearch/elasticsearch:8.15.0
# Копируем сертификат в контейнер
COPY minio/certs/public.crt /usr/share/elasticsearch/config/minio-cert.pem
# сертификат public-192.168.9.53.crt только для доступа к MinIO на сервере DATAHOOK
#COPY minio/certs/public-192.168.9.53.crt /usr/share/elasticsearch/config/minio-cert.pem
# Импортируем сертификат в JDK truststore
USER root
RUN /usr/share/elasticsearch/jdk/bin/keytool -importcert \
    -alias minio-cert \
    -file /usr/share/elasticsearch/config/minio-cert.pem \
    -keystore /usr/share/elasticsearch/jdk/lib/security/cacerts \
    -storepass changeit -noprompt && \
    chown elasticsearch:elasticsearch /usr/share/elasticsearch/jdk/lib/security/cacerts
USER elasticsearch