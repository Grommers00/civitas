import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:http/http.dart' as http;
import 'package:mobile/services/season.dart';
import 'package:mobile/services/standing.dart';

Future fetchNewsItems(int seasonId) async {
  String ip = DotEnv().env['IP'];
  String port = DotEnv().env['PORT'];

  final response =
      await http.get("http://$ip$port/standing/$seasonId");

  if (response.statusCode == 200) {
    List<StandingItem> news = (json.decode(response.body) as List)
        .map((data) => StandingItem.fromJson(data))
        .toList();

    return news;
  }
}

Widget ListBuilder(BuildContext context, int seasonId) {
  return FutureBuilder(
      future: fetchNewsItems(seasonId),
      builder: (context, snapshot) {
        if (snapshot.connectionState == ConnectionState.none ||
            !snapshot.hasData) {
          return CircularProgressIndicator();
        }

        return ListView.builder(
            itemCount: snapshot.data.length,
            itemBuilder: (context, index) {
              final StandingItem item = snapshot.data[index];
              return Card(
                  shape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(8.0),
                  ),
                  // color: Theme.of(context).accentColor,
                  child: Container(
                      margin: EdgeInsets.all(12),
                      child: ListBody(children: [
                       
                        Text(
                          item.playerId.toString(),
                        ),
                        Text(
                          "Matches: ${item.matches.toString()}",
                        ),
                      ])));
            });
      });
}

class StandingsListPage extends StatelessWidget {
  final SeasonItem season;

  StandingsListPage(this.season);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Game: ${season.game}, Season: {season.title}"),
      ),
      body: Center(child: ListBuilder(context, season.id)),
    );
  }
}
