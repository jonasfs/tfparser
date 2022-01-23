<script>
export default {
	name: 'GreetingEnd',
	data() {
		return {
			transition: 0,
		}
	},
	created() {
		var interval = setInterval(() => {
			if (this.transition == 2) {
				clearInterval(interval)
				this.$emit('toggleFirstTime')
			} else {
				this.transition += 1
			}
		}, 2000)
	},
};
</script>

<template>
	<div>
		<transition name="fade" mode="out-in">
			<p class="my-4">
				Done!
			</p>
		</transition>
		<transition name="fade" mode="out-in">
			<p v-if="transition >= 1">
				You can change the player later in case you want to.
			</p>
		</transition>
		<transition name="fade" mode="out-in">
			<p v-if="transition >= 2" class="my-4">
				Now let's finally get to these statistics...
			</p>
		</transition>
	</div>
</template>

<style scoped>
.fade-enter-active
{
	transition: all 1s;
}

.fade-leave-active
{
	transition: all 1s;
}

.fade-enter
{
	opacity: 0
}

.fade-leave-to
{
	opacity: 1
}
</style>
