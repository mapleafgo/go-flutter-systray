// You have generated a new plugin project without
// specifying the `--platforms` flag. A plugin project supports no platforms is generated.
// To add platforms, run `flutter create -t plugin --platforms <platforms> .` under the same
// directory. You can also find a detailed instruction on how to add platforms in the `pubspec.yaml` at https://flutter.dev/docs/development/packages-and-plugins/developing-packages#plugin-platforms.

import 'dart:async';
import 'dart:typed_data';

import 'package:dart_json_mapper/dart_json_mapper.dart';
import 'package:flutter/services.dart';

import 'model/menu_item.dart';

class GoFlutterSystray {
  static const mainMenuKey = "main";
  static const quitCallMethod = "systray_quit_call";

  static const MethodChannel _channel =
      const MethodChannel('go_flutter_systray');

  static final Map<String, Function> _callHanders = {};

  static void registerCallBack(String name, Function callback) =>
      _callHanders[name] = callback;

  static void removeCallBack(String name) => _callHanders.remove(name);

  static void initSystray() {
    _channel.setMethodCallHandler((MethodCall call) async {
      if (_callHanders.containsKey(call.method)) {
        _callHanders[call.method]!();
      }
    });
  }

  static Future<void> hideWindow() => _channel.invokeMethod<void>('hideWindow');

  static Future<void> showWindow() => _channel.invokeMethod<void>('showWindow');

  static Future<void> exitWindow() => _channel.invokeMethod<void>('exitWindow');

  static Future<void> runSystray(MenuItem menu) =>
      _channel.invokeMethod<void>('runSystray', JsonMapper.serialize(menu));

  static Future<void> quitSystray() =>
      _channel.invokeMethod<void>('quitSystray');

  static Future<void> setIcon({
    required String key,
    required Uint8List iconBytes,
  }) =>
      _channel.invokeMethod<void>('setIcon', [key, iconBytes]);

  static Future<void> setTitle({
    required String key,
    required String title,
  }) =>
      _channel.invokeMethod<void>('setTitle', [key, title]);

  static Future<void> setTooltip({
    required String key,
    required String tooltip,
  }) =>
      _channel.invokeMethod<void>('setTooltip', [key, tooltip]);

  static Future<void> itemCheck(String key) =>
      _channel.invokeMethod<void>('itemCheck', key);

  static Future<void> itemUncheck(String key) =>
      _channel.invokeMethod<void>('itemUncheck', key);

  static Future<bool?> itemChecked(String key) =>
      _channel.invokeMethod<bool>('itemChecked', key);

  static Future<void> itemDisable(String key) =>
      _channel.invokeMethod<void>('itemDisable', key);

  static Future<void> itemEnable(String key) =>
      _channel.invokeMethod<void>('itemEnable', key);

  static Future<bool?> itemDisabled(String key) =>
      _channel.invokeMethod<bool>('itemDisabled', key);

  static Future<void> itemHide(String key) =>
      _channel.invokeMethod<void>('itemHide', key);

  static Future<void> itemShow(String key) =>
      _channel.invokeMethod<void>('itemShow', key);
}
