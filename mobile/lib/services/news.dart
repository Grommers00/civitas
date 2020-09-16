class NewsItem {
  final String title;
  final String description;
  final String date;
  final String author;

  NewsItem(this.title, this.date, this.author, this.description);
}

final NewsMockData = [
  NewsItem('Grommers Takes Down Gato!', 'Aug 24, 2020', 'Recham', 'lmao'),
  NewsItem('Ark Takes Down Gato!', 'Aug 25, 2020', 'Recham', 'lmao'),
  NewsItem('Ark Takes Down Gato!', 'Aug 26, 2020', 'Recham', 'lmao')
];
