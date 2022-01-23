// 0=draw, 1=team1, 2=team2
export const getWinnerSide = (match) => {
	if (match.surrender == 1) {
		return 2
	} else if (match.surrender == 2) {
		return 1
	} else {
		if (match.finalScore1 > match.finalScore2) {
			return 1
		} else if (match.finalScore2 > match.finalScore1) {
			return 2
		} else {
			return 0
		}
	}
}

// 0=none, 1=team1, 2=team2
export const getPlayerSide = (steamid, match) => {
	if (Object.hasOwnProperty.call(match.team1, steamid)) {
		return 1
	} else if (Object.hasOwnProperty.call(match.team2, steamid)) {
		return 2
	} else {
		return 0
	}
}

export const getMatchScoreString = (steamid, match) => {
	let playerSide = 0
	let playerSideScore = 0
	let otherSideScore = 0
	playerSide = getPlayerSide(steamid, match)
	let scoreString = 'finalScore'+playerSide
	playerSideScore = match[scoreString].padStart(2, '0')
	if (playerSide == 1) {
		otherSideScore = match['finalScore2'].padStart(2, '0',)
	} else if (playerSide == 2) {
		otherSideScore = match['finalScore1'].padStart(2, '0')
	} else {
		console.log('player side error')
	}
	return(playerSideScore + ' - ' + otherSideScore)
}

export const getPlayerList = (match) => {
	let playerList = []
	Object.values(match.team1).forEach((player) => playerList.push(player))
	Object.values(match.team2).forEach((player) => playerList.push(player))

	return playerList;
}
