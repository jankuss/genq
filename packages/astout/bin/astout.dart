import 'package:analyzer/dart/analysis/features.dart';
import 'package:analyzer/dart/analysis/utilities.dart';
import 'package:args/args.dart';
import 'package:analyzer/dart/ast/ast.dart';

import 'package:astout/ast_writer.dart';

 main(List<String> arguments) {
  final ast = parseFile(path: arguments[0], featureSet: FeatureSet.latestLanguageVersion());

  final buffer = StringBuffer();
  final html = AstWriter(buffer);
  ast.unit.visitChildren(html); 

  print(buffer.toString());
}
