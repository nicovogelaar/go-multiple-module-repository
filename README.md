# go-multiple-module-repository

```shell
$ cat files.txt | go run scripts/affected.go | ./scripts/ci.sh
$ git diff --name-only HEAD~1...HEAD | go run scripts/affected.go | ./scripts/ci.sh
```
