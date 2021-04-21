var todo = []
var doing = []
var done = []
var trash = []

var list_items = document.querySelectorAll(".item");
var lists = document.querySelectorAll(".list-items")
var container = document.getElementById('container')

function checkDblClickHeader(list) {
	for(let i=0;i<list.length;i++) {
		const element = list[i];
		const header = element.parentElement.firstElementChild

		header.addEventListener('dblclick', function(e) {
			e.preventDefault()
			header.setAttribute("contenteditable", "true")
			console.log("SA")

			header.addEventListener("dblclick", function(e) {
				e.preventDefault()
				header.removeAttribute("contenteditable", "true")
			})
		})
	}
}
checkDblClickHeader(lists)

for(let i = 0; i < lists.length; i++) {
	if(lists[i].childElementCount === 0) {
		lists[i].classList.add("wh")
	} else {
		lists[i].classList.remove("wh")
	}

}

var draggedItem = null;

var list_items_length = list_items.length;

function removeItem(item, list) {
	if(list == 'todo') {
		if(doing.includes(item)) {
			let x = doing.indexOf(item)
			doing.splice(x, 1)
		}

		if(done.includes(item)) {
			let x = done.indexOf(item)
			done.splice(x, 1)
		}

		if(trash.includes(item)) {
			let x = trash.indexOf(item)
			trash.splice(x, 1)
		}
	}

	else if(list == 'doing') {
		if(todo.includes(item)) {
			let x = todo.indexOf(item)
			todo.splice(x, 1)
		}

		if(done.includes(item)) {
			let x = done.indexOf(item)
			done.splice(x, 1)
		}

		if(trash.includes(item)) {
			let x = trash.indexOf(item)
			trash.splice(x, 1)
		}
	}

	else if(list == 'done') {
		if(doing.includes(item)) {
			let x = doing.indexOf(item)
			doing.splice(x, 1)
		}

		if(todo.includes(item)) {
			let x = todo.indexOf(item)
			todo.splice(x, 1)
		}

		if(trash.includes(item)) {
			let x = trash.indexOf(item)
			trash.splice(x, 1)
		}
	}

	else if(list == 'trash') {
		if(doing.includes(item)) {
			let x = doing.indexOf(item)
			doing.splice(x, 1)
		}

		if(done.includes(item)) {
			let x = done.indexOf(item)
			done.splice(x, 1)
		}

		if(todo.includes(item)) {
			let x = todo.indexOf(item)
			todo.splice(x, 1)
		}
	}

	localStorage.setItem("todo", todo)
	localStorage.setItem("doing", doing)
	localStorage.setItem("done", done)
	localStorage.setItem("trash", trash)
}

function addList(list) {
	if(list.classList.contains("todo-list")) {
		if(!todo.includes(draggedItem.textContent)) {
			todo.push(draggedItem.textContent);
			removeItem(draggedItem.textContent, 'todo');
		}
	} else if(list.classList.contains("doing-list")) {
		if(!doing.includes(draggedItem.textContent)) {
			doing.push(draggedItem.textContent)
			removeItem(draggedItem.textContent, 'doing');
		}
	} else if(list.classList.contains("done-list")) {
		if(!done.includes(draggedItem.textContent)) {
			done.push(draggedItem.textContent)
			removeItem(draggedItem.textContent, 'done');
		}
	} else if(list.classList.contains("trash-list")) {
		if(!trash.includes(draggedItem.textContent)) {
			trash.push(draggedItem.textContent)
			removeItem(draggedItem.textContent, 'trash');
		}
	}
	localStorage.setItem("todo", todo)
	localStorage.setItem("doing", doing)
	localStorage.setItem("done", done)
	localStorage.setItem("trash", trash)
}

function check_drag() {
	for(let i = 0; i < list_items_length; i++) {
		const item = list_items[i];

		item.addEventListener("dblclick", function(e) {
			e.preventDefault()
			item.setAttribute("contenteditable", "true")

			item.addEventListener("dblclick", function(e) {
				e.preventDefault()
				item.removeAttribute("contenteditable", "true")
			})
		})

		item.addEventListener("dragstart", function() {
			draggedItem = item;
			if(draggedItem.parentElement.childElementCount === 1) {
				draggedItem.parentElement.classList.add("wh")
			}
			addList(draggedItem.parentElement)
			console.log("Added")
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


// Handles Input
const inputHandler = () => {
	;(function(){

	  var tTime = 100  // transition transform time from #register in ms
	  var wTime = 200  // transition width time from #register in ms
	  var eTime = 1000 // transition width time from inputLabel in ms

	  var toggle = document.getElementById('toggle')

	  // init
	  // --------------
	  var position = 0

	  putQuestion()

	  progressButton.addEventListener('click', done)
	  inputField.addEventListener('keyup', function(e){
	    transform(0, 0) // ie hack to redraw
	    if(e.keyCode == 13) done()
	  })

	  // functions
	  // --------------

	  // load the next question
	  function putQuestion() {
	  	inputLabel.innerHTML = "What do you want to do today?";
	  	toggle.addEventListener('click', function() {
	  		if(toggle.checked) {
	    		inputLabel.innerHTML = "What do you want to do today?";
	  		} else {
	  			inputLabel.innerHTML = "What Column do you want to add?"
	  		}
	  	})
	    inputField.value = ''
	    inputField.type = 'text'  
	    inputField.focus()
	    showCurrent()
	  }
	  
	  // when all the questions have been answered
	  async function done() {
  		if(toggle.checked == true) {
		    let appendDatas = `<div class="item" draggable="true">${inputField.value}</div>`
			lists[0].innerHTML += appendDatas;
			todo.push(inputField.value);
			localStorage.setItem("todo", todo)
			lists[0].classList.remove("wh")

			list_items = document.querySelectorAll(".item");
		} else {
			container.innerHTML += `<div class="${inputField.value} list">
				<div class="header">${inputField.value}</div>
				<div class="list-items ${inputField.value}-list" id="${inputField.value}">
				</div>
			</div>`
			lists = document.querySelectorAll(".list-items")
			checkDblClickHeader(lists)
		}
	  }

	  // helper
	  // --------------

	  function showCurrent(callback) {
	    inputContainer.style.opacity = 1
	    inputProgress.style.transition = ''
	    inputProgress.style.width = '100%'
	    setTimeout(callback, wTime)
	  }

	  function transform(x, y) {
	    register.style.transform = 'translate(' + x + 'px ,  ' + y + 'px)'
	  }

	}())
}

inputHandler();

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
