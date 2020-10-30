var ws
var userelem
var msgelem
var logelem
var events = []

setup = (function() {
	userelem = document.getElementById('user')
	msgelem = document.getElementById('msg')
	logelem = document.getElementById('log')
})

connect = (function(user) {
	if (ws) {
		ws.close()
	}

	ws = new WebSocket('ws://' + document.location.host + '/talk/' + user + document.location.pathname)
	ws.onmessage = onevent
})

connectbtn = (function() {
	connect(userelem.value)
})

onevent = (function(event) {
	event = JSON.parse(event.data)
	event.when = Date.parse(event.when)
	console.log(event)
	events.push(event)

	p = document.createElement("p")
	p.textContent = JSON.stringify(event)
	logelem.prepend(p)
})

send = (function(event) {
	ws.send(JSON.stringify(event))
})

msgevent = (function(msg) {
	return {
		'type': 'msg',
		'msg': msg,
	}
})

sendbtn = (function() {
	send(msgevent(msgelem.value))
})
