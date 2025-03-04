package main

import (
	"fmt"
	"unsafe"
)

/* Explanation:

Each field in a struct is aligned according to its type's alignment requirements.
If necessary, Go adds padding between fields to ensure that each field starts at an address that is a multiple of its required alignment.
Padding in golang could be: 1, 2, 4, 8 bytes

---

// padding 8
struct 'A' size 32 bytes
- align attr1 (int8): 1 -> [x0000000](8)
- align attr2 (string): 8 -> [xxxxxxxx][xxxxxxxx](16)
- align attr3 (int8): 1 -> [x0000000](8)

// so, as you see, golang adding new padding (waste space)

// padding -8
struct 'B' size 24 bytes (padding - 8)
- align attr1 (string): 8 -> [xxxxxxxx][xxxxxxxx](16)
- align attr2 (int8): 1 -> [x0000000](8)
- align attr3 (int8): 1 -> [xx000000](8) (no new padding)

// in struct B attr2 and attr3 are int8 (1 byte) and they come one after another
// that is why align is less than in A struct, because same types that are
// come one after another adding their space to actual padding if it has needed space

---

Interesting things come out with type struct:

// padding 8
struct 'C' size 16 bytes (addr: 0x14000010030)
- align attr1 (struct {} - 0x14000010030): 1 -> [] (no new padding)
	- empty struct takes 0 bytes (and no new padding!), but still exists in memory
- align attr2 (string - 0x14000010030): 8 -> [xxxxxxxx][xxxxxxxx]
	- will be placed right after attr1 on same address, so o padding is needed because the empty struct does not introduce alignment constraints

// padding 8
struct 'D' size 24 bytes (addr: 0x1400000c018)
- align attr1 (string - 0x1400000c018): 8 -> [xxxxxxxx][xxxxxxxx](16)
- align attr2 (struct {} - 0x1400000c028): 1 -> [00000000](8)

// the difference is due to how golang treats the placement of an empty struct (struct{}):
// - struct C - struct{} as a first field doesn't affect alignment and takes zero bytes in space
// - struct D - golang ensures it gets a memory address, making it appear as if it takes space

*/

func main() {
	a := A{}

	fmt.Printf("struct 'A' size %+v bytes\n", int(unsafe.Sizeof(a)))
	fmt.Printf("- align attr1 (%T): %+v\n", a.attr1, int(unsafe.Alignof(a.attr1)))
	fmt.Printf("- align attr2 (%T): %+v\n", a.attr2, int(unsafe.Alignof(a.attr2)))
	fmt.Printf("- align attr3 (%T): %+v\n", a.attr3, int(unsafe.Alignof(a.attr3)))

	fmt.Println()

	b := B{}

	fmt.Printf("struct 'B' size %+v bytes\n", int(unsafe.Sizeof(b)))
	fmt.Printf("- align attr1 (%T): %+v\n", b.attr1, int(unsafe.Alignof(b.attr1)))
	fmt.Printf("- align attr2 (%T): %+v\n", b.attr2, int(unsafe.Alignof(b.attr2)))
	fmt.Printf("- align attr3 (%T): %+v\n", b.attr3, int(unsafe.Alignof(b.attr3)))

	fmt.Println()

	c := C{}

	fmt.Printf("struct 'C' size %+v bytes (addr: %p)\n", int(unsafe.Sizeof(c)), &c)
	fmt.Printf("- align attr1 (%T - %p): %+v\n", c.attr1, &c.attr1, int(unsafe.Alignof(c.attr1)))
	fmt.Printf("- align attr2 (%T - %p): %+v\n", c.attr2, &c.attr2, int(unsafe.Alignof(c.attr2)))

	fmt.Println()

	d := D{}

	fmt.Printf("struct 'D' size %+v bytes (addr: %p)\n", int(unsafe.Sizeof(d)), &d)
	fmt.Printf("- align attr1 (%T - %p): %+v\n", d.attr1, &d.attr1, int(unsafe.Alignof(d.attr1)))
	fmt.Printf("- align attr2 (%T - %p): %+v\n", d.attr2, &d.attr2, int(unsafe.Alignof(d.attr2)))
}

type A struct {
	attr1 int8   // 1 byte
	attr2 string // 16 byte
	attr3 int8   // 1 byte
}

type B struct {
	attr1 string
	attr2 int8
	attr3 int8
}

type C struct {
	attr1 struct{}
	attr2 string
}

type D struct {
	attr1 string
	attr2 struct{}
}
