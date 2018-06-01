$("#logout").click(function(event){
    event.preventDefault();
    del_cookie("admin_id");
    window.location.href = "/login/index";
})

function del_cookie(name)
{
    document.cookie = name + '=; expires=Thu, 01 Jan 1970 00:00:01 GMT;path=/;';
}

$("form[data-type=formAction]").submit(function(event){
    event.preventDefault();
    var target = event.target;
    var action = $(target).attr("action");
    $.post(action, $(target).serialize(), function(ret){
        if(ret.Ret == "0") {
            alert(ret.Reason);
        } else {
            location.href = $(target).attr("form-rediret");
        }
    },"json")
})

$("input[name='ethcall']").click(function(event){
    event.preventDefault();
    var target = event.target;
    var r=$(this).parents('.row');
 	var funcname=r.find('[name=funcname]').text();
	var retObj=r.find('[name=funcret]');
    //alert(retObj.html());
  
					
    var action = "/eth/?id=0"+"&"+"func="+funcname+"&"+"param1=123"+"&"+"param2=abc";//$(target).attr("action");
	// alert(action);
    $.get(action, function(ret){
        if(ret.Ret == "0") {
			//alert(ret.Reason+':'+ret.Res);
			//$("input[name='funcret']").val(ret.Reason+':'+ret.Res);
			retObj.val(ret.Reason+':'+ret.Res);
            //alert(ret.Reason);
        } else {
			//$("input[name='funcret']").val(ret.Reason+':'+ret.Res);
			retObj.val(ret.Reason+':'+ret.Res);
			//alert(ret.Reason);
            //location.href = $(target).attr("form-rediret");
        }
    },"json")
})


$("input[name='ipfscall_notuse']").click(function(event){
    event.preventDefault();
    var target = event.target;
    var action = "/ipfs/?id=1"+"&"+"func=web3_clientVersion"+"&"+"param1=123"+"&"+"param2=abc";//$(target).attr("action");
	 alert(action);
	
    $.get(action, function(ret){
        if(ret.Ret == "0") {
			$("input[name='web3_clientVersion_ret']").val(ret.Reason+':'+ret.Res);
            //alert(ret.Reason);
        } else {
			$("input[name='web3_clientVersion_ret']").val(ret.Reason+':'+ret.Res);
			//alert(ret.Reason);
            //location.href = $(target).attr("form-rediret");
        }
    },"json")
})
