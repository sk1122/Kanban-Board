
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
		headers: {"Authorization": `Bearer ${localStorage.getItem('access_token')}`},

		success: function(response) {
			let data = response["data"]
			let len = response["data"].length

			console.log(data)
		}
	})
}

function postData(todo, doing, done, trash) {
	$.ajax({
		url: "/api/todo/list",
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