package main; import ( "fmt"; "html"; "lo
"net/http"; "strconv"; "strings"; "time" 
trolMessage struct { Target string; Count
func main() { controlChannel := make(chan
sage);workerCompleteChan := make(chan boo
PollChannel := make(chan chan bool); work
false;go admin(controlChannel, statusPoll
{ select { case respChan := <- statusPoll
Chan <- workerActive; case msg := <-contr
workerActive = true; go doStuff(msg, work
teChan); case status := <- workerComplete
Active = ststus; }}}; func admin(cc chan 
sage, statusPollChannel chan chan bool) {
Func("/admin", func(w http.ResponseWriter
quest) { hostTokens := strings.Split(r.Ho
r.ParseForm(); count, err := strconv.Pars
Value("count"), 10, 32); if err != nil { 
err.Error()); return; }; msg := ControlMe
r.FormValue("target"), Count: count}; cc 
printf(w, "Control message issued for Tar
%d", html.EscapeString(r.FormValue("targe
}); http.HandleFunc("/status",func(w http
er, r *http.Request) { reqChan := make(ch
tusPollChannel <- reqChan;timeout := time
time.Second); select { case result := <- 
result { fmt.Fprint(w, "ACTIVE"); } else 
print(w, "INACTIVE"); }; retuPEACE FOR ALL