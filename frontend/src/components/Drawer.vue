<script>
import GreetingSetPlayer from '@/components/GreetingSetPlayer'

export default {
	name: 'Drawer',
	components: {
		GreetingSetPlayer,
	},
	data() {
		return {
			user: {},
			setPlayerDialog: false,
		}
	},
	methods: {
		getUser() {
			global.backend.getPlayerProfile().then((message) => {
				this.user = message.payload
			})
		},
		done() {
			this.getUser()
			this.$emit("refreshMain")
			this.setPlayerDialog = false
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
				<v-list-item link @click.stop="setPlayerDialog = true">
					<v-list-item-icon>
					</v-list-item-icon>
					<v-list-item-title >
						Change User
					</v-list-item-title>
				</v-list-item>
			</v-list>

		</v-navigation-drawer>
		<v-dialog
			v-model="setPlayerDialog"
			persistent
		>
			<v-card>
				<GreetingSetPlayer @done="done" />
				<v-card-actions>
					<v-btn
						color="red darken-1"
						@click="done"
						text
						dark
					>
						Cancel
					</v-btn>
				</v-card-actions>
			</v-card>
		</v-dialog>
	</v-card>
</template>
