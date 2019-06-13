// Copyright 2019 The TCell Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
// You may obtain a copy of the license at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This command creates a go file (with an init function) that registers
// a terminal description with the tcell terminfo database.  If this is
// placed into a separate package, then that package can be imported at
// run time. If no term values are specified on the command line, then $TERM
// is used.  This command requires mkinfo to be on your path.
//
// Usage is like this:
//
// mkinfo [-go file.go] -P <package name> [-quiet] [<term>...]
//

package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"

	"github.com/gdamore/tcell/terminfo"
	"github.com/gdamore/tcell/terminfo/dynamic"
)

func dotGoAddInt(w io.Writer, n string, i int) {
	if i == 0 {
		// initialized to 0, ignore
		return
	}
	fmt.Fprintf(w, "\t\t%-13s %d,\n", n+":", i)
}
func dotGoAddStr(w io.Writer, n string, s string) {
	if s == "" {
		return
	}
	fmt.Fprintf(w, "\t\t%-13s %q,\n", n+":", s)
}

func dotGoAddArr(w io.Writer, n string, a []string) {
	if len(a) == 0 {
		return
	}
	fmt.Fprintf(w, "\t\t%-13s []string{", n+":")
	did := false
	for _, b := range a {
		if did {
			fmt.Fprint(w, ", ")
		}
		did = true
		fmt.Fprintf(w, "%q", b)
	}
	fmt.Fprintln(w, "},")
}

func dotGoHeader(w io.Writer, packname string) {
	fmt.Fprintln(w, "// Generated automatically.  DO NOT HAND-EDIT.")
	fmt.Fprintln(w, "")
	if packname != "" {
		fmt.Fprintf(w, "package %s\n\n", packname)
		fmt.Fprintln(w, "import \"github.com/gdamore/tcell/terminfo\"")
	} else {
		fmt.Fprintln(w, "package terminfo")
	}
	fmt.Fprintln(w, "")
}

func dotGoTrailer(w io.Writer) {
}

