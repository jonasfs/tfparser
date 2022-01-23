<script>
import GreetingIntro from '@/components/GreetingIntro'
import GreetingSetPath from '@/components/GreetingSetPath'
import GreetingSetPlayer from '@/components/GreetingSetPlayer'
import GreetingEnd from '@/components/GreetingEnd'

export default {
	name: 'FirstTime',
	components: {
		GreetingIntro,
		GreetingSetPath,
		GreetingSetPlayer,
		GreetingEnd,
	},
	data() {
		return {
			step: 0,
			demosFound: 0,
		}
	},
	methods: {
		next(step) {
			this.step = step
		},
		setParsing(num) {
			this.demosFound = num
		},
		toggleFirstTime() {
			this.$emit('toggleFirstTime')
		},
	},
};
</script>

<template>
	<v-row id="row" align="center">
		<v-col align="center">
			<transition name="slide" mode="out-in" class="mx-auto">
				<div v-if="step == 0" key="1">
					<GreetingIntro @next="next" data-test="GreetingIntro" />
				</div>
				<div v-else-if="step == 1" key="2">
					<GreetingSetPath @next="next" @setParsing="setParsing" data-test="GreetingSetPath" />
				</div>
				<div v-else-if="step == 2" key="3">
					<GreetingSetPlayer
						:demosFound="demosFound"
						@next="next"
						data-test="GreetingSetPlayer"
					/>
				</div>
				<div v-else-if="step == 3" key="4">
					<GreetingEnd
						data-test="GreetingEnd"
						@toggleFirstTime="toggleFirstTime"
					/>
				</div>
			</transition>
		</v-col>
	</v-row>
</template>

<style scoped>
#row
{
	height: 100vh;
}

.slide-enter-active
{
	transition: all 0.5s;
}
.slide-leave-active
{
	transition: all 0.5s;
}

.slide-enter
{
  transform: translateX(50%);
	opacity: 0;
}

.slide-leave-to
{
  transform: translateX(-50%);
	opacity: 0;
}

</style>
