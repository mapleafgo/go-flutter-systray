# go_flutter_systray

This Go package implements the host-side of the Flutter [go_flutter_systray](https://github.com/mapleafgo/go-flutter-systray) plugin.

## Usage

Import as:

```go
import go_flutter_systray "github.com/mapleafgo/go-flutter-systray/go"
```

Then add the following option to your go-flutter [application options](https://github.com/go-flutter-desktop/go-flutter/wiki/Plugin-info):

```go
flutter.AddPlugin(&go_flutter_systray.GoFlutterSystrayPlugin{}),
```
