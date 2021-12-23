# Minimal TCP Service in Go to Do Network Testing

This minimal TCP server is written to test routing and services in any
environment, but originally to understand how Kubernetes services work.
Being a `FROM SCRATCH` container makes it a nice size for experimenting
with large ReplicasSets and such. I prefer TCP directly rather than any
HTTP server (as is popular) because I need insight into things that are
not related to Ingress routing at all, specifically to explore
cluster-to-cluster communications without HTTP (or even gRPC, which uses
HTTP).

To start the service modify the `docker-compose.yaml` values as desired
and execute the following:

```
docker-compose up
```

Or, with just `docker`:

```
docker run -it -e NAME=whatever -e PORT=3444 -p 3444:3444 rwxrob/gotcptest
```

Then connect to it with `netcat`/`nc` (or whatever):

```
nc localhost 3444
```
