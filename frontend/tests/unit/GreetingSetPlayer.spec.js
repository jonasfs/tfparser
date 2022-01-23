import Vue from 'vue'
import Vuetify from 'vuetify'

// Components
import GreetingSetPlayer from '@/components/GreetingSetPlayer.vue'

// Utilities
import {
  mount,
	createLocalVue,
} from '@vue/test-utils'

Vue.use(Vuetify)

describe('GreetingSetPlayer.vue', () => {
	let ctx = {}
	beforeEach(() => {
		const localVue = createLocalVue()
		global.astilectron = {
			onMessage: () => {},
			sendMessage: () => {},
		}
		ctx.wrapper = mount(GreetingSetPlayer, {
			localVue,
			vuetify: new Vuetify()
		})
	})

	it('renders a table', () => {
		expect(ctx.wrapper.find('.v-data-table').exists()).toBe(true)
		expect(ctx.wrapper.vm.parsing).toBe(0)
	})

	it('add a new player', () => {
		ctx.wrapper.vm.addPlayer('121212', 'player 1', '0', 1601859925935)
		let players = Object.values(ctx.wrapper.vm.players)
		expect(players[0].steamid).toBe('121212')
	})

	it('add a new player with non string steamid', () => {
		ctx.wrapper.vm.addPlayer(121212, 'player 1', '0', 1601859925935)
		let players = Object.values(ctx.wrapper.vm.players)
		expect(players[0].steamid).toBe('121212')
	})

	it('add existing player', () => {
		ctx.wrapper.vm.addPlayer('121212', 'player 1', '0', 1601859915935)
		let players = Object.values(ctx.wrapper.vm.players)
		expect(players[0].steamid).toBe('121212')
		expect(players[0].demos).toBe(1)
		expect(players[0].last).toBe(1601859915935)
		expect(players[0].rank).toBe('0')
		ctx.wrapper.vm.addPlayer('121212', 'player 1 new', '1', 1601859925935)
		players = Object.values(ctx.wrapper.vm.players)
		expect(players[0].demos).toBe(2)
		expect(players[0].name).toBe('player 1 new')
		expect(players[0].rank).toBe('1')
		expect(players[0].last).toBe(1601859925935)
	})

	it('add existing player older stat', () => {
		ctx.wrapper.vm.addPlayer('121212', 'player 1 new', '1', 1601859925935)
		ctx.wrapper.vm.addPlayer('121212', 'player 1', '0', 1601859915935)
		let players = Object.values(ctx.wrapper.vm.players)
		expect(players[0].demos).toBe(2)
		expect(players[0].name).toBe('player 1 new')
		expect(players[0].rank).toBe('1')
		expect(players[0].last).toBe(1601859925935)
	})

	it('add multiple players at once', () => {
		let players = [
			{
				steamid: '000',
				name: 'player 0',
				rank: '0',
				timestamp: 1601859925935,
			},
			{
				steamid: '111',
				name: 'player 1',
				rank: '2',
				timestamp: 1601859925935,
			},
			{
				steamid: '222',
				name: 'player 2',
				rank: '1',
				timestamp: 1601859925935,
			},
		]
		ctx.wrapper.vm.addPlayers(players)
		players = ctx.wrapper.vm.players
		expect(players['000'].demos).toBe(1)
		expect(players['000'].steamid).toBe('000')
		expect(players['111'].steamid).toBe('111')
		expect(players['222'].rank).toBe('1')
	})

	it('add multiple players with repetitions', () => {
		let players = [
			{
				steamid: '000',
				name: 'player 0',
				rank: '0',
				timestamp: 1601859925935,
			},
			{
				steamid: '111',
				name: 'player 1',
				rank: '2',
				timestamp: 1601859925935,
			},
			{
				steamid: '222',
				name: 'player 2',
				rank: '1',
				timestamp: 1601859925935,
			},
		]
		ctx.wrapper.vm.addPlayers(players)
		players = [
			{
				steamid: '000',
				name: 'player 0 new',
				rank: '1',
				timestamp: 1601859955935,
			},
			{
				steamid: '333',
				name: 'player 3',
				rank: '2',
				timestamp: 1601859955935,
			},
			{
				steamid: '222',
				name: 'player 2',
				rank: '0',
				timestamp: 1601859955935,
			},
		]
		ctx.wrapper.vm.addPlayers(players)
		let playersLength = Object.keys(ctx.wrapper.vm.players).length
		expect(playersLength).toBe(4)
		players = ctx.wrapper.vm.players
		expect(players['000'].demos).toBe(2)
		expect(players['000'].name).toBe('player 0 new')
		expect(players['000'].rank).toBe('1')
		expect(players['111'].rank).toBe('2')
		expect(players['222'].rank).toBe('0')
		expect(players['222'].demos).toBe(2)
		expect(players['333'].demos).toBe(1)

	})

})
