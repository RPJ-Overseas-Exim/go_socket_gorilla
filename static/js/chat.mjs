function connectToSocket() {
    const protocol = location.protocol
    const host = location.host

    let wsURL = "ws://"
    if(protocol === "https:"){
        wsURL = "wss://"
    }

    wsURL += host + "/admin/ws"

    const conn = new WebSocket(wsURL)
    conn.onopen = function(){
        console.log("socket connected")
    }

    conn.onclose = function(){
        console.log("socket disconnected")
    }
}

if( location.pathname === "/admin"){
    connectToSocket()
}

//const chatBtns = document.querySelectorAll(".chat-button")
//
//if(chatBtns?.length>0){
//    chatBtns.forEach(chatBtn=>{
//        chatBtn.addEventListener("click", async ()=>{
//            const chatId = chatBtn.getAttribute("data-chatId")
//            const res = await fetch(getBaseURL() + "/admin/joinChat/" + chatId)
//            console.log(res)
//        })
//    })
//
//}
