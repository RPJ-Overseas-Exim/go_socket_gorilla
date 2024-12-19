export function connectToSocketUser() {
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

export function getBaseURL(){
    return location.protocol + location.host
}
