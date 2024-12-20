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

    return conn
}

function sendMessage(conn){
    const messageForm = document.querySelector(".message-form")
    const messageInput = document.querySelector(".message-input")

    if(messageForm){
        messageForm.addEventListener("submit", async (e)=>{
            e.preventDefault()
            const data = new FormData(e.target)
            const dataObject = Object.fromEntries(data.entries())
            
            if(dataObject?.message.length > 0){
                conn.send(dataObject.message)
                appendMessage(dataObject.message)
            }

            messageInput.value = ""
        })
    }
}

function appendMessage(msg){

    const messageBox = document.createElement("div")
    messageBox.className = "pb-4 px-4 flex justify-end"
    const message = document.createElement("div")
    message.className = "border border-border px-4 py-1 rounded-lg rounded-br-[0px]"
    message.innerText = msg
    messageBox.appendChild(message)

    const messages = document.querySelector("#messages")
    if(messages){
        messages.appendChild(messageBox)
    }

}

if( location.pathname === "/admin"){
    const conn = connectToSocket()

    document.addEventListener("htmx:afterRequest", function (e){
        if(e.detail.target.id == "chat-messages"){
            sendMessage(conn)
        }
    })
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
