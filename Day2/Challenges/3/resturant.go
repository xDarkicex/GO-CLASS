package main
  import(
    "log"
    "os"
    "text/template"
  )

  var tpl *template.Template
  type menu struct{
    Name string
    Price string
  }

  type meal struct {
  	Menus []menu
  }
  type resturant struct {
    Name string
    Breakfast, Lunch, Dinner meal
  }
  func init() {
  	tpl = template.Must(template.ParseFiles("index.gohtml"))

  }
  func main() {
    resturants := []resturant{
  		resturant{
        Name: "Dennys",
  			Breakfast: meal{
  				Menus: []menu{
  					menu{"Pankcake", "2.99"},
  					menu{"Eggs", "4.99"},
  				},
  			},
  			Lunch: meal{
  					Menus: []menu{
  					menu{"Sandwitch", "7.99"},
            menu{"Hamburger", "11.99"},
  				},
  			},
        Dinner: meal{
          	Menus: []menu{
            menu{"Hamburger", "11.99"},
            menu{"Salad", "5.99"},
          },
        },
      },
  	resturant{
        Name: "Ruth Chris",
  			Dinner: meal{
  					Menus: []menu{
            menu{"Porter Steak", "39.99"},
  					menu{"RibEye", "76.83"},
  				},
  			},
      },
    }

  	err := tpl.Execute(os.Stdout, resturants)
  	if err != nil {
  		log.Fatalln(err)
  	}
  }
