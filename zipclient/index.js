$( document ).ready(function() {
    console.log("ready");
    $("#submit").click(function() {
        var cityName = ($("#name").val());
        var stateName = ($("#state").val());
        var theUrl = "http://localhost:4000/zips/" + cityName;

        response = $.ajax({
            url: theUrl,
            success: (response) => {
                myFunction(response, stateName);
              }
        });

        
    });
    
    function myFunction(response, stateName) {
        console.log(response);
        // var header = document.createElement("h1");
        
        var zipCodeDiv = document.createElement('div');
        response.forEach((element) => {
            var zip = document.createElement("p");
            if (element.State == stateName) {
                zip.innerText = element.Code;
                zipCodeDiv.appendChild(zip);
            }
        })
        // header.textContent = response;
        var elemDiv = document.createElement('div');
        elemDiv.className += "container";
        elemDiv.appendChild(zipCodeDiv);
        document.body.appendChild(elemDiv);
    }

    
});


