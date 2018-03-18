// ProjectName:: IIITK-VizB-Services ver1.0
//
// ProjectTitle:: Vizualization services for big datasets
//
// ProjectDescription:: See doc folder for more information.
//
// Authors:: i) ETHAKOTA PAVANSAI
//          ii) Shajulin Benedict
//          iii) VANKUDOTH PREMKUMAR
//          iv) SUMALA SAI SIDDARDHA


package bigdata

import (
  "fmt"
  "net/http"
  "github.com/senseyeio/roger"
)

/// bigdata info here.
var Bigdata = `
<!DOCTYPE html>
<html>
<head>
  <title>BigdataInfo</title>
  <style>
    body {background-color: powderblue;}
    h1 {color: red;}
    p {color: blue;}
  </style>
</head>
<body>
<h1>Bigdata Plot for IIITK-VizB-Services!!</h1>
<img src="http://www.clipartsmania.com/gif/star/animation-red-star.GIF"
alt="http://www.clipartsmania.com/gif/star/animation-red-star.GIF" style="width:48px;height:48px;">
<p> Step 1: Collects data for plots <p>
<p> Step 4: Display the graph using service <p>
</body>
</html>
`

/// bigdata service - bigdata service for IIITK-VizB-Services
func Bigdatacompscatter(w http.ResponseWriter, req *http.Request){

   fmt.Fprintf(w, Bigdata)

   //Pursue with the steps
   // connect to RServe using Roger
   rClient, err := roger.NewRClient("127.0.0.1", 6311)
   if err != nil {
   fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
   }
   fmt.Printf("Connected to RServe")

   // call bigdata R function, gathering the response
   returnVar, err := rClient.Eval("bgCompScatter()")
   if err != nil {
     fmt.Printf("bigdata is not completed %s and %s", err.Error(), returnVar)
     return
   }
   //Similarly do it for the other bigdata plots
   //..TODO

}

func Bigdatadistbox(w http.ResponseWriter, req *http.Request){

   fmt.Fprintf(w, Bigdata)

   //Pursue with the steps
   // connect to RServe using Roger
   rClient, err := roger.NewRClient("127.0.0.1", 6311)
   if err != nil {
   fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
   }
   fmt.Printf("Connected to RServe")

   // call bigdata R function, gathering the response
   returnVar, err := rClient.Eval("bgDistBox()")
   if err != nil {
     fmt.Printf("bigdata is not completed %s and %s", err.Error(), returnVar)
     return
   }
   //Similarly do it for the other bigdata plots
   //..TODO

}

func Bigdatacomparea(w http.ResponseWriter, req *http.Request){

   fmt.Fprintf(w, Bigdata)

   //Pursue with the steps
   // connect to RServe using Roger
   rClient, err := roger.NewRClient("127.0.0.1", 6311)
   if err != nil {
   fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
   }
   fmt.Printf("Connected to RServe")

   // call bigdata R function, gathering the response
   returnVar, err := rClient.Eval("bgCompArea()")
   if err != nil {
     fmt.Printf("bigdata is not completed %s and %s", err.Error(), returnVar)
     return
   }
   //Similarly do it for the other bigdata plots
   //..TODO

}

func Bigdatadistdotbox(w http.ResponseWriter, req *http.Request){

   fmt.Fprintf(w, Bigdata)

   //Pursue with the steps
   // connect to RServe using Roger
   rClient, err := roger.NewRClient("127.0.0.1", 6311)
   if err != nil {
   fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
   }
   fmt.Printf("Connected to RServe")

   // call bigdata R function, gathering the response
   returnVar, err := rClient.Eval("bgDistDotBox()")
   if err != nil {
     fmt.Printf("bigdata is not completed %s and %s", err.Error(), returnVar)
     return
   }
   //Similarly do it for the other bigdata plots
   //..TODO

}

func Bigdatadistkernaldensity(w http.ResponseWriter, req *http.Request){

   fmt.Fprintf(w, Bigdata)

   //Pursue with the steps
   // connect to RServe using Roger
   rClient, err := roger.NewRClient("127.0.0.1", 6311)
   if err != nil {
   fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
   }
   fmt.Printf("Connected to RServe")

   // call bigdata R function, gathering the response
   returnVar, err := rClient.Eval("bgDistKernalDensity()")
   if err != nil {
     fmt.Printf("bigdata is not completed %s and %s", err.Error(), returnVar)
     return
   }
   //Similarly do it for the other bigdata plots
   //..TODO

}

func Bigdatadistnotchedbox(w http.ResponseWriter, req *http.Request){

   fmt.Fprintf(w, Bigdata)

   //Pursue with the steps
   // connect to RServe using Roger
   rClient, err := roger.NewRClient("127.0.0.1", 6311)
   if err != nil {
   fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
   }
   fmt.Printf("Connected to RServe")

   // call bigdata R function, gathering the response
   returnVar, err := rClient.Eval("bgDistNotchedBox()")
   if err != nil {
     fmt.Printf("bigdata is not completed %s and %s", err.Error(), returnVar)
     return
   }
   //Similarly do it for the other bigdata plots
   //..TODO

}

