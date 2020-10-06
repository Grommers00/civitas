import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';

Widget DrawerFactory(BuildContext context) {
  return ListView(padding: EdgeInsets.zero, children: <Widget>[
    DrawerHeader(
      child: FittedBox(
          fit: BoxFit.fill,
          child: Image.network(
              'https://cdn.vox-cdn.com/thumbor/w4lhNFQz97ZHBZXZKjR1Z5_qQ8A=/0x0:400x225/1220x813/filters:focal(168x81:232x145):format(webp)/cdn.vox-cdn.com/uploads/chorus_image/image/67416047/1f9249103f371671071532e02e3ab39d2da49cbe_400x225.0.png')),
      decoration: BoxDecoration(
        color: Theme.of(context).primaryColor,
      ),
    ),
    ListTile(
      title: Text('News'),
      leading: Icon(Icons.new_releases),
      onTap: () {},
    ),
    ListTile(
      title: Text('Profile'),
      leading: Icon(Icons.account_circle),
      onTap: () {},
    ),
    ListTile(
      title: Text('Standings'),
      leading: Icon(Icons.list),
      onTap: () {},
    ),
    ListTile(
      title: Text('Calendar'),
      leading: Icon(Icons.calendar_today),
      onTap: () {},
    ),
    ListTile(
      title: Text('Upcoming Matches'),
      leading: Icon(Icons.timelapse),
      onTap: () {},
    )
  ]);
}
