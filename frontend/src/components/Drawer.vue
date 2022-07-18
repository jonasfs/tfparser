<script>
export default {
	name: 'Drawer',
	data() {
		return {
			user: {},
		}
	},
	methods: {
		getUser() {
			global.backend.getPlayerProfile().then((message) => {
				this.user = message.payload
			})
		},
	},
	created() {
		this.getUser()
	},
};
</script>

<template>
	<v-card>
		<v-navigation-drawer
			app
			permanent
			right
			expand-on-hover
			width="200"
			data-test="Drawer"
		>
			<v-list>
				<v-list-item>
					<v-list-item-icon>
						<v-icon>
							mdi-account
						</v-icon>
					</v-list-item-icon>

					<v-list-item-content>
						<v-list-item-title v-text="user.nickname"></v-list-item-title>
					</v-list-item-content>
				</v-list-item>
			</v-list>

			<v-divider></v-divider>

			<v-list nav dense>
				<v-list-item link>
					<v-list-item-icon>
					</v-list-item-icon>
					<v-list-item-title>
						Change User
					</v-list-item-title>
				</v-list-item>
			</v-list>

		</v-navigation-drawer>
	</v-card>
</template>
