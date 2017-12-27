{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">


<div class="login-container">
	<div class="row">
		<div class="col-md-12">
			<div class="text-center m-b-md"><h3>LOGIN</h3></div>
			<div class="hpanel">
				<div class="panel-body">
						<form method=POST action="/login" id="loginForm">
							<div class="form-group">
								<label class="control-label" for="Name">Name</label>
								<input type="text" placeholder="name" title="Please enter you username" required="" value="" name="name" id="name" class="form-control">
								<span id="nameHelp" class="help-block small">Name</span>
							</div>
							<div class="form-group">
								<label class="control-label" for="password">Password</label>
								<input type="password" title="Please enter your password" placeholder="******" required="" value="" name="password" id="password" class="form-control">
								<span id="passwordHelp" class="help-block small">A strong password</span>
							</div>
							<div class="alert alert-danger" id="loginAlert" style="display:none">
							  Info
							</div>
							<button class="btn btn-success btn-block">Login</button>
							<a class="btn btn-default btn-block" href="/login">Register</a>
						</form>
				</div>
			</div>
		</div>
	</div>
</div>
</div>


<script type="text/javascript">
$("#loginForm").submit(function(event) {

	$("#loginAlert").hide()
	var request = $.ajax({
		url: "/api/login",
		type: "POST",
		data: JSON.stringify({
			name: $("#name").val(),
			password: $("#password").val()
		}),
		dataType: "JSON",
		success: function(data){
			console.log(data);
			document.cookie = "apiKey="+data.ApiKey;
			window.location = "/"
	    },
	    error: function(data){
	    	var resp = data.responseJSON;
	    	
	    	for (var key in resp.fields) {
	    		$("#"+key+"Help").text(resp.fields[key]);
	    	}
	    	console.log(resp)
	    	$("#loginAlert").show()
	    	$("#loginAlert").text(resp.message);	        
	    },
	    processData: false,
	});

	event.preventDefault();
});
</script>