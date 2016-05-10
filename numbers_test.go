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

package slice

import "testing"

const (
	SampleIntMissing = 45
	SampleSum        = 7373
	SampleAvg        = 368.65
)

var (
	SampleIntArray = []int{
		296, 112, 380, 243, 336,
		376, 664, 556, 162, 173,
		684, 503, 542, 215, 607,
		2, 132, 539, 646, 205,
	}
	SampleIntMissingArray = []int{
		762, 667, 711, 254, 549,
		379, 385, 944, 481, 402,
	}
	SampleIntGT500 = []int{
		664, 556, 684, 503, 542,
		607, 539, 646,
	}
)

func TestIntAverage(t *testing.T) {
	sample := NInt(SampleIntArray)

	if sample.Average() != SampleAvg {
		t.Error("Unexpected value average from slice elements")
	}
}

func TestIntEqual(t *testing.T) {
	sample := NInt(SampleIntArray)
	sample2 := make([]int, len(sample))
	copy(sample2, sample)

	if !sample.Equal(sample2) {
		t.Error("Slices should be equal to each other")
	}
	if !NInt(sample2).Equal(sample) {
		t.Error("Slices should be equal to each other")
	}

	sample2 = sample2[:len(sample2)-1]
	if sample.Equal(sample2) {
		t.Error("Slices should be different")
	}
}

func TestIntSum(t *testing.T) {
	sample := NInt(SampleIntArray)

	if sample.Sum() != SampleSum {
		t.Error("Unexpected value summing slice elements")
	}
}

func TestIntIndexOf(t *testing.T) {
	sample := NInt(SampleIntArray)

	for i, item := range SampleIntArray {
		if retIdx := sample.IndexOf(item); retIdx != i {
			t.Errorf("Expected index '%d' but got '%d'", i, retIdx)
		}
	}

	if sample.IndexOf(SampleIntMissing) != -1 {
		t.Errorf("The index of '%d' should be -1", SampleIntMissing)
	}
}

func TestIntExists(t *testing.T) {
	sample := NInt(SampleIntArray)

	for _, item := range SampleIntArray {
		if !sample.Exists(item) {
			t.Errorf("The integer '%d' should be found", item)
		}
	}

	if sample.Exists(SampleIntMissing) {
		t.Errorf("The integer '%d' should not exists", SampleIntMissing)
	}
}

func TestIntExistsAll(t *testing.T) {
	sample := NInt(SampleIntArray)
	testSample := make([]int, 6)
	copy(testSample, sample[2:8])

	if !sample.ExistsAll(testSample...) {
		t.Error("All elements of specified sample should exists")
	}

	testSample = append(testSample, SampleIntMissing)
	if sample.ExistsAll(testSample...) {
		t.Errorf("The element '%d' should not exists", SampleIntMissing)
	}
}

func TestIntExistsAny(t *testing.T) {
	sample := NInt(SampleIntArray)
	testSample := make([]int, len(SampleIntMissingArray))
	copy(testSample, SampleIntMissingArray)

	if sample.ExistsAny(testSample...) {
		t.Error("None of elements of specified sample should exists")
	}

	testSample = append(testSample, SampleIntArray[5])
	if !sample.ExistsAny(testSample...) {
		t.Errorf("The element '%d' should exists", SampleIntArray[5])
	}
}

func TestIntTrueForAll(t *testing.T) {
	sample := NInt(SampleIntArray)

	isPositive := func(a int) bool {
		return a > 0
	}

	if !sample.TrueForAll(isPositive) {
		t.Error(
			"Every element of specified sample should be greater than zero")
	}

	isBig := func(a int) bool {
		return a > 1000
	}
	if sample.TrueForAll(isBig) {
		t.Error("None of elements of specified sample should be greater than 1000")
	}
}

func TestIntTrueForAny(t *testing.T) {
	sample := NInt(SampleIntArray)

	hasEvenNumbers := func(a int) bool {
		return a%2 == 0
	}

	if !sample.TrueForAny(hasEvenNumbers) {
		t.Error("Should be found even numbers on specified sample")
	}
}

func TestIntWhere(t *testing.T) {
	sample := NInt(SampleIntArray)

	fSample := sample.Where(func(a int) bool {
		return a > 500
	})
	if len(fSample) != len(SampleIntGT500) {
		t.Errorf("Unexpected filtered sample length: %d", len(fSample))
	}
	if !fSample.Equal(SampleIntGT500) {
		t.Error("Unexpected elements from filtered sample")
	}
}
