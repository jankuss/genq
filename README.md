# genq: Blazing Fast Data Class Generation for Dart

![](./docs/preview.gif)

**Tired of waiting for build_runner to churn through your codebase?**

genq cuts through the wait and generates data classes for your Dart projects in milliseconds, not minutes. ⚡️

Inspired by **freezed**, genq offers a familiar syntax for defining data classes, but with a focus on **lightning-fast performance**. No need to write copyWith, toString, or equals methods by hand - genq does it all for you.

**Here's what genq brings to the table:**

* **Reduce Boilerplate:** genq generates the boilerplate code for you, so you can focus on what matters.
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

Read more about defining your data classes [here](#defining-data-classes).

### 4. Generate the code

Run the genq command in your project directory, and you will have your desired data classes generated in no time:

```
genq
```

## Defining Data Classes

To define a data class, first you need to annotate your class with `@genq`. Also, you need to use the generated mixin with the name `_$<ClassName>`.
Additionally, you need to define a factory constructor with named parameters and a redirecting constructor with the name `_<ClassName>`.

At first, you might get errors from the IDE because the generated file is not yet created. After running `genq`, the errors should disappear.

```dart
import 'package:genq/genq.dart'; // <- Import genq

part 'user.genq.dart'; // <- Add a part directive to the generated file

@genq // <- Annotate the class with @genq
class User with _$User { // <- Add the mixin _$<ClassName>
  factory User({ // <- Define a factory constructor
    required String name, // <- Define fields as named parameters
    required int age,
  }) = _User; // <- Redirecting constructor, _<ClassName>
}
```

The generated class will have the following methods:
- `copyWith`: Create a copy of the data class with modified fields.
- `toString`: Generate a human-readable string representation of the data class.
- `==`: Compare two data classes for equality.

## JSON Serialization/Deserialization

To generate JSON serialization/deserialization code, you need to use the `@Genq(json: true)` annotation instead of `@genq`. 

```dart
import 'package:genq/genq.dart';

part 'user.genq.dart';

@Genq(json: true)
class User with _$User {
  factory User({
    @JsonKey(name: 'full_name')
    required String name,
    required int age,
  }) = _User;
}
```

This will generate two public functions, which you can use to serialize/deserialize the data class to/from JSON:

```dart
$UserFromJson(Map<String, dynamic> json) => /* ... */;
$UserToJson(User value) => /* ... */;
```

### Customize JSON Serialization

You can customize the generated JSON serialization/deserialization code for fields using the `@JsonKey` annotation. 

```dart
import 'package:genq/genq.dart';

part 'user.genq.dart';

@Genq(json: true)
class User with _$User {
  factory User({
    // Customizing the JSON key for the field 'name'. When deserializing, the value of 'full_name' will be assigned to the 'name' field.
    @JsonKey(name: 'full_name')
    required String name,
    // Excluding the field 'age' from JSON serialization/deserialization.
    @JsonKey(includeFromJson: false, includeToJson: false)
    required int age,
  }) = _User;
}
```

### Enums

Enums are also supported for JSON serialization/deserialization. They need to be annotated with `@JsonEnum`.

```dart
import 'package:genq/genq.dart';

part 'user.genq.dart';

@JsonEnum()
enum Role {
  // You can annotate the enum values with @JsonValue to customize the JSON serialization/deserialization.
  // For example, the string 'ADMIN' will get deserialized to the Role.admin value and vice versa.
  @JsonValue('ADMIN')
  admin,
  @JsonValue('USER')
  user,
}
```

Similary to the @Genq(json: true) annotation, this will generate two public functions, which you can use to serialize/deserialize the enum to/from JSON:

```dart
$RoleFromJson(Object json) => /* ... */;
$RoleToJson(Role value) => /* ... */;
```

### Notes

The fundamental idea behind the JSON codegen for it to be fast and efficient is to **publicly** expose the generated functions. Based on this, genq can assume for a type
`T` that the functions `$TFromJson` and `$TToJson` are available, thus avoiding unnecessary traversal of other files.

## How?

genq uses its own subset parser of the dart language and generates code directly from the parsed AST. This allows genq to generate code much faster than `build_runner`, which uses the `analyzer` package. Code generation is also done in parallel for each file, which further speeds up the process.

Also, the code generator only cares about the information within the data class definition, which allows it to ignore the rest of the codebase.

### Notes on the subset parser

The subset parser is written for the specific structures of data classes as defined [here](#defining-data-classes). Thus, there may be parsing errors if the code does not follow the expected structure. While the parser is generally robust when encountering unparsable code, there may be cases where it fails to parse the code correctly. If you encounter such a case, please open an [issue](https://github.com/jankuss/genq/issues/new) with the code that caused the error.

## Downsides of `genq`

- `build_runner` is extensible & pluggable and can be used for a wide variety of tasks, whereas `genq` is focused on data class generation. Freezed for example leverages this to generate JSON Serialization code using `json_serializable`.
- `genq` is written in Go, so it does not neatly integrate with the Dart ecosystem.

## Future Plans

- [ ] JSON Serialization/Deserialization
- [ ] Extensibility
