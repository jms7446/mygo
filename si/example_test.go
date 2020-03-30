package si

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func ExampleDeadline_OverDue() {
	d1 := NewDeadline(time.Now().Add(-4 * time.Hour))
	d2 := NewDeadline(time.Now().Add(4 * time.Hour))
	fmt.Println(d1.OverDue())
	fmt.Println(d2.OverDue())
	// Output:
	// true
	// false
}

func Example_marshalJSON() {
	t := Task{
		"Laundry",
		DONE,
		NewDeadline(time.Date(2015, time.August, 16, 15, 43, 0, 0, time.UTC)),
	}
	b, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
	// Output:
	// {"Title":"Laundry","Status":2,"Deadline":"2015-08-16T15:43:00Z"}
}
