import Vue from 'vue'
import Vuetify from 'vuetify'

// Components
import Main from '@/components/Main.vue'
import FirstTime from '@/components/FirstTime.vue'

// Utilities
import {
  mount,
  createLocalVue
} from '@vue/test-utils'

Vue.use(Vuetify)

describe('Main.vue', () => {

	beforeEach(() => {
		global.backend = require('@/plugins/backendMock')
	})

	it('renders for the first time', done => {
		global.backend.checkFirstTime = () => new Promise(
			resolve => {
				let message = {}
				message.name = 'success'
				message.payload = true
				resolve(message)
			}
		)
		const wrapper = mount(Main)

		setTimeout(() => {
			expect(wrapper.vm.firstTime).toBe(true)
			expect(wrapper.findComponent(FirstTime).exists()).toBe(true)
			done()
		})
	})

	it('not the first time', done => {
		global.backend.checkFirstTime = () => new Promise(
			resolve => {
				let message = {}
				message.name = 'success'
				message.payload = false
				resolve(message)
			}
		)
		const wrapper = mount(Main)
		setTimeout(() => {
			expect(wrapper.vm.firstTime).toBe(false)
			expect(wrapper.findComponent(FirstTime).exists()).toBe(false)
			done()
		})
	})
})
