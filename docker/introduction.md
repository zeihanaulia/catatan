# Docker 

## Mengenal docker

## Snippet

- Hapus semua docker image berdasarkan nama

```bash
docker rmi $(docker images --format '{{.Repository}}:{{.Tag}}' | grep 'elasticsearch')
```

- Hapus semua docker image yang <None>

```bash
docker rmi $(docker images -f "dangling=true" -q)
```