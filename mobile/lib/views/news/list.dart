import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:http/http.dart' as http;
import 'package:mobile/services/news.dart';

import 'details.dart';

Future fetchNewsItems() async {
  String ip = DotEnv().env['IP'];
  String port = DotEnv().env['PORT'];
  
  final response = await http.get("http://$ip$port/news/");

  if (response.statusCode == 200) {
    List<NewsItem> news = (json.decode(response.body) as List)
        .map((data) => NewsItem.fromJson(data))
        .toList();

    return news;
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

class NewsListPage extends StatelessWidget {
  static Route<dynamic> route() => MaterialPageRoute(
        builder: (context) => NewsListPage(),
      );

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: ListBuilder(context),
    );
  }
}
