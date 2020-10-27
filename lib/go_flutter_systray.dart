// You have generated a new plugin project without
// specifying the `--platforms` flag. A plugin project supports no platforms is generated.
// To add platforms, run `flutter create -t plugin --platforms <platforms> .` under the same
// directory. You can also find a detailed instruction on how to add platforms in the `pubspec.yaml` at https://flutter.dev/docs/development/packages-and-plugins/developing-packages#plugin-platforms.

import 'dart:async';
import 'dart:typed_data';

import 'package:dart_json_mapper/dart_json_mapper.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

import 'model/menu_item_bean.dart';

class GoFlutterSystray {
  static const MethodChannel _channel =
      const MethodChannel('go_flutter_systray');

  static Future<void> hideWindow() => _channel.invokeMethod('hideWindow');

  static Future<void> showWindow() => _channel.invokeMethod('showWindow');

  static Future<void> runSystray({
    @required MenuItem menu,
    @required String exitMethod,
  }) =>
      _channel.invokeMethod(
        'runSystray',
        [JsonMapper.serialize(menu), exitMethod],
      );

  static Future<void> setIcon({
    @required String key,
    @required Uint8List iconBytes,
  }) =>
      _channel.invokeMethod('setIcon', [key, iconBytes]);

  static Future<void> setTitle({
    @required String key,
    @required String title,
  }) =>
      _channel.invokeMethod('setTitle', [key, title]);

  static Future<void> setTooltip({
    @required String key,
    @required String tooltip,
  }) =>
      _channel.invokeMethod('setTooltip', [key, tooltip]);

  static Future<void> itemCheck(String key) =>
      _channel.invokeMethod('itemCheck', [key]);

  static Future<void> itemUncheck(String key) =>
      _channel.invokeMethod<bool>('itemUncheck', [key]);

  static Future<bool> itemChecked(String key) =>
      _channel.invokeMethod<bool>('itemChecked', [key]);

  static Future<void> itemDisable(String key) =>
      _channel.invokeMethod<bool>('itemDisable', [key]);

  static Future<void> itemEnable(String key) =>
      _channel.invokeMethod<bool>('itemEnable', [key]);

  static Future<bool> itemDisabled(String key) =>
      _channel.invokeMethod<bool>('itemDisabled', [key]);

  static Future<void> itemHide(String key) =>
      _channel.invokeMethod<bool>('itemHide', [key]);

  static Future<void> itemShow(String key) =>
      _channel.invokeMethod<bool>('itemShow', [key]);
}