func Bigdatadisttuftebox(w http.ResponseWriter, req *http.Request){

   fmt.Fprintf(w, Bigdata)

   //Pursue with the steps
   // connect to RServe using Roger
   rClient, err := roger.NewRClient("127.0.0.1", 6311)
   if err != nil {
   fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
   }
   fmt.Printf("Connected to RServe")

   // call bigdata R function, gathering the response
   returnVar, err := rClient.Eval("bgDistTufteBox()")
   if err != nil {
     fmt.Printf("bigdata is not completed %s and %s", err.Error(), returnVar)
     return
   }
   //Similarly do it for the other bigdata plots
   //..TODO

}

func Bigdatadistviolin(w http.ResponseWriter, req *http.Request){

   fmt.Fprintf(w, Bigdata)

   //Pursue with the steps
   // connect to RServe using Roger
   rClient, err := roger.NewRClient("127.0.0.1", 6311)
   if err != nil {
   fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
   }
   fmt.Printf("Connected to RServe")

   // call bigdata R function, gathering the response
   returnVar, err := rClient.Eval("bgDistViolin()")
   if err != nil {
     fmt.Printf("bigdata is not completed %s and %s", err.Error(), returnVar)
     return
   }
   //Similarly do it for the other bigdata plots
   //..TODO

}
func Bigdatarelbubble(w http.ResponseWriter, req *http.Request){

   fmt.Fprintf(w, Bigdata)

   //Pursue with the steps
   // connect to RServe using Roger
   rClient, err := roger.NewRClient("127.0.0.1", 6311)
   if err != nil {
   fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
   }
   fmt.Printf("Connected to RServe")

   // call bigdata R function, gathering the response
   returnVar, err := rClient.Eval("bgRelBubble()")
   if err != nil {
     fmt.Printf("bigdata is not completed %s and %s", err.Error(), returnVar)
     return
   }
   //Similarly do it for the other bigdata plots
   //..TODO

}

func Bigdatareltable(w http.ResponseWriter, req *http.Request){

   fmt.Fprintf(w, Bigdata)

   //Pursue with the steps
   // connect to RServe using Roger
   rClient, err := roger.NewRClient("127.0.0.1", 6311)
   if err != nil {
   fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
   }
   fmt.Printf("Connected to RServe")

   // call bigdata R function, gathering the response
   returnVar, err := rClient.Eval("bgRelTable()")
   if err != nil {
     fmt.Printf("bigdata is not completed %s and %s", err.Error(), returnVar)
     return
   }
   //Similarly do it for the other bigdata plots
   //..TODO

}

func Bigdatarelcount(w http.ResponseWriter, req *http.Request){

   fmt.Fprintf(w, Bigdata)

   //Pursue with the steps
   // connect to RServe using Roger
   rClient, err := roger.NewRClient("127.0.0.1", 6311)
   if err != nil {
   fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
   }
   fmt.Printf("Connected to RServe")

   // call bigdata R function, gathering the response
   returnVar, err := rClient.Eval("bgRelCount()")
   if err != nil {
     fmt.Printf("bigdata is not completed %s and %s", err.Error(), returnVar)
     return
   }
   //Similarly do it for the other bigdata plots
   //..TODO

}

func Bigdatareljitter(w http.ResponseWriter, req *http.Request){

   fmt.Fprintf(w, Bigdata)

   //Pursue with the steps
   // connect to RServe using Roger
   rClient, err := roger.NewRClient("127.0.0.1", 6311)
   if err != nil {
   fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
   }
   fmt.Printf("Connected to RServe")

   // call bigdata R function, gathering the response
   returnVar, err := rClient.Eval("bgRelJitter()")
   if err != nil {
     fmt.Printf("bigdata is not completed %s and %s", err.Error(), returnVar)
     return
   }
   //Similarly do it for the other bigdata plots
   //..TODO

}

func Bigdatarelcorrelogram(w http.ResponseWriter, req *http.Request){

   fmt.Fprintf(w, Bigdata)

   //Pursue with the steps
   // connect to RServe using Roger
   rClient, err := roger.NewRClient("127.0.0.1", 6311)
   if err != nil {
   fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
   }
   fmt.Printf("Connected to RServe")

   // call bigdata R function, gathering the response
   returnVar, err := rClient.Eval("bgRelCorrelogram()")
   if err != nil {
     fmt.Printf("bigdata is not completed %s and %s", err.Error(), returnVar)
     return
   }
   //Similarly do it for the other bigdata plots
   //..TODO

}

