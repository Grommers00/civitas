import 'package:flutter/material.dart';
import 'package:mobile/theme.dart';

import 'components/drawer.dart';
import 'components/tabs.dart';

void main() => runApp(const MyApp());

class MyApp extends StatefulWidget {
  const MyApp({Key key}) : super(key: key);

  @override
  _MyAppState createState() => _MyAppState();
}

class _MyAppState extends State<MyApp> {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Civitas',
      darkTheme: darkTheme,
      home: Scaffold(
        appBar: AppBar(
          title: const Text('Civitas'),
        ),
        drawer: DrawerFactory(context, null, null),
        body: Center(child: TabsPage()),
      ),
    );
  }
}
