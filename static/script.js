function convertToJson(){
    event.preventDefault();
    const name = document.forms["myForm"]["name"].value;
    const password = document.forms["myForm"]["psw"].value;
    console.log(name);
    var myHeader = new Headers();
    myHeader.append("content-type","application/json");
    var myInit = {
        method: "POST",
        headers: myHeader,
        body: JSON.stringify({name:name,password:password})
    };
    const req = new Request(".",myInit);
    fetch(req,myInit).then((response)=>{
        const text =response.text().then((str)=>{document.write(str)})
        return
    })
}
//   const xhttp = new XMLHttpRequest()
//   xhttp.open("POST",'/login',false)
//   xhttp.setRequestHeader("content-type","application/json")
//   xhttp.send(JSON.stringify({name:name,password:password}))