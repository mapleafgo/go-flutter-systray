import 'dart:typed_data';

import 'package:dart_json_mapper/dart_json_mapper.dart';
import 'package:go_flutter_systray/go_flutter_systray.dart';

import '../go_flutter_systray.dart';

@JsonSerializable()
class MenuItem {
  String key;
  Uint8List? _icon;
  String? _title;
  String? _tooltip;
  late bool _isCheckbox;
  List<MenuItem>? child;

  MenuItem({
    required this.key,
    this.child,
    String? title,
    Uint8List? icon,
    String? tooltip,
    bool? isCheckbox,
  }) {
    _icon = icon;
    _title = title;
    _tooltip = tooltip ?? title;
    _isCheckbox = isCheckbox ?? false;
  }

  factory MenuItem.main({
    required Uint8List icon,
    required List<MenuItem> child,
    required String title,
    required String tooltip,
  }) =>
      MenuItem(
        key: "main",
        icon: icon,
        child: child,
        title: title,
        tooltip: tooltip,
      );

  factory MenuItem.separator() => MenuItem(key: "");

  setIcon(Uint8List icon) {
    GoFlutterSystray.setIcon(key: key, iconBytes: icon);
    _icon = icon;
  }

  Uint8List? get icon => _icon;

  setTitle(String title) {
    GoFlutterSystray.setTitle(key: key, title: title);
    _title = title;
  }

  String? get title => _title;

  setTooltip(String tooltip) {
    GoFlutterSystray.setTooltip(key: key, tooltip: tooltip);
    _tooltip = tooltip;
  }

  String? get tooltip => _tooltip;

  setCheckbox(bool isCheckbox) {
    if (isCheckbox) {
      GoFlutterSystray.itemCheck(key);
    } else {
      GoFlutterSystray.itemUncheck(key);
    }
    _isCheckbox = isCheckbox;
  }

  bool get isCheckbox => _isCheckbox;

  Future<void> check() => GoFlutterSystray.itemCheck(key);

  Future<void> uncheck() => GoFlutterSystray.itemUncheck(key);

  Future<bool?> checked() => GoFlutterSystray.itemChecked(key);

  Future<void> disable() => GoFlutterSystray.itemDisable(key);

  Future<void> enable() => GoFlutterSystray.itemEnable(key);

  Future<bool?> disabled() => GoFlutterSystray.itemDisabled(key);

  Future<void> hide() => GoFlutterSystray.itemHide(key);

  Future<void> show() => GoFlutterSystray.itemShow(key);
}
