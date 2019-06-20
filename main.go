// This script demonstrates how to use the apiclient package.
//
// Before running, set the RIOT_APIKEY enviornment variable to your developer
// key from the Riot developer portal.
package main

import (
	"context"
	"encoding/json"
	"strconv"

	//"encoding/json"
	"fmt"
	"log"
	"net/http"

	//"strconv"

	"github.com/AdrianOrtuno/appSumamoos/Application"
	"github.com/AdrianOrtuno/appSumamoos/Domain/region"
	"github.com/AdrianOrtuno/appSumamoos/Persistance/ratelimit"

	"github.com/AdrianOrtuno/appSumamoos/Vendor/gorilla/mux"
)

const (
	league = "6b5c7950-5260-11e7-8125-c81f66dbb56c"
	reg    = region.EUW1
	apiKey = "RGAPI-34f7e6fc-3fcd-4139-bdab-fc678fc6e773"
	/*AddresCosmosDB      = "dev-bigdataesports.documents.azure.com:10255"
	DB                  = "leagueoflegends"
	userName            = "dev-bigdataesports"
	password            = "abWVlAQ73U4WVSD3QUtlkEZPJds2WcCA2OrrGrbl3xHluIp6t571JuSMShP7JeWcmMmJ8oE0ylPNU9OmosYtWg=="
	collectionSummoner  = "summoner"
	collectionListMatch = "listMatch"
	collectionMatch     = "matchs"*/
)

func main() {
	// Se crean las rutas de llamadas de la API
	router := mux.NewRouter()
	router.HandleFunc("/getSummoner/{summoner}", getSummoner).Methods("GET")
	router.HandleFunc("/getMatchList/{summoner}", getMatchList).Methods("GET")
	router.HandleFunc("/getMatch/{idMatch}", getMatch).Methods("GET")
	//router.HandleFunc("/getAllSummonerFromCosmosDB", getAllSummonerFromCosmosDB).Methods("GET")
	//router.HandleFunc("/updateSummoner/{nombreSummoner}", updateSummoner).Methods("GET")
	http.ListenAndServe(":8080", router)
}

