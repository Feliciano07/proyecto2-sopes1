// Paquete principal, acá iniciará la ejecución
package main

// Importar dependencias, notar que estamos en un módulo llamado grpctuiter
import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"strconv"

	"log"
	"net"
	greetpb "tuiterserver/greet.pb"

	"google.golang.org/grpc"

	"math/rand"

	"cloud.google.com/go/pubsub"
	"github.com/joho/godotenv"
)

// Iniciar una estructura que posteriormente gRPC utilizará para realizar un server
type server struct {
}

//estructura del juego
type Juego struct {
	Request_number int    `json:"request_number"`
	Game           int    `json:"game"`
	Gamename       string `json:"gamename"`
	Winner         string `json:"winner"`
	Players        int    `json:"players"`
	Worker         string `json:"worker"`
}

func CargarClaves() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./clave.json")
}

func Parser(game Juego) string {
	b, err := json.Marshal(game)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(b))
	return string(b)
}

func goDotEnvVariable(key string) string {

	// Leer el archivo .env ubicado en la carpeta actual
	err := godotenv.Load(".env")

	// Si existio error leyendo el archivo
	if err != nil {
		log.Fatalf("Error cargando las variables de entorno")
	}

	// Enviar la variable de entorno que se necesita leer
	return os.Getenv(key)
}

func publish(msg string) error {
	projectID := goDotEnvVariable("PROJECT_ID")
	topicID := goDotEnvVariable("TOPIC_ID")
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		fmt.Println("error")
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{Data: []byte(msg)})

	id, err := result.Get(ctx)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("Error: %v", err)
	}
	fmt.Println("Published a message; msg ID: %v\n", id)
	return nil
}

func GameRandom(jugadores int) int {
	val := rand.Intn(jugadores)
	return val
}

func GameMaximo(jugadores int) int {
	return jugadores
}

// Función que será llamada desde el cliente
// Debemos pasarle un contexto donde se ejecutara la funcion
// Y utilizar las clases que fueron generadas por nuestro proto file
// Retornara una respuesta como la definimos en nuestro protofile o un error
func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf(">> SERVER: Función Greet llamada con éxito. Datos: %v\n", req)

	// Todos los datos podemos obtenerlos desde req
	// Tendra la misma estructura que definimos en el protofile
	// Para ello utilizamos en este caso el GetGreeting
	juego := req.GetGreeting().GetJuego()
	nombreJuego := req.GetGreeting().GetNombreJuego()
	jugadores := req.GetGreeting().GetJugadores()

	ganador := 0
	jugadores2 := int(jugadores)

	if juego == 1 {
		ganador = GameRandom(jugadores2)
	} else {
		ganador = GameRandom(jugadores2)
	}

	game := Juego{
		Request_number: 1,
		Game:           int(juego),
		Gamename:       nombreJuego,
		Winner:         strconv.Itoa(ganador),
		Players:        int(jugadores),
		Worker:         "PubSub",
	}

	salida := Parser(game)
	publish(salida)

	result := " ||| No.Juego -> " + strconv.FormatInt(juego, 10) + "| nombreJuego -> " + nombreJuego + "| Jugadores -> " + strconv.FormatInt(jugadores, 10) + "| Ganador-> " + strconv.Itoa(ganador) + " |||"

	fmt.Printf(">> SERVER: %s\n", result)
	// Creamos un nuevo objeto GreetResponse definido en el protofile
	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}

// Funcion principal
func main() {

	//Cargar claves google
	CargarClaves()

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("sin .env")
	}

	// Leer el host de las variables del ambiente
	host := os.Getenv("HOST")
	fmt.Println(">> SERVER: Iniciando en ", host)

	// Primero abrir un puerto para poder escuchar
	// Lo abrimos en este puerto arbitrario
	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf(">> SERVER: Error inicializando el servidor: %v", err)
	}

	fmt.Println(">> SERVER: Empezando server gRPC")

	// Ahora si podemos iniciar un server de gRPC
	s := grpc.NewServer()

	// Registrar el servicio utilizando el codigo que nos genero el protofile
	greetpb.RegisterGreetServiceServer(s, &server{})

	fmt.Println(">> SERVER: Escuchando servicio...")
	// Iniciar a servir el servidor, si hay un error salirse
	if err := s.Serve(lis); err != nil {
		log.Fatalf(">> SERVER: Error inicializando el listener: %v", err)
	}
}
