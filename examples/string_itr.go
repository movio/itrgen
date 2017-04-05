/*
 * CODE GENERATED AUTOMATICALLY WITH github.com/movio/itrgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */
package main

// string Iterator
type StringItr []string

// MapMyStruct Utility Map function
func (itr StringItr) MapMyStruct(fn func(string) (MyStruct, error)) ([]MyStruct, error) {
	results := []MyStruct{}
	for _, i := range itr {
		result, err := fn(i)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

// MapString Utility Map function
func (itr StringItr) MapString(fn func(string) (string, error)) ([]string, error) {
	results := []string{}
	for _, i := range itr {
		result, err := fn(i)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

// MapInt64 Utility Map function
func (itr StringItr) MapInt64(fn func(string) (int64, error)) ([]int64, error) {
	results := []int64{}
	for _, i := range itr {
		result, err := fn(i)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}
