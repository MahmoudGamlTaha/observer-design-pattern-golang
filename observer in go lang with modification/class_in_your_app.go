import (
	"fmt"

)


var handler = customObserver.NewSubjectHandler()

type ClassInYourApp struct {

}


}

func init(){

//fmt.Println("init run sheet")

 runsheet:= ClassInYourApp{}

 handler.AddObserver(runsheet)

}

func NewClassInYourApp() *ClassInYourApp {

	service := ClassInYourApp{}

	return &service

}


func (service ClassInYourApp)  GetId()string{

	return "runSheetID"

}

func (service ClassInYourApp)  functionx(){
param := "test"
	handler.AddParams(service.GetId(), "test" )

}

func (service ClassInYourApp) NotifyEvent( m interface{}){

   // fmt.Println("Event Run",m)

	//class1:= ClassInYourApp{}

 //   handler.RemoveObserver(class1)

}


func (service ClassInYourApp) ValidateData(data interface{}) interface{}{

	//data = "daa"

	//fmt.Println("data validate", data)

	return data

}
