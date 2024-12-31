# Sum Interface Generator

## Usage

```
sig <package> <file contains spec>"
```

`sum.go` and `iface.go` will be generated.

## Sum Spec Format

one spec each line.

```
ExportedInterfaceName: internalSumInterfaceName: Type { | Type }
```

- `ExportedInterfaceName` will be written in `iface.go`. User should use it and user can add additional methods.
- `internalSumInterfaceName` will be written in `sum.go`. It can not be edited.

## Example

For the following string, generate interface code.

```
GlobalDef: _sum: *A | *B
```

The following code is generated and should not be changed.

```go
// in sum.go
// Code generated DO NOT EDIT.
package <package>
type _sum interface {
    _sum() // not exported dummy method
}
func (*A) _sum() {}
func (*B) _sum() {}

// in iface.go
package <package>
type GlobalDef interface {
    _sum
    // client can add additional methods
}
```

modify `GlobalDef` still need to add method definition.

`_sum` is for internal package usage(for switch case).

## Design Reason

To be more extensible, the `_sum` interface should not be exported. Instead use a wrapper public interface that is not generated.

- add some additional methods to the interface. Define a new interface that embeds the `_sum` interface.


