package customObserver


import (

	//"fmt"

	"sync"

	"time"

)

type Observer interface {

    GetId() string

    ValidateData(interface{}) interface{}

    NotifyEvent( m interface{}) // handel event

}

type subjectHandler struct {

	 RegisteredObserved  map[string][]Observer // map[string][]Observer

	 ObserverData       map[string] interface{}

	// RegisteredObserved [] chan *Observer

	 mutex *sync.Mutex

}

func init(){

	// init observer

}

var observeHandler * subjectHandler

func NewSubjectHandler() *subjectHandler {

	if observeHandler == nil{

		observeHandler = new(subjectHandler)

		observeHandler.RegisteredObserved = make(map[string][] Observer)

		observeHandler.ObserverData = make(map[string] interface{})

		observeHandler.mutex = &sync.Mutex{}

	}

	return observeHandler

}

func NewObserver() *Observer{

	return new(Observer)

}

func(handler* subjectHandler) AddObserver(m Observer){

	handler.mutex.Lock()

	// handler.RegisteredObserved = append(handler.RegisteredObserved, m)

	handler.RegisteredObserved[(m).GetId()] =  append(handler.RegisteredObserved[(m).GetId()], m)

	//handler.Notify()

    defer handler.mutex.Unlock()

}

func (handler *subjectHandler) RemoveObserver(m Observer) {

	handler.mutex.Lock()

	defer handler.mutex.Unlock()

	delete(handler.RegisteredObserved, (m).GetId())

}


func (handler *subjectHandler) Notify(){

	for {

		for _, observers := range handler.RegisteredObserved {

		//	fmt.Println("from loop", v.GetId())

			for _, observer := range observers {

				(observer).NotifyEvent((observer).ValidateData(handler.ObserverData[(observer.GetId())]))	

			}

		}

		time.Sleep(5 * time.Second)

	}

}


func (handler *subjectHandler) NotifyOne(key string){

			//	fmt.Println("from loop", v.GetId())

	        observers := handler.RegisteredObserved[key]

	        for _,v := range observers {

				v.ValidateData(v.ValidateData(handler.ObserverData[v.GetId()]))

			}

}

func (handler *subjectHandler) StaticNotifyForAllKeysObservers(m interface{}) { // same event for all

	for {

		for _, observers := range handler.RegisteredObserved {

			//	fmt.Println("from loop", v.GetId())

			for _, observer := range observers {

				(observer).NotifyEvent((observer).ValidateData(m))

			}

		}

		time.Sleep(5 * time.Second)

	}

}

func (handler *subjectHandler) AddParams(id string,m interface{}){

	handler.ObserverData[id] = m

}
