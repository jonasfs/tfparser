/// <reference types="cypress" />

context('Main screen tests', () => {
	const fetchMatchesResult = {
		name: 'success',
		steamid: 1,
		payload: {
			'matchHash1': {
				fileHash: 'matchHash1',
				timestamp: 1604966400,
				mapName: 'de_dust2',
				finalScore1: '13',
				finalScore2: '16',
				surrender: 0,
				team1: {
					'1': {
						steamid: '1',
						nickname: 'name1',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'2': {
						steamid: '2',
						nickname: 'name2',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'3': {
						steamid: '3',
						nickname: 'name3',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'4': {
						steamid: '4',
						nickname: 'name4',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'5': {
						steamid: '5',
						nickname: 'name5',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
				},
				team2: {
					'6': {
						steamid: '6',
						nickname: 'name6',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'7': {
						steamid: '7',
						nickname: 'name7',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'8': {
						steamid: '8',
						nickname: 'name8',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'9': {
						steamid: '9',
						nickname: 'name9',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'10': {
						steamid: '10',
						nickname: 'name10',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
				},
			},
			'matchHash2': {
				fileHash: 'matchHash2',
				timestamp: 1605986400,
				mapName: 'de_dust2',
				finalScore1: '16',
				finalScore2: '10',
				surrender: 0,
				team1: {
					'1': {
						steamid: '1',
						nickname: 'name1',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'2': {
						steamid: '2',
						nickname: 'name2',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'3': {
						steamid: '3',
						nickname: 'name3',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'4': {
						steamid: '4',
						nickname: 'name4',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'5': {
						steamid: '5',
						nickname: 'name5',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
				},
				team2: {
					'6': {
						steamid: '6',
						nickname: 'name6',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'7': {
						steamid: '7',
						nickname: 'name7',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'8': {
						steamid: '8',
						nickname: 'name8',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'9': {
						steamid: '9',
						nickname: 'name9',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'10': {
						steamid: '10',
						nickname: 'name10',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
				},
			},
			'matchHash3': {
				fileHash: 'matchHash3',
				timestamp: 1605686400,
				mapName: 'de_overpass',
				finalScore1: '15',
				finalScore2: '15',
				surrender: 0,
				team1: {
					'1': {
						steamid: '1',
						nickname: 'name1',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'2': {
						steamid: '2',
						nickname: 'name2',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'3': {
						steamid: '3',
						nickname: 'name3',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'4': {
						steamid: '4',
						nickname: 'name4',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'5': {
						steamid: '5',
						nickname: 'name5',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
				},
				team2: {
					'6': {
						steamid: '6',
						nickname: 'name6',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'7': {
						steamid: '7',
						nickname: 'name7',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'8': {
						steamid: '8',
						nickname: 'name8',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'9': {
						steamid: '9',
						nickname: 'name9',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'10': {
						steamid: '10',
						nickname: 'name10',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
				},
			},
			'matchHash4': {
				fileHash: 'matchHash4',
				timestamp: 1502996400,
				mapName: 'de_nuke',
				finalScore1: '6',
				finalScore2: '10',
				surrender: 0,
				team1: {
					'6': {
						steamid: '6',
						nickname: 'name6',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'2': {
						steamid: '2',
						nickname: 'name2',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'3': {
						steamid: '3',
						nickname: 'name3',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'4': {
						steamid: '4',
						nickname: 'name4',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'5': {
						steamid: '5',
						nickname: 'name5',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
				},
				team2: {
					'1': {
						steamid: '1',
						nickname: 'name1',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'7': {
						steamid: '7',
						nickname: 'name7',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'8': {
						steamid: '8',
						nickname: 'name8',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'9': {
						steamid: '9',
						nickname: 'name9',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
					'10': {
						steamid: '10',
						nickname: 'name10',
						kills: '0',
						assists: '0',
						deaths: '0',
					}, 
				},
			},


		}
	}


  it('home screen renders', () => {

		cy.visit('/', {
			onBeforeLoad(win) {
				win.backend = require('../../../src/plugins/backendMock')
				cy.stub(win.backend, 'checkFirstTime').resolves({name: 'success', payload: false}).as('checkFirstTime')
				cy.stub(win.backend, 'fetchMatches').withArgs(1).resolves(fetchMatchesResult).as('fetchMatches')
			}
		}).then(() => {
			cy.get('[data-test=Main]').should('be.visible')
			cy.get('[data-test=Drawer]').should('be.visible')
			cy.get('[data-test=MainContent]').should('be.visible')
			cy.get('[data-test=container').should('have.attr', 'data-test-length', '4')
		})
		cy.get('[data-test=container').should('have.attr', 'data-test-parsing', '0')
		cy.window().then((win) => {
			win.backend.ee.emit('backend-msg', {
				name: 'parsing',
				payload: 1,
			})
		})
		cy.get('[data-test=container').should('have.attr', 'data-test-parsing', '1')


		cy.window().then((win) => {
			win.backend.ee.emit('backend-msg', {
				name: 'parsed',
				payload: {
					fileHash: 'matchHash5',
					timestamp: 1602966400,
					mapName: 'de_dust2',
					finalScore1: '14',
					finalScore2: '16',
					surrender: 0,
					team1: {
						'1': {
							steamid: '1',
							nickname: 'name1',
							kills: '0',
							assists: '0',
							deaths: '0',
						}, 
						'2': {
							steamid: '2',
							nickname: 'name2',
							kills: '0',
							assists: '0',
							deaths: '0',
						}, 
						'3': {
							steamid: '3',
							nickname: 'name3',
							kills: '0',
							assists: '0',
							deaths: '0',
						}, 
						'4': {
							steamid: '4',
							nickname: 'name4',
							kills: '0',
							assists: '0',
							deaths: '0',
						}, 
						'5': {
							steamid: '5',
							nickname: 'name5',
							kills: '0',
							assists: '0',
							deaths: '0',
						}, 
					},
					team2: {
						'6': {
							steamid: '6',
							nickname: 'name6',
							kills: '0',
							assists: '0',
							deaths: '0',
						}, 
						'7': {
							steamid: '7',
							nickname: 'name7',
							kills: '0',
							assists: '0',
							deaths: '0',
						}, 
						'8': {
							steamid: '8',
							nickname: 'name8',
							kills: '0',
							assists: '0',
							deaths: '0',
						}, 
						'9': {
							steamid: '9',
							nickname: 'name9',
							kills: '0',
							assists: '0',
							deaths: '0',
						}, 
						'10': {
							steamid: '10',
							nickname: 'name10',
							kills: '0',
							assists: '0',
							deaths: '0',
						}, 
					},
				},
			})
		})
		cy.get('[data-test=container').should('have.attr', 'data-test-length', '5')
		cy.get('[data-test=container').should('have.attr', 'data-test-parsing', '0')
	})
})
