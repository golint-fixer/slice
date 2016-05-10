/*
 * Copyright 2015 Fabrício Godoy
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

import (
	"strings"
	"testing"
)

const (
	SampleTextMissing = "Maecenas"
)

var (
	SampleTextArray = []string{
		"Lorem", "ipsum", "dolor", "sit", "amet,", "consectetur", "adipiscing",
		"elit", "Sed", "tortor", "justo", "dui", "iaculis", "molestie",
		"Integer",
	}
	SampleTextMissingArray = []string{
		"Maecenas", "cursus", "enim", "venenatis", "venenatis", "nisi", "vitae",
		"fermentum", "velit",
	}
	StopWords = []string{
		"ab", "ac", "ad", "adhic", "aliqui", "aliquis", "an", "ante", "apud",
		"at", "atque", "aut", "autem", "cum", "cur", "de", "deinde", "dum",
		"ego", "enim", "ergo", "es", "est", "et", "etiam", "etsi", "ex", "fio",
		"haud", "hic", "iam", "idem", "igitur", "ille", "in", "infra", "inter",
		"interim", "ipse", "is", "ita", "magis", "modo", "mox", "nam", "ne",
		"nec", "necque", "neque", "nisi", "non", "nos", "o", "ob", "per",
		"possum", "post", "pro", "quae", "quam", "quare", "qui", "quia",
		"quicumque", "quidem", "quilibet", "quis", "quisnam", "quisquam",
		"quisque", "quisquis", "quo", "quoniam", "sed", "si", "sic", "sive",
		"sub", "sui", "sum", "super", "suus", "tam", "tamen", "trans", "tu",
		"tum", "ubi", "uel", "uero",
	}
	SampleTextNoStopWords = []string{
		"Lorem", "ipsum", "dolor", "sit", "amet,", "consectetur", "adipiscing",
		"elit", "tortor", "justo", "dui", "iaculis", "molestie", "Integer",
	}
	SampleTextExcept = []string{"Sed"}
)

func TestStringEqual(t *testing.T) {
	sample := String(SampleTextArray)
	sample2 := make([]string, len(sample))
	copy(sample2, sample)

	if !sample.Equal(sample2) {
		t.Error("Slices should be equal to each other")
	}
	if !String(sample2).Equal(sample) {
		t.Error("Slices should be equal to each other")
	}

	sample2 = sample2[:len(sample2)-1]
	if sample.Equal(sample2) {
		t.Error("Slices should be different")
	}
}

func TestStringExcept(t *testing.T) {
	sample := String(SampleTextArray)

	out := sample.Except(SampleTextNoStopWords, true)
	if !out.Equal(SampleTextExcept) {
		t.Error("Unexpected result from except operation")
	}

	out = sample.Except(SampleTextArray, false)
	if len(out) != 0 {
		t.Error("Except from all elements should result in a zero-length slice")
	}
}

func TestStringIndexOf(t *testing.T) {
	sample := String(SampleTextArray)

	for i, item := range SampleTextArray {
		if retIdx := sample.IndexOf(item, false); retIdx != i {
			t.Errorf("Expected index '%d' but got '%d'", i, retIdx)
		}
		if retIdx := sample.IndexOf(strings.ToUpper(item), true); retIdx != i {
			t.Errorf("Expected index '%d' but got '%d'", i, retIdx)
		}
	}

	if sample.IndexOf(SampleTextMissing, false) != -1 {
		t.Errorf("The index of '%s' should be -1", SampleTextMissing)
	}

	missingUpper := strings.ToUpper(SampleTextMissing)
	if sample.IndexOf(missingUpper, true) != -1 {
		t.Errorf("The index of '%s' should be -1", missingUpper)
	}
}

func TestStringExists(t *testing.T) {
	sample := String(SampleTextArray)

	for _, item := range SampleTextArray {
		if !sample.Exists(item, false) {
			t.Errorf("The text '%s' should be found", item)
		}
		itemUpper := strings.ToUpper(item)
		if !sample.Exists(itemUpper, true) {
			t.Errorf("The text '%s' should be found", itemUpper)
		}
	}

	if sample.Exists(SampleTextMissing, false) {
		t.Errorf("The text '%s' should not exists", SampleTextMissing)
	}

	missingUpper := strings.ToUpper(SampleTextMissing)
	if sample.Exists(missingUpper, true) {
		t.Errorf("The text '%s' should not exists", missingUpper)
	}
}

func TestStringExistsAll(t *testing.T) {
	sample := String(SampleTextArray)
	testSample := make([]string, 6)
	copy(testSample, sample[2:8])

	if !sample.ExistsAll(testSample, false) {
		t.Error("All elements of specified sample should exists")
	}

	testSample = append(testSample, SampleTextMissing)
	if sample.ExistsAll(testSample, false) {
		t.Errorf("The element '%s' should not exists", SampleTextMissing)
	}
}

func TestStringExistsAny(t *testing.T) {
	sample := String(SampleTextArray)
	testSample := make([]string, len(SampleTextMissingArray))
	copy(testSample, SampleTextMissingArray)

	if sample.ExistsAny(testSample, false) {
		t.Error("None of elements of specified sample should exists")
	}

	testSample = append(testSample, SampleTextArray[5])
	if !sample.ExistsAny(testSample, false) {
		t.Errorf("The element '%s' should exists", SampleTextArray[5])
	}
}

func TestStringTrueForAll(t *testing.T) {
	sample := String(SampleTextArray)

	hasVowel := func(s string) bool {
		return strings.IndexAny(s, "aeiou") >= 0
	}

	if !sample.TrueForAll(hasVowel) {
		t.Error(
			"Every element of specified sample should have at least one vowel")
	}

	isBig := func(s string) bool {
		return len(s) > 50
	}
	if sample.TrueForAll(isBig) {
		t.Error(
			"None of elements of specified sample should have more than " +
				"50 characters")
	}
}

func TestStringTrueForAny(t *testing.T) {
	sample := String(SampleTextArray)
	stopWords := String(StopWords)

	hasStopWords := func(s string) bool {
		return stopWords.Exists(s, true)
	}

	if !sample.TrueForAny(hasStopWords) {
		t.Error("Should be found stop words on specified sample")
	}
}

func TestStringWhere(t *testing.T) {
	sample := String(SampleTextArray)

	fSample := sample.Where(func(s string) bool {
		return !String(StopWords).Exists(s, true)
	})
	if len(fSample) != len(SampleTextNoStopWords) {
		t.Errorf("Unexpected filtered sample length: %d", len(fSample))
	}
	if !fSample.Equal(SampleTextNoStopWords) {
		t.Error("Unexpected elements from filtered sample")
	}

	fSample = sample.Where(func(string) bool {
		return false
	})
	if len(fSample) != 0 {
		t.Error("Rejecting all elements on where operation should result in a " +
			"zero-length slice")
	}
}
