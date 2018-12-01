# go-multiple-module-repository

```shell
$ go run scripts/affected.go -files `cat files.txt | tr '\n' ','`
$ go run scripts/affected.go -files `git diff --name-only HEAD~1...HEAD | tr '\n' ','`
```
