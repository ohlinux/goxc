package core

/*
   Copyright 2013 Am Laher

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

func StringSliceDelIndex(a []string, i int) []string {
	a = append(a[:i], a[i+1:]...)
	// OR  a = a[:i+copy(a[i:], a[i+1:])]
	return a
}
func StringSliceDel(slice []string, value string) []string {
	p := StringSlicePos(slice, value)
	if p > -1 {
		return StringSliceDelIndex(slice, p)
	}
	return slice
}

func StringSlicePos(slice []string, value string) int {
	for p, v := range slice {
		if v == value {
			return p
		}
	}
	return -1
}