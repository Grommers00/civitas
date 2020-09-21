import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:mobile/services/news.dart';

class NewsDetailPage extends StatelessWidget {
  final NewsItem newsItem;

  NewsDetailPage(this.newsItem) : super();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("News"),
      ),
      body: Container(
          color: Theme.of(context).primaryColor,
          child: Column(
            children: [
              Container(
                width: MediaQuery.of(context).size.width,
                height: 250,
                decoration: BoxDecoration(
                  image: DecorationImage(
                    fit: BoxFit.fill,
                    image: NetworkImage(
                      'https://airentertainment.biz/wp-content/uploads/2020/08/Among-Us-17_08_2020-20_02_24-1.png',
                    ),
                  ),
                ),
              ),
              ListTile(
                title: Text(
                  newsItem.title,
                  style: TextStyle(
                      color: Theme.of(context).accentColor, fontSize: 18),
                ),
                subtitle: Text(
                  "By ${newsItem.author} on ${newsItem.date}",
                  style: TextStyle(
                      color: Theme.of(context).accentColor.withAlpha(150)),
                ),
              ),
              Container(
                  child: Text(newsItem.article,
                      style: TextStyle(color: Colors.white)))
            ],
          )),
      floatingActionButton: FloatingActionButton(
        onPressed: () => print("They want to share this!"),
        tooltip: 'Increment',
        child: Icon(Icons.share),
      ),
    );
  }
}
