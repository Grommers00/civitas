class CalendarItem {
  final int id;
  final DateTime date;
  final String description;
  final String competitor1;
  final String competitor2;

  CalendarItem(
      this.id, this.date, this.description, this.competitor1, this.competitor2);

  factory CalendarItem.fromJson(Map<String, dynamic> json) {
    return CalendarItem(
      json['id'],
      json['date'],
      json['description'],
      json['competitor1'],
      json['competitor2'],
    );
  }
}
