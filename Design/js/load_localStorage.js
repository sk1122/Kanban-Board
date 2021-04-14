
// setInterval(function() {
// 	let todo = localStorage.getItem("todo")
// 	let doing = localStorage.getItem("doing")
// 	let done = localStorage.getItem("done")
// 	let trash = localStorage.getItem("trash")

// 	postData(todo, doing, done, trash)
// }, 10000)

function requestData() {
	$.ajax({
		url: "/api/list",
		type: "GET",

		success: function(response) {
			let data = response["data"]
			let len = response["data"].length
		}
	})
}

function postData(todo, doing, done, trash) {
	$.ajax({
		url: "/api/api/todo/list",
		type: "POST",
		data: {
			"todo": todo,
			"doing": done,
			"done": done,
			"trash": trash,
		},
		dataType: "json",

		success: function(response) {
			console.log(response)
		},

		error: function(response) {
		}
	})
}

function login(username, password) {
	$.ajax({
		url: "/api/login",
		type: "POST",
		data: '{"username": "sk1122", "password": "satyam#789"}',
		dataType: "text",

		success: function(response) {
			var a = JSON.parse(response)
			console.log(a.token.access_token)
		},

		error: function(response) {
			console.log(response)
		}
	})
}