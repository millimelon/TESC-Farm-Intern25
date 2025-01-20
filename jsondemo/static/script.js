const endpoint = document.getElementById("endpoint");
endpoint.addEventListener("change", (e) => {
    let val = e.target.value;
    document.getElementById("json").innerText = val;
});
const url = document.getElementById("url").value;
function submitapi() {
    const val = endpoint.value;
    if (endpoint == "") {
		document.getElementById("response").innerText = "Invalid Endpoint";
        return;
    }
    const payload = document.getElementById("json").value;
	fetch(url+'8078/'+endpoint.value, {
        	method: 'POST',
            mode: 'no-cors',
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
	});
}
