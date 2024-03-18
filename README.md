# logwrap
Various wrappers for golang logger libraries

## TODO wrapper for
 - [x] logrus
 - [ ] slog
 - [ ] zap

## Usage
To add logwrap to your project, run:

```shell
go get github.com/qrave1/logwrap
```

```go
func main() {
	log := logwrap.NewDefaultPrettyWrapper()
	
	log.Info("Hello, Wolrd!")
}
```
