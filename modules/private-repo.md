
**how to go get from a private repo: **

ref: 
- https://stackoverflow.com/questions/27500861/whats-the-proper-way-to-go-get-a-private-repository
- https://go.dev/doc/faq#git_https


step 1: `vi ~/.gitconfig` and add the below code
```
[url "ssh://git@github.com/"]
    insteadOf = https://github.com/
```


step 2: `vi ~/.zshrc` and add below env var to your terminal profile

```
export GOPRIVATE=github.com/muly/*
```

step 3: go get the private repo
```
go get github.com/muly/ecommerce/db@latest
```
