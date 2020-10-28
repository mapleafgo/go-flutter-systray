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
        tooltip: "GoFlutterSystray",
        child: [
          MenuItem(key: "showWindow", title: "Show", tooltip: "Show"),
          MenuItem(key: "quit", title: "退出", tooltip: "退出"),
        ],
      );
      await GoFlutterSystray.runSystray(menu: menu, exitMethod: "quit");
      GoFlutterSystray.registerCallBack(
        "showWindow",
        () => GoFlutterSystray.showWindow(),
      );
      GoFlutterSystray.registerCallBack(
        "quit",
        () => GoFlutterSystray.quit(),
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
          child: Column(
            children: [
              Text("这是一个托盘菜单的实例应用"),
              RaisedButton(
                child: Text("HideWindow"),
                onPressed: GoFlutterSystray.hideWindow,
              ),
            ],
          ),
        ),
      ),
    );
  }
}
