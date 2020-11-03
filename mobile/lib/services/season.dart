class SeasonItem {
  final String title;
  final String desc;
  final String game;
  final int id;

  SeasonItem(this.title, this.id, this.desc, this.game);

  factory SeasonItem.fromJson(Map<String, dynamic> json) {
    return SeasonItem(
      json['title'],
      json['id'],
      json['desc'],
      json['game']
    );
  }
}
