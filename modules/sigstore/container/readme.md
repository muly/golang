


```
docker build -t vasureddym/sigempty .
docker push vasureddym/sigempty
    Note: the output in the terminal should have the digest infromation. it can also be retrieved using the below command.
docker inspect vasureddym/sigempty:latest | jq '.[] | .RepoDigests'
```


```
cosign generate-key-pair

cosign sign --key cosign.key vasureddym/sigempty@sha256:ACTUAL_SHA
    Note: replace ACTUAL_SHA with the actual digest sha
    This command creates and uploads associated metadata to the OCI registry where your image is stored. At this point when you go to docker hub and click on your image, it should have 2 items. the image itself and the metadata file.

    https://hub.docker.com/repository/docker/vasureddym/sigempty/general


cosign verify --key cosign.key vasureddym/sigempty
```