func dotGoInfo(w io.Writer, t *terminfo.Terminfo, desc string) {

	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "func init() {")
	fmt.Fprintf(w, "\t// %s\n", desc)
	fmt.Fprintln(w, "\tAddTerminfo(&Terminfo{")
	dotGoAddStr(w, "Name", t.Name)
	dotGoAddArr(w, "Aliases", t.Aliases)
	dotGoAddInt(w, "Columns", t.Columns)
	dotGoAddInt(w, "Lines", t.Lines)
	dotGoAddInt(w, "Colors", t.Colors)
	dotGoAddStr(w, "Bell", t.Bell)
	dotGoAddStr(w, "Clear", t.Clear)
	dotGoAddStr(w, "EnterCA", t.EnterCA)
	dotGoAddStr(w, "ExitCA", t.ExitCA)
	dotGoAddStr(w, "ShowCursor", t.ShowCursor)
	dotGoAddStr(w, "HideCursor", t.HideCursor)
	dotGoAddStr(w, "AttrOff", t.AttrOff)
	dotGoAddStr(w, "Underline", t.Underline)
	dotGoAddStr(w, "Bold", t.Bold)
	dotGoAddStr(w, "Dim", t.Dim)
	dotGoAddStr(w, "Blink", t.Blink)
	dotGoAddStr(w, "Reverse", t.Reverse)
	dotGoAddStr(w, "EnterKeypad", t.EnterKeypad)
	dotGoAddStr(w, "ExitKeypad", t.ExitKeypad)
	dotGoAddStr(w, "SetFg", t.SetFg)
	dotGoAddStr(w, "SetBg", t.SetBg)
	dotGoAddStr(w, "SetFgBg", t.SetFgBg)
	dotGoAddStr(w, "PadChar", t.PadChar)
	dotGoAddStr(w, "AltChars", t.AltChars)
	dotGoAddStr(w, "EnterAcs", t.EnterAcs)
	dotGoAddStr(w, "ExitAcs", t.ExitAcs)
	dotGoAddStr(w, "EnableAcs", t.EnableAcs)
	dotGoAddStr(w, "SetFgRGB", t.SetFgRGB)
	dotGoAddStr(w, "SetBgRGB", t.SetBgRGB)
	dotGoAddStr(w, "SetFgBgRGB", t.SetFgBgRGB)
	dotGoAddStr(w, "Mouse", t.Mouse)
	dotGoAddStr(w, "MouseMode", t.MouseMode)
	dotGoAddStr(w, "SetCursor", t.SetCursor)
	dotGoAddStr(w, "CursorBack1", t.CursorBack1)
	dotGoAddStr(w, "CursorUp1", t.CursorUp1)
	dotGoAddStr(w, "KeyUp", t.KeyUp)
	dotGoAddStr(w, "KeyDown", t.KeyDown)
	dotGoAddStr(w, "KeyRight", t.KeyRight)
	dotGoAddStr(w, "KeyLeft", t.KeyLeft)
	dotGoAddStr(w, "KeyInsert", t.KeyInsert)
	dotGoAddStr(w, "KeyDelete", t.KeyDelete)
	dotGoAddStr(w, "KeyBackspace", t.KeyBackspace)
	dotGoAddStr(w, "KeyHome", t.KeyHome)
	dotGoAddStr(w, "KeyEnd", t.KeyEnd)
	dotGoAddStr(w, "KeyPgUp", t.KeyPgUp)
	dotGoAddStr(w, "KeyPgDn", t.KeyPgDn)
	dotGoAddStr(w, "KeyF1", t.KeyF1)
	dotGoAddStr(w, "KeyF2", t.KeyF2)
	dotGoAddStr(w, "KeyF3", t.KeyF3)
	dotGoAddStr(w, "KeyF4", t.KeyF4)
	dotGoAddStr(w, "KeyF5", t.KeyF5)
	dotGoAddStr(w, "KeyF6", t.KeyF6)
	dotGoAddStr(w, "KeyF7", t.KeyF7)
	dotGoAddStr(w, "KeyF8", t.KeyF8)
	dotGoAddStr(w, "KeyF9", t.KeyF9)
	dotGoAddStr(w, "KeyF10", t.KeyF10)
	dotGoAddStr(w, "KeyF11", t.KeyF11)
	dotGoAddStr(w, "KeyF12", t.KeyF12)
	dotGoAddStr(w, "KeyF13", t.KeyF13)
	dotGoAddStr(w, "KeyF14", t.KeyF14)
	dotGoAddStr(w, "KeyF15", t.KeyF15)
	dotGoAddStr(w, "KeyF16", t.KeyF16)
	dotGoAddStr(w, "KeyF17", t.KeyF17)
	dotGoAddStr(w, "KeyF18", t.KeyF18)
	dotGoAddStr(w, "KeyF19", t.KeyF19)
	dotGoAddStr(w, "KeyF20", t.KeyF20)
	dotGoAddStr(w, "KeyF21", t.KeyF21)
	dotGoAddStr(w, "KeyF22", t.KeyF22)
	dotGoAddStr(w, "KeyF23", t.KeyF23)
	dotGoAddStr(w, "KeyF24", t.KeyF24)
	dotGoAddStr(w, "KeyF25", t.KeyF25)
	dotGoAddStr(w, "KeyF26", t.KeyF26)
	dotGoAddStr(w, "KeyF27", t.KeyF27)
	dotGoAddStr(w, "KeyF28", t.KeyF28)
	dotGoAddStr(w, "KeyF29", t.KeyF29)
	dotGoAddStr(w, "KeyF30", t.KeyF30)
	dotGoAddStr(w, "KeyF31", t.KeyF31)
	dotGoAddStr(w, "KeyF32", t.KeyF32)
	dotGoAddStr(w, "KeyF33", t.KeyF33)
	dotGoAddStr(w, "KeyF34", t.KeyF34)
	dotGoAddStr(w, "KeyF35", t.KeyF35)
	dotGoAddStr(w, "KeyF36", t.KeyF36)
	dotGoAddStr(w, "KeyF37", t.KeyF37)
	dotGoAddStr(w, "KeyF38", t.KeyF38)
	dotGoAddStr(w, "KeyF39", t.KeyF39)
	dotGoAddStr(w, "KeyF40", t.KeyF40)
	dotGoAddStr(w, "KeyF41", t.KeyF41)
	dotGoAddStr(w, "KeyF42", t.KeyF42)
	dotGoAddStr(w, "KeyF43", t.KeyF43)
	dotGoAddStr(w, "KeyF44", t.KeyF44)
	dotGoAddStr(w, "KeyF45", t.KeyF45)
	dotGoAddStr(w, "KeyF46", t.KeyF46)
	dotGoAddStr(w, "KeyF47", t.KeyF47)
	dotGoAddStr(w, "KeyF48", t.KeyF48)
	dotGoAddStr(w, "KeyF49", t.KeyF49)
	dotGoAddStr(w, "KeyF50", t.KeyF50)
	dotGoAddStr(w, "KeyF51", t.KeyF51)
	dotGoAddStr(w, "KeyF52", t.KeyF52)
	dotGoAddStr(w, "KeyF53", t.KeyF53)
	dotGoAddStr(w, "KeyF54", t.KeyF54)
	dotGoAddStr(w, "KeyF55", t.KeyF55)
	dotGoAddStr(w, "KeyF56", t.KeyF56)
	dotGoAddStr(w, "KeyF57", t.KeyF57)
	dotGoAddStr(w, "KeyF58", t.KeyF58)
	dotGoAddStr(w, "KeyF59", t.KeyF59)
	dotGoAddStr(w, "KeyF60", t.KeyF60)
	dotGoAddStr(w, "KeyF61", t.KeyF61)
	dotGoAddStr(w, "KeyF62", t.KeyF62)
	dotGoAddStr(w, "KeyF63", t.KeyF63)
	dotGoAddStr(w, "KeyF64", t.KeyF64)
	dotGoAddStr(w, "KeyCancel", t.KeyCancel)
	dotGoAddStr(w, "KeyPrint", t.KeyPrint)
	dotGoAddStr(w, "KeyExit", t.KeyExit)
	dotGoAddStr(w, "KeyHelp", t.KeyHelp)
	dotGoAddStr(w, "KeyClear", t.KeyClear)
	dotGoAddStr(w, "KeyBacktab", t.KeyBacktab)
	dotGoAddStr(w, "KeyShfLeft", t.KeyShfLeft)
	dotGoAddStr(w, "KeyShfRight", t.KeyShfRight)
	dotGoAddStr(w, "KeyShfUp", t.KeyShfUp)
	dotGoAddStr(w, "KeyShfDown", t.KeyShfDown)
	dotGoAddStr(w, "KeyCtrlLeft", t.KeyCtrlLeft)
	dotGoAddStr(w, "KeyCtrlRight", t.KeyCtrlRight)
	dotGoAddStr(w, "KeyCtrlUp", t.KeyCtrlUp)
	dotGoAddStr(w, "KeyCtrlDown", t.KeyCtrlDown)
	dotGoAddStr(w, "KeyMetaLeft", t.KeyMetaLeft)
	dotGoAddStr(w, "KeyMetaRight", t.KeyMetaRight)
	dotGoAddStr(w, "KeyMetaUp", t.KeyMetaUp)
	dotGoAddStr(w, "KeyMetaDown", t.KeyMetaDown)
	dotGoAddStr(w, "KeyAltLeft", t.KeyAltLeft)
	dotGoAddStr(w, "KeyAltRight", t.KeyAltRight)
	dotGoAddStr(w, "KeyAltUp", t.KeyAltUp)
	dotGoAddStr(w, "KeyAltDown", t.KeyAltDown)
	dotGoAddStr(w, "KeyAltShfLeft", t.KeyAltShfLeft)
	dotGoAddStr(w, "KeyAltShfRight", t.KeyAltShfRight)
	dotGoAddStr(w, "KeyAltShfUp", t.KeyAltShfUp)
	dotGoAddStr(w, "KeyAltShfDown", t.KeyAltShfDown)
	dotGoAddStr(w, "KeyMetaShfLeft", t.KeyMetaShfLeft)
	dotGoAddStr(w, "KeyMetaShfRight", t.KeyMetaShfRight)
	dotGoAddStr(w, "KeyMetaShfUp", t.KeyMetaShfUp)
	dotGoAddStr(w, "KeyMetaShfDown", t.KeyMetaShfDown)
	dotGoAddStr(w, "KeyCtrlShfLeft", t.KeyCtrlShfLeft)
	dotGoAddStr(w, "KeyCtrlShfRight", t.KeyCtrlShfRight)
	dotGoAddStr(w, "KeyCtrlShfUp", t.KeyCtrlShfUp)
	dotGoAddStr(w, "KeyCtrlShfDown", t.KeyCtrlShfDown)
	dotGoAddStr(w, "KeyShfHome", t.KeyShfHome)
	dotGoAddStr(w, "KeyShfEnd", t.KeyShfEnd)
	dotGoAddStr(w, "KeyCtrlHome", t.KeyCtrlHome)
	dotGoAddStr(w, "KeyCtrlEnd", t.KeyCtrlEnd)
	dotGoAddStr(w, "KeyMetaHome", t.KeyMetaHome)
	dotGoAddStr(w, "KeyMetaEnd", t.KeyMetaEnd)
	dotGoAddStr(w, "KeyAltHome", t.KeyAltHome)
	dotGoAddStr(w, "KeyAltEnd", t.KeyAltEnd)
	dotGoAddStr(w, "KeyCtrlShfHome", t.KeyCtrlShfHome)
	dotGoAddStr(w, "KeyCtrlShfEnd", t.KeyCtrlShfEnd)
	dotGoAddStr(w, "KeyMetaShfHome", t.KeyMetaShfHome)
	dotGoAddStr(w, "KeyMetaShfEnd", t.KeyMetaShfEnd)
	dotGoAddStr(w, "KeyAltShfHome", t.KeyAltShfHome)
	dotGoAddStr(w, "KeyAltShfEnd", t.KeyAltShfEnd)
	fmt.Fprintln(w, "\t})")
	fmt.Fprintln(w, "}")
}

