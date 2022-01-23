<script>
export default {
	name: 'GreetingIntro',
	methods: {
		click() {
			this.$emit('next', 1)
		}
	},
};
</script>

<template>
	<div>
		<p class="my-4">
			It looks like you're running the parser for the first time.
		</p>
		<p>
			We need to setup a few things before we start. Click on the GO button to proceed.
		</p>
		<p class="my-4">
			<v-btn color="success" @click="click" data-test="button">
				GO
			</v-btn>
		</p>
	</div>
</template>
