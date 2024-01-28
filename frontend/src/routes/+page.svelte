<script lang="ts">
	import { Button } from "$lib/components/ui/button";
	import { request } from "$lib/services/api-service/apiService";

	let resp: string = "";
	// Get response from test api
	const test = async () => {
		const response = await request("");
		console.log(response);
		resp = response.message;
	};
</script>

<svelte:head>
	<title>Here We Go</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<section>
	<Button on:click={test}>Hello {resp}</Button>
</section>
