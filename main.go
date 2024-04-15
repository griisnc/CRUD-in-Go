package main

import (
	"database/sql" // permite hacer las funciones de mysql
	"net/http"

	//"log"
	"fmt"
	"text/template"

	_ "github.com/go-sql-driver/mysql" //se conecta a un driver de github
)

func conexionBD() (conexion *sql.DB) { //conexión a la base de datos
	Driver := "mysql"
	Usuario := "root"
	Contra := ""
	Nombre := "proyecto"
	

	conexion, err := sql.Open(Driver, Usuario+":"+Contra+"@tcp(127.0.0.1)/"+Nombre)
	if err != nil {
		panic(err.Error())
	}
	return conexion
}

//se selecciona la carpeta de las plantillas y se usan los templates
var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() { //declaración de las funciones y los templates a utilizar
	http.HandleFunc("/", Inicio)
	http.HandleFunc("/editar_borrarU", editar_borrarU)
	http.HandleFunc("/editar_borrarProveedor", editar_borrarProveedor)
	http.HandleFunc("/ConsulUsuarios", ConsulUsuarios)
	http.HandleFunc("/inventory", inventory)
	http.HandleFunc("/proveedores", proveedores)
	http.HandleFunc("/proveedores2", proveedores2)
	http.HandleFunc("/crear", Crear)
	http.HandleFunc("/insertar", Insertar)
	http.HandleFunc("/borrar", Borrar)
	http.HandleFunc("/editar", Editar)
	http.HandleFunc("/borrarProveedor", borrarProveedor)
	http.HandleFunc("/editar2", Editar2)
	http.HandleFunc("/editar3", Editar3)
	http.HandleFunc("/actualizar2", Actualizar2)
	http.HandleFunc("/actualizar", Actualizar)
	http.HandleFunc("/enviar3", Enviar3)
	http.HandleFunc("/actualizar3", Actualizar3)
	http.HandleFunc("/borrar3", Borrar3)
	http.HandleFunc("/productos", Productos)
	http.HandleFunc("/enviar", Enviar)
	http.HandleFunc("/enviar2", Enviar2)
	http.HandleFunc("/login", Login)
	fmt.Println("Servidor Corriendo....") //mensaje en consola
	http.ListenAndServe(":8080", nil)     //abre un puerto a usaar para el localhost
}

func Borrar(w http.ResponseWriter, r *http.Request) { //funcion que permite borrar a cierto usaurio
	//fmt.Println(Id_u)
	Id_u := r.URL.Query().Get("id_u")                                                        //del formulario se obtiene el dato que se ingresó en caja de texto id_u
	conexionestablecida := conexionBD()                                                      // se establece la conexion a la base de datos
	borrarregistros, err := conexionestablecida.Prepare("DELETE FROM usuarios WHERE Id_u=?") //se realiza la consulta a la base de datos y se especifica la tabla
	if err != nil {
		panic(err.Error())
	}
	borrarregistros.Exec(Id_u) //se indica el usario que se eliminará deacuerdo a su id

	http.Redirect(w, r, "/", 301) //al terminar regresa al menu principal

}
func borrarProveedor(w http.ResponseWriter, r *http.Request) { //funcion que permite borrar proveedor
	id_pro := r.URL.Query().Get("id_pro") //del formulario se obtiene el dato que se ingresó en caja de texto id_pro
	fmt.Println("kkk", id_pro)

	conexionestablecida := conexionBD() //se establece la conexion a la base de datos
	borrarregistros, err := conexionestablecida.Prepare("DELETE FROM proveedores WHERE id_pro=?")
	//se realiza la consulta a la base de datos y se especifica la tabla, se ejecuta el comando delete
	if err != nil {
		panic(err.Error())
	}
	borrarregistros.Exec(id_pro) //se indica el usario que se eliminará deacuerdo a su id

	http.Redirect(w, r, "/", 301) //al terminar regresa al menu principal

}
func Borrar3(w http.ResponseWriter, r *http.Request) { //funcion para borrar algun producto
	id_p := r.URL.Query().Get("id_p") //del formulario se obtiene el dato que se ingresó en caja de texto id_p
	fmt.Println("id: ", id_p)

	conexionestablecida := conexionBD() //se establece la conexion a la base de datos
	//se realiza la consulta a la base de datos y se especifica la tabla, se ejecuta el comando delete
	borrarregistros, err := conexionestablecida.Prepare("DELETE FROM productos WHERE id_p=?")
	if err != nil {
		panic(err.Error())
	}
	borrarregistros.Exec(id_p) //se indica el usario que se eliminará deacuerdo a su id

	http.Redirect(w, r, "/inventory", 301) //al terminar regresa al menu principal

}

