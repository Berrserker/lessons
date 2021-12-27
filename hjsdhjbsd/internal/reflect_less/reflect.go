package reflectless
//go:generate
import (
	"fmt"
	`html/template`
	"log"
	`os`
	"reflect"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type MyType int

type MyEmpty interface {}

type Speaker interface {
	Hello()
}

type Walker interface {
    Walk()
}

type Stringer interface {
    String() string
}

type Human struct {
	Greeting string `bson:"asa"`
	Meters int 
}

func (h Human) Walk() {
    fmt.Println(h.Meters)
}

func (h Human) Hello() {
	fmt.Println(h.Greeting)
}

func (h Human) String() string {
    return fmt.Sprintf("Greeting: %s, Meters: %d", h.Greeting, h.Meters)
}

func Lesson(v interface{}) {

	// var	one int = 10

	// var	two MyType = 20
	
	// one = two
	//
	// h := Human{Greeting: "Hello"}
	// s := Speaker(h)
	// s.Hello()
	//
	// w, ok := s.(Walker)
    // if !ok{
    //     fmt.Println("type assertion problem")
    //     return
    // }
    //
    // w.Walk()
	//
	rv := reflect.ValueOf(v)
	tv := reflect.TypeOf(v)
	log.Println(rv)
	log.Println(tv)
	log.Println(rv.Type())
	i := 124.4
	rvi := reflect.ValueOf(i)
	log.Println(rvi)
	log.Println(rvi.Type().Bits())
	var myint64 int64 = 1248
	var myint8 int8 = 12
	log.Println(reflect.ValueOf(myint64).Type().ConvertibleTo(reflect.TypeOf(myint8)))
	if reflect.ValueOf(myint64).Type().ConvertibleTo(reflect.TypeOf(myint8)){
		crv := reflect.ValueOf(myint64).Convert(reflect.TypeOf(myint8))
		log.Println(crv)
	}
	if rv.Kind() == reflect.Ptr {
		rv.Elem()
	}
	if rv.Kind() == reflect.Struct {
		// log.Println(rv.Addr())
		log.Println(rv.NumField())
		log.Println(rv.NumMethod())
	}
	
	reflectTest()
}
//
// func ToString(any interface{}) string {
// 	switch v := any.(type) {
// 	case Stringer:
// 		return v.String()
// 	case int:
// 		return strconv.Itoa(v)
// 	case string:
// 		return any.(string)
// 	case float32:
// 	case float64:
// 	case TEST:
// 		return
// 	default:
// 		return "???"
// 	}
// }

type TEST struct {
	Uuid uuid.UUID
	Code int32
	Name string
}

func reflectTest() {
	m := map[string]interface{}{
		"uuid":"uuid",
		"code":"1",
		"name": "ploas",
	}
	
	p := &TEST{}
	
	// rvp := reflect.ValueOf(p)
	rvp := reflect.Indirect(reflect.ValueOf(p))
	if rvp.Kind() == reflect.Ptr {
		rvp.Elem()
	}
	if rvp.Kind() == reflect.Struct {
		log.Println(rvp.Addr())
		log.Println(rvp.NumField())
		log.Println(rvp.NumMethod())
	}
	for i := 0; i < rvp.Type().NumField(); i++ {
		fld := rvp.Type().Field(i)
		fldn := strings.ToLower(fld.Name)
		fldv := m[fldn]
		ff := rvp.FieldByName(fld.Name)
		if fldvs, ok := fldv.(string); ok && ff.Type().AssignableTo(reflect.TypeOf(uuid.UUID{})) {
			uu, _ := uuid.Parse(fldvs)
			ff.Set(reflect.ValueOf(uu))
		} else if fldvs, ok := fldv.(string); ok && ff.Type().AssignableTo(reflect.TypeOf(int32(0))){
			fInt, _ := strconv.Atoi(fldvs)
			ff.Set(reflect.ValueOf(int32(fInt)))
		}	else {
			ff.Set(reflect.ValueOf(fldv))
		}		
	}
	
	log.Printf("%+v", p)
	
	var constTemplate  = `//code generated
	package {{.Package}}
	type {{.Name}}s []{{.Name}}
	func (c {{.Name}}s)List() []{{.Name}} {
	  return []{{.Name}}{{"{"}}{{.List}}{{"}"}}
	`
	typeName := "color"
	consts := make([]string, 0)
	
	templateData := struct {
		Package string
		Name string
		List string
	}{
		Package: "TestGenerate",
		Name:    typeName,
		List:    strings.Join(consts, ", "),
	}
	t := template.Must(template.New("const-list").Parse(constTemplate))
	if err := t.Execute(os.Stdout, templateData); err != nil {
		log.Println(err)
	}
}