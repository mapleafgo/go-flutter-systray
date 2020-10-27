import 'dart:typed_data';

import 'package:dart_json_mapper/dart_json_mapper.dart';
import 'package:go_flutter_systray/go_flutter_systray.dart';

@JsonSerializable()
class MenuItem {
  String key;
  Uint8List _icon;
  String _title;
  String _tooltip;
  List<MenuItem> child;

  MenuItem({icon, title, tooltip}) {
    _icon = icon;
    _title = title;
    _tooltip = tooltip;
  }

  factory MenuItem.separator() => MenuItem();

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
}
