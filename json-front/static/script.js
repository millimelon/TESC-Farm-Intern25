alert("test");
const endpoint = document.getElementById("endpoint");
endpoint.addEventListener("change", (e) => {
    let val = e.target.value;
    document.getElementById("response").innerText = val;
    alert(val);
});
function submitapi() {
    const val = endpoint.value;
    if (endpoint == "") {
		document.getElementById("response").innerText = "Invalid Endpoint";
        return;
    }
    const payload = document.getElementById("payload").value;
    alert(payload);
	fetch('localhost:8078/'+endpoint, {
        	method: 'POST',
        	headers: { 'Content-Type': 'application/json' },
        	body: JSON.stringify(payload)
	}).then(response => {
		if (!response.ok) {
			console.log(response);
			document.getElementById("response").innerText = response;
        	}
        	return response.json();
    	}).then(data => {
		document.getElementById("response").innerText = data;
	}
}

