package main
  import(
    "html/template"
    "os"
    "fmt"

  )
  var tpl *template.Template

  func init(){
    tpl = template.Must(template.ParseFiles("formLetter.html"))
  }
  func main(){
    friends := []string{"Nina", "Daniel", "David", "Adrian"}

    err = tpl.ExecuteTemplate(os.Stdout, "formLetter.html")
    if err != nil{
      fmt.Println("There was an error in template", err)
    }

  }
