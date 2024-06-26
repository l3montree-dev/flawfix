// Copyright (C) 2024 Tim Bastin, l3montree UG (haftungsbeschränkt)
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package utils

func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func Map[T, U any](s []T, f func(T) U) []U {
	r := make([]U, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func Reduce[T, U any](s []T, f func(U, T) U, init U) U {
	r := init
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

func Flat[T any](s [][]T) []T {
	res := make([]T, 0)
	for _, subslice := range s {
		res = append(res, subslice...)
	}
	return res
}

type CompareResult[T any] struct {
	OnlyInA []T
	OnlyInB []T
	InBoth  []T
}

func CompareSlices[T any, K comparable](a, b []T, serializer func(T) K) CompareResult[T] {
	res := CompareResult[T]{}
	inA := make(map[K]bool)
	inB := make(map[K]bool)

	for _, v := range a {
		inA[serializer(v)] = true
	}

	for _, v := range b {
		inB[serializer(v)] = true

		if _, ok := inA[serializer(v)]; ok {
			res.InBoth = append(res.InBoth, v)
		} else {
			res.OnlyInB = append(res.OnlyInB, v)
		}
	}

	for _, v := range a {
		if _, ok := inB[serializer(v)]; !ok {
			res.OnlyInA = append(res.OnlyInA, v)
		}
	}

	return res
}
