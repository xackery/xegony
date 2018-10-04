var path = require('path');
var config = require('../../config/');

module.exports.task = function(gulp, plugins, paths) {

	gulp.src(paths.app.styles)
		.pipe(plugins.concat('app.scss'))
		.pipe(
			plugins.sass({
				includePaths: [
					path.resolve( config.srcDir ),
					path.resolve( config.npmDir ),
				]
			})
			.on('error', plugins.sass.logError)
		)
		.pipe(plugins.autoprefixer())
		.pipe(gulp.dest(config.destDir + '/css'))
		.pipe(plugins.connect.reload()).on( 'error', plugins.sass.logError );
};
