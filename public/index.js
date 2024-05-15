function alertClicked(alert) {
    window.alert(`${alert.RouteNumber != "" ? (alert.RouteNumber > 500 ? "Station affected: " : "Route Affected: ") + alert.RouteNumber : ""}\n${alert.DateEffective}\nDescription: ${alert.Description}\n`);
}

function on_ready() {
    // first we need to call api :o
    $("body").hide();
    fetch("/api").then((body) => {
        body.json().then((json) => {
            if (json.length == 0) { 
                $("#kjsdf-1").hide();   
            } else {
                $("#kjsdf-2").hide();
                let div = document.createElement("div");
                div.className = "alerts row p-2";
                div.style = "margin-left: 2rem;margin-right: 2rem;";
                for (let i = 0; i < json.length; i++) {
                    let p = document.createElement("p");
                    p.className = "alert bg-white col-12";
                    p.style = "border: 1px solid black;";
                    p.onclick = () => {alertClicked(json[i])};
                    let tn = document.createTextNode(json[i].Title);
                    p.appendChild(tn);
                    div.appendChild(p);
                }
                document.body.appendChild(div);
            }
            /*
            <div class="alerts row p-2">
                {{ range $val := . }}
                    <p class="alert bg-info col-12">{{ $val }}</p>
                {{ end }}
            </div>
            */

            $("body").show();
        });
    });
}