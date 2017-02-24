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

// $("#inputdata").submit(function(){
//         alert($("#inputid").val());
// });