type users struct { //estrutura de la tabla de usuarios
	Id_u    int    //id de usuario
	Usuario string // nombre del usuario
	Pass    string // contraseña
	Tipo_u  string //tipo de usuario
}
type prod struct { //se hace una estructura en la cual se guardan la información de cada producto
	Id_p             int    // su ID
	Descripcion      string //Una breve descripcion
	Tipo_de_producto string //tipo, puede ser lacteos, limpieza
	Stock            int    //cantidad existente
	Id_pro           int    //id del proveedor
	Precio_unidad    int    //precio por cada producto
	Precio_t         int    //se estima el valor total de los productos
}
type prove struct { //se hace una estructura en la cual se guardan la información de cada proveedor
	Id_pro    int    //id del proveedor
	Proveedor string //nombre del proveedor
}

func editar_borrarProveedor(w http.ResponseWriter, r *http.Request) { //funcion que nos permite abrir la plantilla y enviarle datos
	conexionestablecida := conexionBD()                                      //conexion a la base de datos
	registros, err := conexionestablecida.Query("SELECT * FROM proveedores") //se ejecuta la query a la base de datos
	if err != nil {                                                          //en caso de que halla algun erro en al query nos lo marcará
		panic(err.Error())
	}
	us := prove{}              //variable del tipo estrucutra de proveedores
	arregloprovee := []prove{} //arreglo que guarda la informacion que hay en la estructura he imprime en consola
	for registros.Next() {     //ciclo para obtener los datos de la consulta
		var Id_pro int       //guarda el id del proveedor
		var Proveedor string //guarda el nombre del proveedor

		err = registros.Scan(&Id_pro, &Proveedor) //guarda en las variables los datos obtenidos

		if err != nil { //en caso de que no se pueda obtener nos indicara el error
			panic(err.Error())
		}
		us.Id_pro = Id_pro //al elemento tipo id_pro se le asigna su valor en la base datos
		us.Proveedor = Proveedor

		arregloprovee = append(arregloprovee, us) //al arreglo se le agregará los valores
		//de la estructura y de la estructura con la informacion de la base de datos

	}
	//ejecuta la plantilla se le indica que será de lectura, el nombre del archivo
	//y se envia el arreglo con los datos obtenidos
	plantillas.ExecuteTemplate(w, "editar_borrarProveedor", arregloprovee)
}
func editar_borrarU(w http.ResponseWriter, r *http.Request) { //funcion que nos permite abrir la plantilla y enviarle datos
	conexionestablecida := conexionBD()                                   //conexion a la base de datos
	registros, err := conexionestablecida.Query("SELECT * FROM usuarios") //se ejecuta la query a la base de datos
	if err != nil {                                                       //en caso de que halla algun erro en al query nos lo marcará
		panic(err.Error())
	}
	us := users{}             //variable del tipo estrucutra de proveedores
	arreglousers := []users{} //arreglo que guarda la informacion que hay en la estructura he imprime en consola
	for registros.Next() {    //ciclo para obtener los datos de la consulta
		var Id_u int //guarda el id del usuario
		var Usuario, Pass, Tipo_u string
		err = registros.Scan(&Id_u, &Usuario, &Pass, &Tipo_u) //guarda en las variables los datos obtenidos

		if err != nil { //en caso de que no se pueda obtener nos indicara el error
			panic(err.Error())
		}
		us.Id_u = Id_u //a cada elemento se le asigna su valor en la base datos
		us.Usuario = Usuario
		us.Pass = Pass
		us.Tipo_u = Tipo_u
		arreglousers = append(arreglousers, us) //al arreglo se le agregará los valores
		//de la estructura y de la estructura con la informacion de la base de datos

	}
	//fmt.Println(arreglousers)
	//ejecuta la plantilla se le indica que será de lectura, el nombre del archivo
	//y se envia el arreglo con los datos obtenidos
	plantillas.ExecuteTemplate(w, "editar_borrarU", arreglousers)
}
func Inicio(w http.ResponseWriter, r *http.Request) { //función que abre la pagina principal del inventario
	plantillas.ExecuteTemplate(w, "inicio", nil)
}
func Crear(w http.ResponseWriter, r *http.Request) { //función que abre la pagina para crear un usuario
	plantillas.ExecuteTemplate(w, "crear", nil)
}
func Insertar(w http.ResponseWriter, r *http.Request) { //función que recibe los datos del formulario de nuevo usuario
	if r.Method == "POST" { //se verifica si los datos se mandaron por el metodo POST
		//recibe los datos
		usuario := r.FormValue("usuario")
		pass := r.FormValue("pass")
		tipo_u := r.FormValue("tipo_u")

		conexionestablecida := conexionBD() //se manda llamar la conexion
		// se hace la consulta para agregar al usuario
		insertarregistros, err := conexionestablecida.Prepare("INSERT INTO usuarios(usuario,pass,tipo_u) VALUES (?,?,?)")
		if err != nil { //se verifica que no haya errores
			panic(err.Error())
		}
		insertarregistros.Exec(usuario, pass, tipo_u) //se mandan los datos recibidos por parametros a la bd

		http.Redirect(w, r, "/", 301) //despues de la consulta se redireciona a la pagina principal
	}
}
func Actualizar(w http.ResponseWriter, r *http.Request) { //función que actualiza los datos del usuario despues de editarlos
	if r.Method == "POST" { //se verifica si los datos se mandaron por el metodo POST
		//recibe los datos
		id_u := r.FormValue("id_u")
		usuario := r.FormValue("usuario")
		pass := r.FormValue("pass")
		tipo_u := r.FormValue("tipo_u")

		conexionestablecida := conexionBD() //se manda llamar la conexion
		// se hace la consulta para actualizar al usuario
		modificarregistros, err := conexionestablecida.Prepare("UPDATE usuarios SET usuario=?, pass=?, tipo_u=? WHERE Id_u=? ")
		if err != nil {
			panic(err.Error()) //se verifica que no haya errores
		}
		modificarregistros.Exec(usuario, pass, tipo_u, id_u) //se mandan los datos recibidos por parametros a la bd

		http.Redirect(w, r, "/", 301) //despues de la consulta se redireciona a la pagina principal
	}
}
func Actualizar2(w http.ResponseWriter, r *http.Request) { //función que actualiza la tabla de los proveedores
	if r.Method == "POST" { //se verifica si los datos se mandaron por el metodo POST
		//recibe los datos
		id_pro := r.FormValue("id_pro")
		proveedor := r.FormValue("proveedor")
		//fmt.Print(id_pro)
		conexionestablecida := conexionBD() //se manda llamar la conexion
		// se hace la consulta para actualizar a los proveedores
		modificarregistros, err := conexionestablecida.Prepare("UPDATE proveedores SET Proveedor=? WHERE id_pro=? ")
		if err != nil {
			panic(err.Error()) //se verifica que no haya errores
		}
		modificarregistros.Exec(proveedor, id_pro) //se mandan los datos recibidos por parametros a la bd

		http.Redirect(w, r, "/", 301) //despues de la consulta se redireciona a la pagina principal
	}
}
func Actualizar3(w http.ResponseWriter, r *http.Request) { //función para actualizar los productos
	if r.Method == "POST" { //se verifica si los datos se mandaron por el metodo POST
		//recibe los datos
		id_p := r.FormValue("id_p")
		descripcion := r.FormValue("descripcion")
		Tipo_de_producto := r.FormValue("Tipo_de_producto")
		stock := r.FormValue("stock")
		//Id_pro := r.FormValue("Id_pro")
		precio_unidad := r.FormValue("precio_unidad")
		precio_t := r.FormValue("precio_t")

		fmt.Print(id_p)

		conexionestablecida := conexionBD() //se manda llamar la conexion
		// se hace la consulta para actualizar a los productos
		modificarregistros, err := conexionestablecida.Prepare("UPDATE productos SET descripcion=?,Tipo_de_Producto=?,stock=?,precio_unidad=?,precio_t=? WHERE id_p=? ")
		if err != nil {
			panic(err.Error()) //se verifica que no haya errores
		}
		modificarregistros.Exec(descripcion, Tipo_de_producto, stock, precio_unidad, precio_t, id_p) //se mandan los datos recibidos por parametros a la bd

		http.Redirect(w, r, "/inventory", 301) //despues de la consulta se redireciona a la pagina del inventario
	}
}
func Editar(w http.ResponseWriter, r *http.Request) { //función que manda el id del usuario a editar mediante un boton
	//se reciben los datos por parametro con el metodo GET
	idu := r.URL.Query().Get("id_u")
	fmt.Println(idu)

	conexionestablecida := conexionBD()                                                    //se manda llamar la conexion
	registro, err := conexionestablecida.Query("SELECT * FROM usuarios WHERE id_u=?", idu) //se verifica el usuario recibido
	if err != nil {
		panic(err.Error()) //se verifica que no haya errores
	}
	us := users{} //se manda llamr la estructura de usuarios

	for registro.Next() { //se recorre la tabla de usuarios hasta encontrar el usuario recibido
		//se declaran las variables
		var Id_u int
		var Usuario, Pass, Tipo_u string
		err = registro.Scan(&Id_u, &Usuario, &Pass, &Tipo_u) //se reciben

		if err != nil {
			panic(err.Error()) //se verifican que no haya errores
		}
		//se pasan las variables a la estructura
		us.Id_u = Id_u
		us.Usuario = Usuario
		us.Pass = Pass
		us.Tipo_u = Tipo_u

	}
	//fmt.Println(us)
	plantillas.ExecuteTemplate(w, "editar", us) //se reciben los datos actualizados

}
func Editar2(w http.ResponseWriter, r *http.Request) { //función de boton de editar proveedores
	//recibe el id por parametro
	idu := r.URL.Query().Get("id_pro")
	fmt.Println(idu)

	conexionestablecida := conexionBD()                                                         //se manda llamar la conexion
	registro, err := conexionestablecida.Query("SELECT * FROM proveedores WHERE id_pro=?", idu) //se verifica el usuario recibido
	if err != nil {
		panic(err.Error()) //se verifica que no haya errores
	}
	us := prove{}              //se manda llamar la estructura de los proveedores
	arregloprovee := []prove{} //se declara un arrglo de proveedores
	for registro.Next() {      //se recorre la tabla de proveedores hasta encontrar el id recibido
		//se declaran las variables
		var Id_pro int
		var Proveedor string

		err = registro.Scan(&Id_pro, &Proveedor) //se verifican que no haya errores

		if err != nil {
			panic(err.Error()) //se verifican que no haya errores
		}
		us.Id_pro = Id_pro
		us.Proveedor = Proveedor

		arregloprovee = append(arregloprovee, us) //se copian los datos del arreglo a la estructura

	}
	//fmt.Println(us)
	plantillas.ExecuteTemplate(w, "editar2", us) //se reciben los datos actualizados

}
func Editar3(w http.ResponseWriter, r *http.Request) { //función de boton de editar productos
	//recibe el id por parametro
	idu := r.URL.Query().Get("id_p")
	fmt.Println(idu)
	conexionestablecida := conexionBD()                                                      //se manda llamar la conexion
	registros, err := conexionestablecida.Query("SELECT * FROM productos WHERE id_p=?", idu) //se verifica el usuario recibido
	if err != nil {
		panic(err.Error()) //se verifican que no haya errores
	}
	us := prod{}            //se manda llamar la estructura de los productos
	arregloprod := []prod{} //se declara el arreglo segun la estructura
	for registros.Next() {  //se recorre la tabla de proveedores hasta encontrar el id recibido
		var Id_p int
		var Descripcion string
		var Tipo_de_producto string
		var Stock int
		var Id_pro int
		var Precio_unidad int
		var Precio_t int
		//reciben los datos segun el id
		err = registros.Scan(&Id_p, &Descripcion, &Tipo_de_producto, &Stock, &Id_pro, &Precio_unidad, &Precio_t)

		if err != nil {
			panic(err.Error()) //se verifican que no haya errores
		}
		us.Id_p = Id_p
		us.Descripcion = Descripcion
		us.Id_pro = Id_pro
		us.Precio_t = Precio_t
		us.Precio_unidad = Precio_unidad
		us.Stock = Stock
		us.Tipo_de_producto = Tipo_de_producto
		arregloprod = append(arregloprod, us) //se copian los datos del arreglo a la estructura

	}
	//fmt.Println(arreglousers)
	plantillas.ExecuteTemplate(w, "editar3", us) //se reciben los datos actualizados
}
func Enviar3(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		user := r.FormValue("usuario")
		//pass := r.FormValue("pass")
		//fmt.Println("user: ?", user, "pass: ?", pass)
		conexionestablecida := conexionBD()

		enviarregistros, err := conexionestablecida.Prepare("SELECT * FROM usuarios WHERE usuario=?")
		//	fmt.Println(consulta)

		if err != nil {
			panic(err.Error())
			return
		}
		if err == nil {
			fmt.Print("exito")
		}
		enviarregistros.Exec(user)
		http.Redirect(w, r, "/", 301)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	
	plantillas.ExecuteTemplate(w, "login", nil)
}
func Enviar(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		descripcion := r.FormValue("descripcion")
		Tipo_de_producto := r.FormValue("Tipo_de_producto")
		stock := r.FormValue("stock")
		Id_pro := r.FormValue("Id_pro")
		precio_unidad := r.FormValue("precio_unidad")
		precio_t := r.FormValue("precio_t")

		conexionestablecida := conexionBD()
		insertarproducts, err := conexionestablecida.Prepare("INSERT INTO productos (descripcion, Tipo_de_producto, stock, Id_pro, precio_unidad, precio_t) VALUES (?,?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insertarproducts.Exec(descripcion, Tipo_de_producto, stock, Id_pro, precio_unidad, precio_t)

		http.Redirect(w, r, "/inventory", 301)
	}
}

