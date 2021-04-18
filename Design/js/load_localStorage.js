
list = ["todo", "doing", "done", "trash"]
list1 = ["Todo", "Doing", "Done", "Trash"]

function a() {
	let todo = localStorage.getItem("todo")
	todo = todo.split(",")
	let doing = localStorage.getItem("doing")
	doing = doing.split(",")
	let done = localStorage.getItem("done")
	done = done.split(",")
	let trash = localStorage.getItem("trash")
	trash = trash.split(",")

	postData(todo, doing, done, trash)
}

setInterval(a, 10000)

function getHtmlForDrag(data) {
	return `<div class="item" draggable="true">${data}</div>`
}

function appendData(data) {	
	if(data == "")
		return
	document.getElementById(data.Category.toLowerCase()).innerHTML += getHtmlForDrag(data.Todo)
}

function requestData() {
	$.ajax({
		url: "/api/get",
		type: "GET",
		headers: {"Authorization": `Bearer ${localStorage.getItem('access_token')}`},

		success: function(response) {
			let data = response["data"]
			let len = response["data"].length

			console.log(data)

			for(let i=0;i<len;i++)
				appendData(data[i])
		},

		error: function(response) {
			console.log(response)
		}
	})
}

function postData(todo, doing, done, trash) {
	
	let dataToPost = {
		"todo": todo,
		"doing": doing,
		"done": done,
		"trash": trash
	}

	$.ajax({
		url: "/api/todo/list",
		type: "POST",
		data: JSON.stringify(dataToPost),
		headers: {"Authorization": `Bearer ${localStorage.getItem('access_token')}`},
		dataType: "text",

		success: function(response) {
			console.log("Successful")
		},

		error: function(response) {
			console.log("Error")
		}
	})
}

// Document Loads
requestData()