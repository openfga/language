# Self Error

**Error Code:** `self-error`

**Category:** Naming Validation

## Summary

`self-error` is a reserved error code in the `ValidationError` enum. It is **not currently emitted** by the validator.

## Description

Using `self` or `this` as a type or relation name is a real validation error, but it is reported under the dedicated keyword codes rather than `self-error`:

- A reserved **type** name (`self`/`this`) raises [`reserved-type-keywords`](./reserved-type-keywords.md) with the message `a type cannot be named 'self' or 'this'.`
- A reserved **relation** name (`self`/`this`) raises [`reserved-relation-keywords`](./reserved-relation-keywords.md) with the message `a relation cannot be named 'self' or 'this'.`

The `self-error` code is defined in both the Go and JavaScript error enums for forward compatibility but is not raised by any validation path today. This page exists to document that fact and point to the codes that actually fire.

## Related Errors

- [`reserved-type-keywords`](./reserved-type-keywords.md) - Raised when `self`/`this` is used as a type name
- [`reserved-relation-keywords`](./reserved-relation-keywords.md) - Raised when `self`/`this` is used as a relation name
- [`invalid-name`](./invalid-name.md) - General invalid naming issues

## Implementation Notes

The `self-error` code is declared in:
- Go: `pkg/go/validation/errors.go`
- JavaScript: `pkg/js/errors.ts`

No code path in either implementation currently raises it; the `self`/`this` naming checks live in the name-validation phase and emit the `reserved-*-keywords` codes above.
