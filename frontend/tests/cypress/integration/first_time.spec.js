/// <reference types="cypress" />

context('Start up app for the first time', () => {
  beforeEach(() => {
		cy.window().then((win) => {
			win.backend = require('../../../src/plugins/backendMock')
		})
		cy.visit('/')
  })

  it('actually not the first time', () => {
		cy.window().then((win) => {
			win.backend = require('../../../src/plugins/backendMock')
			cy.stub(win.backend, 'checkFirstTime').resolves({name: 'success', payload: false})
		})
		cy.wait(3000)
		cy.get('[data-test=FirstTime]').should('not.be.visible')
		cy.get('[data-test=GreetingIntro]').should('not.be.visible')
	})

  it('returns error when invalid path is entered', () => {
		cy.window().then((win) => {
			win.backend = require('../../../src/plugins/backendMock')
			cy.stub(win.backend, 'validatePath').resolves({name: 'error', payload: 'os.IsNotExist'})
		})
		cy.get('[data-test=GreetingIntro]').should('be.visible')
		cy.get('[data-test=button]').click()
		cy.get('[data-test=GreetingSetPath]').should('be.visible')
		cy.get('[data-test=button]').click()
		cy.get('[data-test=GreetingSetPlayer]').should('not.be.visible')
	})

  it('returns error when no demos were found', () => {
		cy.window().then((win) => {
			win.backend = require('../../../src/plugins/backendMock')
			cy.stub(win.backend, 'validatePath').resolves({name: 'error', payload: 'no demo files'})
		})
		cy.get('[data-test=GreetingIntro]').should('be.visible')
		cy.get('[data-test=button]').click()
		cy.get('[data-test=GreetingSetPath]').should('be.visible')
		cy.get('[data-test=button]').click()
		cy.get('[data-test=GreetingSetPlayer]').should('not.be.visible')
	})

  it('set up first session successfully', () => {
		cy.window().then((win) => {
			win.backend = require('../../../src/plugins/backendMock')
			cy.stub(win.backend, 'validatePath').resolves({name: 'success', payload: 5, path: 'testPath'})
		})
		cy.get('[data-test=button]').click()
		cy.get('[data-test=GreetingSetPath]').should('be.visible')
		cy.get('[data-test=button]').click()
		cy.get('[data-test=GreetingSetPlayer]').should('be.visible')
		cy.get('[data-test=parsing-msg]').should('be.visible')
		cy.get('[data-test=parsing-left]').then(($span) => {
			const num = parseInt($span.text())
			expect(num).to.eq(5)
		})
		cy.get('[data-test=card').should('have.attr', 'data-test-length', '0')
		cy.window().then((win) => {
			cy.fixture('parsedMatch').then((json) => {
				win.backend.ee.emit('backend-msg', {
					name: 'parsed',
					payload: json,
				})
			})
			cy.get('[data-test=parsed-player]').should('have.length', 10)
		})
		cy.get('[data-test=card').should('have.attr', 'data-test-length', '10')
		cy.window().then((win) => {
			cy.fixture('parsedMatch').then((json) => {
				const team1 = [
						{steamid: 1, nickname: 'name1', rank: '0'},
						{steamid: 11, nickname: 'name11', rank: '0'},
						{steamid: 12, nickname: 'name12', rank: '0'},
						{steamid: 13, nickname: 'name13', rank: '0'},
						{steamid: 14, nickname: 'name14', rank: '0'},
				]
				const team2 = [
						{steamid: 15, nickname: 'name15', rank: '0'},
						{steamid: 16, nickname: 'name16', rank: '0'},
						{steamid: 17, nickname: 'name17', rank: '0'},
						{steamid: 18, nickname: 'name18', rank: '0'},
						{steamid: 7, nickname: 'name7', rank: '0'},
				]
				json.team1 = team1
				json.team2 = team2
				win.backend.ee.emit('backend-msg', {
					name: 'parsed',
					payload: json,
				})
			})
		})
		cy.wait(500)
		cy.get('[data-test-id=1] > [data-test=demos]').then(($td) => {
			const num = parseInt($td.text())
			expect(num).to.eq(2)
		})
		cy.get('[data-test=card').should('have.attr', 'data-test-length', '18')
		cy.window().then((win) => {
			cy.fixture('parsedMatch').then((json) => {
				json.timestamp += 1
				const team1 = [
						{steamid: 18, nickname: 'new name18', rank: '0'},
						{steamid: 20, nickname: 'name20', rank: '0'},
						{steamid: 21, nickname: 'name21', rank: '0'},
						{steamid: 22, nickname: 'name22', rank: '0'},
						{steamid: 23, nickname: 'name23', rank: '0'},
				]
				const team2 = [
						{steamid: 24, nickname: 'name24', rank: '0'},
						{steamid: 25, nickname: 'name25', rank: '0'},
						{steamid: 26, nickname: 'name26', rank: '0'},
						{steamid: 27, nickname: 'name27', rank: '0'},
						{steamid: 28, nickname: 'name28', rank: '0'},
				]
				json.team1 = team1
				json.team2 = team2

				win.backend.ee.emit('backend-msg', {
					name: 'parsed',
					payload: json,
				})
			})
			cy.get('[data-test=card').should('have.attr', 'data-test-length', '27')
		})
		cy.get('[data-test-id=18]').within(() => {
			cy.get('td').eq(0).then(($td) => {
				const name = $td.text()
				expect(name).to.equal('new name18')
			})
			cy.get('td').eq(1).then(($td) => {
				const num = parseInt($td.text())
				expect(num).to.eq(2)
			})
		})
		cy.window().then((win) => {
			cy.fixture('parsedMatch').then((json) => {
				json.timestamp -= 1
				const team1 = [
						{steamid: 1, nickname: 'old name1', rank: '1'},
						{steamid: 18, nickname: 'name18', rank: '0'},
						{steamid: 31, nickname: 'name31', rank: '0'},
						{steamid: 32, nickname: 'name32', rank: '0'},
						{steamid: 33, nickname: 'name33', rank: '0'},
				]
				const team2 = [
						{steamid: 34, nickname: 'name34', rank: '0'},
						{steamid: 35, nickname: 'name35', rank: '0'},
						{steamid: 36, nickname: 'name36', rank: '0'},
						{steamid: 37, nickname: 'name37', rank: '0'},
						{steamid: 38, nickname: 'name38', rank: '0'},
				]
				json.team1 = team1
				json.team2 = team2

				win.backend.ee.emit('backend-msg', {
					name: 'parsed',
					payload: json,
				})
			})
		})
		cy.get('[data-test=card').should('have.attr', 'data-test-length', '35')
		cy.get('[data-test-id=1]').within(() => {
			cy.get('td').eq(0).then(($td) => {
				const name = $td.text()
				expect(name).to.equal('name1')
			})
		})

		cy.window().then((win) => {
			cy.fixture('parsedMatch').then((json) => {
				json.timestamp += 3
				const team1 = [
						{steamid: 1, nickname: 'name1', rank: '1'},
						{steamid: 18, nickname: 'name18', rank: '1'},
						{steamid: 31, nickname: 'name31', rank: '0'},
						{steamid: 32, nickname: 'name32', rank: '0'},
						{steamid: 33, nickname: 'name33', rank: '0'},
				]
				const team2 = [
						{steamid: 34, nickname: 'name34', rank: '0'},
						{steamid: 35, nickname: 'name35', rank: '0'},
						{steamid: 36, nickname: 'name36', rank: '0'},
						{steamid: 37, nickname: 'name37', rank: '0'},
						{steamid: 38, nickname: 'name38', rank: '0'},
				]
				json.team1 = team1
				json.team2 = team2

				win.backend.ee.emit('backend-msg', {
					name: 'parsed',
					payload: json,
				})
			})
			cy.get('[data-test=parsing-msg]').should('not.be.visible')
		})

		cy.get('[data-test-id=18]').within(() => {
			cy.get('td').eq(0).then(($td) => {
				const name = $td.text()
				expect(name).to.equal('name18')
			})
			cy.get('td').eq(1).then(($td) => {
				const num = parseInt($td.text())
				expect(num).to.eq(4)
			})
			cy.get('td').eq(2).then(($td) => {
				const num = parseInt($td.text())
				expect(num).to.eq(1)
			})
		})

		cy.get('[data-test=card').within(() => {
			cy.get('[data-test=search-bar]').type('name18')
			cy.get('[data-test-id=1]').should('not.be.visible')
			cy.get('[data-test-id=18]').click()
		})

		cy.get('[data-test=GreetingEnd]').should('be.visible')
		cy.wait(6000)
		cy.get('[data-test=FirstTime]').should('not.be.visible')
		cy.get('[data-test=GreetingEnd]').should('not.be.visible')
	})
})
