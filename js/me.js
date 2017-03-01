function ajaxtest(){
  alert("ajax start");
    // $.post("/get_customs",
    // {
    //   name: "Donald Duck",
    //   city: "Duckburg"
    // },
    // function(data,status){
    //     alert("Data: " + data.address + "\nStatus: " + status);
    // });

    $.getJSON( "/get_customs", function( data ) {
	  var items = [];
	  $.each( data, function( key0, val0 ) {
	    // items.push( "<li id='" + key0 + "'>" + key0 + ":" + val0 + "</li>" );
	    $.each( val0, function( key1, val1 ) {
		  items.push( "<li id='" + key1 + "'>" + key1 + ":" + val1 + "</li>" );
		});
		items.push("<br/>")
	  });
	 
	  $( "<ul/>", {
	    "class": "my-new-list",
	    html: items.join( "" )
	  }).appendTo( "body" );
	});
}