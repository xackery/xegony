var path = require('path');
var pathToPhaser = path.join(__dirname, '/node_modules/phaser/');
var phaser = path.join(pathToPhaser, 'dist/phaser.js');
var webpack = require('webpack');

module.exports = {
  entry: './src/main.ts',
  output: {
    path: path.resolve(__dirname, '../../www/js/'),
    filename: 'main.js',
  },
  module: {
    rules: [
      { test: /\.ts$/, loader: 'ts-loader', exclude: '/node_modules/' }
    ]
  },
  devServer: {
    contentBase: path.resolve(__dirname, '../../www'),
    publicPath: '../../www/',
    host: '127.0.0.1',
    port: 8080,
    open: true
  },
  resolve: {
    extensions: ['.ts', '.js'],
    alias: {
      phaser: phaser
    }
  },
  plugins: [
    new webpack.DefinePlugin({
      BASE_URL: JSON.stringify("http://localhost:8083/v1/"),
      IS_DEBUG: JSON.stringify(true)
    })
  ]
};