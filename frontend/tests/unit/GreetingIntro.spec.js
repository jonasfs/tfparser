import Vue from 'vue'
import Vuetify from 'vuetify'

// Components
import GreetingIntro from '@/components/GreetingIntro.vue'

// Utilities
import {
  mount,
  createLocalVue
} from '@vue/test-utils'

Vue.use(Vuetify)

describe('GreetingIntro.vue', () => {

	it('renders a button', () => {
		const wrapper = mount(GreetingIntro)
		
		expect(wrapper.find('button').exists()).toBe(true)
	})

	it('clicking the button emmits a "next" event', () => {
		const wrapper = mount(GreetingIntro)
		const event = jest.fn()
		const button = wrapper.find('.v-btn')

		wrapper.vm.$on('next', event)

		expect(event).toHaveBeenCalledTimes(0)

		button.trigger('click')
		expect(event).toHaveBeenCalledTimes(1)
	})

})
