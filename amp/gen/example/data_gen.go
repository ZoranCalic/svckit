// Code generated by go generate; DO NOT EDIT.
package example

import "time"

type PersonDiff struct {
	Name         *string            `json:"name,omitempty" bson:"name,omitempty" msg:"name"`
	Age          *int               `json:"a,omitempty" bson:"a,omitempty" msg:"a"`
	NotInAdapter *string            `json:"notInAdapter,omitempty" bson:"notInAdapter,omitempty" msg:"notInAdapter"`
	Children     map[int]*ChildDiff `json:"c,omitempty" bson:"c,omitempty" msg:"c"`
}

func (o PersonDiff) empty() bool {
	return o.Name == nil &&
		o.Age == nil &&
		o.NotInAdapter == nil &&
		(o.Children == nil || len(o.Children) == 0)
}

type ChildDiff struct {
	DateOfBirth  *time.Time `json:"dob,omitempty" bson:"dob,omitempty" msg:"dob"`
	NotInAdapter *string    `json:"notInAdapter,omitempty" bson:"notInAdapter,omitempty" msg:"notInAdapter"`
}

func (o ChildDiff) empty() bool {
	return o.DateOfBirth == nil &&
		o.NotInAdapter == nil
}

func (o *Person) MergeDiff(d *PersonDiff) bool {
	changed := false
	if d.Name != nil && *d.Name != o.Name {
		o.Name = *d.Name
		changed = true
	}
	if d.Age != nil && *d.Age != o.Age {
		o.Age = *d.Age
		changed = true
	}
	if d.NotInAdapter != nil && *d.NotInAdapter != o.NotInAdapter {
		o.NotInAdapter = *d.NotInAdapter
		changed = true
	}
	for k, dc := range d.Children {
		c, ok := o.Children[k]
		if dc == nil {
			if ok {
				delete(o.Children, k)
			}
			changed = true
			continue
		}
		if !ok {
			c = &Child{}
			if o.Children == nil {
				o.Children = make(map[int]*Child)
			}
			o.Children[k] = c
		}
		if c.MergeDiff(dc) {
			changed = true
		}
	}
	return changed
}

func (o *Child) MergeDiff(d *ChildDiff) bool {
	changed := false
	if d.DateOfBirth != nil && *d.DateOfBirth != o.DateOfBirth {
		o.DateOfBirth = *d.DateOfBirth
		changed = true
	}
	if d.NotInAdapter != nil && *d.NotInAdapter != o.NotInAdapter {
		o.NotInAdapter = *d.NotInAdapter
		changed = true
	}
	return changed
}

// Creates diff (i) between new (n) and old (o) Person.
// So that diff applyed to old will produce new.
func (o *Person) createDiff(n *Person) *PersonDiff {
	i := &PersonDiff{}
	if n.Name != o.Name {
		i.Name = &n.Name
	}
	if n.Age != o.Age {
		i.Age = &n.Age
	}
	if n.NotInAdapter != o.NotInAdapter {
		i.NotInAdapter = &n.NotInAdapter
	}
	i.Children = make(map[int]*ChildDiff)
	for k, nc := range n.Children {
		oc, ok := o.Children[k]
		if !ok {
			oc = &Child{}
		}
		ip := oc.createDiff(nc)
		if ip != nil {
			i.Children[k] = ip
		}
	}
	for k, _ := range o.Children {
		if _, ok := n.Children[k]; !ok {
			i.Children[k] = nil // signal delete
		}
	}
	if len(i.Children) == 0 {
		i.Children = nil
	}
	if i.empty() {
		return nil
	}
	return i
}

// Creates diff (i) between new (n) and old (o) Child.
// So that diff applyed to old will produce new.
func (o *Child) createDiff(n *Child) *ChildDiff {
	i := &ChildDiff{}
	if n.DateOfBirth != o.DateOfBirth {
		i.DateOfBirth = &n.DateOfBirth
	}
	if n.NotInAdapter != o.NotInAdapter {
		i.NotInAdapter = &n.NotInAdapter
	}
	if i.empty() {
		return nil
	}
	return i
}

type PersonAdapter interface {
	Name() string
	Age() int
	Children() map[int]ChildAdapter
}

// Creates diff (i) between new (n) and old (o) Person.
// So that diff applyed to old will produce new.
func (o *Person) AdapterDiff(n PersonAdapter) *PersonDiff {
	i := &PersonDiff{}
	if v := n.Name(); v != o.Name {
		i.Name = &v
	}
	if v := n.Age(); v != o.Age {
		i.Age = &v
	}
	i.Children = make(map[int]*ChildDiff)
	nChildren := n.Children()
	for k, nc := range nChildren {
		oc, ok := o.Children[k]
		if !ok {
			oc = &Child{}
		}
		ic := oc.AdapterDiff(nc)
		if ic != nil {
			i.Children[k] = ic
		}
	}
	for k, _ := range o.Children {
		if _, ok := nChildren[k]; !ok {
			i.Children[k] = nil // signal delete
		}
	}
	if len(i.Children) == 0 {
		i.Children = nil
	}
	if i.empty() {
		return nil
	}
	return i
}

type ChildAdapter interface {
	DateOfBirth() time.Time
}

// Creates diff (i) between new (n) and old (o) Child.
// So that diff applyed to old will produce new.
func (o *Child) AdapterDiff(n ChildAdapter) *ChildDiff {
	i := &ChildDiff{}
	if v := n.DateOfBirth(); v != o.DateOfBirth {
		i.DateOfBirth = &v
	}
	if i.empty() {
		return nil
	}
	return i
}

func (o *Person) Copy() *Person {
	o2 := &Person{}
	*o2 = *o
	o2.Children = make(map[int]*Child)
	for k, oc := range o.Children {
		o2.Children[k] = oc.Copy()
	}
	return o2
}

func (o *Child) Copy() *Child {
	o2 := &Child{}
	*o2 = *o
	return o2
}
