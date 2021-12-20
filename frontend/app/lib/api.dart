import 'package:dio/dio.dart';
import 'package:quizstar/word_detail.dart';

Dio dio = Dio();

// String basicUrl = "http://localhost:8006";
String basicUrl = "http://120.24.163.74:8006";

Future<dynamic> wordList(word) {
  return dio.get('$basicUrl/word-english/word-list/$word').then((res) {
    return res.data["data"].cast<String>();
  });
}

Future<Detail> wordDetial(word) {
  return dio.get('$basicUrl/word-english/word-detail/$word').then((res) {
    if (res.data["status"] == "failure") {
      var data = new Detail();
      data.voice = "";
      data.meaning = "";
      data.data = [];
      return data;
    }
    return Detail.fromJSon(res.data);
  });
}
