docker run -d \
--name=mycockroahdb \
--hostname=mycockroahdb \
-p 26257:26257 -p 8088:8080  \
cockroachdb/cockroach start --insecure
