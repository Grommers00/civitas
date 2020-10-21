import 'package:flutter/material.dart';
import 'package:mobile/views/news/list.dart';
import 'package:mobile/views/profile/details.dart';
import 'package:mobile/views/schedule/details.dart';
import 'package:mobile/views/seasons/list.dart';

import '../views/profile/details.dart';
import '../views/schedule/details.dart';

class TabNavigationItem {
  final Widget page;
  final Widget title;
  final Icon icon;

  TabNavigationItem({
    @required this.page,
    @required this.title,
    @required this.icon,
  });

  static List<TabNavigationItem> get items => [
        TabNavigationItem(
          page: DateState(),
          icon: Icon(Icons.calendar_today),
          title: Text("Schedule"),
        ),
        TabNavigationItem(
          page: NewsListPage(),
          icon: Icon(Icons.home),
          title: Text("Home"),
        ),
        TabNavigationItem(
            page: SeasonsListPage(),
            icon: Icon(Icons.list),
            title: Text("Standings")),
        TabNavigationItem(
          page: ProfilePage(),
          icon: Icon(Icons.people),
          title: Text("Profile"),
        ),
      ];
}

class TabsPage extends StatefulWidget {
  @override
  _TabsPageState createState() => _TabsPageState();
}

class _TabsPageState extends State<TabsPage> {
  int _currentIndex = 0;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: IndexedStack(
        index: _currentIndex,
        children: <Widget>[
          for (final tabItem in TabNavigationItem.items) tabItem.page,
        ],
      ),
      bottomNavigationBar: BottomNavigationBar(
         type: BottomNavigationBarType.fixed,
        currentIndex: _currentIndex,
        onTap: (int index) => setState(() => _currentIndex = index),
        items: <BottomNavigationBarItem>[
          for (final tabItem in TabNavigationItem.items)
            BottomNavigationBarItem(
              icon: tabItem.icon,
              title: tabItem.title,
            ),
        ],
      ),
    );
  }
}
