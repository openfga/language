# Schema Version Unsupported

**Error Code:** `schema-version-unsupported`

**Category:** Schema Validation

## Summary

The specified schema version is not supported by the current OpenFGA implementation, preventing proper model validation and execution.

## Description

OpenFGA supports specific schema versions that define the available features, syntax rules, and validation behavior. Each schema version represents a specific set of capabilities:

- **Schema 1.0**: Original feature set with basic relation definitions, no type restrictions (no longer supported by language or OpenFGA, transformation only supported in [npm:@openfga/syntax-transformer@v0.1.6](https://www.npmjs.com/package/@openfga/syntax-transformer/v/0.1.6) and below)
- **Schema 1.1**: Enhanced features including conditions, advanced wildcard patterns, improved validation
- **Schema 1.2**: Interchangeable with 1.1, additionally allows modules

Using an unsupported schema version prevents the system from knowing which validation rules to apply and which features are available.

## Example 1: Using a version that is not yet supported

The following model would trigger this error:

```
model
  schema 2.0  # Error: Version 2.0 not yet supported

type user

type document
  relations
    define viewer: [user]
```

**Error Location:** Line 2, `schema 2.0` declaration.

**Error Message:** `Schema version '2.0' is not supported. Supported versions are: 1.1, 1.2`

## Resolution

Use a supported schema version:

### Option 1: Use latest supported version (recommended)

```
model
  schema 1.1  # Use latest supported version

type user

type document
  relations
    define viewer: [user]
```

### Steps to fix:

1. **Check supported versions:**
   - Review the error message for list of supported versions
   - Check OpenFGA documentation for version compatibility

2. **Choose appropriate version:**
   - Do not use `1.0` it is no longer supported by OpenFGA
   - Use `1.1` or `1.2` for new models (recommended)
   - Use `1.2` when using modules

3. **Update schema declaration:**
   - Change to a supported version
   - Ensure your model features are compatible with chosen version

4. **Verify model compatibility:**
   - Check that all features used are supported in chosen schema version
   - Test model validation with updated version

## Feature Compatibility Matrix

| Feature                              | Schema 1.0 | Schema 1.1 | Schema 1.2 |
|--------------------------------------|------------|------------|------------|
| Supported                            | ❌          | ✅          | ✅          |
| Basic relations                      | ✅          | ✅          | ✅          |
| Simple wildcards                     | ✅          | ✅          | ✅          |
| Wildcards                            | ✅          | ✅          | ✅          |
| Basic operations                     | ✅          | ✅          | ✅          |
| Type Restrictions                    | ❌          | ✅          | ✅          |
| Conditions                           | ❌          | ✅          | ✅          |
| Operator grouping `(a or (b and c))` | ❌          | ✅          | ✅          |
| Modules                              | ❌          | ❌          | ✅          |

## Version Migration Guidelines

### Migrating from Schema 1.0 to 1.1

The main changes when migrating from Schema 1.0 to 1.1 include:
1. [Adding model schema version field](#-model-schema-versions)
2. [Adding type enforcements and removing need to specify `as self`](#type-enforcements--removing-as-self)
3. [Disallowing string literals in user_ids](#disallowing-string-literals-in-user_ids)
4. [Enforcing userset type restrictions](#enforcing-userset-type-restrictions)
5. [Requiring you to specify for which relations you can write tuples with public access (using ‘`*`’)](#public-access)
6. [Changes in query evaluation behavior with type restrictions](#query-evaluation-behavior-with-type-restrictions)

To facilitate migration to the new DSL schema, you will need to update tuples that are no longer valid. In particular, all tuples with wildcard (`*` or `user:*`) user field defined with model schema 1.0 <u><strong>MUST</strong></u> be deleted and re-added back.

#### Model Schema Versions

Since the changes in the DSL are significant, we have decided to add a schema version to the DSL. The previous version of the DSL’s schema was `1.0`, and the new schema version will be `1.1`. To use the new syntax please add the following to the top of the model:

```
model
  schema 1.1
```

#### Type Enforcements & Removing as self

We’ll use the following version 1.0 model and tuples to illustrate the changes we’ll need to make:

```python
model
  schema 1.0

type user

type group
  relations
    define member as self # `as self` is no longer supported and must be replaced with type restrictions

type folder
  relations
    define parent as self
    define viewer as self or viewer from parent
    
type document
  relations
    define parent as self
    define viewer as self
    define can_read as viewer or viewer from parent
```

With the above model, we can write tuples like:
```yaml
- user: 'user:bob'
  relation: 'member'
  object: 'group:sales'
- user: 'folder:sales'
  relation: 'parent'
  object: 'document:pricing'
- user: 'group:sales#member'
  relation: 'viewer'
  object: 'folder:sales'
- user: 'user:john'
  relation: 'viewer'
  object: 'document:pricing'
```

Those tuples match the intent of how the model was designed, but without type restrictions (introduced in `1.1`) we can also write tuples that would not make as much sense. For example, we can say that a document is a member of the sales group:
```yaml
- user: 'document:pricing' # This is a valid tuple on schema 1.0, but does not make sense
  relation: 'member'
  object: 'group:sales'
```

To be able to better validate tuples and make the model more readable, version 1.1 requires you to specify type restrictions for all the relations that were previously assignable (e.g. relations defined `as self` in any way), and it removes the `as self` keyword.

The model above needs to be rewritten as below:

> Mote: The `as` has also been replaced with `:` to separate the relation name from its definition.

```python
model
  schema 1.1

type user

type group
  relations
    define member: [user]    # Replace `as self` with type restriction: user

type folder
  relations
    define parent: [folder]  # Replace `as self` with type restriction: folder
    define viewer: [user] or viewer from parent # Replace `as self` with type restriction: user
    
type document
  relations
    define parent: [folder]  # Replace `as self` with type restriction: folder
    define viewer: [user]    # Replace `as self` with type restriction: user
    define can_read: viewer or viewer from parent
```

After making these changes, OpenFGA will start validating the tuples more strictly, for example, you won’t be able to assign a `document` as a member of a `group`. If your application is writing invalid tuples, you’ll start getting errors when invoking the `Write` API.

## Disallowing String Literals in user_ids

With version 1.0 models, you could write a tuple where the user id did not specify a type, for example:

```yaml
- user: 'bob'
  relation: 'member'
  object: 'group:sales'
```

However, with version 1.1 you always need to specify an object, so “bob’” is no longer a valid identifier. If you don’t have a type in your model that defines relations for users, you can add a ‘user’ type with no relations to your model, for example:

```python
model
  schema 1.1
  
type user  # Define a user type with no relations
```
You can then use that type when writing tuples:

```yaml
- user: 'user:bob'  # Specify the user type explicitly
  relation: 'member'
  object: 'group:sales'
```

## Enforcing Userset Type Restrictions

With the model above, the following tuples will be valid according to the type definitions:

```yaml
- user: 'user:bob'
  relation: 'member'
  object: 'group:sales'
- user: 'folder:sales'
  relation: 'parent'
  object: 'document:pricing'
- user: 'user:john'
  relation: 'viewer'
  object: 'document:pricing'
```

However, the one below will not be valid, as we can’t assign `group:sales#member` to the viewer relationship of a folder.

```yaml
- user: 'group:sales#member'
  relation: 'viewer'
  object: 'folder:sales'
```

You might think that given `group:sales#member` are actually users, you should still be able to assign it. OpenFGA calls expressions like `group:sales#member` ["usersets"](https://openfga.dev/docs/concepts#what-is-a-user), and with our model we can only assign users.

The issue is that there are a lot of other usersets that you don't want to be assigned as viewers of a folder. For example, you would not want to add `document:pricing#viewer` as viewers of the folder as conceptually it does not make sense to say “every viewer of this document should be a viewer of this folder”.

To allow these tuples to be written, you need to specify `group#member` as a valid type for the folder’s viewer relationship. You would want to do the same with the document’s viewer relationship if you want to define that the members of a group can be viewers of a document:

```python
model
  schema 1.1

type user

type group
  relations
    define member: [user]                # Replace `as self` with type restriction: user

type folder
  relations
    define parent: [folder]              # Replace `as self` with type restriction: folder
    define viewer: [user, group#member] or viewer from parent # Replace `as self` with type restriction: user, group#member
    
type document
  relations
    define parent: [folder]              # Replace `as self` with type restriction: folder
    define viewer: [user, group#member]  # Replace `as self` with type restriction: user, group#member
    define can_read: viewer or viewer from parent
```

You can identify which usersets you need to add by looking at tuples in your store that have the following structure:

```yaml
- user: 'group:sales#member'
  relation: 'viewer'
  object: 'folder:sales'
```

If you find a tuple like that, you’ll need to add `group#member` in the list of types allowed in the `viewer` relation of the `folder` type.

## Public Access

When using version 1.0, you can indicate public access to specific objects by specifying a wildcard user in a relationship to any object, e.g.:

```yaml
- user: '*'
  relation: 'viewer'
  object: 'document:pricing'
```

When you write the tuple above, all users are granted with the “viewer” relationship for the “pricing" document. You can write those kinds of tuples for any relation that is <ProductConcept section="what-are-direct-and-implied-relationships" linkName="directly assignable" /> in the model.

In version 1.1 we want to be more explicit about the tuples you can write, so you’ll need to declare in the DSL which relations allow wildcards and for which object types. If we want to let any object of type “user” to be a viewer of a specific document, we’ll need to explicitly define it.

```python
model
  schema 1.1

type user

type document
  relations
    define viewer: [user, user:*]  # Allow any user or wildcard user to be a viewer
```

You’ll need to specify `user:*` as the user value in the tuple to enable this:

```yaml
- user: 'user:*'
  relation: 'viewer'
  object: 'document:pricing'
```

Being explicit about the wildcard type restrictions also lets you model scenarios like “all employees can see this document, but not all external users”, “all user accounts can access this document, but not service/machine-to-machine accounts”.

This change implies that you’ll need to change your code to write tuples with this new syntax, and that you’ll need to migrate existing tuples to use the new format.

You might have 3 kinds of tuples in your model that use “*”, with different migration strategies:

1. Tuples that have user = “*”

   You would need to retrieve those tuples and write them using the proper type (e.g. `user:*`). To retrieve them, you’ll need to use the Read endpoint, filter on your side the tuples that have `user = “*”`, and call the Write API for each one, with the proper type, e.g:
   
   ```yaml
   - user: 'user:*'
     relation: 'viewer'
     object: 'document:pricing'
   ```

2. Tuples that have `user = "employee:*”`, where `employee` is NOT a type that is defined in the new iteration of your model.

   If you have tuples with this format, they will be considered invalid because they don’t have a corresponding type in the model. If you need such a type defined, you’ll need to add it to the model, and the scenario will be similar to the one described below.

3. Tuples that have `user = “user:*”`, which would mean "the user with user_id = '*'”, where `user` is type that is defined in the new iteration of your model.

   In this case, the meaning of the tuple will change. If you were intending to specify "a user with user id = *", you will need to encode it in a different way instead of using “*”. If you intended to specify “every user has this relationship with this object” then it’s not the way it would have worked with schema version = 1.0, but it will work with version = 1.1.

> ⚠️ Warning
> If you have any wildcard tuples (i.e., `*` or `user:*`) that were created with model schema 1.0, you _**must**_ delete and re-add these tuples with the appropriate type. This allows OpenFGA to interpret these tuples appropriately with the model schema 1.1 semantics. Failure to delete and re-add may cause OpenFGA to interpret these tuples incorrectly.

## Query Evaluation Behavior with Type Restrictions

When you make changes to a model that already has tuples, those tuples might become invalid. Some cases where this can happen are:

- If you rename/delete a type.
- If you rename/delete a relation.
- If you remove types from the list of allowed types in a relation, including changes for Public Access.
- If OpenFGA introduces a change that makes a tuple invalid.

In these cases, OpenFGA will not consider those invalid tuples when evaluating queries (check, expand, list-objects, etc). However, after any of the changes above happens, you should delete those tuples as having a large number of invalid tuples will negatively affect performance.


## Related Errors

- [`schema-version-required`](./schema-version-required.md) - When no schema version is specified
- [`invalid-schema-version`](./invalid-schema-version.md) - When version format is invalid
- [`invalid-schema`](./invalid-schema.md) - General schema structure issues

## Implementation Notes

This validation is enforced consistently across:
- Go implementation: `pkg/go/validation/schema_validation.go`
- JavaScript implementation: `pkg/js/validator/validate-dsl.ts`
- Java implementation: Java schema validation package

The validation checks against a predefined list of supported schema versions and rejects models using unsupported versions.
