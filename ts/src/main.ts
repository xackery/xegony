import * as $ from 'jquery';
import * as pb from './pb/index';

var xegonyAPI: pb.XegonyApi;

function main() {
	xegonyAPI = new pb.XegonyApi(BASE_URL);
}

$(document).ready(function () {
	main();
});