<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta property="og:title" content="{{.Site.Title}}">
	<meta property="og:type" content="website">
	<meta itemprop="name" content="{{.Site.Name}}">
	<meta itemprop="description" content="{{.Site.Description}}">
	{{/*<meta itemprop="og:url" content="{{.Site.Url}}">*/}}
	<meta itemprop="og:image" content="{{.Site.Image}}">
	<meta itemprop="og:site_name" content="{{.Site.Name}}">
	<meta property="og:description" content="{{.Site.Description}}">
	<meta name="description" content="{{.Site.Description}}">
	<meta name="author" content="{{.Site.Author}}">

	<!-- Page title -->
	<title>{{.Site.Title}}</title>

	<link rel="apple-touch-icon" sizes="180x180" href="/apple-touch-icon.png">
	<link rel="icon" type="image/png" sizes="32x32" href="/favicon-32x32.png">
	<link rel="icon" type="image/png" sizes="16x16" href="/favicon-16x16.png">
	<link rel="manifest" href="/manifest.json">
	<link rel="mask-icon" href="/safari-pinned-tab.svg" color="#5bbad5">
	<meta name="theme-color" content="#ffffff">

	<!-- Vendor styles -->
	<link rel="stylesheet" href="/vendor/fontawesome/css/font-awesome.css" />
	<link rel="stylesheet" href="/vendor/metisMenu/dist/metisMenu.css" />
	<link rel="stylesheet" href="/vendor/animate.css/animate.css" />
	<link rel="stylesheet" href="/vendor/bootstrap/dist/css/bootstrap.css" />
	<link rel="stylesheet" href="/vendor/xeditable/bootstrap3-editable/css/bootstrap-editable.css">
	<!-- App styles -->
	<link rel="stylesheet" href="/fonts/pe-icon-7-stroke/css/pe-icon-7-stroke.css" />
	<link rel="stylesheet" href="/fonts/pe-icon-7-stroke/css/helper.css" />
	<link rel="stylesheet" href="/fonts/xegony/css/xegony-awesome.min.css" />
	
	<link rel="stylesheet" href="/styles/builds.css" />
	<link rel="stylesheet" href="/styles/inventory.css" />
	<link rel="stylesheet" href="/styles/icons.css" />
	<link rel="stylesheet" href="/styles/style.css" />
	
	<link rel="stylesheet" href="/vendor/summernote/dist/summernote.css" />
	<link rel="stylesheet" href="/vendor/summernote/dist/summernote-bs3.css" />
	
	<link rel="stylesheet" href="/vendor/datatables.net-bs/css/dataTables.bootstrap.min.css">
	<link rel="stylesheet" href="/vendor/select2-3.5.2/select2.css">
	<link rel="stylesheet" href="vendor/select2-bootstrap/select2-bootstrap.css">
	

	<script src="/vendor/jquery/dist/jquery.min.js"></script>
	<script src="/vendor/jquery-ui/jquery-ui.min.js"></script>

	<script src="/vendor/datatables/media/js/jquery.dataTables.min.js"></script>
	<script src="/vendor/datatables.net-bs/js/dataTables.bootstrap.min.js"></script>
	<script src="//cdn.jsdelivr.net/velocity/1.5/velocity.min.js"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/velocity/1.2.3/velocity.ui.min.js"></script>
	<script src="vendor/select2-3.5.2/select2.min.js"></script>
	

</head>
<body class="fixed-navbar sidebar-scroll">
	{{ template "body" .}}
<!--[if lt IE 7]>
<p class="alert alert-danger">You are using an <strong>outdated</strong> browser. Please <a href="http://browsehappy.com/">upgrade your browser</a> to improve your experience.</p>
<![endif]-->




<!-- Vendor scripts -->

<script src="/vendor/slimScroll/jquery.slimscroll.min.js"></script>
<script src="/vendor/bootstrap/dist/js/bootstrap.min.js"></script>
<script src="/vendor/metisMenu/dist/metisMenu.min.js"></script>
<script src="/vendor/iCheck/icheck.min.js"></script>
<script src="/vendor/xeditable/bootstrap3-editable/js/bootstrap-editable.min.js"></script>
<!-- App scripts -->
<script src="/scripts/homer.js"></script>


