import 'package:flutter/material.dart';
import 'package:quizstar/word_list.dart';

class WordAZ extends StatefulWidget {
  // AZView({Key? key}) : super(key: key);

  @override
  _WordAZState createState() => _WordAZState();
}

class _WordAZState extends State<WordAZ> {
  List<String> az = <String>[
    "A",
    "B",
    "C",
    "D",
    "E",
    "F",
    "G",
    "H",
    "I",
    "G",
    "K",
    "L",
    "M",
    "N",
    "O",
    "P",
    "Q",
    "R",
    "S",
    "T",
    "U",
    "V",
    "W",
    "X",
    "Y",
    "Z"
  ];

  Widget _createItem(BuildContext context, int index) {
    var word = az[index];
    return Card(
      margin: EdgeInsets.all(8),
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(4.0),
      ),
      elevation: 2,
      child: Row(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          InkWell(
            onTap: () {
              Navigator.of(context).push(MaterialPageRoute(
                builder: (context) => WordList(name: word),
              ));
            },
            child: Text(
              word,
              style: TextStyle(fontSize: 30),
            ),
          ),
        ],
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("单词分组"),
      ),
      body: GridView.builder(
          itemCount: az.length,
          gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
              crossAxisCount: 4, //横轴三个子widget
              childAspectRatio: 2 //宽高比为1时，子widget
              ),
          itemBuilder: _createItem),
    );
  }
}
