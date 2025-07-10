# OpenFGA Validation Error Troubleshooting Guide

This guide provides quick solutions to common OpenFGA validation errors. For detailed documentation on each error, see the individual error documents linked below.

## Quick Reference

### Most Common Errors

| Error | Quick Fix | Link |
|-------|-----------|------|
| `schema-version-required` | Add `schema 1.1` after `model` | [Details](./schema-version-required.md) |
| `undefined-relation` | Define the missing relation or fix typo | [Details](./undefined-relation.md) |
| `undefined-type` | Define the missing type or fix typo | [Details](./undefined-type.md) |
| `invalid-name` | Use lowercase with underscores only | [Details](./invalid-name.md) |
| `duplicated-error` | Remove or rename duplicate definitions | [Details](./duplicated-error.md) |

### Schema and Structure Issues

| Error | Quick Fix                                    | Link |
|-------|----------------------------------------------|------|
| `invalid-syntax` | Check indentation and keyword spelling       | [Details](./invalid-syntax.md) |
| `invalid-schema` | Ensure proper `model` and `schema` structure | [Details](./invalid-schema.md) |
| `schema-version-unsupported` | Use supported version (`1.1` or `1.2`)           | [Details](./schema-version-unsupported.md) |
| `invalid-schema-version` | Use format `X.Y` (e.g., `1.1`)               | [Details](./invalid-schema-version.md) |

### Relationship and Reference Errors

| Error | Quick Fix | Link |
|-------|-----------|------|
| `relation-no-entry-point` | Add `[user]` or direct assignment | [Details](./relation-no-entry-point.md) |
| `cyclic-relation` | Break circular references | [Details](./cyclic-relation.md) |
| `invalid-relation-type` | Fix `type#relation` syntax | [Details](./invalid-relation-type.md) |
| `tupleuserset-not-direct` | Add direct assignment to tupleset relation | [Details](./tupleuserset-not-direct.md) |

### Naming and Keyword Issues

| Error | Quick Fix | Link |
|-------|-----------|------|
| `self-error` | Don't use `self` or `this` as names | [Details](./self-error.md) |
| `reserved-type-keywords` | Use different type names | [Details](./reserved-type-keywords.md) |
| `reserved-relation-keywords` | Use different relation names | [Details](./reserved-relation-keywords.md) |

### Wildcard and Advanced Features

| Error | Quick Fix | Link |
|-------|-----------|------|
| `invalid-wildcard-error` | Use correct `[type:*]` syntax | [Details](./invalid-wildcard-error.md) |
| `type-wildcard-relation` | Don't mix wildcard with relation | [Details](./type-wildcard-relation.md) |

### 🎯 Condition Errors

| Error | Quick Fix | Link |
|-------|-----------|------|
| `condition-not-defined` | Define the missing condition | [Details](./condition-not-defined.md) |
| `condition-not-used` | Remove unused condition or use it | [Details](./condition-not-used.md) |
| `different-nested-condition-name` | Fix condition name consistency | [Details](./different-nested-condition-name.md) |

### Multi-file and Module Errors

| Error | Quick Fix | Link |
|-------|-----------|------|
| `multiple-modules-in-file` | Split into separate files | [Details](./multiple-modules-in-file.md) |

## 🛠️ Step-by-Step Troubleshooting

### 1. **Start with Schema Issues**
```bash
# Check for these first - they prevent other validation
- Missing or invalid schema version
- Incorrect model structure
- Basic syntax errors
```

### 2. **Fix Naming Problems**
```bash
# Common naming issues to check
- Type/relation names with uppercase, spaces, or special characters
- Use of reserved keywords like 'self', 'this', 'model'
- Names starting with numbers
```

### 3. **Resolve Reference Issues**
```bash
# Check all references exist
- Types referenced in relations exist
- Relations referenced in computed usersets exist
- Conditions referenced in relations exist
```

### 4. **Address Structural Problems**
```bash
# Check authorization model structure
- Relations have entry points (direct assignments)
- No circular dependencies
- Proper tuple-to-userset structure
```

## Emergency Fixes

### **Model Won't Parse at All**
```
1. Check schema declaration: `model` followed by `schema 1.1`
2. Verify indentation (2 spaces per level)
3. Check for typos in keywords: `type`, `relations`, `define`
```

### **Multiple Undefined Errors**
```
1. Look for typos in type/relation names
2. Check case sensitivity (use lowercase)
3. Verify all referenced types are defined
```

### **Authorization Not Working**
```
1. Check for `relation-no-entry-point` errors
2. Verify direct assignments exist: `[user]`
3. Look for circular dependencies
```

## Validation Checklist

Before deploying your authorization model, verify:

- [ ] Schema version is specified (`schema 1.1`)
- [ ] All type names use lowercase and underscores
- [ ] All relation names use lowercase and underscores
- [ ] No reserved keywords used as names
- [ ] All referenced types exist
- [ ] All referenced relations exist
- [ ] Every relation has an entry point
- [ ] No circular dependencies
- [ ] Conditions are defined if used
- [ ] Multi-file modules are properly organized

## Best Practices

### **Naming Conventions**
- Use descriptive, business-oriented names
- Stick to lowercase with underscores
- Avoid technical jargon in favor of domain terms

### **Model Structure**
- Start simple, add complexity gradually
- Use clear hierarchical relationships
- Document complex permission logic

### **Testing Strategy**
- Validate model after each change
- Test with sample authorization data
- Use YAML test cases for regression testing

## Advanced Debugging

### **Use The FGA CLI**

You can get the FGA CLI from: https://github.com/openfga/cli

1. [Quick model validation](https://github.com/openfga/cli/#validate-an-authorization-model)
    ```bash
    fga model validate your-model.fga
    ```

2. [Proper validation with tests and expectations](https://github.com/openfga/cli/#run-tests-on-an-authorization-model) 
    ```bash
    fga model test your-model.fga --test-file tests.yaml
    ```

### **Common Pattern Issues**
- Overly complex nested relationships
- Missing direct assignments in hierarchies
- Inconsistent permission patterns

## Getting Help

- **Documentation**: See individual error documents for detailed explanations
- **Examples**: Check the validation test cases for working examples as well as the OpenFGA [sample stores](https://github.com/openfga/sample-stores).
- **Community**: [OpenFGA community](https://openfga.dev/community) for advanced use cases

---

*This troubleshooting guide covers the most common validation scenarios. For detailed information about any specific error, click the links to view the comprehensive error documentation.*