func Bigdatarelmarginalhistogram(w http.ResponseWriter, req *http.Request){

   fmt.Fprintf(w, Bigdata)

   //Pursue with the steps
   // connect to RServe using Roger
   rClient, err := roger.NewRClient("127.0.0.1", 6311)
   if err != nil {
   fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
   }
   fmt.Printf("Connected to RServe")

   // call bigdata R function, gathering the response
   returnVar, err := rClient.Eval("bgRelMarginalHistogram()")
   if err != nil {
     fmt.Printf("bigdata is not completed %s and %s", err.Error(), returnVar)
     return
   }
   //Similarly do it for the other bigdata plots
   //..TODO

}

func Bigdatarelscatterencircling(w http.ResponseWriter, req *http.Request){

   fmt.Fprintf(w, Bigdata)

   //Pursue with the steps
   // connect to RServe using Roger
   rClient, err := roger.NewRClient("127.0.0.1", 6311)
   if err != nil {
   fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
   }
   fmt.Printf("Connected to RServe")

   // call bigdata R function, gathering the response
   returnVar, err := rClient.Eval("bgRelScatterEncircling()")
   if err != nil {
     fmt.Printf("bigdata is not completed %s and %s", err.Error(), returnVar)
     return
   }
   //Similarly do it for the other bigdata plots
   //..TODO

}

func Bigdatacomppie(w http.ResponseWriter, req *http.Request){

   fmt.Fprintf(w, Bigdata)

   //Pursue with the steps
   // connect to RServe using Roger
   rClient, err := roger.NewRClient("127.0.0.1", 6311)
   if err != nil {
   fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
   }
   fmt.Printf("Connected to RServe")

   // call bigdata R function, gathering the response
   returnVar, err := rClient.Eval("bgCompPie()")
   if err != nil {
     fmt.Printf("bigdata is not completed %s and %s", err.Error(), returnVar)
     return
   }
   //Similarly do it for the other bigdata plots
   //..TODO

}

func Bigdatacompstackedarea(w http.ResponseWriter, req *http.Request){

   fmt.Fprintf(w, Bigdata)

   //Pursue with the steps
   // connect to RServe using Roger
   rClient, err := roger.NewRClient("127.0.0.1", 6311)
   if err != nil {
   fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
   }
   fmt.Printf("Connected to RServe")

   // call bigdata R function, gathering the response
   returnVar, err := rClient.Eval("bgCompStackedArea()")
   if err != nil {
     fmt.Printf("bigdata is not completed %s and %s", err.Error(), returnVar)
     return
   }
   //Similarly do it for the other bigdata plots
   //..TODO

}

func Bigdatacompwaffle(w http.ResponseWriter, req *http.Request){

   fmt.Fprintf(w, Bigdata)

   //Pursue with the steps
   // connect to RServe using Roger
   rClient, err := roger.NewRClient("127.0.0.1", 6311)
   if err != nil {
   fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
   }
   fmt.Printf("Connected to RServe")

   // call bigdata R function, gathering the response
   returnVar, err := rClient.Eval("bgCompWaffle()")
   if err != nil {
     fmt.Printf("bigdata is not completed %s and %s", err.Error(), returnVar)
     return
   }
   //Similarly do it for the other bigdata plots
   //..TODO

}

func Bigdatacmphistogram(w http.ResponseWriter, req *http.Request){

   fmt.Fprintf(w, Bigdata)

   //Pursue with the steps
   // connect to RServe using Roger
   rClient, err := roger.NewRClient("127.0.0.1", 6311)
   if err != nil {
   fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
   }
   fmt.Printf("Connected to RServe")

   // call bigdata R function, gathering the response
   returnVar, err := rClient.Eval("bgCmpHistogram()")
   if err != nil {
     fmt.Printf("bigdata is not completed %s and %s", err.Error(), returnVar)
     return
   }
   //Similarly do it for the other bigdata plots
   //..TODO

}

func Bigdatacmpbdbar(w http.ResponseWriter, req *http.Request){

   fmt.Fprintf(w, Bigdata)

   //Pursue with the steps
   // connect to RServe using Roger
   rClient, err := roger.NewRClient("127.0.0.1", 6311)
   if err != nil {
   fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
   }
   fmt.Printf("Connected to RServe")

   // call bigdata R function, gathering the response
   returnVar, err := rClient.Eval("bgCmpBDBar()")
   if err != nil {
     fmt.Printf("bigdata is not completed %s and %s", err.Error(), returnVar)
     return
   }
   //Similarly do it for the other bigdata plots
   //..TODO

}

func Bigdatacmpheat(w http.ResponseWriter, req *http.Request){

   fmt.Fprintf(w, Bigdata)

   //Pursue with the steps
   // connect to RServe using Roger
   rClient, err := roger.NewRClient("127.0.0.1", 6311)
   if err != nil {
   fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
   }
   fmt.Printf("Connected to RServe")

   // call bigdata R function, gathering the response
   returnVar, err := rClient.Eval("bgCmpHeat()")
   if err != nil {
     fmt.Printf("bigdata is not completed %s and %s", err.Error(), returnVar)
     return
   }
   //Similarly do it for the other bigdata plots
   //..TODO

}
