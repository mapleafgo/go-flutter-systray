import 'dart:typed_data';

import 'package:dart_json_mapper/dart_json_mapper.dart';
import 'package:flutter/material.dart';
import 'package:go_flutter_systray/go_flutter_systray.dart';

import '../go_flutter_systray.dart';

@JsonSerializable()
class MenuItem {
  String key;
  Uint8List _icon;
  String _title;
  String _tooltip;
  List<MenuItem> child;

  MenuItem({
    @required String key,
    Uint8List icon,
    String title,
    String tooltip,
    List<MenuItem> child,
  }) {
    this.key = key;
    this.child = child;
    _icon = icon;
    _title = title;
    _tooltip = tooltip;
  }

  factory MenuItem.main({
    @required Uint8List icon,
    @required List<MenuItem> child,
    String title,
    String tooltip,
  }) =>
      MenuItem(
        key: "main",
        icon: icon,
        child: child,
        title: title,
        tooltip: tooltip,
      );

  factory MenuItem.separator() => MenuItem(key: "");

  set icon(Uint8List iconBytes) {
    GoFlutterSystray.setIcon(key: key, iconBytes: iconBytes);
    _icon = iconBytes;
  }

  get icon => _icon;

  set title(String titleStr) {
    GoFlutterSystray.setTitle(key: key, title: titleStr);
    _title = titleStr;
  }

  get title => _title;

  set tooltip(String tooltipStr) {
    GoFlutterSystray.setTooltip(key: key, tooltip: tooltip);
    _tooltip = tooltipStr;
  }

  get tooltip => _tooltip;

  Future<void> check() => GoFlutterSystray.itemCheck(key);

  Future<void> uncheck() => GoFlutterSystray.itemUncheck(key);

  Future<bool> checked() => GoFlutterSystray.itemChecked(key);

  Future<void> disable() => GoFlutterSystray.itemDisable(key);

  Future<void> enable() => GoFlutterSystray.itemEnable(key);

  Future<bool> disabled() => GoFlutterSystray.itemDisabled(key);

  Future<void> hide() => GoFlutterSystray.itemHide(key);

  Future<void> show() => GoFlutterSystray.itemShow(key);
}