<script>



	$(function () {

		$("[title]").tooltip();

		$("[item]").mouseenter(function(event, ui) {
			if ($(this).attr("title") == undefined) {
				console.log("do");
				$.ajax({
					url: "/api/item/"+$(this).attr("item")+"/tooltip",
					type: "GET",                
					dataType: "JSON",
					success: function(data){
						$("[item="+data.id+"]").attr("data-html", true);
						$("[item="+data.id+"]").attr("title", data.content);                    
						$("[item="+data.id+"]").tooltip({
							'selector': '',
							'placement': 'top',
							'container': 'body',
						});
					},
					error: function(data){
						var resp = data.responseJSON;
						console.log(resp);
					},
					processData: false,
				});
			}       
		});
		

		$('#fixed-navbar').click(function(){
			if($('body').hasClass('fixed-navbar')) {
				$('body').removeClass('fixed-navbar');
				$('.status-fixed-navbar').html("Off");

				if($('body').hasClass('sidebar-scroll')) {
					$('body').removeClass('sidebar-scroll');
					$('#navigation').slimScroll({destroy: true});
					$('#navigation').attr('style', '');
					$('.status-fixed-sidebar').html("Off")
				}

				if($('body').hasClass('fixed-small-header')) {
					$('body').removeClass('fixed-small-header');
					$('.status-fixed-small-header').html("Off")
				}

			} else {
				$('body').addClass('fixed-navbar');
				$('.status-fixed-navbar').html("<span class='text-success font-bold'>On</span>");
				$('body').removeClass('boxed');
				$('.status-boxed-layout').html("Off")
			}
		});

		$('#fixed-sidebar').click(function(){

			if($('body').hasClass('sidebar-scroll')) {
				$('body').removeClass('sidebar-scroll');
				$('#navigation').slimScroll({destroy: true});
				$('#navigation').attr('style', '');
				$('.status-fixed-sidebar').html("Off")
			} else {
				$('body').addClass('fixed-navbar');
				$('body').addClass('sidebar-scroll');
				$('#navigation').slimScroll({
					height: '100%',
					opacity: 0.3,
					size : 0,
					wheelStep : 10
				});
				$('.status-fixed-navbar').html("<span class='text-success font-bold'>On</span>");
				$('.status-fixed-sidebar').html("<span class='text-success font-bold'>On</span>");
				$('body').removeClass('boxed');
				$('.status-boxed-layout').html("Off")
			}
		});

		$('#fixed-footer').click(function(){
			if($('body').hasClass('fixed-footer')) {
				$('body').removeClass('fixed-footer');
				$('.status-fixed-footer').html("Off");
			} else {
				$('body').addClass('fixed-footer');
				$('.status-fixed-footer').html("<span class='text-success font-bold'>On</span>");
				$('body').removeClass('boxed');
				$('.status-boxed-layout').html("Off")
			}
		});

		$('#fixed-small-header').click(function(){

			if($('body').hasClass('fixed-small-header')) {
				$('body').removeClass('fixed-small-header');
				$('body').removeClass('sidebar-scroll');
				$('#navigation').slimScroll({destroy: true});
				$('#navigation').attr('style', '');
				$('body').removeClass('fixed-navbar');
				$('.status-fixed-small-header').html("Off");
				$('.status-fixed-sidebar').html("Off");
				$('.status-fixed-navbar').html("Off")
			} else {
				$('body').addClass('fixed-navbar');
				$('body').addClass('sidebar-scroll');
				$('#navigation').slimScroll({
					height: '100%',
					opacity: 0.3,
					size : 0,
					wheelStep : 10
				});
				$('body').addClass('fixed-small-header');
				$('.status-fixed-navbar').html("<span class='text-success font-bold'>On</span>");
				$('.status-fixed-sidebar').html("<span class='text-success font-bold'>On</span>");
				$('.status-fixed-small-header').html("<span class='text-success font-bold'>On</span>");
				$('body').removeClass('boxed');
				$('.status-boxed-layout').html("Off")
			}
		});

		$('#boxed-layout').click(function(){
			if($('body').hasClass('boxed')) {
				$('body').removeClass('boxed');
				$('.status-boxed-layout').html("Off")
			} else {
				$('body').addClass('boxed');
				$('.status-boxed-layout').html("<span class='text-success font-bold'>On</span>");
				$('body').removeClass('fixed-small-header');
				$('body').removeClass('sidebar-scroll');
				$('#navigation').slimScroll({destroy: true});
				$('#navigation').attr('style', '');
				$('body').removeClass('fixed-navbar');
				$('body').removeClass('fixed-footer');
				$('.status-fixed-small-header').html("Off");
				$('.status-fixed-footer').html("Off");
				$('.status-fixed-sidebar').html("Off");
				$('.status-fixed-navbar').html("Off");
			}
		});

	});

</script>
</body>
</html>