<script>
import * as matchUtils from '@/plugins/matchUtils'

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
				{ text: 'Rank', value: 'rank' },
				{ text: 'Last Played', value: 'last' },
			],
			players: {},
			sortBy: 'demos',
			search: '',
			parsing: 0,
		}
	},
	methods: {
		addPlayer(steamid, nickname, rank, timestamp) {
			let newPlayer = !(Object.prototype.hasOwnProperty.call(this.players, steamid))
			let player
			if (newPlayer) {
				player = {
					steamid: steamid,
					nickname: nickname,
					rank: rank,
					demos: 1,
					last: timestamp,
				}
			} else {
				player = this.players[steamid]
				if (timestamp > player.last) {
					player.last = timestamp
					player.nickname = nickname
					player.rank = rank
				}
				player.demos += 1
			}
			this.players[steamid] = player
		},
		addPlayers(playerArr, timestamp) {
			playerArr.forEach((player) => {
				const { steamid, nickname, rank } = player
				this.addPlayer(steamid, nickname, rank, timestamp)
			})
		},
		startParser() {
			global.backend.initializeParser().then((message) => {
				console.log("debug parserInitialized")
				console.log(message)
			})
		},
		pickPlayer(steamid) {
			global.backend.setPlayerProfile(steamid)
		},
	},
	created() {
		this.parsing = this.demosFound
		global.backend.ee.on('backend-msg', (msg) => {
			if (msg.name === 'parsing') {
				this.parsing = parseInt(msg.payload)
			} else if (msg.name === 'parsed') {
				let playerList = matchUtils.getPlayerList(msg.payload)
				this.addPlayers(playerList, msg.payload.timestamp)
				this.parsing -= 1
			} else if (msg.name === 'playerSet') {
				this.$emit('next', 3)
			}
		})
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
				:headers="headers"
				:items="Object.values(players)"
				:sort-by.sync="sortBy"
				:sort-desc="true"
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
						<td>{{item.rank}}</td>
						<td>{{item.last}}</td>
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
