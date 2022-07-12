<script>
import { DateTime } from 'luxon'

export default {
	name: 'GreetingSetPlayer',
	props: {
		demosFound: {
			type: Number,
		},
	},
	data() {
		return {
			headers: [
				{ text: 'Name', value: 'nickname' },
				{ text: 'Demos', value: 'demos' },
				{ text: 'Last Played', value: 'last' },
			],
			players: {},
			sortBy: ['demos', 'last'],
			search: '',
			parsing: 0,
		}
	},
	methods: {
		startParser() {
			global.backend.initializeParser().then((message) => {
				console.log("debug parserInitialized")
				console.log(message)
			})
		},
		getPlayers() {
			console.log("debug getPlayerList")
			global.backend.getPlayerList().then((message) => {
				this.players = Object.values(message.payload)
				console.log(message)
			})
		},
		pickPlayer(steamid) {
			var self = this
			global.backend.setPlayerProfile(steamid).then((message) => {
				const {payload} = message
				if (payload) {
					localStorage.setItem("steamid", steamid)
					self.$emit('next', 3)
				}
			})
		},
		formatDate(timestamp) {
			return DateTime.fromISO(timestamp).toRelativeCalendar() 
		},
	},
	created() {
		this.parsing = this.demosFound
		console.log("debug - listening")
		global.backend.ee.on('backend-msg', (msg) => {
			console.log("backend-msg:")
			console.log(msg)
			if (msg.name === 'parsing') {
				this.parsing = parseInt(msg.payload)
			} else if (msg.name === 'parsed') {
				console.log("parsed triggered")
				this.parsing -= 1
				this.getPlayers()
			} else if (msg.name === 'playerSet') {
				this.$emit('next', 3)
			}
		})
		this.getPlayers()
		this.startParser()
	},
};
</script>

<template>
	<div>
		<p v-if="parsing > 0" data-test="parsing-msg">
			We are currently parsing
			<span data-test="parsing-left">{{ parsing }}</span>
			match{{ parsing > 1? 'es' : '' }}
		</p>
		<p>
			Pick a player from the list below
			<span v-if="parsing > 0"> or wait until they appear in it after parsing is done </span>
		</p>
		<v-card data-test="card" :data-test-length="Object.values(players).length">
			<v-card-title>
				Players
				<v-spacer></v-spacer>
				<v-text-field
					v-model="search"
					label="Search"
					append-icon="mdi-magnify"
					single-line
					hide-details
					data-test="search-bar"
				></v-text-field>
			</v-card-title>
			<v-data-table
				multi-sort
				:headers="headers"
				:items="Object.values(players)"
				:sort-by.sync="sortBy"
				:sort-desc="[true, true]"
				:search="search"
				:loading="parsing > 0"
				:footer-props="{
					disableItemsPerPage: true,
					itemsPerPage: 10,
					itemsPerPageText: '',
				}"
			>
				<template v-slot:item="{ item }">
					<tr
						class="player"
						:key="item.steamid"
						data-test="parsed-player"
						:data-test-id="item.steamid"
						@click="pickPlayer(item.steamid)"
					>
						<td>{{item.nickname}}</td>
						<td data-test="demos">{{item.demos}}</td>
						<td> {{ formatDate(item.last) }} </td>

					</tr>
				</template>
			</v-data-table>
		</v-card>
	</div>
</template>

<style scope>
.v-data-footer__select {
	display: none !important;
}

.player {
	cursor: pointer;
}
</style>
