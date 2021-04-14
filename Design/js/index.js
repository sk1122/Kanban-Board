var todo = []
var doing = []
var done = []
var trash = []

var list_items = document.querySelectorAll(".item");
var lists = document.querySelectorAll(".list-items")

for(let i = 0; i < lists.length; i++) {
	if(lists[i].childElementCount === 0) {
		lists[i].classList.add("wh")
	} else {
		lists[i].classList.remove("wh")
	}
}

var draggedItem = null;

var list_items_length = list_items.length;

function addList(list) {
	if(list.classList.contains("todo-list")) {
		if(!todo.includes(draggedItem.textContent))
			todo.push(draggedItem.textContent);
	} else if(list.classList.contains("doing-list")) {
		if(!doing.includes(draggedItem.textContent))
			doing.push(draggedItem.textContent)
	} else if(list.classList.contains("done-list")) {
		if(!done.includes(draggedItem.textContent))
			done.push(draggedItem.textContent)
	} else if(list.classList.contains("trash-list")) {
		if(!trash.includes(draggedItem.textContent))
			trash.push(draggedItem.textContent)
	}
	localStorage.setItem("todo", todo)
	localStorage.setItem("doing", doing)
	localStorage.setItem("done", done)
	localStorage.setItem("trash", trash)
}

function check_drag() {
	for(let i = 0; i < list_items_length; i++) {
		const item = list_items[i];

		item.addEventListener("dragstart", function() {
			draggedItem = item;
			if(draggedItem.parentElement.childElementCount === 1) {
				draggedItem.parentElement.classList.add("wh")
			}
			if(draggedItem.parentElement.classList.contains("todo-list")) {
				var x = todo.indexOf(draggedItem.textContent);
				todo.splice(x, 1);
			} else if(draggedItem.parentElement.classList.contains("doing-list")) {
				var x = doing.indexOf(draggedItem.textContent);
				doing.splice(x, 1);
			} else if(draggedItem.parentElement.classList.contains("done-list")) {
				var x = done.indexOf(draggedItem.textContent);
				done.splice(x, 1);
			} else if(draggedItem.parentElement.classList.contains("trash-list")) {
				var x = trash.indexOf(draggedItem.textContent);
				trash.splice(x, 1);
			}
			localStorage.setItem("todo", todo)
			localStorage.setItem("doing", doing)
			localStorage.setItem("done", done)
			localStorage.setItem("trash", trash)
			setTimeout(function() {
				item.style.display = "none";
			}, 0)
		})

		item.addEventListener("dragend", function() {
			setTimeout(function() {
				if(draggedItem != null)
					draggedItem.style.display = 'block';
				draggedItem = null;
			}, 0)
		})

		for(let j = 0; j < lists.length; j++) {
			const list = lists[j];

			list.addEventListener('dragover', function(e) {
				e.preventDefault();
			})

			list.addEventListener('dragenter', function(e) {
				e.preventDefault();
			})

			list.addEventListener('dragleave', function(e) {
				e.preventDefault();
			})

			list.addEventListener('drop', function() {
				list.append(draggedItem)
				list.classList.remove("wh")
				addList(list)
			})
		}
	}
}

check_drag()

var formData = ""

var form = document.getElementById('form');
form.addEventListener("submit", function(e) {
	e.preventDefault()

	var formElements=document.getElementById("form").elements;
	for (var i=0; i<formElements.length; i++)
	    if (formElements[i].type!="submit")
	        formData = formElements[i].value

	let appendDatas = `<div class="item" draggable="true">${formData}</div>`
	lists[0].innerHTML += appendDatas;
	todo.push(formData);
	localStorage.setItem("todo", todo)
	lists[0].classList.remove("wh")

	list_items = document.querySelectorAll(".item");
})

var target = document.querySelector('.todo-list')
// create an observer instance
var observer = new MutationObserver(function(mutations) {
  list_items = document.querySelectorAll(".item");
  list_items_length = list_items.length;
  check_drag()
});
// configuration of the observer:
var config = { attributes: true, childList: true, characterData: true };
// pass in the target node, as well as the observer options
observer.observe(target, config);