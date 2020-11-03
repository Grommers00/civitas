class StandingItem {
  final int id;
  final int seasonId;
  final int playerId;
  final String playerName;
  final int wins;
  final int loses;
  final int matches;

  StandingItem(this.id, this.seasonId, this.playerName, this.playerId,
      this.wins, this.loses, this.matches);

  factory StandingItem.fromJson(Map<String, dynamic> json) {
    return StandingItem(json['id'], json['seasonId'], json['playerName'],
        json["playerId"], json["wins"], json["loses"], json["matches"]);
  }
}
