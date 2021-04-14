var todo = []
var doing = []
var done = []
var trash = []


var list_items = document.querySelectorAll(".list-item");
var lists = document.querySelectorAll(".list")

var draggedItem = null;

var list_items_length = list_items.length;

function check_drag() {
	for(let i = 0; i < list_items_length; i++) {
		const item = list_items[i];

		item.addEventListener("dragstart", function() {
			draggedItem = item;
			if(draggedItem.parentElement.classList.contains("todo")) {
				var x = todo.indexOf(draggedItem.textContent);
				todo.splice(x, 1);
			} else if(draggedItem.parentElement.classList.contains("doing")) {
				var x = doing.indexOf(draggedItem.textContent);
				doing.splice(x, 1);
			} else if(draggedItem.parentElement.classList.contains("done")) {
				var x = done.indexOf(draggedItem.textContent);
				done.splice(x, 1);
			} else if(draggedItem.parentElement.classList.contains("trash")) {
				var x = trash.indexOf(draggedItem.textContent);
				trash.splice(x, 1);
			}
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
				list.style.background = "rgba(0, 0, 0, 0.4)"
			})

			list.addEventListener('dragleave', function(e) {
				e.preventDefault();
				list.style.background = "rgba(0, 0, 0, 0.0)"
			})

			list.addEventListener('drop', function() {
				list.append(draggedItem)
				if(list.classList.contains("todo")) {
					if(!todo.includes(draggedItem.textContent))
						todo.push(draggedItem.textContent);
				} else if(list.classList.contains("doing")) {
					if(!doing.includes(draggedItem.textContent))
						doing.push(draggedItem.textContent)
				} else if(list.classList.contains("done")) {
					if(!done.includes(draggedItem.textContent))
						done.push(draggedItem.textContent)
				} else if(list.classList.contains("trash")) {
					if(!trash.includes(draggedItem.textContent))
						trash.push(draggedItem.textContent)
				}
				list.style.background = "rgba(0, 0, 0, 0.0)"
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

	let appendDatas = `<div class="list-item" draggable="true">${formData}</div>`

	lists[0].innerHTML += appendDatas;
	todo.push(formData);

	list_items = document.querySelectorAll(".list-item");
})

var target = document.querySelector('.todo')
// create an observer instance
var observer = new MutationObserver(function(mutations) {
  list_items = document.querySelectorAll(".list-item");
  list_items_length = list_items.length;
  check_drag()
});
// configuration of the observer:
var config = { attributes: true, childList: true, characterData: true };
// pass in the target node, as well as the observer options
observer.observe(target, config);