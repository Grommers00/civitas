import 'package:flutter/material.dart';

// https://coolors.co/ffffff-ffded6-ffbcac-eb898d-a16c76-453e53

final silk = Color.fromRGBO(255, 218, 199, 1.0);
final melon = Color.fromRGBO(255, 188, 172, 1.0);
final coral = Color.fromRGBO(235, 137, 141, 1.0);
final rose = Color.fromRGBO(161, 108, 118, 1.0);
final violet = Color.fromRGBO(69, 62, 83, 1.0);

final ThemeData darkTheme = ThemeData(
  visualDensity: VisualDensity.adaptivePlatformDensity,
  scaffoldBackgroundColor: Colors.white,
  primaryColor: violet,
  accentColor: silk,
);
