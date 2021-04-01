import 'dart:typed_data';

import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:go_flutter_systray/go_flutter_systray.dart';
import 'package:go_flutter_systray/model/menu_item.dart';
import 'main.mapper.g.dart' show initializeJsonMapper;

void main() {
  initializeJsonMapper();
  runApp(MyApp());
}

class MyApp extends StatefulWidget {
  @override
  _MyAppState createState() => _MyAppState();
}

class _MyAppState extends State<MyApp> {
  @override
  void initState() {
    super.initState();
    GoFlutterSystray.initSystray();
    rootBundle
        .load("assets/icon.ico")
        .then((file) => Uint8List.view(file.buffer))
        .then((icon) async {
      var menu = MenuItem.main(
        icon: icon,
        title: "GoFlutterSystray",
        child: [
          MenuItem(key: "window", title: "Show"),
          MenuItem(key: "check", title: "Check", isCheckbox: true),
          MenuItem.separator(),
          MenuItem(key: GoFlutterSystray.quitCallMethod, title: "退出"),
        ],
      );
      await GoFlutterSystray.runSystray(menu);
      GoFlutterSystray.registerCallBack(
        "window",
        GoFlutterSystray.showWindow,
      );
      GoFlutterSystray.registerCallBack(
        "check",
        () async {
          const key = "check";
          if (await GoFlutterSystray.itemChecked(key) == true) {
            GoFlutterSystray.itemUncheck(key);
          } else {
            GoFlutterSystray.itemCheck(key);
          }
        },
      );
      GoFlutterSystray.registerCallBack(
        GoFlutterSystray.quitCallMethod,
        GoFlutterSystray.exitWindow,
      );
    });
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(
          title: const Text('Plugin example app'),
        ),
        body: Center(
          child: Text("这是一个 go_flutter_systray 的示例应用"),
        ),
      ),
    );
  }
}
