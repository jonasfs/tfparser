import Vue from 'vue'
import Vuetify from 'vuetify'

// Components
import GreetingSetPath from '@/components/GreetingSetPath.vue'

// Utilities
import {
  mount,
} from '@vue/test-utils'

Vue.use(Vuetify)

describe('GreetingSetPath.vue', () => {
	let ctx = {}
	beforeEach(() => {
		global.backend = require('@/plugins/backendMock')
		ctx.wrapper = mount(GreetingSetPath)
		ctx.button = ctx.wrapper.find('.v-btn')
		ctx.event = jest.fn()
		ctx.wrapper.vm.$on('next', ctx.event)
	})

	it('renders a button', () => {
		expect(ctx.button.exists()).toBe(true)
	})

	it('loading on button click', async () => {
		expect(ctx.wrapper.vm.loading).toBe(false)
		global.backend.validatePath = () => new Promise(
			() => {
			}
		)
		await ctx.button.trigger('click')
		expect(ctx.wrapper.vm.loading).toBe(true)
		expect(ctx.event).toHaveBeenCalledTimes(0)
	})

	it('backend returned an error', done => {
		global.backend.validatePath = path => new Promise(
			resolve => {
				let message = {}
				message.name = 'error'
				message.payload = 'os.IsNotExist'
				message.path = path
				resolve(message)
			}
		)
		ctx.button.trigger('click')
		setTimeout(() => {
			expect(ctx.wrapper.vm.loading).toBe(false)
			expect(ctx.wrapper.vm.error).toBe(true)
			expect(ctx.event).toHaveBeenCalledTimes(0)
			done()
		})
	})

	it('path was valid', done => {
		global.backend.validatePath = path => new Promise(
			resolve => {
				let message = {}
				message.name = ''
				message.payload = ''
				message.path = path
				resolve(message)
			}
		)
		ctx.button.trigger('click')
		setTimeout(() => {
			expect(ctx.wrapper.vm.loading).toBe(false)
			expect(ctx.wrapper.vm.error).toBe(false)
			expect(ctx.event).toHaveBeenCalledTimes(1)
			done()
		})
	})
})
