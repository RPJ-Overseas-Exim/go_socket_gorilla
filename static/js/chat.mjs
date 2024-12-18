function connectToSocket() {
    const link = document.location
    const protocol = link.protocol
    const host = link.host

    let wsURL = "ws://"
    if(protocol === "https:"){
        wsURL = "wss://"
    }

    wsURL += host + "/ws?email=abc@gmail.com"

    const conn = new WebSocket(wsURL)
    conn.onopen = function(){
        console.log("socket connected")
    }

    conn.onclose = function(){
        console.log("socket disconnected")
    }
}

window.onload = function(){
    connectToSocket()
}