func Productos(w http.ResponseWriter, r *http.Request) { //función que redirige a la pagina de productos
	conexionestablecida := conexionBD() //llama la conexión de la bd
	//hace una consulta para desplegar los proveedores
	mprov, err := conexionestablecida.Query("SELECT * FROM proveedores")
	if err != nil {
		panic(err.Error()) //verifica los errores
	}
	us := prove{}              //manda llamar la estructura de proveedores
	arregloprovee := []prove{} //crea un arreglo sobre la estructura
	for mprov.Next() {         //se recorre la tabla de proveedores hasta encontrar el id recibido
		var Id_pro int
		var Proveedor string
		err = mprov.Scan(&Id_pro, &Proveedor) //recibe los datos a desplegar
		if err != nil {
			panic(err.Error()) //verifica los errores
		}
		us.Id_pro = Id_pro
		us.Proveedor = Proveedor
		arregloprovee = append(arregloprovee, us) //copia los datos recibido al arreglo

	}
	fmt.Println("Proveedores: ", arregloprovee)
	plantillas.ExecuteTemplate(w, "productos", nil) //manda la consulta a la pagina
}
func Proveedores(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "proveedores", nil)
}
func Enviar2(w http.ResponseWriter, r *http.Request) {
	Proveedor := r.FormValue("Proveedor")

	conexionestablecida := conexionBD()
	insertarprovedores, err := conexionestablecida.Prepare("INSERT INTO proveedores (Proveedor) VALUES (?)")
	if err != nil {
		panic(err.Error())
	}
	insertarprovedores.Exec(Proveedor)

	http.Redirect(w, r, "/", 301)
}
func Admin(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Holi") //mensaje en el navegador
	plantillas.ExecuteTemplate(w, "login", nil)
}
func inventory(w http.ResponseWriter, r *http.Request) {
	conexionestablecida := conexionBD()
	registros, err := conexionestablecida.Query("SELECT * FROM productos")
	if err != nil {
		panic(err.Error())
	}
	us := prod{}
	arregloprod := []prod{}
	for registros.Next() {
		var Id_p int
		var Descripcion string
		var Tipo_de_producto string
		var Stock int
		var Id_pro int
		var Precio_unidad int
		var Precio_t int

		err = registros.Scan(&Id_p, &Descripcion, &Tipo_de_producto, &Stock, &Id_pro, &Precio_unidad, &Precio_t)

		if err != nil {
			panic(err.Error())
		}
		us.Descripcion = Descripcion
		us.Id_p = Id_p
		us.Id_pro = Id_pro
		us.Precio_t = Precio_t
		us.Precio_unidad = Precio_unidad
		us.Stock = Stock
		us.Tipo_de_producto = Tipo_de_producto
		arregloprod = append(arregloprod, us)

	}
	//fmt.Println(arreglousers)
	plantillas.ExecuteTemplate(w, "inventory", arregloprod)
}
func proveedores(w http.ResponseWriter, r *http.Request) {
	conexionestablecida := conexionBD()
	registros, err := conexionestablecida.Query("SELECT * FROM proveedores")
	if err != nil {
		panic(err.Error())
	}
	us := prove{}
	arregloprovee := []prove{}
	for registros.Next() {
		var Id_pro int
		var Proveedor string

		err = registros.Scan(&Id_pro, &Proveedor)

		if err != nil {
			panic(err.Error())
		}
		us.Id_pro = Id_pro
		us.Proveedor = Proveedor

		arregloprovee = append(arregloprovee, us)

	}
	//fmt.Println(arreglousers)
	plantillas.ExecuteTemplate(w, "proveedores", arregloprovee)
}
func ConsulUsuarios(w http.ResponseWriter, r *http.Request) {
	conexionestablecida := conexionBD()
	registros, err := conexionestablecida.Query("SELECT * FROM usuarios")
	if err != nil {
		panic(err.Error())
	}
	us := users{}
	arreglousers := []users{}
	for registros.Next() {
		var Id_u int
		var Usuario, Pass, Tipo_u string
		err = registros.Scan(&Id_u, &Usuario, &Pass, &Tipo_u)
		if err != nil {
			panic(err.Error())
		}
		us.Id_u = Id_u
		us.Usuario = Usuario
		us.Pass = Pass
		us.Tipo_u = Tipo_u
		arreglousers = append(arreglousers, us)
	}
	//fmt.Println(arreglousers)
	plantillas.ExecuteTemplate(w, "ConsulUsuarios", arreglousers)
}
func proveedores2(w http.ResponseWriter, r *http.Request) {
	conexionestablecida := conexionBD()
	registros, err := conexionestablecida.Query("SELECT * FROM proveedores")
	if err != nil {
		panic(err.Error())
	}
	us := prove{}
	arregloprovee := []prove{}
	for registros.Next() {
		var Id_pro int
		var Proveedor string

		err = registros.Scan(&Id_pro, &Proveedor)

		if err != nil {
			panic(err.Error())
		}
		us.Id_pro = Id_pro
		us.Proveedor = Proveedor

		arregloprovee = append(arregloprovee, us)

	}
	//fmt.Println(arreglousers)
	plantillas.ExecuteTemplate(w, "proveedores2", arregloprovee)
}
