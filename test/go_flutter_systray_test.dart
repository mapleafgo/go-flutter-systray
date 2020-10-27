import 'package:flutter/services.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:go_flutter_systray/go_flutter_systray.dart';

void main() {
  const MethodChannel channel = MethodChannel('go_flutter_systray');

  TestWidgetsFlutterBinding.ensureInitialized();

  setUp(() {
    channel.setMockMethodCallHandler((MethodCall methodCall) async {
      return '42';
    });
  });

  tearDown(() {
    channel.setMockMethodCallHandler(null);
  });

  test('getPlatformVersion', () async {
    expect(await GoFlutterSystray.platformVersion, '42');
  });
}
