package main

import "fmt"

type Dummy struct {
	Name     string
	IsSet    bool
	IntValue int
}

func testOutput(d []Dummy, msg string, dd []Dummy) {
	fmt.Println(msg)
	for i, e := range d {
		response := "     "
		if dd[i] != d[i] {
			response = "  :-D"
		}
		fmt.Printf("%d:%s -> %v\n", i, response, e)
	}
}

func testMapOutput(d map[string]Dummy, msg string, dd map[string]Dummy) {
	fmt.Println(msg)
	for k, v := range d {
		response := "     "
		if dd[k] != d[k] {
			response = "  :-D"
		}
		fmt.Printf("%s:%s -> %v\n", k, response, v)
	}
}

func modifyArray(d []Dummy) {
	for i := 0; i < len(d); i++ {
		d[i].Name = d[i].Name + "_2"
		d[i].IsSet = true
		d[i].IntValue++
	}
}

func arrays() {
	d := make([]Dummy, 0)
	dd := make([]Dummy, 0)
	ddd := make([]Dummy, 0)

	d = append(d, Dummy{Name: "name1", IntValue: 1})
	d = append(d, Dummy{Name: "name2", IntValue: 2})
	d = append(d, Dummy{Name: "name3", IntValue: 3})
	d = append(d, Dummy{Name: "name4", IntValue: 4})

	dd = append(dd, Dummy{Name: "name1", IntValue: 1})
	dd = append(dd, Dummy{Name: "name2", IntValue: 2})
	dd = append(dd, Dummy{Name: "name3", IntValue: 3})
	dd = append(dd, Dummy{Name: "name4", IntValue: 4})

	for _, e := range d {
		e.Name = e.Name + "_1"
		e.IsSet = true
		e.IntValue++
	}

	testOutput(d, "Array: Changes after range-loop", dd)

	for i := 0; i < len(d); i++ {
		e := d[i]
		e.Name = e.Name + "_1"
		e.IsSet = true
		e.IntValue++
	}

	testOutput(d, "Array: Changes index-loop", dd)

	for i := 0; i < len(d); i++ {
		e := &d[i]
		e.Name = e.Name + "_1"
		e.IsSet = true
		e.IntValue++
	}

	testOutput(d, "Array: Changes after index-pointer-loop", dd)

	for i := 0; i < len(d); i++ {
		d[i].Name = d[i].Name + "_1"
		d[i].IsSet = true
		d[i].IntValue++
	}

	testOutput(d, "Array: Changes after for-w-ref-loop", dd)

	ddd = append(d, Dummy{Name: "name1_1", IntValue: 2})
	ddd = append(d, Dummy{Name: "name2_1", IntValue: 3})
	ddd = append(d, Dummy{Name: "name3_1", IntValue: 4})
	ddd = append(d, Dummy{Name: "name4_1", IntValue: 5})

	modifyArray(d)

	testOutput(d, "Array: Changes in a function", ddd)
}

func modifyMap(d map[string]Dummy) {
	for k, v := range d {
		v.IntValue++
		d[k] = v
	}
}

func maps() {
	d := make(map[string]Dummy)
	d["eins"] = Dummy{Name: "name1", IntValue: 1}
	d["zwei"] = Dummy{Name: "name2", IntValue: 2}
	d["drei"] = Dummy{Name: "name3", IntValue: 3}
	d["vier"] = Dummy{Name: "name4", IntValue: 4}

	dd := make(map[string]Dummy)
	dd["eins"] = Dummy{Name: "name1", IntValue: 1}
	dd["zwei"] = Dummy{Name: "name2", IntValue: 2}
	dd["drei"] = Dummy{Name: "name3", IntValue: 3}
	dd["vier"] = Dummy{Name: "name4", IntValue: 4}

	for k, v := range d {
		v.Name = v.Name + "_" + k
		v.IntValue++
		d[k] = v
	}

	testMapOutput(d, "Map: Changes in a range loop", dd)

	ddd := make(map[string]Dummy)
	ddd["eins"] = Dummy{Name: "name1_eins", IntValue: 1}
	ddd["zwei"] = Dummy{Name: "name2_zwei", IntValue: 2}
	ddd["drei"] = Dummy{Name: "name3_drei", IntValue: 3}
	ddd["vier"] = Dummy{Name: "name4_vier", IntValue: 4}

	modifyMap(d)

	testMapOutput(d, "Map: Changes in a function", ddd)
}

func main() {
	arrays()

	fmt.Println("")

	maps()

}
