$("table").hide();

$("a:eq(1)").click(function(){
	$("table").hide();
	$("form:eq(1)").show();
});

$("a:eq(2)").click(function(){
	$("form:eq(1)").hide();
	$("table").show();
});

$("form:eq(0)").submit(function(){
    alert("search function is creating ... ");
});

$("form:eq(1)").submit(function(){
    $.post("/submit_input_data", function(data, status){
        alert("Data: " + data + "\nStatus: " + status);
    });
});