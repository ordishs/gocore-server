<script>
	import { onMount } from 'svelte';

	import Header from '../components/Header.svelte';
	import Service from '../components/Service.svelte';
	import { Circle } from 'svelte-loading-spinners';

	const apiURL = '/api';
	let banner = 'Loading data from gocore_server at ' + apiURL;
	let data = [];
	let filter = '';
	let details = false;
	let filteredData = [];
	let goodComms = true;

	$: filteredData = data.filter((d) => {
		const searchText = `${d.serviceName}|${d.host}|${d.context}|${d.commit}|${d.address}`;
		const terms = filter.trim().split(/\s+/); // Split by one or more spaces

		for (let i = 0; i < terms.length; i++) {
			if (!searchText.includes(terms[i])) {
				return false;
			}
		}
		return true;
	});

	async function getDataWithTimeout(url, timeoutMillis = 5000) {
		const controller = new AbortController();
		const timeoutId = setTimeout(() => controller.abort(), timeoutMillis);

		console.log(`Fetching data from ${url}...`);
		try {
			const response = await fetch(url, {
				method: 'GET',
				signal: controller.signal
			});
			goodComms = true;
			return await response.json();
		} catch (err) {
			goodComms = false;
			throw err;
		} finally {
			clearTimeout(timeoutId);
			banner = '';
		}
	}

	onMount(function () {
		console.log(`OnMount: Fetching data from ${apiURL}...`);
		setInterval(async () => {
			data = await getDataWithTimeout(apiURL);
		}, 2000);
	});

	async function removeMe(json) {
		try {
			const url = `${apiURL}/${json.host}/${json.address}/${json.serviceName}`;
			await fetch(url, { method: 'delete' });
			data = await getDataWithTimeout(apiURL);
		} catch (err) {
			console.log(err);
		}
	}
</script>

<svelte:head>
	<link
		rel="stylesheet"
		href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css"
	/>
	<link
		rel="stylesheet"
		href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap-theme.min.css"
	/>
</svelte:head>

<main>
	{#if banner}
		<div class="outer">
			<div class="popup">
				<div>{banner}</div>
				<div class="spinwrap">
					<Circle size="50" color="#A82124" />
				</div>
			</div>
		</div>
	{/if}

	<Header bind:filter bind:details />

	<div class="container-fluid">
		{#each filteredData as d}
			<Service {removeMe} json={d} {details} />
		{/each}

		{#if !goodComms}
			<div class="outer">
				<div class="popup bg-danger">
					<div id="msg">Cannot connect to gocore_server on {apiURL}</div>
				</div>
			</div>
		{/if}
	</div>
</main>

<style>
	main {
		margin-top: 160px;
	}

	@media (min-width: 768px) {
		main {
			margin-top: 51px;
		}
	}

	.outer {
		position: fixed;
		left: 0;
		right: 0;
		height: 100vh;
		min-width: 100%;
		display: flex;
		align-items: center;
		justify-content: center;
		justify-items: center;
		z-index: 9999;
	}
	.popup {
		/* display: flex; */
		/* align-items: center; */
		margin-bottom: 100px;
	}

	.spinwrap {
		display: flex;
		justify-content: center;
	}

	#msg {
		padding: 60px;
	}
</style>
