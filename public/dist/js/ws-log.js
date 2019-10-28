var ws = new WebSocket('ws://localhost:8080/echo');
ws.onmessage = function (msg) {
    var li = createLIFromWSMessage(msg);
    var logOutlet = document.querySelector('.log-outlet')
    logOutlet.insertBefore(li, logOutlet.firstChild)
}

var wsConnectSetIntervalId = setInterval(() => {
    if (_isWSOpenAndReady()) {
        ws.send('flashscreen')
        clearInterval(wsConnectSetIntervalId);
        console.log('+')
    }
}, 1000);

function sendQuery(query) {
    var query = document.querySelector('.query-input').value;
    ws.send(query)
    setHighlightText(query);
}

var highlightText = "";
function createLIFromWSMessage(msg) {
    var li = document.createElement('li')
    li.style.listStyle = "none"
    li.style.fontSize = 'larger'
    li.style.background = "#c7c6c0"
    li.style.marginTop = '2px'
    li.style.paddingTop = '4px'
    li.style.paddingBottom = '4px'
    li.style.marginBottom = '2px'

    li.innerHTML = (msg.data.replace(highlightText, `<b style='color: red'>${highlightText}</b>`) || "");

    return li;
}

function _isWSOpenAndReady() {
    return 1 === ws.readyState;
}


function setHighlightText(query) {
    const infoIndex = query.indexOf('info=');
    if (-1 != infoIndex){
        window.highlightText = query.split('info=')[1]
    }
}