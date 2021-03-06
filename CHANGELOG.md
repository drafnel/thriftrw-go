Releases
========

v1.5.0 (unreleased)
-------------------

-   No changes yet.


v1.4.0 (2017-07-21)
-------------------

-   Added support for `go.tag` annotations on struct fields. Corresponding
    fields of the generated Go structs will be tagged with these values.
-   AST: Added a `LineNumber` function to get the line on which an AST Node was
    defined.


v1.3.0 (2017-07-05)
-------------------

-   Plugins: Added support for overriding the communication channels used by
    plugins.
-   Plugins: Template output is now compliant with gofmt.
-   Enums now implement the `encoding.TextUnmarshaler` interface. This allows
    retrieving enum values by name and integrates better with other encoding
    formats.
-   Enums now have a `<EnumName>_Values()` function which returns all known
    values for that enum.


v1.2.0 (2017-04-17)
-------------------

-   The Thrift IDL is now embedded withing the generated package. It is
    accessible via the package global ThriftModule.


v1.1.0 (2017-03-23)
-------------------

-   `Equals()` methods are now generated for all custom types.
-   Added flags `--no-types`, `--no-constants`, `--no-service-helpers`.
-   Plugins: Fail code generation if communication with a plugin fails to
    disconnect properly.
-   Handle `nil` values in generated `String()` methods.
-   Fixed conflicts in helper functions when imported types had names similar
    to locally defined types.
-   AST: All type references and complex constant values now record the line
    numbers on which they were specified.
-   AST: Added a Node interface to unify different AST object types.
-   AST: Added Walk to traverse the AST using a Visitor.


v1.0.0 (2016-11-14)
-------------------

-   Field names `ToWire`, `FromWire`, `String`, and for exceptions `Error` are
    now reserved. Override the names of these fields in the generated code with
    the `go.name` annotation.
-   Plugins: The version of ThriftRW used to compile the plugin is now matched
    against the version actually generating code.


v0.5.0 (2016-11-10)
-------------------

-   **Breaking**: Generated enums now have first class JSON support. Enums are
    (un)marshalled from/to strings if possible with fallback to integer for
	unrecognized values.
-   **Breaking**: `Args`, `Result`, and `Helper` types for service functions
    are now generated in the same package as the user-defined types. These
    types are now named similarly to `$service_$function_Args` where `$service`
    and `$function` are the names of the Thrift service and function normalized
    based on Go naming conventions.
-   Code generation will abort if struct fields, after conversion to Go style
    names, are not unique in the structure.
-   A `go.name` annotation may now be specified to override the names of
    entities in the generated Go code. The annotation is supported for struct,
    union, and exception types, and their fields, enum types and enum items,
    and parameters of functions.
-   Plugins: Renamed `Service.Name` to `Service.ThriftName` since it contains
    the name of the service as it appeared in the Thrift file.
-   Plugins: Added a new `Service.Name` field which contains the name of field
    normalized per Go naming conventions. This, along with `Function.Name` may
    be used to build the names of the `Args`, `Result`, and `Helper` types for
    a function.
-   Plugins: Removed `Service.Directory` and `Service.ImportPath` because these
    are now same as the corresponding module.
-   Plugins: Constructors for `Plugin` and `ServiceGenerator` clients and
    handlers are now exposed in the same package as the interfaces.
-   Non-primitive types constants are now inlined in the generated Go code
    instead of being referenced in an effort to reduce the impact of user
    errors on the generated code. This is because non-primitive constants were
    previously implemented as global `var`s which might be modified by
    user code.


v0.4.0 (2016-11-01)
-------------------

-   **Breaking**: Remove the `--yarpc` flag. Install the ThriftRW YARPC plugin
    from `go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc` and use
    `--plugin=yarpc` instead.
-   **Breaking**: The `compile` API now exposes annotations made while
    referencing native Thrift types. This changes the `TypeSpec`s for primitive
    types from values to types.
-   The `compile` API now also exposes annotations for `typedef` declarations.
-   Generate args structs and helpers for oneway functions.
-   Expose whether a function is oneway to plugins.
-   Expose the version of the library under `go.uber.org/thriftrw/version.Version`.
-   Generated code will test for version compatibility with the current version
    of ThriftRW during initialization.


v0.3.2 (2016-10-05)
-------------------

-   Fix import paths for code generated using `--yarpc`. Note that this flag
    will be removed in a future version.


v0.3.1 (2016-09-30)
-------------------

-   Fix missing canonical import path to `go.uber.org/thriftrw`.


v0.3.0 (2016-09-29)
-------------------

-   **Breaking**: Renamed project to `go.uber.org/thriftrw`.
-   **Breaking**: Keywords reserved by Apache Thrift are now disallowed as
    identifiers in Thrift files.
-   **Breaking**: The `Package` field of the `plugin.TypeReference`,
    `plugin.Service`, and `plugin.Module` structs was renamed to `ImportPath`.


v0.2.1 (2016-09-27)
-------------------

-   Plugin templates: Fixed a bug where imports in templates would use the base
    name of the package even if it had a hyphen in it if it wasn't available on
    the `GOPATH`.
-   Plugin templates: Imports in generated code are now always qualified if the
    package name doesn't match the base name.


v0.2.0 (2016-09-09)
-------------------

-   Added a `-v`/`--version` flag.
-   Added a plugin system.

    ThriftRW now provides a plugin system to allow customizing code generation.
    Initially, only the generated code for `service` declarations is
    customizable. Check the documentation for more details.
-   **Breaking**: Fixed a bug where all-caps attributes that are not known
    abbreviations were changed to PascalCase.
-   **Breaking**: The `String()` method for `enum` types now returns the name
    of the item as specified in the Thrift file.


v0.1.0 (2016-08-31)
-------------------

-   Initial release.
