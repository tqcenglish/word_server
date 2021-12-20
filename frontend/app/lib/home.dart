import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:logging/logging.dart';
import 'package:quizstar/word_az.dart';

class HomePage extends StatefulWidget {
  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  final log = Logger('MyClassName');
  List<String> images = [
    "images/study.png",
    "images/test.png",
    "images/save.png",
  ];

  List<String> des = [
    "常见单词, 图片记忆, 真人发音",
    "测试复习，强化记忆，超脱众人",
    "收藏重现，巩固强化，游刃有余",
  ];

  Widget customcard(String tag, String image, String des) {
    return Padding(
      padding: EdgeInsets.symmetric(
        vertical: 20.0,
        horizontal: 30.0,
      ),
      child: InkWell(
        onTap: () {
          log.info("info");
          switch (tag) {
            case "记单词":
              Navigator.of(context).push(MaterialPageRoute(
                builder: (context) => WordAZ(),
              ));
              break;
            default:
          }
          // Navigator.of(context).pushReplacement(MaterialPageRoute(
          //   // in changelog 1 we will pass the tag name to ther other widget class
          //   // this name will be used to open a particular JSON file
          //   // for a particular language
          //   builder: (context) => GetJson(tag),
          // ));
        },
        child: Material(
          color: Colors.indigoAccent,
          elevation: 10.0,
          borderRadius: BorderRadius.circular(25.0),
          child: Container(
            child: Column(
              children: <Widget>[
                Padding(
                  padding: EdgeInsets.symmetric(
                    vertical: 10.0,
                  ),
                  child: Material(
                    elevation: 5.0,
                    borderRadius: BorderRadius.circular(100.0),
                    child: Container(
                      // changing from 200 to 150 as to look better
                      height: 150.0,
                      width: 150.0,
                      child: ClipOval(
                        child: Image(
                          fit: BoxFit.cover,
                          image: AssetImage(
                            image,
                          ),
                        ),
                      ),
                    ),
                  ),
                ),
                Center(
                  child: Text(
                    tag,
                    style: TextStyle(
                      fontSize: 20.0,
                      color: Colors.white,
                      fontFamily: "Quando",
                      fontWeight: FontWeight.w700,
                    ),
                  ),
                ),
                Container(
                  padding: EdgeInsets.all(20.0),
                  child: Text(
                    des,
                    style: TextStyle(
                        fontSize: 18.0,
                        color: Colors.white,
                        fontFamily: "Alike"),
                    maxLines: 5,
                    textAlign: TextAlign.justify,
                  ),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    SystemChrome.setPreferredOrientations(
        [DeviceOrientation.portraitDown, DeviceOrientation.portraitUp]);
    return Scaffold(
      appBar: AppBar(
        title: Text(
          "记单词",
          style: TextStyle(
            fontFamily: "Quando",
          ),
        ),
      ),
      body: ListView(
        children: <Widget>[
          customcard("记单词", images[0], des[0]),
          customcard("测试", images[1], des[1]),
          customcard("生词本", images[2], des[2]),
        ],
      ),
    );
  }
}