func getSummoner(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Start getSummoner")
	//Variables iniciales
	var summonerName string
	params := mux.Vars(r)
	key := apiKey
	httpClient := http.DefaultClient
	ctx := context.Background()
	limiter := ratelimit.NewLimiter()
	client := Application.New(key, httpClient, limiter)

	summonerName = params["summoner"]

	//fmt.Println("GetBySummonerName")
	summoner, err := client.GetBySummonerName(ctx, reg, summonerName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//COSMOS DB

	/*dialInfo := &mgo.DialInfo{
		Addrs:    []string{AddresCosmosDB},
		Timeout:  60 * time.Second,
		Database: DB,
		Username: userName,
		Password: password,
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		},
	}
	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	session, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}

	defer session.Close()

	// SetSafe changes the session safety mode.
	// If the safe parameter is nil, the session is put in unsafe mode, and writes become fire-and-forget,
	// without error checking. The unsafe mode is faster since operations won't hold on waiting for a confirmation.
	// http://godoc.org/labix.org/v2/mgo#Session.SetMode.
	session.SetSafe(&mgo.Safe{})

	// get collection
	collection := session.DB(DB).C(collectionSummoner)

	// insert Document in collection
	err = collection.Insert(summoner)

	if err != nil {
		log.Fatal("Problem inserting data: ", err)
		return
	}*/

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(summoner)
	fmt.Println("End getSummoner")
}

func getMatchList(w http.ResponseWriter, r *http.Request) {
	//Variables iniciales
	var summonerName string
	params := mux.Vars(r)
	key := apiKey
	httpClient := http.DefaultClient
	ctx := context.Background()
	limiter := ratelimit.NewLimiter()
	client := Application.New(key, httpClient, limiter)

	summonerName = params["summoner"]

	//fmt.Println("GetBySummonerName")
	summoner, err := client.GetBySummonerName(ctx, reg, summonerName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//fmt.Println(summoner)

	fmt.Println("Start GetMatchlist")
	matchList, err := client.GetMatchlist(ctx, reg, summoner.AccountID, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//COSMOS DB

	/*dialInfo := &mgo.DialInfo{
		Addrs:    []string{AddresCosmosDB},
		Timeout:  60 * time.Second,
		Database: DB,
		Username: userName,
		Password: password,
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		},
	}

	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	session, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}

	defer session.Close()

	session.SetSafe(&mgo.Safe{})

	// get collection
	collection := session.DB(DB).C(collectionListMatch)

	// insert Document in collection
	err = collection.Insert(matchList)*/

	if err != nil {
		log.Fatal("Problem inserting data: ", err)
		return
	}

	//fmt.Println(tmpl.Execute(w, matchList))
	//retun json
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(matchList)
	fmt.Println("End GetMatchlist")
}

func getMatch(w http.ResponseWriter, r *http.Request) {
	//Variables iniciales
	var idMatch int64
	params := mux.Vars(r)
	key := apiKey
	httpClient := http.DefaultClient
	ctx := context.Background()
	limiter := ratelimit.NewLimiter()
	client := Application.New(key, httpClient, limiter)

	idMatch, err := strconv.ParseInt(params["idMatch"], 10, 64)

	fmt.Println("GetMatch")
	myMatch, err := client.GetMatch(ctx, reg, idMatch)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//COSMOS DB

	/*dialInfo := &mgo.DialInfo{
		Addrs:    []string{AddresCosmosDB},
		Timeout:  60 * time.Second,
		Database: DB,
		Username: userName,
		Password: password,
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		},
	}

	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	session, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}

	defer session.Close()
	session.SetSafe(&mgo.Safe{})

	// get collection
	collection := session.DB(DB).C(collectionMatch)

	// insert Document in collection
	err = collection.Insert(myMatch)

	if err != nil {
		log.Fatal("Problem inserting data: ", err)
		return
	}*/

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(myMatch)

	fmt.Println("End GetMatch")
}

/*func getAllSummonerFromCosmosDB(w http.ResponseWriter, r *http.Request) {
	fmt.Println("START Obtain summoner from CosmosDB")
	//Conectamos con Cosmos DB
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{AddresCosmosDB},
		Timeout:  60 * time.Second,
		Database: DB,
		Username: userName,
		Password: password,
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		},
	}

	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	session, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}

	defer session.Close()
	session.SetSafe(&mgo.Safe{})

	// get collection
	collection := session.DB(DB).C(collectionSummoner)

	// Get a Document from the collection
	result := []Application.Summoner{}
	err = collection.Find(bson.M{}).All(&result)
	if err != nil {
		log.Fatal("Error finding record: ", err)
		return
	}

	fmt.Println("Invocadores:", result)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(result)

	fmt.Println("END Obtain summoner from CosmosDB")
}

func updateSummoner(w http.ResponseWriter, r *http.Request) {
	fmt.Println("START UPDATE summoner from CosmosDB")

	params := mux.Vars(r)
	summonerName := params["nombreSummoner"]

	//Conectamos con Cosmos DB
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{AddresCosmosDB},
		Timeout:  60 * time.Second,
		Database: DB,
		Username: userName,
		Password: password,
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		},
	}

	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	session, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		fmt.Printf("No se pudo conectar con Azure Cosmos DB, go error %v\n", err)
		os.Exit(1)
	}

	defer session.Close()
	session.SetSafe(&mgo.Safe{})

	// get collection
	collection := session.DB(DB).C(collectionSummoner)

	updateQuery := bson.M{"accountid": "xe7bCB1uBLXcO_-d6pAHtyF9xpDnxM4dQnmOalVYuv-oXFIm7HhehaVw"}
	change := bson.M{"$set": bson.M{"name": summonerName}}

	err = collection.Update(updateQuery, change)
	if err != nil {
		log.Fatal("Error updating record: ", err)
		return
	}

	fmt.Println("END UPDATE summoner from CosmosDB")
	getAllSummonerFromCosmosDB(w, r)
}*/
