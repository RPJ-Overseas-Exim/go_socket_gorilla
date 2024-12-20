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

    conn.onmessage = function(msg){
        console.log(msg.data)
        let data

        try{
            data = JSON.parse(msg.data)
            const chat = document.querySelector(`[data-chatId=${data?.ChatId}]`)
            //console.log("for chat Id: ", chat)

            if(chat){
                const indicator = chat.querySelector(".online-indicator")
                //console.log(indicator)

                if(data?.event === "online"){
                    // make the indicator visible
                    console.log("online event")
                    indicator.classList.remove("hidden")
                }
                if(data?.event === "offline"){
                    // make the indicator invisible
                    console.log("offline event")
                    indicator.classList.add("hidden")
                }
                if(data?.event === "reload"){
                    // show the new message
                    const currentChat = document.getElementById("#message-chatId")
                    console.log("reload event: ", currentChat.dataset.chatId)
                }
            }

            return // exit 
        }catch(err){
            data = msg.data
        }

        appendMessage(data, false)
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
                appendMessage(dataObject.message, true)
            }

            messageInput.value = ""
        })
    }
}

function appendMessage(msg, admin){

    const messageBox = document.createElement("div")
    messageBox.className = `pb-4 px-4 flex ${admin ? "justify-end" : "justify-start"}`
    const message = document.createElement("div")
    message.className = `border border-border px-4 py-1 rounded-lg ${admin ? "rounded-br-[0px]" : "rounded-bl-[0px]"}`
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
            const messages = document.querySelector("#messages")

            sendMessage(conn)
            messages.scrollTo({top: messages.scrollHeight, behavior: "smooth"})
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
