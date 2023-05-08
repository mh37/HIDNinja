// Get the Sidebar
var mySidebar = document.getElementById("mySidebar");

// Get the DIV with overlay effect
var overlayBg = document.getElementById("myOverlay");

// Toggle between showing and hiding the sidebar, and add overlay effect
function w3_open() {
    if (mySidebar.style.display === 'block') {
    mySidebar.style.display = 'none';
    overlayBg.style.display = "none";
    } else {
    mySidebar.style.display = 'block';
    overlayBg.style.display = "block";
    }
}

// Close the sidebar with the close button
function w3_close() {
    mySidebar.style.display = "none";
    overlayBg.style.display = "none";
}

//WebSocket communication
let input = document.getElementById("cmdInput");
let output = document.getElementById("output");
let socket = new WebSocket("ws://localhost:8080/echo");

socket.onopen = () => {
    output.innerHTML += "Connection established\n";
};

socket.onmessage = (e) => {
    output.innerHTML += "Delivered: " + e.data + "\n";
};

socket.onclose = event => {
    output.innerHTML += "Connection Closed\n";
    socket.send("Client Closed!")
};

socket.onerror = error => {
    console.log("Socket Error: ", error);
    output.innerHTML += "Error " + error + "\n";
};


function send() {
    socket.send(input.value);
    input.value = "";
}