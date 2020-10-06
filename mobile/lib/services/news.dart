class NewsItem {
  final String title;
  final String article;
  final String date;
  final String author;
  final int id;
  final String img;

  NewsItem(this.title, this.date, this.author, this.article, this.id, this.img);

  factory NewsItem.fromJson(Map<String, dynamic> json) {
    return NewsItem(
      json['title'],
      json['date'],
      json['author'],
      json['article'],
      json['id'],
      json['img'],
    );
  }
}
