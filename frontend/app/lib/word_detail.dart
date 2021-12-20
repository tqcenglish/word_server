import 'dart:math';

import 'package:flutter/material.dart';
import 'package:quizstar/api.dart';

class WordDetail extends StatefulWidget {
  WordDetail({Key? key, required this.words, required this.word})
      : super(key: key);

  final List<String> words;
  final String word;
  @override
  _WordDetailState createState() => _WordDetailState();
}

class _WordDetailState extends State<WordDetail> {
  String _word = "";
  bool showDetail = false;
  late Detail _wordDetail;

  Future<void> _loadMoreData(String word) async {
    this._word = word;
    var res = await wordDetial(word);
    if (res.voice != "") {
      if (mounted) {
        setState(() {
          _wordDetail = res;
        });
      }
    } else {
      random();
    }
  }

  void random() {
    var rng = new Random();
    var newWord = widget.words[rng.nextInt(widget.words.length)];
    _loadMoreData(newWord);
  }

  Widget _createItem(BuildContext context, int index) {
    var path = _wordDetail.data[index];
    return Card(
      margin: EdgeInsets.all(40),
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(10.0),
      ),
      elevation: 2,
      child: Row(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          InkWell(
              onTap: () {},
              child: Image.network('$basicUrl/word-static/$path')),
        ],
      ),
    );
  }

  @override
  void initState() {
    _loadMoreData(widget.word);
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
          actions: [
            IconButton(
                onPressed: () {
                  // 随机获取新单词
                  random();
                },
                icon: Icon(Icons.skip_next_sharp))
          ],
          title: InkWell(
            onTap: () {
              setState(() {
                showDetail = !showDetail;
              });
            },
            child: Text(_word),
          )),
      body: showDetail
          ? Center(
              child: Text(_wordDetail.meaning),
            )
          : Container(
              child: _wordDetail.data.length > 0
                  ? GridView.builder(
                      itemCount: _wordDetail.data.length,
                      gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
                          crossAxisCount: 2, //横轴2个子widget
                          childAspectRatio: 1 //宽高比为1时，子widget
                          ),
                      itemBuilder: _createItem)
                  : Center(
                      child: Text("图片被一群老鼠偷走了"),
                    )),
    );
  }
}

class Detail {
  late List<String> data;
  late String voice;
  late String meaning;

  static Detail fromJSon(Map<String, dynamic> input) {
    return new Detail()
      ..data = input["data"].cast<String>()
      ..voice = input["voice"]
      ..meaning = input["meaning"];
  }
}
