# glog
This is a helper module for simple logging.

```golang
type Logger interface {
	Log(args ...interface{})
	Fatal(args ...interface{})
}
```


