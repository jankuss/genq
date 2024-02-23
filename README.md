# genq: Blazing Fast Data Class Generation for Dart

![](./docs/preview.gif)

**Tired of waiting for build_runner to churn through your codebase?**

genq cuts through the wait and generates data classes for your Dart projects in seconds, not minutes. ⚡️

Inspired by **freezed**, genq offers a familiar syntax for defining data classes, but with a focus on **lightning-fast performance**. No need to write copyWith, toString, or equals methods by hand - genq does it all for you.

**Here's what genq brings to the table:**

* **Reduce Boilerplate:** genq generates the boilerplate code for you, so you can focus on what matters:
    * **copyWith:** Create copies of your data classes with modified fields.
    * **toString:** Generate human-readable string representations of your data classes.
    * **==**: Compare data classes for equality.
* **Speed Demon:** Generate data classes in a flash, even for large projects.
* **Simple and Familiar:** Syntax similar to freezed, making it easy to learn and use.

## `genq` vs `build_runner` + `freezed`

`build_runner` + `freezed` 🐌 | `genq` 🚀
:-------------------------:|:-------------------------:
![build_runner](./docs/freezed.png) | ![genq](./docs/genq.png)

In this benchmark (located in `./benchmarks`), _count_ is the number of files in the benchmark, containing 250 annotated classes each. So for example, _count=1_ means 250 classes, _count=2_ means 500 classes, and so on. For count 10, `build_runner` and `freezed` took around 46 seconds, while `genq` took 0.11 seconds. **This is a >400x speedup!**

#### Notes

1. Never trust a benchmark you didn't falsify yourself.
2. genq is optimized to perform one task and one task only - data class generation, whereas build_runner is built to do a lot more. Take this into account when choosing between the two.

## Getting started is easy

### 1. Install

Install genq via brew using the following command:

```
brew install jankuss/genq/genq
```

Or download the latest release from the [releases page](https://github.com/jankuss/genq/releases).

### 2. Add `genq` to your project

```
dependencies:
  # ... Other dependencies ...
  genq: ^0.2.0
```

### 3. Define your data classes

```dart
import 'package:genq/genq.dart';

part 'user.genq.dart';

@genq
class User with _$User {
  factory User({
    required String name,
    required int age,
  }) = _User;
}
```

### 4. Generate the code

Run the genq command in your project directory, and you will have your desired data classes generated in no time:

```
genq
```

## Future Plans

- [ ] JSON Serialization/Deserialization
- [ ] Extensibility
