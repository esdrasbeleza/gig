# gig

`gig` is a `.gitignore` generator.

Dependencies:
- [gitignore files](https://github.com/dvcs/gitignore)
- [packr](github.com/gobuffalo/packr)
- [go-prompt](github.com/c-bata/go-prompt)
- [testify](github.com/stretchr/testify)

## Using

### From command line

```
$ gig golang code
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with `go test -c`
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

/vendor/
/Godeps/

.vscode/*
!.vscode/settings.json
!.vscode/tasks.json
!.vscode/launch.json
!.vscode/extensions.json

# Ignore all local history of files
.history
```
