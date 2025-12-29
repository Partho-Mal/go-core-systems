// item data model
// Define Item
//
//	Store actual value
//	Store expiration time
package cache

import "time"

// Item stores a value and its expiration time
type Item struct {
	// actual value stored in cache
	value interface{}

	// time when this item expires
	expiration time.Time
}
