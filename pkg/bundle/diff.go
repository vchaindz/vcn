/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package bundle

import (
	"fmt"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/google/go-cmp/cmp"
)

func (m Manifest) Diff(x Manifest, raw bool) (bool, string, error) {
	if raw {
		return m.diffCmp(x)
	}
	return m.diffCustom(x)
}

func (m Manifest) diffCmp(x Manifest) (equal bool, out string, err error) {
	var r diffReporter
	equal = cmp.Equal(m, x, cmp.Reporter(&r))
	out = r.String()
	return
}

type diffReporter struct {
	path  cmp.Path
	diffs []string
}

func (r *diffReporter) PushStep(ps cmp.PathStep) {
	r.path = append(r.path, ps)
}

func (r *diffReporter) Report(rs cmp.Result) {
	if !rs.Equal() {
		vx, vy := r.path.Last().Values()
		var line string
		switch true {
		case !vx.IsValid():
			line = fmt.Sprintf("%#v:\n\t+: %+v\n", r.path, vy)
		case !vy.IsValid():
			line = fmt.Sprintf("%#v:\n\t-: %+v\n", r.path, vx)
		default:
			line = fmt.Sprintf("%#v:\n\t-: %+v\n\t+: %+v\n", r.path, vx, vy)
		}
		r.diffs = append(r.diffs, line)
	}
}

func (r *diffReporter) PopStep() {
	r.path = r.path[:len(r.path)-1]
}

func (r *diffReporter) String() string {
	return strings.Join(r.diffs, "\n")
}

func (m Manifest) diffCustom(x Manifest) (equal bool, out string, err error) {

	type diff struct {
		from Descriptor
		to   Descriptor
	}

	parents := make(map[Descriptor]bool)
	pByPath := make(map[string]Descriptor)
	pByDig := make(map[string][]Descriptor)
	for _, d := range x.Items {
		parents[d] = true
		pByPath[d.Path] = d
		pByDig[d.Digest.String()] = append(pByDig[d.Digest.String()], d)
	}
	children := make(map[Descriptor]bool)
	for _, d := range m.Items {
		children[d] = true
	}

	adds := make([]Descriptor, 0)
	mods := make([]diff, 0)
	rens := make([]diff, 0)
	dels := make([]Descriptor, 0)

mainLoop:
	for _, d := range m.Items {

		// Unchanged
		if parents[d] {
			delete(parents, d)
			continue
		}

		// Content modifications
		if orig, ok := pByPath[d.Path]; ok {
			mods = append(mods, diff{
				from: orig,
				to:   d,
			})
			delete(parents, orig)
			continue
		}

		// File renamings
		if list, ok := pByDig[d.Digest.String()]; ok {
			for i, dd := range list {
				if !children[dd] {
					rens = append(rens, diff{
						from: dd,
						to:   d,
					})
					delete(parents, dd)
					pByDig[d.Digest.String()] = append(list[:i], list[i+1:]...)
					continue mainLoop
				}
			}
		}

		adds = append(adds, d)
	}

	for d := range parents {
		dels = append(dels, d)
	}

	equal = len(adds) == 0 && len(mods) == 0 && len(rens) == 0 && len(dels) == 0

	lines := []string{}

	sprintf := func(format string, a ...interface{}) {
		lines = append(lines, fmt.Sprintf(format, a...))
	}

	for _, d := range adds {
		sprintf(
			"\tnew item:   %s (%s)\n\t            + %s",
			d.Path,
			humanize.Bytes(d.Size),
			d.Digest.String(),
		)
	}
	for _, dd := range mods {
		sprintf(
			"\tmodified:   %s (%s -> %s)\n\t            - %s\n\t            + %s",
			dd.to.Path,
			humanize.Bytes(dd.from.Size),
			humanize.Bytes(dd.to.Size),
			dd.from.Digest.String(),
			dd.to.Digest.String(),
		)
	}
	for _, dd := range rens {
		sprintf(
			"\trenamed:    %s -> %s (%s)\n\t            = %s",
			dd.to.Path,
			dd.to.Path,
			humanize.Bytes(dd.to.Size),
			dd.to.Digest.String(),
		)
	}
	for _, d := range dels {
		sprintf(
			"\tdeleted:    %s (%s)\n\t            - %s",
			d.Path,
			humanize.Bytes(d.Size),
			d.Digest.String(),
		)
	}

	out = strings.Join(lines, "\n\n") + "\n"

	return
}
