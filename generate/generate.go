package generate

import (
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
)

type Generator struct {
	PkgPath, PkgName, OutName string
	Types                     []string
}

func (g *Generator) Run() error {
	if len(g.Types) == 0 {
		// fmt.Println("skipping file ", g)
		return nil
	}
	f, err := os.Create(g.OutName)
	if err != nil {
		return err
	}

	fmt.Fprintln(f, "package ", g.PkgName)
	fmt.Fprintln(f)
	fmt.Fprintln(f, "import (")
	fmt.Fprintln(f, `  "reflect"`)
	fmt.Fprintln(f)
	fmt.Fprintln(f, `  "github.com/bytedance/sonic"`)
	fmt.Fprintln(f, `)`)
	fmt.Fprintln(f)
	fmt.Fprintln(f, `func init() {`)
	for _, s := range g.Types {

		fmt.Fprintf(f, "  var var%s %s", s, s)
		fmt.Fprintln(f)
		fmt.Fprintf(f, "  sonic.Pretouch(reflect.TypeOf(var%s))", s)
		fmt.Fprintln(f)
	}
	fmt.Fprintln(f, `}`)
	f.Close()

	in, err := ioutil.ReadFile(f.Name())
	if err != nil {
		fmt.Println("error in reading prints")
		return err
	}
	out, err := format.Source(in)
	if err != nil {
		fmt.Println("error in formatting")
		return err
	}
	return ioutil.WriteFile(g.OutName, out, 0644)
}
