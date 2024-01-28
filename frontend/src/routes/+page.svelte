<script>
	import { Button } from "$lib/components/ui/button";

	const log = () => {
		count++;
	};
	let count = 0;
</script>

<svelte:head>
	<title>Here We Go</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<section>
	<Button on:click={log}>Hello {count}</Button>
</section>
