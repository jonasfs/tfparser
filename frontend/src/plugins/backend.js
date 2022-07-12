var EventEmitter = require('events')

export const ee = new EventEmitter()

export const checkFirstTime = () => new Promise(
	resolve => {
		astilectron.sendMessage({name:"checkFirstTime"}, (msg) => {
			resolve(msg)
		})
	}
)

export const validatePath = path => new Promise(
	resolve => {
		astilectron.sendMessage({name:"validatePath", payload: path}, (msg) => {
			resolve(msg)
		})
	}
)

export const initializeParser = () => new Promise(
	resolve => {
		astilectron.sendMessage({name:"initializeParser"}, (msg) => {
			resolve(msg)
		})
	}
)

export const setPlayerProfile = steamid => new Promise(
	resolve => {
		astilectron.sendMessage({name:"setPlayerProfile", payload: steamid}, (msg) => {
			resolve(msg)
		})
	}
)

export const getPlayerProfile = () => new Promise(
	resolve => {
		astilectron.sendMessage({name:"getPlayerProfile"}, (msg) => {
			resolve(msg)
		})
	}
)

export const fetchMatches = steamid => new Promise(
	resolve => {
		astilectron.sendMessage({
			name:"getPlayerMatches",
			payload: steamid
		}, (msg) => {
			resolve(msg)
		})
	}
)

export const getPlayerList = () => new Promise(
	resolve => {
		astilectron.sendMessage({name:"getPlayerList"}, (msg) => {
			resolve(msg)
		})
	}
)
