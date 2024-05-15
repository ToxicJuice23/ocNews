function alertClicked(alert) {
    window.alert(`${alert.RouteNumber != "" ? "Route affected: " + alert.RouteNumber : ""}\n${alert.DateEffective}\nDescription: ${alert.Description}\n`);
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
                for (let i = 0; i < json.length; i++) {
                    let p = document.createElement("p");
                    p.className = "alert bg-info col-12";
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