var packname = ""

func dotGoFile(fname string, term *terminfo.Terminfo, desc string, makeDir bool) error {
	w := os.Stdout
	var e error
	if fname != "-" && fname != "" {
		if makeDir {
			dname := path.Dir(fname)
			_ = os.Mkdir(dname, 0777)
		}
		if w, e = os.Create(fname); e != nil {
			return e
		}
	}
	dotGoHeader(w, packname)
	dotGoInfo(w, term, desc)
	dotGoTrailer(w)
	if w != os.Stdout {
		w.Close()
	}
	cmd := exec.Command("go", "fmt", fname)
	cmd.Run()
	return nil
}

func main() {
	gofile := ""
	quiet := false

	flag.StringVar(&gofile, "go", "", "generate go source in named file")
	flag.StringVar(&packname, "P", packname, "package name (go source)")
	flag.BoolVar(&quiet, "quiet", false, "suppress error messages")
	flag.Parse()
	var e error

	args := flag.Args()
	if len(args) == 0 {
		args = []string{os.Getenv("TERM")}
	}

	tdata := make(map[string]*terminfo.Terminfo)
	descs := make(map[string]string)

	for _, term := range args {
		if t, desc, e := dynamic.LoadTerminfo(term); e != nil {
			if !quiet {
				fmt.Fprintf(os.Stderr,
					"Failed loading %s: %v\n", term, e)
			}
		} else {
			tdata[term] = t
			descs[term] = desc
		}
	}

	if len(tdata) == 0 {
		// No data.
		os.Exit(0)
	}

	for term, t := range tdata {
		if t.Name == term {
			e = dotGoFile(gofile, t, descs[term], false)
			if e != nil {
				fmt.Fprintf(os.Stderr, "Failed %s: %v", gofile, e)
				os.Exit(1)
			}
		}
	}

}
