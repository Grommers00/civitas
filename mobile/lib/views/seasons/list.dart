import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:mobile/services/season.dart';
import 'package:mobile/views/standings/list.dart';

Future fetchNewsItems() async {
  final response = await http.get('http://192.168.17.15:3000/seasons/');

  if (response.statusCode == 200) {
    List<SeasonItem> news = (json.decode(response.body) as List)
        .map((data) => SeasonItem.fromJson(data))
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
              return Center(
                  child: Padding(
                      padding: const EdgeInsets.all(12.0),
                      child: InkWell(
                        splashColor:
                            Theme.of(context).accentColor.withAlpha(30),
                        onTap: () => Navigator.push(
                            context,
                            MaterialPageRoute(
                                builder: (context) =>
                                    StandingsListPage(snapshot.data[index]))),
                        child: Container(
                            height: 24,
                            child: Row(
                                children: [Text(snapshot.data[index].game)])),
                      )));
            });
      });
}

class SeasonsListPage extends StatelessWidget {
  static Route<dynamic> route() => MaterialPageRoute(
        builder: (context) => SeasonsListPage(),
      );

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: ListBuilder(context),
    );
  }
}
