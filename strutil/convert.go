package strutil

/*
TypeToString converts a value to its string representation.

If value isn't a string and no custom fn is provided, it returns nil.

The optional `fn` arg allows for custom string conversion.

This function should not throw errs or panics; it must returns nil if a conversion fails.
*/
func TypeToString[T any](value T, fn func(T) *string) *string {
	if fn == nil {
		if s, ok := any(value).(string); ok {
			return &s
		}
		return nil
	}

	return fn(value)
}
