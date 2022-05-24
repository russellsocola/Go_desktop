//uso basico de fyne.io
// package main

// import (
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/widget"
// )

// func main() {
// 	a := app.New()
// 	w := a.NewWindow("DESKTOP ONE BY RUSSELL") // nombre de la ventana emergentee

// 	//contenido de la ventana emergente
// 	w.SetContent(widget.NewLabel("Hello World!"))
// 	w.ShowAndRun()

// }
//-------------------------EJEMPLO 2 DEL USO DE fYNE -------------

// package main

// import (
// 	"fmt"

// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/widget"
// )

// func main() {
// 	myApp := app.New()
// 	myWindow := myApp.NewWindow("Hola")           // nombre de la ventana emergentee
// 	myWindow.SetContent(widget.NewLabel("Hello")) // contenido de la ventana emergentee

// 	myWindow.Show() //abre la ventana
// 	myApp.Run()     //corre la funcion, la siguiente funcion no se ejecutara hasta que se cierre la ventana
// 	tidyUp()        //al cerrar la ventana muestra mensaje en consola
// }

// func tidyUp() {
// 	fmt.Println("Exited") //se imprimira en la consola.
// }

//----------EJEMPLO 3 DEL USO DEL FYNE------

// package main

// import (
// 	"fmt"
// 	"time"

// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/widget"
// )

//la funcion establece el formato de hora HH:MM:ss

// func updateTime(clock *widget.Label) {
// 	formatted := time.Now().Format("Time: 03:04:05")
// 	clock.SetText(formatted)
// }
// func main() {
// 	a := app.New()
// 	w := a.NewWindow("Clock") //nombre de la ventana emergente

// 	clock := widget.NewLabel("")
// 	w.SetContent(clock)
// 	//la go func permite ejecutar en segundo plano, la go rutina
// 	go func() {
// 		for range time.Tick(time.Second) { //llamo a la funcion
// 			updateTime(clock) //mando parametro a la funcion
// 		}
// 	}()

// 	w.Resize(clock.MinSize()) //se ajusta a el tanaño del contenido
// 	w.ShowAndRun()
// 	fmt.Println("Se ejecuto con exito")
// }

//----------PARTE 4 MANEJO DE MULTI VENTANAS-----

// package main

// import (
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/widget"
// )

// func main() {

// 	a := app.New()
// 	//se abriran dos ventanas simultaneamente
// 	w := a.NewWindow("First Window")

// 	w2 := a.NewWindow("Second Window")

// 	//Añadimos contenido a las ventanas

// 	w.SetContent(widget.NewLabel("Hola esta es la ventana uno de Russell"))
//	w.Resize(fyne.NewSize(100, 100)) //Agregamos un tamaño definido a la ventana
// 	//Agrego un master a la ventana principal (si se cierra esta ventana se cerrara todo)
// 	w.SetMaster()

// 	w.Show()

// 	w2.SetContent(widget.NewLabel("Hola esta es la segunda ventana de Russell"))
// 	w2.Show()

// 	a.Run()

// }

//----------------PARTE 5 PRUEBA DDE APLICACACIONES GRAFICAS-----

package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeUI() (*widget.Label, *widget.Entry) {
	out := widget.NewLabel("Hello world!")
	in := widget.NewEntry()

	in.OnChanged = func(content string) {
		out.SetText("Hello " + content + " !")
	}
	return out, in
}

func main() {
	a := app.New()
	w := a.NewWindow("Hello Person")

	w.SetContent(container.NewVBox(makeUI()))
	w.ShowAndRun()
}
