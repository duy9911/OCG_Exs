let buts = document.getElementsByClassName("button"); //buts = button
let datashow = document.getElementById("screen");  // datashows

let theFirstInput = true;
let url = "http://localhost:8000";

function DataToQuery(data) {
    let query = "";
    let op = "";
    let indexOfOp = 0;
    for (let i = 0; i < data.length; i++) {
        let a = data.charCodeAt(i);
        if (a < 48) {
            indexOfOp = i;
            break;
        }
    }
    query = "/"+data[indexOfOp] + "/" + data.slice(0, indexOfOp) + "/" + data.slice(indexOfOp + 1, data.length);
    return query;
    
}

function CallApi(query) {
    return fetch(url + query).then(data => console.log(data));
}

for (let i = 0; i < buts.length; i++) {
    if (buts[i].id == "C" || buts[i].id == "CE") {
        theFirstInput = true;
        buts[i].onclick = function () {
            datashow.innerHTML = "";
        };
    } else {
        if (buts[i].id == "=") {
            buts[i].onclick = function () {
                theFirstInput = true;
                console.log(url + ExpressionToQuery(datashow.innerHTML)); //print to console

                let r = CallApi(ExpressionToQuery(datashow.innerHTML));
                r.then((data) => {
                    console.log(data.Result);
                    datashow.innerHTML = data.Result;
                });
            };
        } else {
            buts[i].onclick = function () {
                if (theFirstInput) {
                    datashow.innerHTML = buts[i].id;
                    theFirstInput = false;
                } else {
                    datashow.innerHTML = datashow.innerHTML + buts[i].id;
                }
            };
        }
    }
}
