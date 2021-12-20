import 'package:flutter/material.dart';
import 'package:quizstar/api.dart';
import 'package:quizstar/word_detail.dart';

class WordList extends StatefulWidget {
  WordList({Key? key, required this.name}) : super(key: key);

  final String name;
  @override
  _WordListState createState() => _WordListState();
}

class _WordListState extends State<WordList> {
  List<String> _list = [];

  Future<void> _loadMoreData() async {
    var data = await wordList(widget.name);
    if (mounted) {
      setState(() {
        _list = data;
      });
    }
  }

  Widget _createItem(BuildContext context, int index) {
    var word = _list[index];
    return Card(
      margin: EdgeInsets.all(8),
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(4.0),
      ),
      elevation: 2,
      child: ListTile(
        onTap: () {
          Navigator.of(context).push(MaterialPageRoute(
            builder: (context) => WordDetail(
              words: _list,
              word: word,
            ),
          ));
        },
        title: Text(word),
      ),
    );
  }

  @override
  void initState() {
    _loadMoreData();
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("列表"),
      ),
      body: ListView.builder(itemCount: _list.length, itemBuilder: _createItem),
    );
  }
}
