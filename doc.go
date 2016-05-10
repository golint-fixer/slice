/*
 * Copyright 2016 Fabrício Godoy
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
Package slice provides operations to work with slices.

String

A String is a named type of slice of strings and provides some useful methods as:

	- Equal to search two slices for equality;
	- Except get the difference between two slices;
	- Exists to determine whether specified string exists on slice;
	- ExistsAll to determine whether all specified strings exists on slice;
	- ExistsAny to determine whether any of specified strings exists on slice;
	- IndexOf to determine index of specified string;
	- TrueForAll to test whether every element of slice matches the conditions
	specified by the predicate function;
	- Where to filter elements from a string slice.

NTypes

NTypes provides named type of slice every number type. Supports the following types:
float32, float64, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64.
Provides the following methods:

	- Average to calculate the average value of elements from slice;
	- Equal to search two slices for equality;
	- Except get the difference between two slices;
	- Exists to determine whether specified string exists on slice;
	- ExistsAll to determine whether all specified numbers exists on slice;
	- ExistsAny to determine whether any of specified numbers exists on slice;
	- IndexOf to determine index of specified number;
	- Sum to calculate the sum of all elements from slice;
	- TrueForAll to test whether every element of slice matches the conditions
	specified by the predicate function;
	- Where to filter elements from a number slice.
*/
package slice
