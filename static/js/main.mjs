import "./lib/htmx.mjs"
let conn,
    msg = document.querySelector(".msg"),
    output = document.querySelector(".output"),
    email = document.querySelector(".email")


const textInput = document.querySelector(".text-input")

if (textInput) {
    textInput.addEventListener("submit", function(e) {
        e.preventDefault()
        if (!conn && email) {
            connectToSocket()
        }
        if (!msg || !msg.value || !conn || !email) {
            return
        }
        conn.send(msg.value)
    })
}

function connectToSocket() {
    if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + document.location.host + "/ws?email=" + email.value)
        conn.onclose = () => {
            var item = document.createElement("div");
            item.textContent = "<b>Connection closed.</b>"
            appendLog(item);
        };

        conn.onmessage = function(evt) {
            var messages = evt.data;
            var item = document.createElement("div")
            item.textContent = "\n" + messages;
            appendLog(item);
        };
    } else {
        var item = document.createElement("div");
        item.textContent = "<b>Your browser does not support WebSockets.</b>";
        appendLog(item);
    }
}

function appendLog(item) {
    output.appendChild(item)
}
