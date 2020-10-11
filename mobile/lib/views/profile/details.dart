import 'package:flutter/material.dart';

class ProfilePage extends StatelessWidget {
  // TODO: Integrate with Auth
  final logoutAction = null;
  final String name = 'Seil';
  final String picture =
      'https://airentertainment.biz/wp-content/uploads/2020/08/Among-Us-17_08_2020-20_02_24-1.png';

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        body: Column(
      mainAxisAlignment: MainAxisAlignment.center,
      children: <Widget>[
        Container(
          color: Theme.of(context).primaryColor,
          height: 200,
        ),
        Container(
          width: 150,
          height: 150,
          decoration: BoxDecoration(
            shape: BoxShape.circle,
            image: DecorationImage(
              fit: BoxFit.fill,
              image: NetworkImage(picture ?? ''),
            ),
          ),
        ),
        SizedBox(height: 24.0),
        Text(
          '$name',
          style: TextStyle(fontSize: 24),
        ),
        Text("Last Active: ",
            style: TextStyle(fontSize: 18, color: Colors.grey.shade600)),
        SizedBox(height: 48.0),
      ],
    ));
  }
}
