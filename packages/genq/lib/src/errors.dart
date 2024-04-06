class InvalidJsonEnumValueError extends Error {
  final Object value;
  final Type enumType;

  InvalidJsonEnumValueError(this.value, this.enumType);

  @override
  String toString() {
    return 'Invalid JSON enum value: $value for enum type: $enumType';
  }
}
