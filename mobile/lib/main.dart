import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:mobile/services/news.dart';
import 'package:mobile/theme.dart';
import 'package:mobile/api.dart';
import 'package:http/http.dart' as http;
import 'package:mobile/views/news/card.dart';

void main() {
  runApp(MyApp());
}

Future<Message> fetchMessage() async {
  final response = await http.get('http://localhost:3000/flutter');

  if (response.statusCode == 200) {
    return Message.fromJson(json.decode(response.body));
  } else {
    throw Exception('Failed to load message');
  }
}

class MyApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Civitas',
      theme: darkTheme,
      home: MyHomePage(title: 'Civitas'),
      debugShowCheckedModeBanner: false,
    );
  }
}

class MyHomePage extends StatefulWidget {
  MyHomePage({Key key, this.title}) : super(key: key);

  final String title;

  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  Future<Message> futureMessage;

  @override
  void initState() {
    super.initState();
    futureMessage = fetchMessage();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title),
      ),
      drawer: Drawer(child: DrawerFactory(context)),
      body: Center(
          child: ListView.builder(
              padding: const EdgeInsets.all(8.0),
              itemCount: NewsMockData.length,
              itemBuilder: (context, index) {
                return NewsCard(NewsMockData[index]);
              })),
    );
  }
}

Widget DrawerFactory(BuildContext context) {
  return ListView(padding: EdgeInsets.zero, children: <Widget>[
    DrawerHeader(

      child: FittedBox(
          fit: BoxFit.fill,
          child: Image.network('https://cdn.vox-cdn.com/thumbor/w4lhNFQz97ZHBZXZKjR1Z5_qQ8A=/0x0:400x225/1220x813/filters:focal(168x81:232x145):format(webp)/cdn.vox-cdn.com/uploads/chorus_image/image/67416047/1f9249103f371671071532e02e3ab39d2da49cbe_400x225.0.png')
      ),
      decoration: BoxDecoration(
        color: Theme.of(context).primaryColor,
      ),
    ),
    ListTile(
      title: Text('News'),
      leading: Icon(Icons.new_releases),
      onTap: () {},
    ),
    ListTile(
      title: Text('Profile'),
      leading: Icon(Icons.account_circle),
      onTap: () {},
    ),
    ListTile(
      title: Text('Standings'),
      leading: Icon(Icons.list),
      onTap: () {},
    ),
    ListTile(
      title: Text('Calendar'),
      leading: Icon(Icons.calendar_today),
      onTap: () {},
    ),
    ListTile(
      title: Text('Upcoming Matches'),
      leading: Icon(Icons.timelapse),

      onTap: () {},
    )
  ]);
}
