import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:mobile/services/news.dart';
import 'package:mobile/theme.dart';
import 'package:http/http.dart' as http;
import 'package:mobile/views/news/card.dart';
import 'package:mobile/views/news/details.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';

import 'components/drawer.dart';

Future main() async {
  await DotEnv().load('.env');

  runApp(MyApp());
}

Future fetchNewsItems() async {
  String ip = DotEnv().env['IP'];
  String port = DotEnv().env['PORT'];
  final response = await http.get('http://' + ip + port + '/news/');
  print("does this still work?");

  if (response.statusCode == 200) {
    List<NewsItem> news = (json.decode(response.body) as List)
        .map((data) => NewsItem.fromJson(data))
        .toList();

    return news;
  }
}

class MyApp extends StatelessWidget {
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
  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          title: Text(widget.title),
        ),
        drawer: Drawer(child: DrawerFactory(context)),
        body: ListBuilder(context));
  }
}

Widget ListBuilder(BuildContext context) {
  return FutureBuilder(
      future: fetchNewsItems(),
      builder: (context, snapshot) {
        if (snapshot.connectionState == ConnectionState.none ||
            !snapshot.hasData) {
          return CircularProgressIndicator();
        }

        return ListView.builder(
            itemCount: snapshot.data.length,
            itemBuilder: (context, index) {
              return Column(children: [
                Card(
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(8.0),
                    ),
                    color: Theme.of(context).primaryColor,
                    child: InkWell(
                        splashColor:
                            Theme.of(context).accentColor.withAlpha(30),
                        onTap: () => Navigator.push(
                            context,
                            MaterialPageRoute(
                                builder: (context) =>
                                    NewsDetailPage(snapshot.data[index]))),
                        child:
                            Stack(alignment: Alignment.bottomRight, children: [
                          Container(
                              decoration: BoxDecoration(
                                  borderRadius: BorderRadius.only(
                                      topLeft: Radius.circular(8.0),
                                      topRight: Radius.circular(8.0))),
                              clipBehavior: Clip.antiAlias,
                              child: Image.network(snapshot.data[index].img)),
                          Container(
                            margin: EdgeInsets.only(bottom: 165),
                            child: Text(
                              snapshot.data[index].date,
                              style: TextStyle(
                                color: Theme.of(context).accentColor,
                                backgroundColor: Theme.of(context).primaryColor,
                                fontSize: 15.0,
                              ),
                            ),
                          ),
                          Container(
                              decoration: BoxDecoration(
                                  color: Theme.of(context).primaryColor),
                              child: ListTile(
                                title: Text(
                                  snapshot.data[index].title,
                                  style: TextStyle(
                                      color: Theme.of(context).accentColor,
                                      fontSize: 18),
                                ),
                                subtitle: Text(
                                  "By ${snapshot.data[index].author}",
                                  style: TextStyle(
                                      color: Theme.of(context)
                                          .accentColor
                                          .withAlpha(150)),
                                ),
                              ))
                        ]))),
              ]);
            });
      });
}
