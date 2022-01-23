var EventEmitter = require('events')
 
export const ee = new EventEmitter()

export const checkFirstTime = () => new Promise(
	resolve => {
		let message = {}
		message.name = 'success'
		message.payload = true
		resolve(message)
	}
)

export const validatePath = path => new Promise(
	resolve => {
		let message = {}
		message.path = path
		if (path === 'replays') {
			message.name = 'success'
			message.payload = ''
		} else {
			message.name = 'error'
			message.payload = 'os.IsNotExist'
		}
		resolve(message)
	}
)

export const initializeParser = () => {
}

export const setPlayerProfile = (steamid) => {
	ee.emit('backend-msg', {name: 'playerSet', payload: steamid})
}

export const fetchMatches = (steamid) => new Promise(
	resolve => {
		let message = {}
		let matches = {}
		message.name = 'success'
		message.payload = matches
		message.steamid = steamid
		console.log('fetch - backendMock version called')
		resolve(message)
	}
)
