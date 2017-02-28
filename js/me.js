$("#ajax").click(fuction(){
	alert("ajax start");
    $.post("/ajax", function(data){
	  // $( "#ajax" ).html( data );
	  alert("ajax success");
	});
});
