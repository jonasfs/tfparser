import * as matchUtils from '@/plugins/matchUtils'

const matchFixture = require('../fixtures/parsedMatch')

describe('matchUtils', () => {
	it ('getPlayerList returns 10 players', () => {
		const playerList = matchUtils.getPlayerList(matchFixture)
		expect(playerList.length).toBe(10)
	})

	it ('player from list should have a steamid field', () => {
		const playerList = matchUtils.getPlayerList(matchFixture)
		const player = playerList[1]
		expect(player).toHaveProperty('steamid')
	})
})
