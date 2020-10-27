import 'package:flutter/material.dart';
import 'package:go_flutter_systray/go_flutter_systray.dart';

void main() => runApp(MyApp());

class MyApp extends StatefulWidget {
  @override
  _MyAppState createState() => _MyAppState();
}

class _MyAppState extends State<MyApp> {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(
          title: const Text('Plugin example app'),
        ),
        body: Center(
          child: RaisedButton(
            child: Text("HideWindow"),
            onPressed: GoFlutterSystray.hideWindow,
          ),
        ),
      ),
    );
  }
}
