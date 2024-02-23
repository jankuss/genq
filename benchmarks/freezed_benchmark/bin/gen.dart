import 'dart:io';

void main(List<String> args) {
  for (var i = 0; i < int.parse(args[0]); i++) { 
    final file = File('./lib/gen/gen$i.bench.dart');
    var str = 'import "package:freezed_annotation/freezed_annotation.dart";\n\n';
    str += "part 'gen$i.bench.freezed.dart';\n\n";

    for (var j = 0; j < 250; j++) {
      str += '''
@freezed
class User$j with _\$User$j {
  factory User$j({
    required String name,
    required int? age,
    required bool registered,
    required void Function() test,
  }) = _User$j;
}\n\n''';

      file.writeAsStringSync(str);
    }
  }
}
