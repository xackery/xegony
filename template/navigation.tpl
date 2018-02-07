<!-- Navigation -->
<aside id="menu">
	<div id="navigation">		
		<ul class="nav" id="side-menu">			
			{{template "navmenu" .}}
		</ul>
	</div>
</aside>

{{if gt .Site.User.ID 0}}
<script src="/vendor/sparkline/index.js"></script>
<script type="text/javascript">
    $("#expLine").sparkline([1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24], {
        type: 'bar',
        barWidth: 4,
        height: '30px',
        barColor: '#62cb31',
        negBarColor: '#53ac2a'
    });
</script>
{{end}}