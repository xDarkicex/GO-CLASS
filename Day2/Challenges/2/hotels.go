package main
  import(
    "log"
    "os"
    "text/template"
  )

  var tpl *template.Template

  type Address struct {
     Number string
     Street string
     City   string
     State  string
     Zip    string
  }

  type region struct {
  	Addresss []Address
  }
  type hotel struct {
    Name string
    Southern, Central, Northern region
  }
  func init() {
  	tpl = template.Must(template.ParseFiles("index.gohtml"))

  }
  func main() {
    hotels := []hotel{
  		hotel{
        Name: "Best Western",
  			Southern: region{
  				Addresss: []Address{
  					Address{"1345", "South West BLV", "San Diago", "California", "78372"},
  					Address{"130", "North Introduction", "San Diago", "California", "76834"},
  					Address{"140", "Vagdenes AVE", "San Diago", "California", "78328"},
  				},
  			},
  			Northern: region{
  				Addresss: []Address{
  					Address{"50", "North Introduction", "San Diago", "California", "76834"},
  				},
  			},
  		},
  	hotel{
        Name: "Big 5",
  			Southern: region{
  				Addresss: []Address{
            Address{"1345", "South West BLV", "San Diago", "California", "78372"},
  					Address{"1342", "West Introduction", "San Diago", "California", "76834"},
  					Address{"1342", "Smith AVE", "San Diago", "California", "78328"},
  				},
  			},
        Northern: region{
  				Addresss: []Address{
  					Address{"5450", "East Introduction Street", "San Diago", "California", "76834"},
  				},
  			},
  		},
    }

  	err := tpl.Execute(os.Stdout, hotels)
  	if err != nil {
  		log.Fatalln(err)
  	}
  }
