import 'package:flutter/material.dart';
import 'package:mobile/services/news.dart';
import 'package:mobile/views/news/details.dart';

class NewsCard extends StatelessWidget {
  final NewsItem newsItem;

  NewsCard(this.newsItem) : super();

  @override
  Widget build(BuildContext context) {
    return Card(
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(8.0),
        ),
        color: Theme.of(context).primaryColor,
        child: InkWell(
            splashColor: Theme.of(context).accentColor.withAlpha(30),
            // Navigation
            onTap: () => Navigator.push(
                context,
                MaterialPageRoute(
                    builder: (context) => NewsDetailPage(newsItem))),
            child: Column(
              children: [
                Stack(
                    alignment: Alignment.topRight,
                    children: [
                      // Article Image
                      Container(
                          decoration: BoxDecoration(
                              borderRadius: BorderRadius.only(
                                  topLeft: Radius.circular(8.0),
                                  topRight: Radius.circular(8.0))),
                          clipBehavior: Clip.antiAlias,
                          child: FittedBox(
                              fit: BoxFit.fitWidth,
                              child: Image.network(
                                'https://airentertainment.biz/wp-content/uploads/2020/08/Among-Us-17_08_2020-20_02_24-1.png',
                              ))),
                      // Article Date
                      Container(
                        margin: EdgeInsets.only(top: 20, bottom: 20),
                        child: Text(
                          newsItem.date,
                          style: TextStyle(
                            color: Theme.of(context).accentColor,
                            backgroundColor: Theme.of(context).primaryColor,
                            fontSize: 15.0,

                          ),
                        ),
                      ),
                    ]),
                // Article Title / Author
                ListTile(
                  title: Text(
                    newsItem.title,
                    style: TextStyle(
                        color: Theme.of(context).accentColor, fontSize: 18),
                  ),
                  subtitle: Text(
                    "By ${newsItem.author}",
                    style: TextStyle(
                        color: Theme.of(context).accentColor.withAlpha(150)),
                  ),
                )
              ],
            )));
  }
}
