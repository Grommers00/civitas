import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';

class DrawerFactory extends StatelessWidget {
  final Function logoutAction;
  final String picture;

  DrawerFactory(BuildContext context, this.logoutAction, this.picture);

  @override
  Widget build(BuildContext context) {
    return Drawer(
        child: ListView(padding: EdgeInsets.zero, children: <Widget>[
      DrawerHeader(
        child: Container(
          margin: EdgeInsets.all(0),
          padding: EdgeInsets.all(0),
          color: Theme.of(context).accentColor,
          child: CircleAvatar(
            child: CircleAvatar(
              radius: 50,
              backgroundImage: NetworkImage(
                picture ??
                    'https://airentertainment.biz/wp-content/uploads/2020/08/Among-Us-17_08_2020-20_02_24-1.png',
              ),
            ),
          ),
        ),
      ),
      ListTile(
        title: Text('Settings'),
        leading: Icon(Icons.settings),
        onTap: () {},
      ),
      ListTile(
        title: Text('Logout'),
        leading: Icon(Icons.exit_to_app),
        onTap: () {
          this.logoutAction();
        },
      )
    ]));
  }
}
